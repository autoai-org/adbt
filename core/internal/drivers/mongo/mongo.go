package mongo

import (
	"context"
	"fmt"
	"io"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

type MongoDB struct {
	connectionString string
	mgoLib           mongoLib
	bsonService      bsonService
	mongoTimeout     time.Duration
	rateLimit        time.Duration
	batchLimit       int
}

func NewMongoService(connectionString string, mgoLib mongoLib, bsonService bsonService, mongoTimeout time.Duration, rateLimit time.Duration, batchLimit int) *MongoDB {
	return &MongoDB{
		connectionString: connectionString,
		mgoLib:           mgoLib,
		bsonService:      bsonService,
		mongoTimeout:     mongoTimeout,
		rateLimit:        rateLimit,
		batchLimit:       batchLimit,
	}
}

func NewDefaultMongoService(connectionString string) *MongoDB {
	rateLimit := 15000000
	batchLimit := 250
	mongoTimeout := 60
	return &MongoDB{
		connectionString: connectionString,
		mgoLib:           &labixMongo{},
		bsonService:      &defaultBsonService{},
		mongoTimeout:     time.Duration(mongoTimeout) * time.Second,
		rateLimit:        time.Duration(rateLimit) * time.Millisecond,
		batchLimit:       batchLimit,
	}
}

func (m *MongoDB) DumpCollectionTo(database, collection string, writer io.Writer) error {
	session, err := m.mgoLib.DialWithTimeout(m.connectionString, m.mongoTimeout)
	if err != nil {
		return fmt.Errorf("Couldn't dial mongodb session: %v", err)
	}
	defer session.Close()
	start := time.Now()
	log.Infof("Backing up %s/%s", database, collection)
	iter := session.SnapshotIter(database, collection, nil)
	err = iter.Err()
	if err != nil {
		return fmt.Errorf("Couldn't obtain iterator over collection=%v/%v: %v", database, collection, err)

	}
	for {
		result, hasNext := iter.Next()
		if !hasNext {
			break
		}
		_, err = writer.Write(result)
		if err != nil {
			return err
		}
	}
	log.Infof("Backing up finished for %s/%s. duration=%v", database, collection, time.Now().Sub(start))
	err = iter.Err()
	if err != nil {
		return fmt.Errorf("Error while iterating over collection=%v/%v noticed only at the end: %v", database, collection, err)
	}
	return nil
}

func (m *MongoDB) RestoreCollectionFrom(database, collection string, reader io.Reader) error {
	session, err := m.mgoLib.DialWithTimeout(m.connectionString, 0)
	if err != nil {
		return fmt.Errorf("Error while dialing mongo session: %v", err)
	}
	defer session.Close()

	err = m.clearCollection(session, database, collection)
	if err != nil {
		return fmt.Errorf("Error while clearing collection=%v/%v: %v", database, collection, err)
	}

	start := time.Now()
	log.Infof("Starting restore of %s/%s", database, collection)

	bulk := session.Bulk(database, collection)

	var batchBytes int
	batchStart := time.Now()

	limiter := rate.NewLimiter(rate.Every(m.rateLimit), 1)

	for {
		next, err := m.bsonService.ReadNextBSON(reader)
		if err != nil {
			return fmt.Errorf("Error while reading bson: %v", err)
		}
		if next == nil {
			break
		}

		// If we have something to write and the next doc would push the batch over
		// the limit, write the batch out now. 15000000 is intended to be within the
		// expected 16MB limit
		if batchBytes > 0 && batchBytes+len(next) > m.batchLimit {
			err = bulk.Run()
			if err != nil {
				return fmt.Errorf("Error while writing bulk: %v", err)
			}

			var duration = time.Since(batchStart)
			log.Infof("Written bulk restore batch for %s/%s. Took %v", database, collection, duration)

			// rate limit between writes to prevent overloading MongoDB
			limiter.Wait(context.Background())

			bulk = session.Bulk(database, collection)
			batchBytes = 0
			batchStart = time.Now()
		}

		bulk.Insert(next)

		batchBytes += len(next)
	}
	err = bulk.Run()
	if err != nil {
		return fmt.Errorf("Error while writing bulk: %v", err)
	}
	log.Infof("Finished restore of %s/%s. Duration: %v", database, collection, time.Since(start))
	return nil
}

func (m *MongoDB) clearCollection(session mongoSession, database, collection string) error {
	start := time.Now()
	log.Infof("Clearing collection %s/%s", database, collection)
	err := session.RemoveAll(database, collection, nil)
	log.Infof("Finished clearing collection %s/%s. Duration : %v", database, collection, time.Now().Sub(start))
	return err
}

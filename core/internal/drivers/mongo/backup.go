package mongo

import (
	"dbtools/drivers"
	"dbtools/storages"
	"dbtools/utility"
	"fmt"
)

type MongoBackupService struct {
	mongo          MongoDB
	storageService storages.StorageService
}

func NewMongoBackupService(mongo MongoDB, storageService storages.StorageService) *MongoBackupService {
	return &MongoBackupService{
		mongo,
		storageService,
	}
}

func (m *MongoBackupService) Backup(colls []drivers.DBColl) error {
	date := utility.FormattedNow()
	for _, coll := range colls {
		err := m.backup(date, coll)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *MongoBackupService) backup(date string, coll drivers.DBColl) error {
	w, err := m.storageService.Writer(date, coll.Database, coll.Collection)
	if err != nil {
		return err
	}
	defer w.Close()
	if err = m.mongo.DumpCollectionTo(coll.Database, coll.Collection, w); err != nil {
		// result := drivers.BackupResult{false, time.Now(), coll}
		return fmt.Errorf("dumping failed for %s/%s: %v", coll.Database, coll.Collection, err)
	}
	if err := w.Close(); err != nil {
		return err
	}
	// result = drivers.BackupResult{true, time.Now(), coll}
	return nil
}

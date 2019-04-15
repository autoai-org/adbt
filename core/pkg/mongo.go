package adbt

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"path/filepath"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	URI      string
	Name     string
	Database string
}

func (m *MongoDB) Backup(identifier string) bool {
	log.Println("Backing up " + m.Name + "...")
	params := m.prepare()
	process := Process{
		Command: "mongodump",
		Params:  params,
	}
	return process.Run(identifier)
}

func (m *MongoDB) Restore() {

}

func (m *MongoDB) prepare() []string {
	createFolderIfNotExist("data")
	createFolderIfNotExist(filepath.Join("adbt", "mongo"))
	currentTime := time.Now()
	timelabel := currentTime.Format("2006-01-02")
	return []string{
		"--uri=" + m.URI,
		"--archive=adbt/mongo/" + m.Name + "." + timelabel + ".gz",
		"--forceTableScan",
		"--gzip",
	}
}

func (m *MongoDB) test() bool {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.URI))
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(m.URI)
		fmt.Println(err)
		return false
	}
	return true
}

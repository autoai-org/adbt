package drivers

import (
	"context"
	"fmt"
	"github.com/unarxiv/adbt/pkg"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"path/filepath"
	"time"
)

type MongoDB struct {
	URI      string
	Name     string
	Database string
}

func (m *MongoDB) Backup(identifier string) bool {
	fmt.Println("Backing up " + m.Name + "...")
	params := m.prepare()
	process := adbt.Process{
		Command: "mongodump",
		Params:  params,
	}
	return process.Run(identifier)
}

func (m *MongoDB) Restore() {

}

func (m *MongoDB) prepare() []string {
	adbt.createFolderIfNotExist(filepath.Join("adbt", "mongo"))
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
		return false
	}
	return true
}

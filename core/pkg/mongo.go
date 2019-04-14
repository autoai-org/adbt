package adbt

import (
	"log"
	"path/filepath"
	"time"
)

type MongoDB struct {
	URI      string
	Name     string
	Database string
}

func (m *MongoDB) Backup() bool {
	log.Println("Backing up " + m.Name + "...")
	params := m.prepare()
	process := Process{
		Command: "mongodump",
		Params:  params,
	}
	return process.Run()
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

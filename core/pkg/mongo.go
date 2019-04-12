package adbt

import "time"

type MongoDB struct {
	URI      string
	Name     string
	Database string
}

func (m *MongoDB) Backup() {
	params := m.prepare()
	process := Process{
		Command: "mongodump",
		Params:  params,
	}
	process.Run()
}

func (m *MongoDB) Restore() {

}

func (m *MongoDB) prepare() []string {
	createFolderIfNotExist("data")
	createFolderIfNotExist("data/mongo")
	currentTime := time.Now()
	timelabel := currentTime.Format("2006-01-02")
	return []string{
		"--uri=" + m.URI,
		"--archive=data/mongo/" + m.Name + "." + timelabel + ".gz",
		"--forceTableScan",
		"--gzip",
	}
}

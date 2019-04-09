package drivers

import (
	"time"
)

// DBColl is a generic db/table/collection map.
// In Mongo, it is database-collection
// In MySQL, it is database-table
type DBColl struct {
	Database   string
	Collection string
}

type BackupService interface {
	Backup(colls []DBColl) error
	Restore(dateDir string, colls []DBColl) error
}

type BackupResult struct {
	Success    bool
	Timestamp  time.Time
	Collection DBColl
}

// DatabaseService is the generic struct for various databases
type DatabaseService struct {
	connectionString string
}

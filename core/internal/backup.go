package dbtools

import (
	"dbtools/drivers"
)

type dbColl struct {
	database   string
	collection string
}

type BackupService interface {
	Backup(colls []dbColl) error
	Restore(dateDir string, colls []dbColl) error
}

func newBackupService(backupable drivers.Backupable, storageService storage.storageService) *BackupService {

}

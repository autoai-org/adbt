package dbtools

type dbColl struct {
	database   string
	collection string
}

type BackupService interface {
	Backup(colls []dbColl) error
	Restore(dateDir string, colls []dbColl) error
}

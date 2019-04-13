package adbt

type Backupable interface {
	Backup()
	Restore()
}

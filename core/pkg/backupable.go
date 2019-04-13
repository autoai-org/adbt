package adbt

type Backupable interface {
	Backup() bool
	Restore() bool
}

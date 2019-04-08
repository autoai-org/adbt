package dbtools

import "io"

type backupable interface {
	backup(database, collection string, writer io.Writer) bool
	restore(database, collection string, reader io.Reader) bool
}

func backup(b backupable, database, collection string, writer io.Writer) {
	b.backup(database, collection, writer)
}

func restore(b backupable, database, collection string, reader io.Reader) {
	b.restore(database, collection, reader)
}

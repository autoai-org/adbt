package storages

import (
	"io"
	"os"
	"path/filepath"

	"github.com/golang/snappy"
	log "github.com/sirupsen/logrus"
)

// FileStorageService provides service that
// save snappy files to the disk
type LocalStorageService struct {
	filepath string
}

func NewLocalStorageService(filepath string) *LocalStorageService {
	return &LocalStorageService{
		filepath,
	}
}

func (f *LocalStorageService) Writer(date, database, collection string) (io.WriteCloser, error) {
	path := filepath.Join(f.filepath, date, database, collection+baseExtension)
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	log.Infof("Saving to path=%s", path)

	return newSnappyWriteCloser(snappy.NewBufferedWriter(file), file), nil
}

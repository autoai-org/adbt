package storages

// FileStorageService provides service that
// save snappy files to the disk
type FileStorageService struct {
	filepath string
}

func newFileStorageService(filepath string) *FileStorageService {
	return &FileStorageService{
		filepath,
	}
}

/*
func (f *FileStorageService) Writer(date, database, collection string) (io.WriteCloser, error) {
	path := filepath.Join(f.filepath, date, database, collection+baseExtension)
	log.Infof("saving to path=%s", path)
}
*/

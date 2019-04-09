package storages

import (
	"io"

	"github.com/klauspost/compress/snappy"
)

const baseExtension = ".snappy"

type StorageService interface {
	Writer(date, database, collection string) (io.WriteCloser, error)
	Reader(date, database, collection string) (io.WriteCloser, error)
}

type snappyWriteCloser struct {
	snappyWriter *snappy.Writer
	writeCloser  io.WriteCloser
}

func newSnappyWriteCloser(snappyWriter *snappy.Writer, writeCloser io.WriteCloser) *snappyWriteCloser {
	return &snappyWriteCloser{
		snappyWriter,
		writeCloser,
	}
}

func (swc *snappyWriteCloser) Write(p []byte) (nRet int, errRet error) {
	return swc.snappyWriter.Write(p)
}

func (swc *snappyWriteCloser) Close() error {
	if err := swc.snappyWriter.Close(); err != nil {
		return err
	}
	return swc.writeCloser.Close()
}

type snappyReadCloser struct {
	snappyReader *snappy.Reader
	readCloser   io.ReadCloser
}

func newSnappyReadCloser(snappyReader *snappy.Reader, readCloser io.ReadCloser) *snappyReadCloser {
	return &snappyReadCloser{
		snappyReader,
		readCloser,
	}
}

func (src *snappyReadCloser) Read(p []byte) (int, error) {
	return src.snappyReader.Read(p)
}

func (src *snappyReadCloser) Close() error {
	return src.readCloser.Close()
}

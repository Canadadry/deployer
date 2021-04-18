package fs

import (
	"fmt"
)

var (
	ErrFileNotFound = fmt.Errorf("File not found")
	ErrClosedFile   = fmt.Errorf("Closed file")
)

type FS interface {
	Open(name string) (File, error)
	Delete(name string) error
	Mkdir(name string) error
	New(name string) (File, error)
	ReadDir(name string) ([]FileInfo, error)
}

type File interface {
	Read([]byte) (int, error)
	Write(p []byte) (n int, err error)
	Close() error
	Stat() FileInfo
}

type FileInfo interface {
	IsDir() bool
}

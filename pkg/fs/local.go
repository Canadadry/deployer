package fs

import (
	// "fmt"
	"os"
)

type local struct {
	base string
}

func NewLocal(base string) *local {
	return &local{base: base}
}
func (l *local) Open(name string) (File, error) {
	stat, err := os.Stat(l.base + name)
	if err != nil {
		return nil, ErrFileNotFound
	}
	if stat.IsDir() {
		return &memoryDirectory{}, nil
	}

	f, err := os.OpenFile(l.base+name, os.O_RDWR, 0644)
	return &localFile{file: f}, err
}

func (l *local) Delete(name string) error {
	return os.RemoveAll(l.base + name)
}

func (l *local) Mkdir(name string) error {
	err := os.Mkdir(l.base+name, 0755)
	if err == nil {
		return nil
	}
	return ErrReservedName
}

func (l *local) New(name string) (File, error) {
	f, err := os.OpenFile(l.base+name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return nil, err
	}
	return &localFile{file: f}, nil
}

func (l *local) ReadDir(name string) ([]FileInfo, error) {
	files, err := os.ReadDir(l.base + name)
	if err != nil {
		return nil, err
	}
	fi := []FileInfo{}
	for _, file := range files {
		fi = append(fi, &fileInfo{name: file.Name(), isDir: file.IsDir()})
	}
	return fi, nil
}

type localFile struct {
	file *os.File
}

func (lf *localFile) Read(b []byte) (int, error) {
	if lf.file == nil {
		return 0, ErrClosedFile
	}
	return lf.file.Read(b)
}

func (lf *localFile) Write(b []byte) (int, error) {
	if lf.file == nil {
		return 0, ErrClosedFile
	}
	return lf.file.Write(b)
}

func (lf *localFile) Close() error {
	if lf.file == nil {
		return ErrClosedFile
	}
	lf.file = nil
	return nil
}

func (lf *localFile) Stat() FileInfo {
	return &fileInfo{}
}

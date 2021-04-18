package fs

import (
	"io"
)

type memory struct {
	files map[string]*memoryFile
}

func (m *memory) Open(name string) (File, error) {
	f, ok := m.files[name]
	if !ok {
		return nil, ErrFileNotFound
	}
	f.open = true
	f.pos = 0
	return f, nil
}

func (m *memory) Delete(name string) error {
	return nil
}

func (m *memory) Mkdir(name string) error {
	return nil
}

func (m *memory) New(name string) (File, error) {
	return nil, nil
}

func (m *memory) ReadDir(name string) ([]FileInfo, error) {
	return nil, nil
}

type memoryFile struct {
	content []byte
	pos     int
	open    bool
}

func (mf *memoryFile) Read(b []byte) (int, error) {
	if mf.open == false {
		return 0, ErrClosedFile
	}
	if mf.pos >= len(mf.content) {
		return 0, io.EOF
	}
	n := copy(b, mf.content[mf.pos:])
	mf.pos += n
	return n, nil
}

func (mf *memoryFile) Write(b []byte) (int, error) {
	if mf.open == false {
		return 0, ErrClosedFile
	}
	mf.content = append(mf.content, b...)
	return len(b), nil
}

func (mf *memoryFile) Close() error {
	if mf.open == false {
		return ErrClosedFile
	}
	mf.open = false
	return nil
}

func (mf *memoryFile) Stat() FileInfo {
	return &fileInfo{}
}

type fileInfo struct {
	isDir bool
}

func (fi *fileInfo) IsDir() bool {
	return fi.isDir
}

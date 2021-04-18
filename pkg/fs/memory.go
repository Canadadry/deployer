package fs

import (
	"io"
	"strings"
)

type memory struct {
	files map[string]File
}

func (m *memory) Open(name string) (File, error) {
	f, ok := m.files[name]
	if !ok {
		return nil, ErrFileNotFound
	}
	mf, ok := f.(*memoryFile)
	if ok {
		mf.open = true
		mf.pos = 0
	}
	return f, nil
}

func (m *memory) Delete(name string) error {
	delete(m.files, name)
	return nil
}

func (m *memory) Mkdir(name string) error {
	f, ok := m.files[name]
	if !ok {
		m.files[name] = &memoryDirectory{}
		return nil
	}
	if f.Stat().IsDir() {
		return nil
	}
	return ErrReservedName
}

func (m *memory) New(name string) (File, error) {
	f := &memoryFile{}
	m.files[name] = f
	return f, nil
}

func (m *memory) ReadDir(name string) ([]FileInfo, error) {
	prefix := name + "/"
	list := []FileInfo{}
	for k := range m.files {
		if len(k) <= len(prefix) {
			continue
		}
		if k[:len(prefix)] != prefix {
			continue
		}
		file := k[len(prefix):]
		if strings.Contains(file, "/") {
			continue
		}
		list = append(list, &fileInfo{name: file})
	}
	return list, nil
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
	name  string
	isDir bool
}

func (fi *fileInfo) IsDir() bool {
	return fi.isDir
}

func (fi *fileInfo) Name() string {
	return fi.name
}

type memoryDirectory struct{}

func (md *memoryDirectory) Read(b []byte) (int, error) {
	return 0, nil
}

func (md *memoryDirectory) Write(b []byte) (int, error) {
	return 0, nil
}

func (md *memoryDirectory) Close() error {
	return nil
}

func (md *memoryDirectory) Stat() FileInfo {
	return &fileInfo{isDir: true}
}

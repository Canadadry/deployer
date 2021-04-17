package fs

import (
	"bytes"
)

type memory struct {
	files map[string]memoryFile
}

func (m *memory) Open(name string) (File, error) {
	f, ok := m.files[name]
	if !ok {
		return nil, ErrFileNotFound
	}
	return &f, nil
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
	content *bytes.Buffer
}

func (mf *memoryFile) Read(b []byte) (int, error) {
	if mf.content == nil {
		return 0, ErrClosedFile
	}
	return mf.content.Read(b)
}
func (mf *memoryFile) Write(b []byte) (int, error) {
	if mf.content == nil {
		return 0, ErrClosedFile
	}
	return mf.content.Write(b)
}
func (mf *memoryFile) Close() error {
	if mf.content == nil {
		return ErrClosedFile
	}
	mf.content = nil
	return nil
}
func (mf *memoryFile) Stat() FileInfo {
	return nil
}

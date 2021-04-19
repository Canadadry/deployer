package fs

import (
	"testing"
)

func TestAllCaseFomMemory(t *testing.T) {
	runAllTestCase(t, func(t *testing.T) FS {
		return &memory{
			files: map[string]File{
				"file_with_content": &memoryFile{
					content: []byte("file content"),
				},
				"empty_file": &memoryFile{
					content: []byte(""),
				},
			},
		}

	})
}

package fs

import (
	"testing"
)

func TestAllCase(t *testing.T) {
	runAllTestCase(t, func() FS {
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

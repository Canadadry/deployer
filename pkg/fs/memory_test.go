package fs

import (
	// "fmt"
	// "reflect"
	// "sort"
	"bytes"
	"testing"
)

func TestOpeningNotExistingFile(t *testing.T) {
	fs := &memory{}
	testOpeningNotExistingFile(t, fs, "fake")
}

func TestOpeningExistingFile(t *testing.T) {
	fs := &memory{
		files: map[string]memoryFile{
			"real_file": memoryFile{
				content: bytes.NewBufferString("file content"),
			},
		},
	}
	testOpeningExistingFile(t, fs, "real_file", "file content")
}

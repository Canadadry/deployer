package fs

import (
	// "fmt"
	// "reflect"
	// "sort"
	"testing"
)

func TestOpeningNotExistingFile(t *testing.T) {
	fs := &memory{}
	testOpeningNotExistingFile(t, fs, "fake")
}

func TestOpeningExistingFile(t *testing.T) {
	fs := &memory{
		files: map[string]memoryFile{
			"real_file": memoryFile{},
		},
	}
	testOpeningExistingFile(t, fs, "real_file")
}

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
		files: map[string]File{
			"real_file": nil,
		},
	}
	testOpeningExistingFile(t, fs, "real_file")
}

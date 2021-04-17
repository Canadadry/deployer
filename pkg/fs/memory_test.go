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

func TestOpeningExistingFileAndWriting(t *testing.T) {
	fs := &memory{
		files: map[string]memoryFile{
			"real_file": memoryFile{
				content: bytes.NewBufferString(""),
			},
		},
	}
	testOpeningExistingFileAndWriting(t, fs, "real_file", "file content")
}

func TestOpeningExistingFile_CannotReadAfterClose(t *testing.T) {
	fs := &memory{
		files: map[string]memoryFile{
			"real_file": memoryFile{
				content: bytes.NewBufferString(""),
			},
		},
	}
	testOpeningExistingFile_CannotReadAfterClose(t, fs, "real_file")
}

func TestOpeningExistingFile_CannotWriteAfterClose(t *testing.T) {
	fs := &memory{
		files: map[string]memoryFile{
			"real_file": memoryFile{
				content: bytes.NewBufferString(""),
			},
		},
	}
	testOpeningExistingFile_CannotWriteAfterClose(t, fs, "real_file")
}

func TestOpeningExistingFile_CannotCloseTwice(t *testing.T) {
	fs := &memory{
		files: map[string]memoryFile{
			"real_file": memoryFile{
				content: bytes.NewBufferString(""),
			},
		},
	}
	testOpeningExistingFile_CannotCloseTwice(t, fs, "real_file")
}

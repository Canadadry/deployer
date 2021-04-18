package fs

import (
	"testing"
)

func TestOpeningNotExistingFile(t *testing.T) {
	fs := &memory{}
	testOpeningNotExistingFile(t, fs, "fake")
}

func TestOpeningExistingFileAndReading(t *testing.T) {
	fs := &memory{
		files: map[string]*memoryFile{
			"real_file": &memoryFile{
				content: []byte("file content"),
			},
		},
	}
	testOpeningExistingFileAndReading(t, fs, "real_file", "file content")
}

func TestOpeningExistingFileAndWriting(t *testing.T) {
	fs := &memory{
		files: map[string]*memoryFile{
			"real_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFileAndWriting(t, fs, "real_file", "file content")
}

func TestOpeningExistingFile_CannotReadAfterClose(t *testing.T) {
	fs := &memory{
		files: map[string]*memoryFile{
			"real_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFile_CannotReadAfterClose(t, fs, "real_file")
}

func TestOpeningExistingFile_CannotWriteAfterClose(t *testing.T) {
	fs := &memory{
		files: map[string]*memoryFile{
			"real_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFile_CannotWriteAfterClose(t, fs, "real_file")
}

func TestOpeningExistingFile_CannotCloseTwice(t *testing.T) {
	fs := &memory{
		files: map[string]*memoryFile{
			"real_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFile_CannotCloseTwice(t, fs, "real_file")
}

func Test_CanOpenTwiceAFile(t *testing.T) {
	fs := &memory{
		files: map[string]*memoryFile{
			"real_file": &memoryFile{
				content: []byte("file content"),
			},
		},
	}
	testOpeningExistingFile_CannotReadAfterClose(t, fs, "real_file")
	testOpeningExistingFileAndReading(t, fs, "real_file", "file content")
}

func Test_CanReadAfterWrite(t *testing.T) {
	fs := &memory{
		files: map[string]*memoryFile{
			"real_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFileAndWriting(t, fs, "real_file", "file content")
	testOpeningExistingFileAndReading(t, fs, "real_file", "file content")
}

func TestOpeningExistingFile_GetSatIsDir(t *testing.T) {
	fs := &memory{
		files: map[string]*memoryFile{
			"real_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFile_GetSatIsDir(t, fs, "real_file", false)
}

func TestDeleteFile(t *testing.T) {
	fs := &memory{
		files: map[string]*memoryFile{
			"real_file": &memoryFile{
				content: []byte("file content"),
			},
		},
	}
	testOpeningExistingFileAndReading(t, fs, "real_file", "file content")
	testDeleteFile(t, fs, "real_file")
	testOpeningNotExistingFile(t, fs, "real_file")
}

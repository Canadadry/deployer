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
		files: map[string]File{
			"real_file": &memoryFile{
				content: []byte("file content"),
			},
		},
	}
	testOpeningExistingFileAndReading(t, fs, "real_file", "file content")
}

func TestOpeningExistingFileAndWriting(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"real_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFileAndWriting(t, fs, "real_file", "file content")
}

func TestOpeningExistingFile_CannotReadAfterClose(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"real_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFile_CannotReadAfterClose(t, fs, "real_file")
}

func TestOpeningExistingFile_CannotWriteAfterClose(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"real_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFile_CannotWriteAfterClose(t, fs, "real_file")
}

func TestOpeningExistingFile_CannotCloseTwice(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"real_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFile_CannotCloseTwice(t, fs, "real_file")
}

func Test_CanOpenTwiceAFile(t *testing.T) {
	fs := &memory{
		files: map[string]File{
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
		files: map[string]File{
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
		files: map[string]File{
			"real_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFile_GetSatIsDir(t, fs, "real_file", false)
}

func TestDeleteFile(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"real_file": &memoryFile{
				content: []byte("file content"),
			},
		},
	}
	testOpeningExistingFileAndReading(t, fs, "real_file", "file content")
	testDeleteFile(t, fs, "real_file")
	testOpeningNotExistingFile(t, fs, "real_file")
}

func TestCreatingFile(t *testing.T) {
	fs := &memory{files: map[string]File{}}

	testOpeningNotExistingFile(t, fs, "real_file")
	testCreateFile(t, fs, "real_file")
	testOpeningExistingFileAndWriting(t, fs, "real_file", "file content")
	testOpeningExistingFileAndReading(t, fs, "real_file", "file content")
}

func TestCreatingFile_WhileFileExist_TruncateIt(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"real_file": &memoryFile{
				content: []byte("file content"),
			},
		},
	}

	testOpeningExistingFileAndReading(t, fs, "real_file", "file content")
	testCreateFile(t, fs, "real_file")
	testOpeningExistingFileAndReading(t, fs, "real_file", "")
}

func TestCreatingDirectory(t *testing.T) {
	fs := &memory{
		files: map[string]File{},
	}

	testCreatingDirectory(t, fs, "real_dir")
}

func TestOpeningDirectory(t *testing.T) {
	fs := &memory{
		files: map[string]File{},
	}

	testCreatingDirectory(t, fs, "real_dir")
	testOpeningDirectory(t, fs, "real_dir")
}

func TestCreatingDirectoryOnExistingFile(t *testing.T) {
	fs := &memory{
		files: map[string]File{},
	}

	testCreateFile(t, fs, "real_file")
	testCreatingDirectoryOnExistingFile(t, fs, "real_file")
}

func TestReadDir(t *testing.T) {
	fs := &memory{
		files: map[string]File{},
	}

	testCreatingDirectory(t, fs, "real_dir")
	testReadDir(t, fs, "real_dir", map[string]bool{})
	testCreateFile(t, fs, "real_dir/real_file")
	testReadDir(t, fs, "real_dir", map[string]bool{"real_file": false})
	testCreatingDirectory(t, fs, "real_dir/sub_dir")
	testReadDir(t, fs, "real_dir", map[string]bool{"real_file": false, "sub_dir": true})
	testCreateFile(t, fs, "real_dir/sub_dir/real_file")
	testReadDir(t, fs, "real_dir", map[string]bool{"real_file": false, "sub_dir": true})
}

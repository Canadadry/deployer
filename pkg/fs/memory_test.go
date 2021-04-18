package fs

import (
	"testing"
)

func TestOpeningNotExistingFile(t *testing.T) {
	fs := &memory{}
	testOpeningNotExistingFile(t, fs, "not_existing_file")
}

func TestOpeningExistingFileAndReading(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"file_with_content": &memoryFile{
				content: []byte("file content"),
			},
		},
	}
	testOpeningExistingFileAndReading(t, fs, "file_with_content", "file content")
}

func TestOpeningExistingFileAndWriting(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"empty_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFileAndWriting(t, fs, "empty_file", "file content")
}

func TestOpeningExistingFile_CannotReadAfterClose(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"empty_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFile_CannotReadAfterClose(t, fs, "empty_file")
}

func TestOpeningExistingFile_CannotWriteAfterClose(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"empty_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFile_CannotWriteAfterClose(t, fs, "empty_file")
}

func TestOpeningExistingFile_CannotCloseTwice(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"empty_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFile_CannotCloseTwice(t, fs, "empty_file")
}

func Test_CanOpenTwiceAFile(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"file_with_content": &memoryFile{
				content: []byte("file content"),
			},
		},
	}
	testOpeningExistingFile_CannotReadAfterClose(t, fs, "file_with_content")
	testOpeningExistingFileAndReading(t, fs, "file_with_content", "file content")
}

func Test_CanReadAfterWrite(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"empty_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFileAndWriting(t, fs, "empty_file", "file content")
	testOpeningExistingFileAndReading(t, fs, "empty_file", "file content")
}

func TestOpeningExistingFile_GetSatIsDir(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"empty_file": &memoryFile{
				content: []byte(""),
			},
		},
	}
	testOpeningExistingFile_GetSatIsDir(t, fs, "empty_file", false)
}

func TestDeleteFile(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"file_with_content": &memoryFile{
				content: []byte("file content"),
			},
		},
	}
	testOpeningExistingFileAndReading(t, fs, "file_with_content", "file content")
	testDeleteFile(t, fs, "file_with_content")
	testOpeningNotExistingFile(t, fs, "file_with_content")
}

func TestCreatingFile(t *testing.T) {
	fs := &memory{files: map[string]File{}}

	testOpeningNotExistingFile(t, fs, "not_existing_file")
	testCreateFile(t, fs, "not_existing_file")
	testOpeningExistingFileAndWriting(t, fs, "not_existing_file", "file content")
	testOpeningExistingFileAndReading(t, fs, "not_existing_file", "file content")
}

func TestCreatingFile_WhileFileExist_TruncateIt(t *testing.T) {
	fs := &memory{
		files: map[string]File{
			"file_with_content": &memoryFile{
				content: []byte("file content"),
			},
		},
	}

	testOpeningExistingFileAndReading(t, fs, "file_with_content", "file content")
	testCreateFile(t, fs, "file_with_content")
	testOpeningExistingFileAndReading(t, fs, "file_with_content", "")
}

func TestCreatingDirectory(t *testing.T) {
	fs := &memory{
		files: map[string]File{},
	}

	testCreatingDirectory(t, fs, "not_exisiting_dir")
}

func TestOpeningDirectory(t *testing.T) {
	fs := &memory{
		files: map[string]File{},
	}

	testCreatingDirectory(t, fs, "not_exisiting_dir")
	testOpeningDirectory(t, fs, "not_exisiting_dir")
}

func TestCreatingDirectoryOnExistingFile(t *testing.T) {
	fs := &memory{
		files: map[string]File{},
	}

	testCreateFile(t, fs, "not_exisiting_dir")
	testCreatingDirectoryOnExistingFile(t, fs, "not_exisiting_dir")
}

func TestReadDir(t *testing.T) {
	fs := &memory{
		files: map[string]File{},
	}

	testCreatingDirectory(t, fs, "not_exisiting_dir")
	testReadDir(t, fs, "not_exisiting_dir", map[string]bool{})
	testCreateFile(t, fs, "not_exisiting_dir/not_exisiting_file")
	testReadDir(t, fs, "not_exisiting_dir", map[string]bool{"not_exisiting_file": false})
	testCreatingDirectory(t, fs, "not_exisiting_dir/not_exisiting_dir")
	testReadDir(t, fs, "not_exisiting_dir", map[string]bool{"not_exisiting_file": false, "not_exisiting_dir": true})
	testCreateFile(t, fs, "not_exisiting_dir/not_exisiting_dir/real_file")
	testReadDir(t, fs, "not_exisiting_dir", map[string]bool{"not_exisiting_file": false, "not_exisiting_dir": true})
}

func TestCannotReadFromDirectory(t *testing.T) {
	fs := &memory{
		files: map[string]File{},
	}

	testCreatingDirectory(t, fs, "not_exisiting_dir")
	testCannotReadFromDirectory(t, fs, "not_exisiting_dir")
}

func TestCannotWriteInDirectory(t *testing.T) {
	fs := &memory{
		files: map[string]File{},
	}

	testCreatingDirectory(t, fs, "not_exisiting_dir")
	testCannotWriteInDirectory(t, fs, "not_exisiting_dir")
}

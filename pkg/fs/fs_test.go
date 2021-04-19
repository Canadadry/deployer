package fs

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func runAllTestCase(t *testing.T, newFs func() FS) {
	tests := map[string]func(t *testing.T, fs FS){
		"TestOpeningNotExistingFile": func(t *testing.T, fs FS) {
			testOpeningNotExistingFile(t, fs, "not_existing_file")
		},
		"TestOpeningExistingFileAndReading": func(t *testing.T, fs FS) {
			testOpeningExistingFileAndReading(t, fs, "file_with_content", "file content")
		},
		"TestOpeningExistingFileAndWriting": func(t *testing.T, fs FS) {
			testOpeningExistingFileAndWriting(t, fs, "empty_file", "file content")
		},
		"TestOpeningExistingFile_CannotReadAfterClose": func(t *testing.T, fs FS) {
			testOpeningExistingFile_CannotReadAfterClose(t, fs, "empty_file")
		},
		"TestOpeningExistingFile_CannotWriteAfterClose": func(t *testing.T, fs FS) {
			testOpeningExistingFile_CannotWriteAfterClose(t, fs, "empty_file")
		},
		"TestOpeningExistingFile_CannotCloseTwice": func(t *testing.T, fs FS) {
			testOpeningExistingFile_CannotCloseTwice(t, fs, "empty_file")
		},
		"Test_CanOpenTwiceAFile": func(t *testing.T, fs FS) {
			testOpeningExistingFile_CannotReadAfterClose(t, fs, "file_with_content")
			testOpeningExistingFileAndReading(t, fs, "file_with_content", "file content")
		},
		"Test_CanReadAfterWrite": func(t *testing.T, fs FS) {
			testOpeningExistingFile_CannotReadAfterClose(t, fs, "file_with_content")
			testOpeningExistingFileAndReading(t, fs, "file_with_content", "file content")
		},
		"TestOpeningExistingFile_GetSatIsDir": func(t *testing.T, fs FS) {
			testOpeningExistingFile_GetSatIsDir(t, fs, "empty_file", false)
		},
		"TestDeleteFile": func(t *testing.T, fs FS) {
			testOpeningExistingFileAndReading(t, fs, "file_with_content", "file content")
			testDeleteFile(t, fs, "file_with_content")
			testOpeningNotExistingFile(t, fs, "file_with_content")
		},
		"TestCreatingFile": func(t *testing.T, fs FS) {
			testOpeningNotExistingFile(t, fs, "not_existing_file")
			testCreateFile(t, fs, "not_existing_file")
			testOpeningExistingFileAndWriting(t, fs, "not_existing_file", "file content")
			testOpeningExistingFileAndReading(t, fs, "not_existing_file", "file content")
		},
		"TestCreatingFile_WhileFileExist_TruncateIt": func(t *testing.T, fs FS) {
			testOpeningExistingFileAndReading(t, fs, "file_with_content", "file content")
			testCreateFile(t, fs, "file_with_content")
			testOpeningExistingFileAndReading(t, fs, "file_with_content", "")
		},
		"TestCreatingDirectory": func(t *testing.T, fs FS) {
			testCreatingDirectory(t, fs, "not_exisiting_dir")
		},
		"TestOpeningDirectory": func(t *testing.T, fs FS) {
			testCreatingDirectory(t, fs, "not_exisiting_dir")
			testOpeningDirectory(t, fs, "not_exisiting_dir")
		},
		"TestCreatingDirectoryOnExistingFile": func(t *testing.T, fs FS) {
			testCreateFile(t, fs, "not_exisiting_dir")
			testCreatingDirectoryOnExistingFile(t, fs, "not_exisiting_dir")
		},
		"TestReadDir": func(t *testing.T, fs FS) {
			testCreatingDirectory(t, fs, "not_exisiting_dir")
			testReadDir(t, fs, "not_exisiting_dir", map[string]bool{})
			testCreateFile(t, fs, "not_exisiting_dir/not_exisiting_file")
			testReadDir(t, fs, "not_exisiting_dir", map[string]bool{"not_exisiting_file": false})
			testCreatingDirectory(t, fs, "not_exisiting_dir/not_exisiting_dir")
			testReadDir(t, fs, "not_exisiting_dir", map[string]bool{"not_exisiting_file": false, "not_exisiting_dir": true})
			testCreateFile(t, fs, "not_exisiting_dir/not_exisiting_dir/real_file")
			testReadDir(t, fs, "not_exisiting_dir", map[string]bool{"not_exisiting_file": false, "not_exisiting_dir": true})
		},
		"TestCannotReadFromDirectory": func(t *testing.T, fs FS) {
			testCreatingDirectory(t, fs, "not_exisiting_dir")
			testCannotReadFromDirectory(t, fs, "not_exisiting_dir")
		},
		"TestCannotWriteInDirectory": func(t *testing.T, fs FS) {
			testCreatingDirectory(t, fs, "not_exisiting_dir")
			testCannotWriteInDirectory(t, fs, "not_exisiting_dir")
		},
	}

	for title, tt := range tests {
		t.Run(title, func(t *testing.T) {
			tt(t, newFs())
		})
	}
}

func testOpeningNotExistingFile(t *testing.T, fs FS, name string) {
	f, err := fs.Open(name)
	if f != nil {
		t.Fatalf("should have returned a nil file got %#v", f)
	}
	if err != ErrFileNotFound {
		t.Fatalf("should have returned %#v got %#v", ErrFileNotFound, err)
	}
}

func testOpeningExistingFileAndReading(t *testing.T, fs FS, name string, content string) {
	f, err := fs.Open(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
	if f == nil {
		t.Fatalf("should not have returned a nil file got %#v", f)
	}

	read, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("cannot read : should not have returned en error got %#v", err)
	}
	if string(read) != content {
		t.Fatalf("not expected result got '%s' want '%s'", string(read), content)
	}
}

func testOpeningExistingFileAndWriting(t *testing.T, fs FS, name string, content string) {
	f, err := fs.Open(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
	if f == nil {
		t.Fatalf("should not have returned a nil file got %#v", f)
	}

	_, err = fmt.Fprintf(f, "%s", content)

	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}

	testOpeningExistingFileAndReading(t, fs, name, content)
}

func testOpeningExistingFile_CannotReadAfterClose(t *testing.T, fs FS, name string) {
	f, err := fs.Open(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
	if f == nil {
		t.Fatalf("should not have returned a nil file got %#v", f)
	}
	err = f.Close()

	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}

	_, err = ioutil.ReadAll(f)
	if err != ErrClosedFile {
		t.Fatalf("can read : should have returned en error got %#v want %#v", err, ErrClosedFile)
	}
}

func testOpeningExistingFile_CannotWriteAfterClose(t *testing.T, fs FS, name string) {
	f, err := fs.Open(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
	if f == nil {
		t.Fatalf("should not have returned a nil file got %#v", f)
	}
	err = f.Close()

	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}

	_, err = fmt.Fprintf(f, "%s", "content")

	if err != ErrClosedFile {
		t.Fatalf("can read : should have returned en error got %#v want %#v", err, ErrClosedFile)
	}
}

func testOpeningExistingFile_CannotCloseTwice(t *testing.T, fs FS, name string) {
	f, err := fs.Open(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
	if f == nil {
		t.Fatalf("should not have returned a nil file got %#v", f)
	}
	err = f.Close()

	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}

	err = f.Close()

	if err != ErrClosedFile {
		t.Fatalf("can read : should have returned en error got %#v want %#v", err, ErrClosedFile)
	}
}

func testOpeningExistingFile_GetSatIsDir(t *testing.T, fs FS, name string, isDir bool) {
	f, err := fs.Open(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
	if f == nil {
		t.Fatalf("should not have returned a nil file got %#v", f)
	}

	if f.Stat().IsDir() != isDir {
		t.Fatalf("file stats IsDir should return %v got %v ", isDir, f.Stat().IsDir())
	}
}

func testDeleteFile(t *testing.T, fs FS, name string) {
	err := fs.Delete(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
}

func testCreateFile(t *testing.T, fs FS, name string) {
	_, err := fs.New(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
}

func testCreatingDirectory(t *testing.T, fs FS, name string) {
	err := fs.Mkdir(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
}

func testOpeningDirectory(t *testing.T, fs FS, name string) {
	testOpeningExistingFile_GetSatIsDir(t, fs, name, true)
}

func testCreatingDirectoryOnExistingFile(t *testing.T, fs FS, name string) {
	err := fs.Mkdir(name)
	if err != ErrReservedName {
		t.Fatalf("should have returned en error got %#v want %#v", err, ErrReservedName)
	}
}

func testReadDir(t *testing.T, fs FS, name string, expectedFiles map[string]bool) {
	infos, err := fs.ReadDir(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
	if len(infos) != len(expectedFiles) {
		t.Fatalf("should have returned %d files  got %v", len(expectedFiles), len(infos))
	}

	for _, i := range infos {
		name := i.Name()
		isDir, ok := expectedFiles[name]
		if !ok {
			t.Fatalf("found file %s that is not in expected files", name)
		}
		if i.IsDir() != isDir {
			t.Fatalf("file stats IsDir should return %v got %v ", isDir, i.IsDir())
		}
	}
}

func testCannotReadFromDirectory(t *testing.T, fs FS, name string) {
	f, err := fs.Open(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
	if f == nil {
		t.Fatalf("should not have returned a nil file got %#v", f)
	}

	_, err = ioutil.ReadAll(f)
	if err != ErrCannotReadOrWriteFromDirectiory {
		t.Fatalf("should have returned en error got %#v want %#v", err, ErrCannotReadOrWriteFromDirectiory)
	}
}

func testCannotWriteInDirectory(t *testing.T, fs FS, name string) {
	f, err := fs.Open(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
	if f == nil {
		t.Fatalf("should not have returned a nil file got %#v", f)
	}

	_, err = fmt.Fprintf(f, "%s", "something")
	if err != ErrCannotReadOrWriteFromDirectiory {
		t.Fatalf("should have returned en error got %#v want %#v", err, ErrCannotReadOrWriteFromDirectiory)
	}
}

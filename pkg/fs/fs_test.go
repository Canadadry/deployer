package fs

import (
	"fmt"
	"io/ioutil"
	"testing"
)

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

package fs

import (
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

func testOpeningExistingFile(t *testing.T, fs FS, name string) {
	_, err := fs.Open(name)
	if err != nil {
		t.Fatalf("should not have returned en error got %#v", err)
	}
}

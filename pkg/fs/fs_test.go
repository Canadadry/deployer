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

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

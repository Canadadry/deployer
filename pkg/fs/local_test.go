package fs

import (
	"os"
	"testing"
)

func TestAllCaseFomLocal(t *testing.T) {
	runAllTestCase(t, func(t *testing.T) FS {
		if _, err := os.Stat("test_local/"); os.IsNotExist(err) {
			err := os.RemoveAll("test_local/")
			if err != nil {
				t.Fatalf("Cannot delete previous test_local directory")
			}
		}
		err := os.Mkdir("test_local/", 0755)
		if err != nil {
			t.Fatalf("Cannot create test_local directory")
		}
		return &local{base: "test_local/"}
	})
}

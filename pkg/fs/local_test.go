package fs

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestAllCaseFomLocal(t *testing.T) {
	runAllTestCase(t, func(t *testing.T) FS {
		_, err := os.Stat("test_local/")
		if err == nil {
			err := os.RemoveAll("test_local/")
			if err != nil {
				t.Fatalf("Cannot delete previous test_local directory")
			}
		}
		err = os.Mkdir("test_local/", 0755)
		if err != nil {
			t.Fatalf("Cannot create test_local directory : %v", err)
		}
		err = ioutil.WriteFile("test_local/file_with_content", []byte("file content"), 0644)
		if err != nil {
			t.Fatalf("Cannot create test_local/file_with_content file : %v", err)
		}
		err = ioutil.WriteFile("test_local/empty_file", []byte{}, 0644)
		if err != nil {
			t.Fatalf("Cannot create test_local/empty_file file : %v", err)
		}
		return &local{base: "test_local/"}
	})
}

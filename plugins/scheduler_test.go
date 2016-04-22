package plugins

import (
	"fmt"
	"os"
	"testing"
)

var filePath string

func TestSetFileMod(t *testing.T) {
	var in_mode os.FileMode = 0744
	setFileMod(filePath, in_mode)

	info, _ := os.Stat(filePath)
	expected_mode := info.Mode()
	if expected_mode != in_mode {
		t.Error("Set file mode fail!")
	}
}

func TestMain(m *testing.M) {
	filePath = "____________testfile_________________"
	_, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Create test file failed:", filePath, err)
	}

	m.Run()

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Remove test file failed:", filePath, err)
	}
}

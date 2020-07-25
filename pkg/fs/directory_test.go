package fs

import (
	"os"
	"testing"
)

// Test that the MakeDir function makes a new directory if one does not exist in current directory..
func TestMakeDir(t *testing.T) {
	testPath := "/tmp/testdir"

	// Test Makedir creates a new directory.
	if ok := MakeDir(testPath); !ok {
		if _, err := os.Stat(testPath); os.IsNotExist(err) {
			t.Errorf("\nExpected:\n\t%v\ngot:\n\tDirectory does not exist", testPath)
		}
		os.Remove(testPath)
	}

	// Test MakeDir returns true, meaning no directory found and will create a new one.
	expectedTrue := MakeDir(testPath)
	if expectedTrue != true {
		t.Errorf("\nExpected:\n\ttrue\ngot:\n\t%v", expectedTrue)
	}

	// Test MakeDir returns false, meaning a directory already exists.
	expectedFalse := MakeDir(testPath)
	if expectedFalse != false {
		t.Errorf("\nExpected:\n\ttrue\ngot:\n\t%v", expectedFalse)
		os.Remove(testPath)
	}
}

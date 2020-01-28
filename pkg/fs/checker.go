package fs

import "os"

// MakeDir will make a new directory if not found.
func MakeDir(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
		return true
	}

	return false
}

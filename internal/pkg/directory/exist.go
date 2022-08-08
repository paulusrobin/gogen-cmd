package directory

import "os"

// Exist function to check a filePath is exists.
func Exist(directoryPath string) bool {
	if _, err := os.Stat(directoryPath); err != nil {
		return false
	}
	return true
}

package file

import (
	"os"
)

// Exist function to check a filePath is exists.
func Exist(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return false
	}
	return true
}

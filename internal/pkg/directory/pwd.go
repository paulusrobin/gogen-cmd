package directory

import "os"

// Pwd function to get current active directory.
func Pwd() string {
	path, _ := os.Getwd()
	return path
}

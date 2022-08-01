package directory

import "os"

// Make function to create a directory.
func Make(path string) error {
	return os.Mkdir(path, os.ModePerm)
}

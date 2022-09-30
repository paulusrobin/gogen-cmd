package directory

import "os"

// Remove function to remove a directory.
func Remove(path string) error {
	return os.RemoveAll(path)
}

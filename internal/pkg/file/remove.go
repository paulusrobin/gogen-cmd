package file

import "os"

// Remove function to remove file.
func Remove(filePath string) error {
	return os.Remove(filePath)
}

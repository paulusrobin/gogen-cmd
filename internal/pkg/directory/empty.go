package directory

import "io/ioutil"

// Empty function to check a directory is empty or not.
func Empty(path string) bool {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false
	}
	return len(files) <= 0
}

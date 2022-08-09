package directory

import (
	"io/fs"
	"io/ioutil"
	"regexp"
)

var defaultFilter = func(info fs.FileInfo) bool {
	return !info.IsDir()
}

// FileNames function to get list file name on directory.
func FileNames(directoryPath string) ([]string, error) {
	return FileNamesWithFilter(directoryPath, "*", defaultFilter)
}

// FileNamesWithFilter function to get list file name with regex filter on directory.
func FileNamesWithFilter(directoryPath, filter string, fn func(info fs.FileInfo) bool) ([]string, error) {
	var fileNames = make([]string, 0)

	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return fileNames, err
	}

	for _, file := range files {
		if defaultFilter(file) {
			if matched, err := regexp.Match(filter, []byte(file.Name())); err != nil || !matched {
				continue
			}
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, nil
}

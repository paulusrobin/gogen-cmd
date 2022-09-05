package directory

import (
	"io/fs"
	"path/filepath"
	"regexp"
)

var (
	fileFilter = func(path string, info fs.FileInfo) bool {
		return !info.IsDir()
	}
	dirFilter = func(path string, info fs.FileInfo) bool {
		return info.IsDir()
	}
)

// FileNames function to get list file name on directory.
func FileNames(directoryPath string) ([]string, error) {
	return FileNamesWithFilter(directoryPath, "^*.*", fileFilter)
}

// DirNames function to get list file directory.
func DirNames(directoryPath string) ([]string, error) {
	return FileNamesWithFilter(directoryPath, "^*.*", dirFilter)
}

// FileNamesWithFilter function to get list file name with regex filter on directory.
func FileNamesWithFilter(directoryPath, filter string, fn func(path string, info fs.FileInfo) bool) ([]string, error) {
	var fileNames = make([]string, 0)
	err := filepath.Walk(directoryPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fn(path, info) {
			return nil
		}
		if matched, err := regexp.Match(filter, []byte(info.Name())); err != nil || !matched {
			return err
		}
		fileNames = append(fileNames, info.Name())
		return nil
	})
	return fileNames, err
}

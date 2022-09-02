package generator

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"path"
)

// Folder function to generate folders.
func Folder(basePath string, generatedFolders []string) error {
	for _, folderPath := range generatedFolders {
		generatedPath := path.Join(basePath, folderPath)
		if directory.Exist(generatedPath) {
			continue
		}
		if err := directory.Make(generatedPath); err != nil {
			return err
		}
	}
	return nil
}

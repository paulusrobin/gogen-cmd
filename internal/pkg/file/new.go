package file

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/errors"
	"os"
)

// New function to create a file.
func New(outputPath string) (*os.File, error) {
	if exist := Exist(outputPath); exist {
		return nil, errors.FileExist{FileName: outputPath}
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

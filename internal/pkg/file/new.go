package file

import (
	"fmt"
	"os"
)

// New function to create a file.
func New(outputPath string) (*os.File, error) {
	if exist := Exist(outputPath); exist {
		return nil, fmt.Errorf("%s is already exist", outputPath)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

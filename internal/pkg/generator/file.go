package generator

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"path"
)

// File function to generate files.
func File(basePath string, generatedFiles map[string]string, parameters map[string]interface{}) error {
	for outputFile, content := range generatedFiles {
		if err := file.Generate(path.Join(basePath, outputFile), content, parameters); err != nil {
			return err
		}
	}
	return nil
}

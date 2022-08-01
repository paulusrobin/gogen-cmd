package helper

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/template"
)

// Generate function to generate file.
func Generate(outputPath, templatePath string, parameters map[string]interface{}) error {
	f, err := file.New(outputPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	return template.Exec(templatePath, parameters, f)
}

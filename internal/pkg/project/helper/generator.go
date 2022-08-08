package helper

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"text/template"
)

// Generate function to generate file.
func Generate(outputPath, content string, parameters map[string]interface{}) error {
	f, err := file.New(outputPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	tmpl, err := template.New("template-parser").Parse(content)
	if err != nil {
		return err
	}

	if err = tmpl.Execute(f, parameters); err != nil {
		return err
	}
	return nil
}

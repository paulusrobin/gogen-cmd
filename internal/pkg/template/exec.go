package template

import (
	"os"
	"path"
	"text/template"
)

var templateBasePath string

func init() {
	templateBasePath, _ = os.Getwd()
	templateBasePath = path.Join(templateBasePath, "internal/pkg/template")
}

// Exec function to execute template to an output io.Writer.
func Exec(templatePath string, parameters map[string]interface{}, output *os.File) error {
	var (
		filePath = path.Join(templateBasePath, templatePath)
		tmpl     *template.Template
		err      error
	)

	tmpl, err = template.ParseFiles(filePath)
	if err != nil {
		return err
	}

	if err = tmpl.Execute(output, parameters); err != nil {
		return err
	}
	return nil
}

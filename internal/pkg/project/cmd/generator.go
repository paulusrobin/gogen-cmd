package cmd

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/helper"
	"path"
)

// Generate function to generate cmd folder files.
func Generate(cfg helper.ProjectConfig) error {
	parameters := map[string]interface{}{
		"ProjectName":   cfg.Name,
		"ProjectModule": cfg.Module,
	}
	generatedFolders := []string{"cmd", "cmd/grpc", "cmd/http", "cmd/subscriber"}
	generatedFiles := map[string]string{
		"cmd/grpc/grpc.go":             "project/cmd/grpc/grpc.go.tmpl",
		"cmd/http/http.go":             "project/cmd/http/http.go.tmpl",
		"cmd/subscriber/subscriber.go": "project/cmd/subscriber/subscriber.go.tmpl",
		"cmd/main.go":                  "project/cmd/main.go.tmpl",
	}

	for _, folderPath := range generatedFolders {
		if err := directory.Make(path.Join(cfg.Path, cfg.Name, folderPath)); err != nil {
			return err
		}
	}

	for outputFile, templatePath := range generatedFiles {
		if err := helper.Generate(path.Join(cfg.Path, cfg.Name, outputFile), templatePath, parameters); err != nil {
			return err
		}
	}
	return nil
}

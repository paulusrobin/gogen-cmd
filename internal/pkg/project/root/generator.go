package root

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/helper"
	"path"
)

// Generate function to generate root folder files.
func Generate(cfg helper.ProjectConfig) error {
	parameters := map[string]interface{}{
		"ProjectName":   cfg.Name,
		"ProjectModule": cfg.Module,
	}
	generatedFiles := map[string]string{
		".env.sample":         "project/root/.env.sample",
		".gitignore":          "project/root/.gitignore",
		"docker-compose.yaml": "project/root/docker-compose.yaml",
		"Dockerfile":          "project/root/Dockerfile",
		"Dockerfile-builder":  "project/root/Dockerfile-builder",
		"go.mod":              "project/root/go.mod.tmpl",
		"Makefile":            "project/root/Makefile",
		"README.md":           "project/root/README.md",
	}

	for outputFile, templatePath := range generatedFiles {
		if err := helper.Generate(path.Join(cfg.Path, cfg.Name, outputFile), templatePath, parameters); err != nil {
			return err
		}
	}
	return nil
}

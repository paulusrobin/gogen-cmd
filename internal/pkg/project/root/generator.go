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
		".env.sample":         string(envSampleTemplate),
		".gitignore":          string(gitIgnoreTemplate),
		"docker-compose.yaml": string(dockerComposeTemplate),
		"Dockerfile":          string(dockerfileTemplate),
		"Dockerfile-builder":  string(dockerfileBuilderTemplate),
		"go.mod":              string(goModTemplate),
		"Makefile":            string(makefileTemplate),
		"README.md":           string(readmeTemplate),
	}

	for outputFile, content := range generatedFiles {
		if err := helper.Generate(path.Join(cfg.Path, cfg.Name, outputFile), content, parameters); err != nil {
			return err
		}
	}
	return nil
}

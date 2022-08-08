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
	generatedFolders := []string{
		"cmd",
		"cmd/grpc",
		"cmd/http",
		"cmd/subscriber",
	}
	generatedFiles := map[string]string{
		"cmd/grpc/grpc.go":             string(grpcTemplate),
		"cmd/http/http.go":             string(httpTemplate),
		"cmd/subscriber/subscriber.go": string(subscriberTemplate),
		"cmd/main.go":                  string(mainTemplate),
	}

	for _, folderPath := range generatedFolders {
		generatedPath := path.Join(cfg.Path, cfg.Name, folderPath)
		if directory.Exist(generatedPath) {
			continue
		}
		if err := directory.Make(generatedPath); err != nil {
			return err
		}
	}

	for outputFile, content := range generatedFiles {
		if err := helper.Generate(path.Join(cfg.Path, cfg.Name, outputFile), content, parameters); err != nil {
			return err
		}
	}
	return nil
}

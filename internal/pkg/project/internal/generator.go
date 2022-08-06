package internal

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/helper"
	"path"
)

// Generate function to generate internal folder files.
func Generate(cfg helper.ProjectConfig) error {
	parameters := map[string]interface{}{
		"ProjectName":            cfg.Name,
		"ProjectModule":          cfg.Module,
		"PackageName":            "greeting",
		"RuntimeBusinessConfigs": []map[string]interface{}{},
	}
	generatedFolders := []string{
		"internal",
		"internal/config",
		"internal/pkg",
		"internal/pkg/greeting",
		"internal/pkg/greeting/endpoint",
		"internal/pkg/greeting/payload",
		"internal/pkg/greeting/usecase",
		"internal/repository",
		"internal/repository/model",
		"internal/server",
		"internal/server/transport",
		"internal/server/transport/grpc",
		"internal/server/transport/http",
		"internal/server/transport/subscriber",
	}
	generatedFiles := map[string]string{
		"internal/config/config.go":                  "project/internal/config/config.go.tmpl",
		"internal/config/provider.go":                "project/internal/config/provider.go.tmpl",
		"internal/pkg/greeting/endpoint/greeting.go": "project/internal/pkg/endpoint/greeting.go.tmpl",
		"internal/pkg/greeting/payload/greeting.go":  "project/internal/pkg/payload/greeting.go.tmpl",
		"internal/pkg/greeting/usecase/greeting.go":  "project/internal/pkg/usecase/greeting.go.tmpl",
		"internal/pkg/greeting/usecase/root.go":      "project/internal/pkg/usecase/root.go.tmpl",
		"internal/pkg/greeting/usecase.go":           "project/internal/pkg/greeting.go.tmpl",
		"internal/server/transport/http/http.go":     "project/internal/server/transport/http/http.go.tmpl",
		"internal/server/server.go":                  "project/internal/server/server.go.tmpl",
		"internal/server/grpc.go":                    "project/internal/server/grpc.go.tmpl",
		"internal/server/http.go":                    "project/internal/server/http.go.tmpl",
		"internal/server/subscriber.go":              "project/internal/server/subscriber.go.tmpl",
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

package cmd

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Init function to generate cmd folder files.
func Init(cfg parameter.ProjectConfig) error {
	parameters := map[string]interface{}{
		"ProjectName":   cfg.Name,
		"ProjectModule": cfg.Module,
	}
	generatedFolders := []string{
		"cmd",
		"cmd/http",
	}
	generatedFiles := map[string]string{
		"cmd/http/http.go": string(httpTemplate),
	}
	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(cfg.Path, generatedFolders)),
		functions.MakeFunc(generator.File(cfg.Path, generatedFiles, parameters)),
		functions.MakeFunc(Generate(cfg)),
	})
}

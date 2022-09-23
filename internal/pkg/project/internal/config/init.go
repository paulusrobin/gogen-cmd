package config

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Init function to initialize config folder.
func Init(cfg parameter.ProjectConfig) error {
	parameters := map[string]interface{}{
		"ProjectName":   cfg.Name,
		"ProjectModule": cfg.Module,
		"PackageName":   "greeting",
	}
	generatedFolders := []string{
		"internal",
		"internal/config",
	}
	generatedFiles := map[string]string{
		"internal/config/config.go":   string(Template),
		"internal/config/provider.go": string(ProviderTemplate),
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(cfg.Path, generatedFolders)),
		functions.MakeFunc(generator.File(cfg.Path, generatedFiles, parameters)),
	})
}

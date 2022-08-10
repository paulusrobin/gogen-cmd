package config

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"path"
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
		if err := file.Generate(path.Join(cfg.Path, cfg.Name, outputFile), content, parameters); err != nil {
			return err
		}
	}
	return nil
}

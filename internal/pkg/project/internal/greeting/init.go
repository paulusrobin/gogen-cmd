package greeting

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
		"internal/pkg",
		"internal/pkg/greeting",
		"internal/pkg/greeting/endpoint",
		"internal/pkg/greeting/dto",
		"internal/pkg/greeting/encoding",
		"internal/pkg/greeting/usecase",
	}
	generatedFiles := map[string]string{
		"internal/pkg/greeting/endpoint/greeting.go": string(greetingEndpointTemplate),
		"internal/pkg/greeting/dto/greeting.go":      string(greetingDataTransferObjectTemplate),
		"internal/pkg/greeting/encoding/greeting.go": string(greetingDataTransferObjectEncodingTemplate),
		"internal/pkg/greeting/usecase/greeting.go":  string(greetingUsecaseFunctionTemplate),
		"internal/pkg/greeting/usecase/root.go":      string(greetingUsecaseObjectTemplate),
		"internal/pkg/greeting/usecase.go":           string(greetingUsecaseTemplate),
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

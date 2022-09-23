package greeting

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

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(cfg.Path, generatedFolders)),
		functions.MakeFunc(generator.File(cfg.Path, generatedFiles, parameters)),
	})
}

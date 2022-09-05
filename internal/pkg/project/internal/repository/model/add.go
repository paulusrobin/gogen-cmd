package model

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Add function to remove repository model.
func Add(request parameter.AddModel) error {
	if request.ModelName == "" {
		if request.SkipIfError {
			return nil
		}
		return fmt.Errorf("model name is required")
	}

	parameters := map[string]interface{}{
		"ModelName": convention.FunctionName(request.ModelName),
	}
	generatedFolders := []string{
		"internal",
		"internal/repository",
		"internal/repository/model",
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/repository/model/%s.go", convention.FileName(request.ModelName)): string(modelTemplate),
	}

	return functions.Walk([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, parameters)),
	})
}

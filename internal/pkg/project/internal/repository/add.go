package repository

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Add function to add repository.
func Add(request parameter.AddRepository) error {
	parameters := map[string]interface{}{
		"ModelName": request.ModelName,
	}
	packageFileName := convention.FileName(request.PackageName)
	modelFileName := convention.FileName(request.ModelName)

	generatedFolders := []string{
		"internal",
		"internal/repository",
		fmt.Sprintf("internal/repository/%s", packageFileName),
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/repository/%s/%s.go", packageFileName, modelFileName): string(repositoryTemplate),
	}

	return functions.Walk([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, parameters)),
	})
}

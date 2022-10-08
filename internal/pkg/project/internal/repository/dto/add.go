package dto

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Add function.
func Add(request parameter.AddRepositoryDataTransferObject) error {
	parameters := map[string]interface{}{
		"Function": convention.FunctionName(request.Name),
	}

	generatedFolders := []string{
		"internal",
		"internal/repository",
		fmt.Sprintf("internal/repository/%s", convention.PackageName(request.RepositoryName)),
		fmt.Sprintf("internal/repository/%s/dto", convention.PackageName(request.RepositoryName)),
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/repository/%s/dto/%s.go",
			convention.PackageName(request.RepositoryName), convention.FileName(request.Name)): string(dtoTemplate),
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, parameters)),
	})
}

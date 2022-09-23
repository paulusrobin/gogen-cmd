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
	name := fmt.Sprintf("%s%s",
		convention.ToUpperFirstLetter(request.RepositoryName),
		convention.ToUpperFirstLetter(request.Name),
	)

	parameters := map[string]interface{}{
		"Name": convention.FunctionName(name),
		"Type": convention.ToUpperFirstLetter(request.Type),
	}

	generatedFolders := []string{
		"internal",
		"internal/repository",
		"internal/repository/model",
		"internal/repository/model/dto",
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/repository/model/dto/%s.go", name): string(dtoTemplate),
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, parameters)),
	})
}

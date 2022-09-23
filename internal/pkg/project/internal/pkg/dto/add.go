package dto

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Add function.
func Add(request parameter.AddDataTransferObject) error {
	parameters := map[string]interface{}{
		"Name": convention.FunctionName(request.Name),
		"Type": request.Type,
	}

	packageFileName := convention.FileName(request.PackageName)
	payloadFileName := convention.FileName(request.Name)

	generatedFolders := []string{
		"internal",
		"internal/pkg",
		fmt.Sprintf("internal/pkg/%s", packageFileName),
		fmt.Sprintf("internal/pkg/%s/dto", packageFileName),
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/pkg/%s/dto/%s%s.go", packageFileName, payloadFileName, convention.ToUpperFirstLetter(request.Type)): string(dtoTemplate),
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, parameters)),
	})
}

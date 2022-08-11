package dto

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"path"
)

// Add function.
func Add(request parameter.AddDataTransferObject) error {
	parameters := map[string]interface{}{
		"Name": request.Name,
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

	for _, folderPath := range generatedFolders {
		generatedPath := path.Join(request.Path, folderPath)
		if directory.Exist(generatedPath) {
			continue
		}
		if err := directory.Make(generatedPath); err != nil {
			return err
		}
	}

	generatedFile := fmt.Sprintf("internal/pkg/%s/dto/%s%s.go",
		packageFileName, payloadFileName, convention.ToUpperFirstLetter(request.Type))
	if err := file.Generate(path.Join(request.Path, generatedFile), string(dtoTemplate), parameters); err != nil {
		return err
	}
	return nil
}

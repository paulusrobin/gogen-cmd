package dto

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"path"
)

// AddUsecase function.
func AddUsecase(request parameter.AddUsecase) error {
	parameters := map[string]interface{}{
		"ProjectName":   request.Name,
		"ProjectModule": request.Module,
		"PackageName":   convention.PackageName(request.PackageName),
		"FunctionName":  convention.FunctionName(request.FunctionName),
	}

	packageFileName := convention.FileName(request.PackageName)
	usecaseFileName := convention.FileName(request.FunctionName)

	generatedFolders := []string{
		"internal",
		"internal/pkg",
		fmt.Sprintf("internal/pkg/%s", packageFileName),
		fmt.Sprintf("internal/pkg/%s/dto", packageFileName),
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/pkg/%s/dto/%sUsecase.go", packageFileName, usecaseFileName): string(usecaseTemplate),
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

	for outputFile, content := range generatedFiles {
		if err := file.Generate(path.Join(request.Path, outputFile), content, parameters); err != nil {
			return err
		}
	}
	return nil
}

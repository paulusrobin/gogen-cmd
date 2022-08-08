package internal

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/dto"
	"path"
)

// AddUsecase function.
func AddUsecase(parameter dto.AddUsecaseParameter) error {
	parameters := map[string]interface{}{
		"ProjectName":   parameter.Name,
		"ProjectModule": parameter.Module,
		"PackageName":   convention.PackageName(parameter.PackageName),
		"FunctionName":  convention.EndpointName(parameter.FunctionName),
	}

	packageFileName := convention.FileName(parameter.PackageName)
	usecaseFileName := convention.FileName(parameter.FunctionName)

	generatedFolders := []string{
		"internal",
		"internal/pkg",
		fmt.Sprintf("internal/pkg/%s", packageFileName),
		fmt.Sprintf("internal/pkg/%s/payload", packageFileName),
		fmt.Sprintf("internal/pkg/%s/usecase", packageFileName),
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/pkg/%s/payload/%s.go", packageFileName, usecaseFileName):  "project/internal/pkg/payload/payload-usecase.go.tmpl",
		fmt.Sprintf("internal/pkg/%s/endpoint/%s.go", packageFileName, usecaseFileName): "project/internal/pkg/usecase/function.go.tmpl",
	}

	for _, folderPath := range generatedFolders {
		if directory.Exist(folderPath) {
			continue
		}
		if err := directory.Make(path.Join(parameter.Path, folderPath)); err != nil {
			return err
		}
	}

	for outputFile, templatePath := range generatedFiles {
		if err := file.Generate(path.Join(parameter.Path, outputFile), templatePath, parameters); err != nil {
			return err
		}
	}
	return nil
}

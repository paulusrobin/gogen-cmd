package internal

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/helper"
	"path"
)

// AddUsecase function.
func AddUsecase(parameter helper.AddUsecaseParameter) error {
	parameters := map[string]interface{}{
		"ProjectName":   parameter.Name,
		"ProjectModule": parameter.Module,
		"PackageName":   helper.PackageName(parameter.PackageName),
		"FunctionName":  helper.EndpointName(parameter.FunctionName),
	}

	packageFileName := helper.FileName(parameter.PackageName)
	usecaseFileName := helper.FileName(parameter.FunctionName)

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
		if err := helper.Generate(path.Join(parameter.Path, outputFile), templatePath, parameters); err != nil {
			return err
		}
	}
	return nil
}

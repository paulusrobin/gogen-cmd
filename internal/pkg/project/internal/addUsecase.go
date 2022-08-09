package internal

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/dto"
	"io/fs"
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
		fmt.Sprintf("internal/pkg/%s/payload/%sUsecase.go", packageFileName, usecaseFileName): string(pkgPayloadUsecaseTemplate),
		fmt.Sprintf("internal/pkg/%s/usecase/%s.go", packageFileName, usecaseFileName):        string(pkgUsecaseFunctionTemplate),
	}

	for _, folderPath := range generatedFolders {
		generatedPath := path.Join(parameter.Path, folderPath)
		if directory.Exist(generatedPath) {
			continue
		}
		if err := directory.Make(generatedPath); err != nil {
			return err
		}
	}

	for outputFile, content := range generatedFiles {
		if err := file.Generate(path.Join(parameter.Path, outputFile), content, parameters); err != nil {
			return err
		}
	}

	_ = GenerateUsecase(dto.GenerateUsecase{
		ProjectConfig: parameter.ProjectConfig,
		PackageName:   packageFileName,
	})

	fileNames, _ := directory.FileNamesWithFilter(
		path.Join(parameter.Path, fmt.Sprintf("internal/pkg/%s/usecase", packageFileName)),
		"*.go",
		func(info fs.FileInfo) bool {
			return info.IsDir()
		})
	_ = GeneratePackage(dto.GeneratePackage{
		ProjectConfig:    parameter.ProjectConfig,
		PackageName:      packageFileName,
		UsecaseFunctions: fileNames,
	})

	return nil
}

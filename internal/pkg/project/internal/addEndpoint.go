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

// AddEndpoint function.
func AddEndpoint(parameter dto.AddEndpointParameter) error {
	parameters := map[string]interface{}{
		"ProjectName":   parameter.Name,
		"ProjectModule": parameter.Module,
		"PackageName":   convention.PackageName(parameter.PackageName),
		"EndpointName":  convention.EndpointName(parameter.EndpointName),
	}

	packageFileName := convention.FileName(parameter.PackageName)
	endpointFileName := convention.FileName(parameter.EndpointName)

	generatedFolders := []string{
		"internal",
		"internal/pkg",
		fmt.Sprintf("internal/pkg/%s", packageFileName),
		fmt.Sprintf("internal/pkg/%s/endpoint", packageFileName),
		fmt.Sprintf("internal/pkg/%s/payload", packageFileName),
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/pkg/%s/payload/%sEndpoint.go", packageFileName, endpointFileName): string(pkgPayloadEndpointTemplate),
		fmt.Sprintf("internal/pkg/%s/endpoint/%s.go", packageFileName, endpointFileName):        string(pkgEndpointTemplate),
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

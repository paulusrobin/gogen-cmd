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

	generatedFolders := []dto.ProjectPath{
		{"internal", false},
		{"internal/pkg", false},
		{fmt.Sprintf("internal/pkg/%s", packageFileName), true},
		{fmt.Sprintf("internal/pkg/%s/payload", packageFileName), false},
		{fmt.Sprintf("internal/pkg/%s/usecase", packageFileName), false},
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/pkg/%s/payload/%sUsecase.go", packageFileName, usecaseFileName): string(pkgPayloadUsecaseTemplate),
		fmt.Sprintf("internal/pkg/%s/usecase/%s.go", packageFileName, usecaseFileName):        string(pkgUsecaseFunctionTemplate),
	}

	for _, folder := range generatedFolders {
		generatedPath := path.Join(parameter.Path, folder.Path)
		if directory.Exist(generatedPath) {
			continue
		}
		if err := directory.Make(generatedPath); err != nil {
			return err
		}
		if folder.IsPackage {
			if err := GeneratePackage(dto.GeneratePackage{
				ProjectConfig: parameter.ProjectConfig,
				PackageName:   packageFileName,
			}); err != nil {
				return err
			}
		}
	}

	for outputFile, content := range generatedFiles {
		if err := file.Generate(path.Join(parameter.Path, outputFile), content, parameters); err != nil {
			return err
		}
	}
	return nil
}

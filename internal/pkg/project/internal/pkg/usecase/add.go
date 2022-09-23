package usecase

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Add function.
func Add(request parameter.AddUsecase) error {
	parameters := map[string]interface{}{
		"ProjectName":   request.Name,
		"ProjectModule": request.Module,
		"PackageName":   convention.PackageName(request.PackageName),
		"FunctionName":  convention.EndpointName(request.FunctionName),
	}

	packageFileName := convention.FileName(request.PackageName)
	usecaseFileName := convention.FileName(request.FunctionName)

	generatedFolders := []string{
		"internal",
		"internal/pkg",
		fmt.Sprintf("internal/pkg/%s", packageFileName),
		fmt.Sprintf("internal/pkg/%s/usecase", packageFileName),
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/pkg/%s/usecase/%s.go", packageFileName, usecaseFileName): string(usecaseFunctionTemplate),
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, parameters)),
		functions.MakeFunc(Generate(parameter.ProjectConfigWithPackage{
			ProjectConfig: request.ProjectConfig,
			PackageName:   request.PackageName,
		})),
	})
}

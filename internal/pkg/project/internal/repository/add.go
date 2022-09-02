package repository

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository/dto"
	"path"
)

// Add function to add repository.
func Add(request parameter.AddRepository) error {
	parameters := map[string]interface{}{
		"PackageName":    convention.PackageName(request.PackageName),
		"ProjectModule":  request.ProjectConfig.Module,
		"RepositoryName": convention.ToUpperFirstLetter(request.RepositoryName),
	}
	packageFileName := convention.FileName(request.PackageName)
	modelFileName := convention.FileName(request.ModelName)

	generatedFolders := []string{
		"internal",
		"internal/repository",
		fmt.Sprintf("internal/repository/%s", packageFileName),
	}
	generatedFiles := map[string]string{
		fmt.Sprintf("internal/repository/%s/%s.go", packageFileName, modelFileName): string(rootTemplate),
	}

	return functions.Walk([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, parameters)),
		functions.MakeFunc(dto.Add(parameter.AddDataTransferObject{
			ProjectConfigWithPackage: parameter.ProjectConfigWithPackage{
				ProjectConfig: request.ProjectConfig,
				PackageName:   convention.PathName(path.Join(request.PackageName, request.RepositoryName)),
			},
			Name: request.FunctionName,
			Type: "Repository",
		})),
		functions.MakeFunc(Generate(request.ProjectConfig)),
	})
}

package repository

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Add function to add repository.
func Add(request parameter.AddRepository) error {
	var (
		generatedFolders = []string{
			"internal",
			"internal/repository",
		}
		generatedFiles = make(map[string]string)
		repositoryPath = request.RepositoryPath()
	)

	generatedFolders = append(generatedFolders, request.GeneratedFolders()...)
	generatedFiles = map[string]string{
		request.FunctionPath(): string(repositoryFunctionTemplate),
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, map[string]interface{}{
			"PackageName":   convention.PackageName(request.RepositoryName),
			"ProjectModule": request.ProjectConfig.Module,
			"FunctionName":  convention.FunctionName(request.FunctionName),
			"InterfaceName": convention.InterfaceName(request.FunctionName),
		})),
		functions.MakeFunc(Generate(parameter.ProjectConfigWithRepository{
			ProjectConfig:  request.ProjectConfig,
			RepositoryPath: repositoryPath,
			RepositoryName: request.RepositoryName,
		})),
	})
}

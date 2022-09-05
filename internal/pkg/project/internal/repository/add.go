package repository

import (
	"fmt"
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
		repositoryPath string
	)

	if request.PackageName != "" {
		packagePath := fmt.Sprintf("internal/repository/%s", convention.PackageName(request.PackageName))
		repositoryPath = fmt.Sprintf("%s/%s", packagePath, convention.PackageName(request.RepositoryName))
		generatedFolders = append(generatedFolders, packagePath, repositoryPath)
	} else {
		repositoryPath = fmt.Sprintf("internal/repository/%s", convention.PackageName(request.RepositoryName))
		generatedFolders = append(generatedFolders, repositoryPath)
	}

	generatedFiles = map[string]string{
		fmt.Sprintf("%s/%s.go", repositoryPath, convention.FileName(request.FunctionName)): string(repositoryFunctionTemplate),
	}

	return functions.Walk([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, map[string]interface{}{
			"PackageName":    convention.PackageName(request.PackageName),
			"ProjectModule":  request.ProjectConfig.Module,
			"RepositoryName": convention.ToUpperFirstLetter(request.RepositoryName),
		})),
		functions.MakeFunc(Generate(parameter.ProjectConfigWithRepository{
			ProjectConfig:  request.ProjectConfig,
			RepositoryPath: repositoryPath,
			PackageName:    request.PackageName,
			RepositoryName: request.RepositoryName,
		})),
	})
}

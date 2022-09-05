package repository

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Remove function to remove repository.
func Remove(request parameter.RemoveRepository) error {
	var repositoryPath = request.RepositoryPath()
	_ = file.Remove(fmt.Sprintf("%s/%s.go",
		repositoryPath, convention.FileName(request.FunctionName)))
	return Generate(parameter.ProjectConfigWithRepository{
		ProjectConfig:  request.ProjectConfig,
		PackageName:    request.PackageName,
		RepositoryName: request.RepositoryName,
		RepositoryPath: repositoryPath,
	})
}

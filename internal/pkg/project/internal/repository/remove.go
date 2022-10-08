package repository

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/file"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Remove function to remove repository.
func Remove(request parameter.RemoveRepository) error {
	var repositoryPath = request.RepositoryPath()
	_ = file.Remove(request.FunctionPath())
	return Generate(parameter.ProjectConfigWithRepository{
		ProjectConfig:  request.ProjectConfig,
		RepositoryName: request.RepositoryName,
		RepositoryPath: repositoryPath,
	})
}

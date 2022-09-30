package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository/dto"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository/model"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/root"
)

// RemoveRepository function to remove a repository.
func RemoveRepository(request parameter.RemoveRepository) error {
	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(model.Remove(parameter.RemoveModel{
			ProjectConfig:  request.ProjectConfig,
			ModelName:      request.ModelName,
			RepositoryName: request.RepositoryName,
			SkipIfError:    true,
		})),
		functions.MakeFunc(dto.Remove(parameter.RemoveRepositoryDataTransferObject{
			ProjectConfigWithPackage: parameter.ProjectConfigWithPackage{
				ProjectConfig: request.ProjectConfig,
				PackageName:   request.PackageName,
			},
			RepositoryName: request.RepositoryName,
			Name:           request.FunctionName,
			Type:           "Repository",
		})),
		functions.MakeFunc(repository.Remove(request)),
		functions.MakeFunc(root.Tidy()),
		functions.MakeFunc(root.Fmt()),
	})
}

package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository/dto"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository/model"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/root"
)

// AddRepository function to add a repository.
func AddRepository(request parameter.AddRepository) error {
	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(model.Add(parameter.AddModel{
			ProjectConfig:  request.ProjectConfig,
			RepositoryName: request.RepositoryName,
			ModelName:      request.ModelName,
			SkipIfError:    true,
		})),
		functions.MakeFunc(dto.Add(parameter.AddRepositoryDataTransferObject{
			ProjectConfigWithPackage: parameter.ProjectConfigWithPackage{
				ProjectConfig: request.ProjectConfig,
				PackageName:   request.PackageName,
			},
			RepositoryName: request.RepositoryName,
			Name:           request.FunctionName,
			Type:           "Repository",
		})),
		functions.MakeFunc(repository.Add(request)),
		functions.MakeFunc(root.Tidy()),
	})
}

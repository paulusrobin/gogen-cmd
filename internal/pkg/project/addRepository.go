package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository/dto"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository/model"
	"path"
)

// AddRepository function to add a repository.
func AddRepository(request parameter.AddRepository) error {
	return functions.Walk([]functions.Func{
		functions.MakeFunc(model.Add(parameter.AddModel{
			ProjectConfig: request.ProjectConfig,
			ModelName:     request.ModelName,
			SkipIfError:   true,
		})),
		functions.MakeFunc(dto.Add(parameter.AddDataTransferObject{
			ProjectConfigWithPackage: parameter.ProjectConfigWithPackage{
				ProjectConfig: request.ProjectConfig,
				PackageName:   convention.PackageName(path.Join(request.PackageName, request.RepositoryName)),
			},
			Name: request.FunctionName,
			Type: "Repository",
		})),
		functions.MakeFunc(repository.Add(request)),
	})
}

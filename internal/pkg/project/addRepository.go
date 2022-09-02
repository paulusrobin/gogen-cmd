package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository/model"
)

// AddRepository function to add a repository.
func AddRepository(request parameter.AddRepository) error {
	return functions.Walk([]functions.Func{
		functions.MakeFunc(model.Add(parameter.AddModel{
			ProjectConfig: request.ProjectConfig,
			PackageName:   request.PackageName,
			ModelName:     request.ModelName,
		})),
		functions.MakeFunc(repository.Add(request)),
	})
}

package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/repository/model"
)

// RemoveRepository function to remove a repository.
func RemoveRepository(request parameter.RemoveRepository) error {
	return functions.Walk([]functions.Func{
		functions.MakeFunc(model.Remove(parameter.RemoveModel{
			ProjectConfig: request.ProjectConfig,
			PackageName:   request.PackageName,
			ModelName:     request.ModelName,
		})),
		functions.MakeFunc(repository.Remove(request)),
	})
}

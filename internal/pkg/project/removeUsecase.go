package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/dto"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/usecase"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/root"
)

// RemoveUsecase function to remove an endpoint.
func RemoveUsecase(request parameter.RemoveUsecase) error {
	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(dto.Remove(parameter.RemoveDataTransferObject{
			ProjectConfigWithPackage: parameter.ProjectConfigWithPackage{
				ProjectConfig: request.ProjectConfig,
				PackageName:   request.PackageName,
			},
			Name: request.FunctionName,
			Type: "Usecase",
		})),
		functions.MakeFunc(usecase.Remove(request)),
		functions.MakeFunc(root.Tidy()),
	})
}

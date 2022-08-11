package pkg

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/dto"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/usecase"
)

// AddUsecase function.
func AddUsecase(request parameter.AddUsecase) error {
	return functions.Walk([]functions.Func{
		functions.MakeFunc(dto.Add(parameter.AddDataTransferObject{
			ProjectConfigWithPackage: parameter.ProjectConfigWithPackage{
				ProjectConfig: request.ProjectConfig,
				PackageName:   request.PackageName,
			},
			Name: request.FunctionName,
			Type: "Usecase",
		})),
		functions.MakeFunc(usecase.Add(request)),
	})
}

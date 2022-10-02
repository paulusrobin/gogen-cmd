package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/dto"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/encoding"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/endpoint"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/usecase"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/root"
)

// AddEndpoint function to add an endpoint.
func AddEndpoint(request parameter.AddEndpoint) error {
	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(dto.Add(parameter.AddDataTransferObject{
			ProjectConfigWithPackage: parameter.ProjectConfigWithPackage{
				ProjectConfig: request.ProjectConfig,
				PackageName:   request.PackageName,
			},
			Name: request.EndpointName,
			Type: "Endpoint",
		})),
		functions.MakeFunc(usecase.Generate(parameter.ProjectConfigWithPackage{
			ProjectConfig: request.ProjectConfig,
			PackageName:   request.PackageName,
		})),
		functions.MakeFunc(encoding.Add(request)),
		functions.MakeFunc(endpoint.Add(request)),
		functions.MakeFunc(root.Tidy()),
		functions.MakeFunc(root.Fmt()),
	})
}

package pkg

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/dto"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/encoding"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/endpoint"
)

// RemoveEndpoint function to remove an endpoint.
func RemoveEndpoint(request parameter.RemoveEndpoint) error {
	return functions.Walk([]functions.Func{
		functions.MakeFunc(dto.Remove(parameter.RemoveDataTransferObject{
			ProjectConfigWithPackage: parameter.ProjectConfigWithPackage{
				ProjectConfig: request.ProjectConfig,
				PackageName:   request.PackageName,
			},
			Name: request.EndpointName,
			Type: "Endpoint",
		})),
		functions.MakeFunc(encoding.Remove(request)),
		functions.MakeFunc(endpoint.Remove(request)),
	})
}

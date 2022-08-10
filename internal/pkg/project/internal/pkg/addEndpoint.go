package pkg

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/dto"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/encoding"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/endpoint"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg/usecase"
)

// AddEndpoint function.
func AddEndpoint(request parameter.AddEndpoint) error {
	return functions.Walk([]functions.Func{
		functions.MakeFunc(dto.AddEndpoint(request)),
		functions.MakeFunc(encoding.Add(request)),
		functions.MakeFunc(endpoint.Add(request)),
		functions.MakeFunc(usecase.Generate(parameter.ProjectConfigWithPackage{
			ProjectConfig: request.ProjectConfig,
			PackageName:   request.PackageName,
		})),
	})
}

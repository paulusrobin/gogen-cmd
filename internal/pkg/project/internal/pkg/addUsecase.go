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
		functions.MakeFunc(dto.AddUsecase(request)),
		functions.MakeFunc(usecase.Add(request)),
	})
}

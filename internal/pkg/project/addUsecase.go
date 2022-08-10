package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal"
)

// AddUsecase function to add an usecase.
func AddUsecase(parameter parameter.AddUsecase) error {
	return internal.AddUsecase(parameter)
}

package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/dto"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal"
)

// AddUsecase function to add an usecase.
func AddUsecase(parameter dto.AddUsecaseParameter) error {
	return internal.AddUsecase(parameter)
}

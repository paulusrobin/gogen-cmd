package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/helper"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal"
)

// AddUsecase function to add an usecase.
func AddUsecase(parameter helper.AddUsecaseParameter) error {
	return internal.AddUsecase(parameter)
}

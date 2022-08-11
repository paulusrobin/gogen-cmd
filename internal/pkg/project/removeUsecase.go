package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal"
)

// RemoveUsecase function to remove an endpoint.
func RemoveUsecase(parameters parameter.RemoveUsecase) error {
	return internal.RemoveUsecase(parameters)
}

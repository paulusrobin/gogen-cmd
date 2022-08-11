package internal

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg"
)

// RemoveUsecase function to remove an endpoint.
func RemoveUsecase(parameter parameter.RemoveUsecase) error {
	return pkg.RemoveUsecase(parameter)
}

package internal

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg"
)

// AddUsecase function.
func AddUsecase(request parameter.AddUsecase) error {
	return pkg.AddUsecase(request)
}

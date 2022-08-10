package internal

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg"
)

// AddEndpoint function.
func AddEndpoint(request parameter.AddEndpoint) error {
	return pkg.AddEndpoint(request)
}

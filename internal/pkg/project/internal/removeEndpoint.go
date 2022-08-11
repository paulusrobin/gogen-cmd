package internal

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/pkg"
)

// RemoveEndpoint function to remove an endpoint.
func RemoveEndpoint(parameter parameter.RemoveEndpoint) error {
	return pkg.RemoveEndpoint(parameter)
}

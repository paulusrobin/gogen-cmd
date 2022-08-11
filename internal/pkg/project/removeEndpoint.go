package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal"
)

// RemoveEndpoint function to remove an endpoint.
func RemoveEndpoint(parameters parameter.RemoveEndpoint) error {
	return internal.RemoveEndpoint(parameters)
}

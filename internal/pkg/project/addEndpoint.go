package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal"
)

// AddEndpoint function to add an endpoint.
func AddEndpoint(parameter parameter.AddEndpoint) error {
	return internal.AddEndpoint(parameter)
}

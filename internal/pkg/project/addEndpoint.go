package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/dto"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal"
)

// AddEndpoint function to add an endpoint.
func AddEndpoint(parameter dto.AddEndpointParameter) error {
	return internal.AddEndpoint(parameter)
}

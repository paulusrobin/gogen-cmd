package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/helper"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal"
)

// AddEndpoint function to add an endpoint.
func AddEndpoint(parameter helper.AddEndpointParameter) error {
	return internal.AddEndpoint(parameter)
}

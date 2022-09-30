package cmd

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
)

// Remove function to remove command.
func Remove(request parameter.RemoveCommand) error {
	_ = directory.Remove(fmt.Sprintf("cmd/%s", convention.CommandName(request.CommandName)))
	return Generate(request.ProjectConfig)
}

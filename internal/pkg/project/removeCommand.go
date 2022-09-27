package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/cmd"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/root"
)

// RemoveCommand function to remove a command.
func RemoveCommand(request parameter.RemoveCommand) error {
	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(cmd.Remove(request)),
		functions.MakeFunc(root.Tidy()),
	})
}

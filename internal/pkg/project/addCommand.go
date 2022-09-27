package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/cmd"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/root"
)

// AddCommand function to add a command.
func AddCommand(request parameter.AddCommand) error {
	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(cmd.Add(request)),
		functions.MakeFunc(root.Tidy()),
	})
}

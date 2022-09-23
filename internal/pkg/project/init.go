package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/cmd"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/root"
)

// Init function to initialize project folder files.
func Init(cfg parameter.ProjectConfig) error {
	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(root.Init(cfg)),
		functions.MakeFunc(cmd.Init(cfg)),
		functions.MakeFunc(internal.Init(cfg)),
		functions.MakeFunc(root.Tidy()),
	})
}

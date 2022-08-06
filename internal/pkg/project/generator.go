package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/cmd"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/helper"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/root"
)

// Generate function to generate project folder files.
func Generate(cfg helper.ProjectConfig) error {
	functions := []func(cfg helper.ProjectConfig) error{
		root.Generate,
		cmd.Generate,
		internal.Generate,
	}

	for _, fn := range functions {
		if err := fn(cfg); err != nil {
			return err
		}
	}
	return nil
}

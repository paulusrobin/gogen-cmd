package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/cmd"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/dto"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/root"
)

// Generate function to generate project folder files.
func Generate(cfg dto.ProjectConfig) error {
	functions := []func(cfg dto.ProjectConfig) error{
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

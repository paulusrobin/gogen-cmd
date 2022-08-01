package project

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/helper"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/root"
)

// Generate function to generate project folder files.
func Generate(cfg helper.ProjectConfig) error {
	return root.Generate(cfg)
}

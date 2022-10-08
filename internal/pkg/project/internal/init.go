package internal

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/config"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/greeting"
)

// Init function to generate internal folder files.
func Init(cfg parameter.ProjectConfig) error {
	generatedFolders := []string{
		"internal",
		"internal/config",
		"internal/pkg",
		"internal/repository",
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(cfg.Path, generatedFolders)),
		functions.MakeFunc(config.Init(cfg)),
		functions.MakeFunc(greeting.Init(cfg)),
	})
}

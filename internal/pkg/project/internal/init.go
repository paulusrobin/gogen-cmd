package internal

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/config"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/greeting"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/server"
)

// Init function to generate internal folder files.
func Init(cfg parameter.ProjectConfig) error {
	generatedFolders := []string{
		"internal",
		"internal/config",
		"internal/pkg",
		"internal/repository",
		"internal/server",
	}

	return functions.Walk([]functions.Func{
		functions.MakeFunc(generator.Folder(cfg.Path, generatedFolders)),
		functions.MakeFunc(config.Init(cfg)),
		functions.MakeFunc(greeting.Init(cfg)),
		functions.MakeFunc(server.Init(cfg)),
	})
}

package internal

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/config"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/greeting"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/internal/server"
	"path"
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

	for _, folderPath := range generatedFolders {
		generatedPath := path.Join(cfg.Path, folderPath)
		if directory.Exist(generatedPath) {
			continue
		}
		if err := directory.Make(generatedPath); err != nil {
			return err
		}
	}

	return functions.Walk([]functions.Func{
		functions.MakeFunc(config.Init(cfg)),
		functions.MakeFunc(greeting.Init(cfg)),
		functions.MakeFunc(server.Init(cfg)),
	})
}

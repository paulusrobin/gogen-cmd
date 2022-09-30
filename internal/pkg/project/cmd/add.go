package cmd

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/convention"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/functions"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/generator"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"path"
)

// Add function to add command.
func Add(request parameter.AddCommand) error {
	parameters := map[string]interface{}{
		"ProjectName":   request.Name,
		"ProjectModule": request.Module,
		"ServerPackage": convention.PackageName(request.CommandName),
		"ServerName":    convention.CommandName(request.CommandName),
	}

	commandPath := fmt.Sprintf("cmd/%s", convention.CommandName(request.CommandName))
	commandFilePath := path.Join(commandPath, fmt.Sprintf("%s.go", convention.CommandName(request.CommandName)))
	generatedFolders := []string{
		"cmd",
		commandPath,
	}
	generatedFiles := map[string]string{
		commandFilePath: string(serverTemplate),
	}

	return functions.WalkSkipErrors([]functions.Func{
		functions.MakeFunc(generator.Folder(request.Path, generatedFolders)),
		functions.MakeFunc(generator.File(request.Path, generatedFiles, parameters)),
		functions.MakeFunc(Generate(request.ProjectConfig)),
	})
}

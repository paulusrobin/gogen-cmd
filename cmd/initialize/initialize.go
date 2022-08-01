package initialize

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/helper"
	"github.com/spf13/cobra"
	"log"
	"path"
)

var (
	projectName   string
	projectModule string
	inFolder      bool
	requiredFlags []string
)

func init() {
	requiredFlags = []string{"name", "module"}
}

func initProject(cmd *cobra.Command, args []string) error {
	log.Printf("project helper:\nproject name: %s\nproject module: %s", projectName, projectModule)
	log.Printf("start initialiazing project")

	var basePath = directory.Pwd()
	if inFolder {
		if !directory.Empty(basePath) {
			return fmt.Errorf("current folder is not empty")
		}
	} else {
		if err := directory.Make(path.Join(basePath, projectName)); err != nil {
			return err
		}
	}

	if err := project.Generate(helper.ProjectConfig{
		Path:   basePath,
		Name:   projectName,
		Module: projectModule,
	}); err != nil {
		return err
	}

	log.Printf("project initialized")
	return nil
}

// Cmd expose command runner
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "gogen initialize is command to generate gogen project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			for _, required := range requiredFlags {
				if err := cmd.MarkFlagRequired(required); err != nil {
					return err
				}
			}
			return nil
		},
		RunE: initProject,
	}
	cmd.Flags().StringVarP(&projectName, "name", "n", "", "generated project name (required)")
	cmd.Flags().StringVarP(&projectModule, "module", "m", "", "generated project module (required)")
	cmd.Flags().BoolVarP(&inFolder, "in-folder", "f", false, "generate project in current folder, current folder must be empty")
	return cmd
}

package remove

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project"
	"github.com/spf13/cobra"
	"log"
)

var (
	commandName   string
	projectName   string
	projectModule string
	requiredFlags []string
)

func init() {
	requiredFlags = []string{"cmd"}
}

func runner(cmd *cobra.Command, args []string) error {
	log.Printf("removing a command")

	var basePath = directory.Pwd()
	if err := project.RemoveCommand(parameter.RemoveCommand{
		ProjectConfig: parameter.ProjectConfig{
			Path:   basePath,
			Name:   projectName,
			Module: projectModule,
		},
		CommandName: commandName,
	}); err != nil {
		return err
	}

	log.Printf("succesfully removing %s command", commandName)
	return nil
}

// Cmd expose command runner
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove a command on gogen project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			for _, required := range requiredFlags {
				if err := cmd.MarkFlagRequired(required); err != nil {
					return err
				}
			}
			var basePath = directory.Pwd()
			p, m, err := project.ValidateProject(basePath)
			if err != nil {
				return err
			}

			if projectName == "" {
				projectName = p
			}

			if projectModule == "" {
				projectModule = m
			}

			if projectName == "" {
				return fmt.Errorf("project name is required")
			}

			if projectModule == "" {
				return fmt.Errorf("project name is required")
			}
			return nil
		},
		RunE: runner,
	}
	cmd.Flags().StringVarP(&projectName, "name", "n", "", "generated project name")
	cmd.Flags().StringVarP(&projectModule, "module", "m", "", "generated project module")
	cmd.Flags().StringVarP(&commandName, "cmd", "c", "", "generated command name (required)")
	return cmd
}

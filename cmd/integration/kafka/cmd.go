package kafka

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project"
	"github.com/spf13/cobra"
	"log"
)

var (
	projectName   string
	projectModule string
)

func runner(cmd *cobra.Command, args []string) error {
	log.Printf("adding a kafka integration")

	var basePath = directory.Pwd()
	if err := project.AddDependencies(parameter.AddDependency{
		ProjectConfig: parameter.ProjectConfig{
			Path:   basePath,
			Name:   projectName,
			Module: projectModule,
		},
		Dependency: parameter.Kafka,
	}); err != nil {
		return err
	}

	log.Printf("succesfully add kafka integration")
	return nil
}

// Cmd expose command runner
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kafka",
		Short: "add kafka publisher dependency on gogen project",
		PreRunE: func(cmd *cobra.Command, args []string) error {
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
	return cmd
}

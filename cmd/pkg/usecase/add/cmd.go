package add

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
	packageName   string
	functionName  string
	requiredFlags []string
)

func init() {
	requiredFlags = []string{"package", "function"}
}

func runner(cmd *cobra.Command, args []string) error {
	log.Printf("adding a usecase")

	var basePath = directory.Pwd()
	if err := project.AddUsecase(parameter.AddUsecase{
		ProjectConfig: parameter.ProjectConfig{
			Path:   basePath,
			Name:   projectName,
			Module: projectModule,
		},
		PackageName:  packageName,
		FunctionName: functionName,
	}); err != nil {
		return err
	}

	log.Printf("succesfully added %s usecase on %s package", functionName, packageName)
	return nil
}

// Cmd expose command runner
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add a usecase on gogen project internal package",
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
	cmd.Flags().StringVarP(&projectName, "name", "n", "", "generated project name (required)")
	cmd.Flags().StringVarP(&projectModule, "module", "m", "", "generated project module (required)")
	cmd.Flags().StringVarP(&packageName, "package", "p", "", "generated package name (required)")
	cmd.Flags().StringVarP(&functionName, "function", "f", "", "generated function name (required)")
	return cmd
}

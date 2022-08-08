package add

import (
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project/helper"
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
	requiredFlags = []string{"name", "module", "package", "function"}
}

func runner(cmd *cobra.Command, args []string) error {
	log.Printf("adding a usecase")

	var basePath = directory.Pwd()
	if err := project.AddUsecase(helper.AddUsecaseParameter{
		ProjectConfig: helper.ProjectConfig{
			Path:   basePath,
			Name:   projectName,
			Module: projectModule,
		},
		PackageName:  packageName,
		FunctionName: functionName,
	}); err != nil {
		return err
	}

	log.Printf("succesfully added a new usecase")
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

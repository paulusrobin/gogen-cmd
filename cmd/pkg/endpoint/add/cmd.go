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
	endpointName  string
	requiredFlags []string
)

func init() {
	requiredFlags = []string{"name", "module", "package", "endpoint"}
}

func runner(cmd *cobra.Command, args []string) error {
	log.Printf("adding an endpoint")

	var basePath = directory.Pwd()
	if err := project.AddEndpoint(helper.AddEndpointParameter{
		ProjectConfig: helper.ProjectConfig{
			Path:   basePath,
			Name:   projectName,
			Module: projectModule,
		},
		PackageName:  packageName,
		EndpointName: endpointName,
	}); err != nil {
		return err
	}

	log.Printf("succesfully added a new endpoint")
	return nil
}

// Cmd expose command runner
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add an endpoint on gogen project internal package",
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
	cmd.Flags().StringVarP(&endpointName, "endpoint", "e", "", "generated endpoint name (required)")
	return cmd
}

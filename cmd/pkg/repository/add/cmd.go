package add

import (
	"fmt"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/directory"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/parameter"
	"github.com/paulusrobin/gogen-cmd/internal/pkg/project"
	"github.com/spf13/cobra"
	"log"
	"path"
)

var (
	projectName    string
	projectModule  string
	repositoryName string
	packageName    string
	functionName   string
	withModel      bool
	requiredFlags  []string
)

func init() {
	requiredFlags = []string{"repository", "function"}
}

func runner(cmd *cobra.Command, args []string) error {
	log.Printf("adding a repository")

	modelName := ""
	if withModel {
		modelName = repositoryName
	}

	var basePath = directory.Pwd()
	if err := project.AddRepository(parameter.AddRepository{
		ProjectConfig: parameter.ProjectConfig{
			Path:   basePath,
			Name:   projectName,
			Module: projectModule,
		},
		RepositoryName: repositoryName,
		PackageName:    packageName,
		FunctionName:   functionName,
		ModelName:      modelName,
	}); err != nil {
		return err
	}

	log.Printf("succesfully added %s on %s repository package", functionName, path.Join(packageName, repositoryName))
	return nil
}

// Cmd expose command runner
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add a repository on gogen project internal package",
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
	cmd.Flags().StringVarP(&repositoryName, "repository", "r", "", "generated repository name (required)")
	cmd.Flags().StringVarP(&functionName, "function", "f", "", "generated function name (required)")
	cmd.Flags().StringVarP(&packageName, "package", "p", "", "generated package name (optional)")
	cmd.Flags().BoolVarP(&withModel, "model", "o", false, "generate model (optional)")
	return cmd
}

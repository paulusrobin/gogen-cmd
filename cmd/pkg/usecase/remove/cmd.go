package remove

import (
	"github.com/spf13/cobra"
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
	panic("implement me")
}

// Cmd expose command runner
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove a usecase on gogen project internal package",
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

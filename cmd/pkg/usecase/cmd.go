package usecase

import (
	"github.com/paulusrobin/gogen-cmd/cmd/pkg/usecase/add"
	"github.com/paulusrobin/gogen-cmd/cmd/pkg/usecase/remove"
	"github.com/spf13/cobra"
)

// Cmd expose command runner
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "usecase",
		Short: "generate gogen project internal package usecase",
	}
	cmd.AddCommand(add.Cmd(), remove.Cmd())
	return cmd
}

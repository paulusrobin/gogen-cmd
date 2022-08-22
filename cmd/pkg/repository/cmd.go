package repository

import (
	"github.com/paulusrobin/gogen-cmd/cmd/pkg/repository/add"
	"github.com/paulusrobin/gogen-cmd/cmd/pkg/repository/remove"
	"github.com/spf13/cobra"
)

// Cmd expose command runner
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repository",
		Short: "generate gogen project internal package repository",
	}
	cmd.AddCommand(add.Cmd(), remove.Cmd())
	return cmd
}

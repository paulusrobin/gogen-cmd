package command

import (
	"github.com/paulusrobin/gogen-cmd/cmd/command/add"
	"github.com/paulusrobin/gogen-cmd/cmd/command/remove"
	"github.com/spf13/cobra"
)

// Cmd expose command runner
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command",
		Short: "generate gogen project cmd package",
	}
	cmd.AddCommand(add.Cmd(), remove.Cmd())
	return cmd
}

package endpoint

import (
	"github.com/paulusrobin/gogen-cmd/cmd/pkg/endpoint/add"
	"github.com/paulusrobin/gogen-cmd/cmd/pkg/endpoint/remove"
	"github.com/spf13/cobra"
)

// Cmd expose command runner
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pkg-endpoint",
		Short: "generate gogen project internal package endpoint",
	}
	cmd.AddCommand(add.Cmd(), remove.Cmd())
	return cmd
}

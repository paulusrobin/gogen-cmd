package integration

import (
	"github.com/paulusrobin/gogen-cmd/cmd/integration/postgres"
	"github.com/paulusrobin/gogen-cmd/cmd/integration/redis"
	"github.com/spf13/cobra"
)

// Cmd expose command runner
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "integration",
		Short: "generate gogen project integration package",
	}
	cmd.AddCommand(postgres.Cmd(), redis.Cmd())
	return cmd
}

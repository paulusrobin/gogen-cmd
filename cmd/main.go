package main

import (
	"github.com/paulusrobin/gogen-cmd/cmd/initialize"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "gogen",
	Short: "gogen is golang gogen project generator",
}

func main() {
	cmd.AddCommand(
		initialize.Cmd(),
	)

	// execute command
	if err := cmd.Execute(); err != nil {
		panic("can't execute cmd")
	}
}

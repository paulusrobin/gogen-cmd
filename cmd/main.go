package main

import (
	"github.com/paulusrobin/gogen-cmd/cmd/command"
	"github.com/paulusrobin/gogen-cmd/cmd/initialize"
	"github.com/paulusrobin/gogen-cmd/cmd/integration"
	"github.com/paulusrobin/gogen-cmd/cmd/pkg/endpoint"
	"github.com/paulusrobin/gogen-cmd/cmd/pkg/repository"
	"github.com/paulusrobin/gogen-cmd/cmd/pkg/usecase"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "gogen",
	Short: "gogen is golang gogen project generator",
}

func main() {
	cmd.AddCommand(
		initialize.Cmd(),
		endpoint.Cmd(),
		usecase.Cmd(),
		repository.Cmd(),
		command.Cmd(),
		integration.Cmd(),
	)

	// execute command
	if err := cmd.Execute(); err != nil {
		panic("can't execute cmd")
	}
}

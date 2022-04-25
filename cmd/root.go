package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/trangmaiq/short/cmd/serve"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "short",
	}

	cmd.AddCommand(serve.NewServeCmd())

	return cmd
}

func Execute() {
	cmd := NewRootCmd()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

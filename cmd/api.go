package cmd

import (
	"github.com/mrayone/golang-learn/cmd/webserver"
	"github.com/spf13/cobra"
)

var (
	apiCommand = &cobra.Command{
		Use:   "api",
		Short: "Startup HTTP API",
		Long:  "Startup HTTP API",
		RunE:  apiExecute,
	}
)

func init() {
	rootCmd.AddCommand(apiCommand)
}

func apiExecute(_ *cobra.Command, _ []string) error {

	webserver.Serve()

	return nil
}

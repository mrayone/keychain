package cmd

import (
	http "github.com/mrayone/golang-learn/protocol"
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

func apiExecute(cmd *cobra.Command, args []string) error {

	http.Serve()

	return nil
}

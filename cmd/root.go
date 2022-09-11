package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Root short command",
	Long:  "Root long command",
}

// Execute root command
func Execute() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("unexpected error while executing command %v", err)
		}
	}()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("error while executing command %v", err)
	}
}

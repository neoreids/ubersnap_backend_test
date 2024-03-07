package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "",
	Short: "app command",
}

func Execute() error {
	rootCmd.AddCommand(
		httpCommand,
	)

	return rootCmd.Execute()
}
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitrats",
	Short: "CLI to interact with gitrats API ",
	Long:  "Manage your check-ins and groups, and track your weekly commits, directly from the terminal",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

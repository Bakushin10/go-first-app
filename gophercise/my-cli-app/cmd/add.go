package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd is the root command of the CLI application
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called.")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

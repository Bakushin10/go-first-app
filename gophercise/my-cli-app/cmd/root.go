package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd is the root command of the CLI application
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Welcome to Task CLI! Run your tasks here.")
	},
}

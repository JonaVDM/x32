package cmd

import (
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log all the messages send by the console",
	Run: func(cmd *cobra.Command, args []string) {
		client.Logger()
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}

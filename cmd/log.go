package cmd

import (
	"github.com/jonavdm/x32/pkg/x32"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log all the messages send by the console",
	Run: func(cmd *cobra.Command, args []string) {
		client.Send(x32.Message{Message: x32.MessageInfo})
		client.Logger()
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
)

var setOnFaderCmd = &cobra.Command{
	Use:   "sendOnFader [status]",
	Short: "The send on fader button",
	Long:  "The send on fader button\n\nStatus should be either a 1 (on) or 0 (off)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "0" {
			client.SetSendOnFader(false)
		}

		if args[0] == "1" {
			client.SetSendOnFader(true)
		}
	},
}

func init() {
	setCmd.AddCommand(setOnFaderCmd)
}

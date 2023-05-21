package cmd

import (
	"github.com/spf13/cobra"
)

// SetCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value on the board",
}

func init() {
	rootCmd.AddCommand(setCmd)
}

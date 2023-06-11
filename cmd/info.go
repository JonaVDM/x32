package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd represents the log command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about the console",
	Run: func(cmd *cobra.Command, args []string) {
		xinfo, err := client.GetXInfo()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Ip Address: %s\n", xinfo.Values[0].String)
			fmt.Printf("Network Name: %s\n", xinfo.Values[1].String)
			fmt.Printf("Console Model: %s\n", xinfo.Values[2].String)
			fmt.Printf("Console version: %s\n", xinfo.Values[3].String)
		}

		status, err := client.GetStatus()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Console status: %s\n", status.Values[0].String)
			fmt.Printf("Server name: %s\n", status.Values[2].String)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}

package cmd

import (
	"fmt"

	"github.com/jonavdm/x32/pkg/x32"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log all the messages send by the console",
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan x32.Message)
		client.Subscribe(ch)

		defer func() {
			client.UnSubscribe(ch)
			close(ch)
		}()

		for {
			msg := <-ch

			fmt.Printf("%s", msg.Message)
			for _, value := range msg.Values {
				fmt.Printf(" [%c %s]", value.Type, valueToText(value))
			}
			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}

func valueToText(value x32.Value) string {
	switch value.Type {
	case 'f':
		return fmt.Sprintf("%f", value.Float)

	case 'i':
		return fmt.Sprintf("%d", value.Int)

	case 's':
		return value.String

	default:
		return string(value.Value)
	}
}

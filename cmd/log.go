package cmd

import (
	"fmt"

	"github.com/jonavdm/x32/pkg/osc"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log all the messages send by the console",
	Run: func(cmd *cobra.Command, args []string) {
		ch := make(chan osc.Message)
		client.Connection.Subscribe(ch)

		defer func() {
			client.Connection.UnSubscribe(ch)
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

func valueToText(value osc.Value) string {
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

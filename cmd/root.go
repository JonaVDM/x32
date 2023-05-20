package cmd

import (
	"fmt"
	"os"

	"github.com/jonavdm/x32/pkg/x32"
	"github.com/spf13/cobra"
)

var client x32.Client
var ip, port string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "x32",
	Short: "Some fun things that can be done on the x32 console",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer client.Close()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initClient)

	rootCmd.PersistentFlags().StringVar(&ip, "ip", "", "The IP address of the console")
	rootCmd.MarkPersistentFlagRequired("ip")

	rootCmd.PersistentFlags().StringVar(&port, "port", "10023", "The port the console is using")
}

func initClient() {
	var err error
	client, err = x32.NewClient(ip + ":" + port)
	if err != nil {
		fmt.Println("Error: Could not connect to the console")
		fmt.Println("")
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var socket string

var rootCmd = &cobra.Command{
	Use:   "hostsetup",
	Short: "Set hostname and change the list of dns servers.",
	Long: `Hostsetup is a utility for Linux that set hostname and
change (add and delete) the list of DNS servers.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&socket, "server", "s", "0.0.0.0:8081", "Server and port of the service")
}

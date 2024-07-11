package cmd

import (
	"context"
	"fmt"
	"hostsetup-service/internal/client/app"
	hsv1 "hostsetup-service/protos/gen/hostsetup"
	"log"

	"github.com/spf13/cobra"
)

var listDnsServersCmd = &cobra.Command{
	Use:   "list-dns-servers",
	Short: "Return list of dns-servers",
	Long: `The command returns a list of dns-servers in a Linux system.
It takes no arguments. If success returns list, otherwise an error.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli, err := app.NewClient(socket)

		if err != nil {
			log.Fatal("connection creation error")
		}

		list, err := cli.ListDNSServers(context.Background(), &hsv1.EmptyRequest{})
		if err != nil {
			log.Fatal(err)
		}

		for _, dnsServer := range list.List {
			fmt.Println(dnsServer)
		}
	},
}

func init() {
	rootCmd.AddCommand(listDnsServersCmd)
}

package cmd

import (
	"context"
	"fmt"
	"hostsetup-service/internal/client/app"
	"hostsetup-service/pkg/config"
	hsv1 "hostsetup-service/protos/gen/hostsetup"
	"log"

	"github.com/spf13/cobra"
)

var deleteDnsServerCmd = &cobra.Command{
	Use:   "delete-dns-server",
	Short: "Delete dns server",
	Long: `The command delete the dns-server from a Linux system.
	It takes one argument. If success returns "success", otherwise an error.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Panic("incorrect number of arguments")
		}

		dnsServer := args[0]

		cfg := config.MustLoadConfig()

		cli, err := app.NewClient(fmt.Sprint(cfg.GRPCServer, ":", cfg.GRPCPort))
		if err != nil {
			log.Fatal("connection creation error")
		}

		success, err := cli.DeleteDNSServer(
			context.Background(),
			&hsv1.DNSServerRequest{DnsServer: dnsServer},
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(success)
	},
}

func init() {
	rootCmd.AddCommand(deleteDnsServerCmd)
}

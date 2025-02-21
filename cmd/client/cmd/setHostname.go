package cmd

import (
	"context"
	"fmt"
	"hostsetup-service/internal/client/app"
	hsv1 "hostsetup-service/protos/gen/hostsetup"
	"log"

	"github.com/spf13/cobra"
)

var setHostnameCmd = &cobra.Command{
	Use:   "set-hostname",
	Short: "Set new hostname",
	Long: `The command sets hostname for Linux. Takes only one argument -
the new hostname. If success returns "success", otherwise an error.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Panic("incorrect number of arguments")
		}

		hostname := args[0]

		cli, err := app.NewClient(socket)
		if err != nil {
			log.Fatal("connection creation error")
		}

		success, err := cli.SetHostname(
			context.Background(),
			&hsv1.HostnameRequest{Name: hostname},
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(success)
	},
}

func init() {
	rootCmd.AddCommand(setHostnameCmd)
}

package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mywg",
	Short: "My Wireguard VPN from CLI",
	Long:  `Deploy & connect to Wireguard VPN in under 5 minutes`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf(err.Error())
		os.Exit(1)
	}
}

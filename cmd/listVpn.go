package cmd

import (
	"log"

	"github.com/spf13/cobra"
	action "github.com/tiptophelmet/mywireguard/pkg/action"
)

var listVPNCmd = &cobra.Command{
	Use:   "list-vpn",
	Short: "List all VPNs",
	Long:  `This command lists all VPNs that have been created.`,
	Run: func(cmd *cobra.Command, args []string) {
		action := action.InitListVpnAction()
		if err := action.List(); err != nil {
			log.Fatalf(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(listVPNCmd)
}

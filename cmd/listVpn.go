package cmd

import (
	"github.com/spf13/cobra"
	action "github.com/tiptophelmet/mywireguard/pkg/action"
)

var listVPNCmd = &cobra.Command{
	Use:   "list-vpn",
	Short: "List all VPNs",
	Long:  `This command lists all VPNs that have been created.`,
	Run: func(cmd *cobra.Command, args []string) {
		listVpnAction := action.InitListVpnAction()
		listVpnAction.List()
	},
}

func init() {
	rootCmd.AddCommand(listVPNCmd)
}

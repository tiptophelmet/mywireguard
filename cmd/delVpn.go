package cmd

import (
	"github.com/spf13/cobra"
	action "github.com/tiptophelmet/mywireguard/pkg/action"
)

var vpnID string

var delVPNCmd = &cobra.Command{
	Use:   "del-vpn",
	Short: "Delete a VPN with the given VPN ID",
	Long:  `This command deletes a VPN with the specified VPN ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		getVpnAction := action.InitGetVpnAction()
		vpn := getVpnAction.Get(vpnID)

		delVpnAction := action.InitDeleteVpnAction(vpn)
		delVpnAction.DestroyInfra()
		delVpnAction.Delete()
	},
}

func init() {
	delVPNCmd.Flags().StringVar(&vpnID, "vpnid", "", "VPN ID of the VPN to delete")
	delVPNCmd.MarkFlagRequired("vpnid")
	rootCmd.AddCommand(delVPNCmd)
}

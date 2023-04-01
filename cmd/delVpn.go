package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	action "github.com/tiptophelmet/mywireguard/pkg/action"
)

var vpnID string

var delVPNCmd = &cobra.Command{
	Use:   "del-vpn",
	Short: "Delete a VPN with the given VPN ID",
	Long:  `This command deletes a VPN with the specified VPN ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		getAction := action.InitGetVpnAction()
		vpnEntry, err := getAction.Get(vpnID)
		if err != nil {
			fmt.Printf("[ERROR] %s", err.Error())
			return
		}

		delAction, err := action.InitDeleteVpnAction(vpnEntry)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return
		}
		
		delAction.DestroyInfra()
		delAction.Delete()
	},
}

func init() {
	delVPNCmd.Flags().StringVar(&vpnID, "vpnid", "", "VPN ID of the VPN to delete")
	delVPNCmd.MarkFlagRequired("vpnid")
	rootCmd.AddCommand(delVPNCmd)
}

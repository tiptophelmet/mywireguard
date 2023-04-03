package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tiptophelmet/mywireguard/pkg/action"
)

var clientID string

var delClientCmd = &cobra.Command{
	Use:   "del-client",
	Short: "Delete a VPN client with the given client ID",
	Long:  `This command deletes a VPN client with the specified client ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		getClientAction := action.InitGetClientAction()
		client := getClientAction.Get(vpnID, clientID)

		getVpnAction := action.InitGetVpnAction()
		vpn := getVpnAction.Get(client.VPNID)

		deleteClientAction := action.InitDeleteClientAction(client, vpn)
		deleteClientAction.Delete()
	},
}

func init() {
	delClientCmd.Flags().StringVar(&clientID, "clientid", "", "Client ID of the VPN client to delete")
	delClientCmd.MarkFlagRequired("clientid")
	delClientCmd.Flags().StringVar(&vpnID, "vpnid", "", "VPN ID of the client to delete")
	delClientCmd.MarkFlagRequired("vpnid")
	rootCmd.AddCommand(delClientCmd)
}

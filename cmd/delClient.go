package cmd

import (
	"fmt"

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
		client, err := getClientAction.Get(vpnID, clientID)
		if err != nil {
			fmt.Printf("[ERROR] %s", err.Error())
			return
		}

		getVpnAction := action.InitGetVpnAction()
		vpn, err := getVpnAction.Get(client.VPNID)
		if err != nil {
			fmt.Printf("[ERROR] %s", err.Error())
			return
		}

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

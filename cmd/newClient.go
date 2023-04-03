package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tiptophelmet/mywireguard/pkg/action"
)

var vpnid, conf string

var newClientCmd = &cobra.Command{
	Use:   "new-client",
	Short: "Create a new VPN client",
	Long:  `This command creates a new VPN client associated with the specified VPN ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("[INFO] Proceeding to create a new client for VPN: %s\n", vpnid)

		getVpnAction := action.InitGetVpnAction()
		vpn := getVpnAction.Get(vpnid)

		newClientAction := action.InitNewClientAction(vpn, conf)
		newClientAction.Prepare()
		newClientAction.Save()
		newClientAction.GenerateWireguardClientConf()
	},
}

func init() {
	newClientCmd.Flags().StringVar(&vpnid, "vpnid", "", "ID of the VPN to associate the client with")
	newClientCmd.Flags().StringVar(&conf, "conf", "", "Path to the VPN client configuration file")
	newClientCmd.MarkFlagRequired("vpnid")
	newClientCmd.MarkFlagRequired("conf")
	rootCmd.AddCommand(newClientCmd)
}

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tiptophelmet/mywireguard/pkg/action"
)

var listClientCmd = &cobra.Command{
	Use:   "list-client",
	Short: "List all clients for a VPN",
	Long:  `This command lists all clients associated with the specified VPN ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		listClientAction := action.InitListClientAction()
		listClientAction.List(vpnid)
	},
}

func init() {
	listClientCmd.Flags().StringVar(&vpnid, "vpnid", "", "ID of the VPN to list clients for")
	listClientCmd.MarkFlagRequired("vpnid")
	rootCmd.AddCommand(listClientCmd)
}

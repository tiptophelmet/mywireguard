package cmd

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	action "github.com/tiptophelmet/mywireguard/pkg/action"
	"github.com/tiptophelmet/mywireguard/pkg/cloud"
)

var tomlPath string

var newVPNCmd = &cobra.Command{
	Use:   "new-vpn",
	Short: "Create a new VPN using the provided TOML configuration",
	Long:  `This command creates a new VPN using the configuration specified in the provided TOML file.`,
	Run: func(cmd *cobra.Command, args []string) {
		cc := cloud.NewCloudConfig()
		cc.ImportToml(tomlPath)

		resolvedCloud, err := cc.InitCloud()
		if err != nil {
			log.Fatalf(err.Error())
		}

		vpnID := strings.TrimSuffix(filepath.Base(tomlPath), path.Ext(tomlPath))
		action, err := action.InitNewVpnAction(vpnID, resolvedCloud)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return
		}

		action.Prepare()
		action.ApplyInfra()
		action.Save()

		fmt.Println("[OK] VPN successfully deployed!")
	},
}

func init() {
	newVPNCmd.Flags().StringVar(&tomlPath, "toml", "", "Path to the VPN TOML configuration file")
	newVPNCmd.MarkFlagRequired("toml")
	rootCmd.AddCommand(newVPNCmd)
}

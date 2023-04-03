package action

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/tiptophelmet/mywireguard/paths"
	"github.com/tiptophelmet/mywireguard/pkg/entry"
	"github.com/tiptophelmet/mywireguard/pkg/utils"
)

type NewClientAction struct {
	vpn          *entry.VpnEntry
	client       *entry.ClientEntry
	confFilePath string
}

func InitNewClientAction(vpn *entry.VpnEntry, confFilePath string) *NewClientAction {
	fmt.Println("[INFO] Initializing new client action ...")

	client := entry.NewClientEntry()

	// Assign vpn ID
	client.VPNID = vpn.ID

	// Assign client ID from .conf filename
	client.ID = strings.TrimSuffix(filepath.Base(confFilePath), path.Ext(confFilePath))

	// Check if client exists
	_, err := os.Stat(paths.BuildVpnClientFilePath(vpn.ID, client.ID, paths.MkDirAllPath))
	if err == nil {
		log.Fatalf("this VPN client already exists: %s", client.ID)
	}

	return &NewClientAction{vpn, client, confFilePath}
}

func (act *NewClientAction) Prepare() {
	fmt.Println("[INFO] Preparing to set up Wireguard peer ...")

	// Wireguard keys
	var err error
	act.client.WgClientPrivateKey, act.client.WgClientPublicKey, err = utils.GenerateWireguardKeyPair()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("[OK] Wireguard client keys successfully generated!")

	// Create wireguard peer & assign allowed IP
	peerAllowedIP, err := act.vpn.CreateWireguardPeer(act.client.WgClientPublicKey)
	if err != nil {
		log.Fatalf(err.Error())
	}

	act.client.WgClientAllowedIP = peerAllowedIP

	fmt.Println("[OK] Wireguard peer allowed IP:", peerAllowedIP)
}

func (act *NewClientAction) Save() error {
	clientFilePath := paths.BuildVpnClientFilePath(act.vpn.ID, act.client.ID, paths.MkDirAllPath)

	err := utils.WriteBinaryFile(clientFilePath, act.client)
	if err != nil {
		return err
	}

	fmt.Println("[OK] Client entry saved!")

	return nil
}

func (act *NewClientAction) GenerateWireguardClientConf() {
	fmt.Println("[INFO] Preparing to generate Wireguard client .conf ...")

	confFileBytes, err := os.ReadFile("static/wireguard/raw/client.conf")
	if err != nil {
		log.Fatalf(err.Error())
	}

	vpnEntryTags, err := utils.ExtractTagMap("wgclient", act.vpn)
	if err != nil {
		log.Fatalf(err.Error())
	}

	clientEntryTags, err := utils.ExtractTagMap("wgclient", act.client)
	if err != nil {
		log.Fatalf(err.Error())
	}

	values := utils.MergeStrMaps(vpnEntryTags, clientEntryTags)
	if len(values) == 0 {
		log.Fatalf("client conf values are absent")
	}

	composed := utils.StrCompose(string(confFileBytes), values)

	err = os.WriteFile(act.confFilePath, []byte(composed), 0755)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("[OK] Client .conf for VPN %s saved to %s\n", act.vpn.ID, act.confFilePath)
}

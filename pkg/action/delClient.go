package action

import (
	"fmt"
	"log"
	"os"

	"github.com/tiptophelmet/mywireguard/paths"
	"github.com/tiptophelmet/mywireguard/pkg/entry"
)

type DeleteClientAction struct {
	client *entry.ClientEntry
	vpn    *entry.VpnEntry
}

func InitDeleteClientAction(client *entry.ClientEntry, vpn *entry.VpnEntry) *DeleteClientAction {
	return &DeleteClientAction{client, vpn}
}

func (act *DeleteClientAction) Delete() {
	err := act.vpn.DeleteWireguardPeer(act.client.WgClientPublicKey)

	if err != nil {
		log.Fatalf(err.Error())
	}

	vpnClientPath := paths.BuildVpnClientFilePath(act.client.VPNID, act.client.ID, paths.GetPath)

	// Delete VPN client .mywg
	if err := os.Remove(vpnClientPath); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("[OK] Deleted client %s from VPN %s\n", act.client.ID, act.client.VPNID)
}

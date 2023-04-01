package action

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/tiptophelmet/mywireguard/paths"
	"github.com/tiptophelmet/mywireguard/pkg/cloud"
	"github.com/tiptophelmet/mywireguard/pkg/entry"
	"github.com/tiptophelmet/mywireguard/pkg/utils"
)

type GetVpnAction struct {
}

func InitGetVpnAction() *GetVpnAction {
	fmt.Println("[INFO] Initializing get VPN action ...")

	gob.Register(&cloud.DigitalOceanCloud{})

	return &GetVpnAction{}
}

func (act *GetVpnAction) Get(vpnID string) (*entry.VpnEntry, error) {
	vpnFilePath := paths.BuildVpnFilePath(vpnID, paths.GetPath)

	_, err := os.Stat(vpnFilePath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("this VPN does not exist: %s", vpnID)
	}

	fmt.Println("[INFO] Getting VPN entry from", vpnFilePath)

	vpnEntry := entry.NewVpnEntry()
	err = utils.ReadBinaryFile(vpnFilePath, &vpnEntry)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return vpnEntry, nil
}

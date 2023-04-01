package action

import (
	"fmt"
	"log"
	"os"

	"github.com/tiptophelmet/mywireguard/paths"
	"github.com/tiptophelmet/mywireguard/pkg/entry"
	"github.com/tiptophelmet/mywireguard/pkg/utils"
)

type GetClientAction struct {
}

func InitGetClientAction() *GetClientAction {
	fmt.Println("[INFO] Initializing get client action ...")

	return &GetClientAction{}
}

func (act *GetClientAction) Get(vpnID string, clientID string) (*entry.ClientEntry, error) {
	clientFilePath := paths.BuildVpnClientFilePath(vpnID, clientID, paths.GetPath)

	_, err := os.Stat(clientFilePath)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("%s VPN does not have this client: %s", vpnID, clientID)
	}

	fmt.Println("[INFO] Getting client entry from", clientFilePath)

	clientEntry := entry.NewClientEntry()
	err = utils.ReadBinaryFile(clientFilePath, &clientEntry)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return clientEntry, nil
}

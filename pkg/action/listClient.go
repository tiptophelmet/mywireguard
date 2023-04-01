package action

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tiptophelmet/mywireguard/paths"
)

type ListClientAction struct {
}

func InitListClientAction() *ListClientAction {
	return &ListClientAction{}
}

func (act *ListClientAction) List(vpnID string) {
	vpnClientsDirPath := paths.BuildVpnClientsDirPath(vpnID, paths.GetPath)
	_, err := os.Stat(vpnClientsDirPath)
	if os.IsNotExist(err) {
		fmt.Printf("this VPN does not exist: %s\n", vpnID)
		return
	}

	// Open the directory
	entries, err := os.ReadDir(vpnClientsDirPath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if len(entries) == 0 {
		fmt.Printf("No clients found for VPN %s\n", vpnID)
		return
	}

	fmt.Printf("(%d) clients found for VPN %s\n", len(entries), vpnID)

	// Loop through the directory entries and print the names of directories
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) == ".mywg" {
			fmt.Println(strings.TrimSuffix(entry.Name(), ".mywg"))
		}
	}
}

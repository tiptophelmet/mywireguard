package action

import (
	"fmt"
	"log"
	"os"

	"github.com/tiptophelmet/mywireguard/paths"
)

type ListVpnAction struct {
}

func InitListVpnAction() *ListVpnAction {
	return &ListVpnAction{}
}

func (act *ListVpnAction) List() {
	// Open the directory
	dirEntries, err := os.ReadDir(paths.BuildVpnsDirPath(paths.MkDirAllPath))
	if err != nil {
		log.Fatalf(err.Error())
	}

	if len(dirEntries) == 0 {
		log.Fatalf("no deployed VPNs found")
	}

	fmt.Printf("(%d) deployed VPNs found:\n", len(dirEntries))

	// Loop through the directory entries and print the names of directories
	for _, entry := range dirEntries {
		if entry.IsDir() {
			fmt.Println(entry.Name())
		}
	}
}

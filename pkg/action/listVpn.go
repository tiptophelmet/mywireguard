package action

import (
	"fmt"
	"os"

	"github.com/tiptophelmet/mywireguard/paths"
)

type ListVpnAction struct {
}

func InitListVpnAction() *ListVpnAction {
	return &ListVpnAction{}
}

func (act *ListVpnAction) List() error {
	// Open the directory
	dirEntries, err := os.ReadDir(paths.BuildVpnsDirPath(paths.MkDirAllPath))
	if err != nil {
		return err
	}

	if len(dirEntries) == 0 {
		fmt.Println("No deployed VPNs found.")
		return nil
	}

	fmt.Printf("(%d) deployed VPNs found:\n", len(dirEntries))

	// Loop through the directory entries and print the names of directories
	for _, entry := range dirEntries {
		if entry.IsDir() {
			fmt.Println(entry.Name())
		}
	}

	return nil
}

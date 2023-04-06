package action

import (
	"fmt"
	"log"
	"os"

	"github.com/tiptophelmet/mywireguard/paths"
	"github.com/tiptophelmet/mywireguard/pkg/entry"
	"github.com/tiptophelmet/mywireguard/pkg/infra"
)

type DeleteVpnAction struct {
	entry *entry.VpnEntry
}

func InitDeleteVpnAction(entry *entry.VpnEntry) *DeleteVpnAction {
	fmt.Println("[INFO] Initializing delete VPN action ...")

	// Check if clients are present
	entries, err := os.ReadDir(paths.BuildVpnClientsDirPath(entry.ID, paths.GetPath))
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf(err.Error())
	}

	if len(entries) != 0 {
		log.Fatalf("failed to delete %s VPN with existing clients", entry.ID)
	}

	fmt.Println("[INFO] No clients found, proceeding ...")

	return &DeleteVpnAction{entry}
}

func (act *DeleteVpnAction) DestroyInfra() {
	fmt.Println("[INFO] Preparing to destroy terraform ...")

	_, executor, err := infra.From(act.entry.Cloud)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = executor.Destroy(paths.GetTerraformDirPath(act.entry.ID, paths.GetPath))
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("[OK] Terraform destroyed!")
}

// on delete - terraform destroy
func (act *DeleteVpnAction) Delete() {
	// Delete associated SSH keys
	publicKeyPath := paths.GetSshKeyFilePath(fmt.Sprintf("%s.pub", act.entry.ID), paths.GetPath)
	if err := os.Remove(publicKeyPath); err != nil {
		log.Fatalf(err.Error())
	}

	privateKeyPath := paths.GetSshKeyFilePath(act.entry.ID, paths.GetPath)
	if err := os.Remove(privateKeyPath); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("[OK] SSH keys deleted!")

	// Delete Terraform
	if err := os.RemoveAll(paths.GetTerraformDirPath(act.entry.ID, paths.GetPath)); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("[OK] Terraform files deleted!")

	// Delete VPN .mywg
	if err := os.RemoveAll(paths.BuildVpnDirPath(act.entry.ID, paths.GetPath)); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("[OK] Deleted VPN entry!")

	fmt.Printf("[OK] Deleted VPN with ID: %s\n", act.entry.ID)
}

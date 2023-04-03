package action

import (
	"fmt"
	"log"
	"os"

	"github.com/tiptophelmet/mywireguard/paths"
	"github.com/tiptophelmet/mywireguard/pkg/cloud"
	"github.com/tiptophelmet/mywireguard/pkg/entry"
	"github.com/tiptophelmet/mywireguard/pkg/infra"
	"github.com/tiptophelmet/mywireguard/pkg/utils"
)

type NewVpnAction struct {
	entry *entry.VpnEntry
}

func InitNewVpnAction(ID string, cloud cloud.Cloud) *NewVpnAction {
	_, err := os.Stat(paths.BuildVpnFilePath(ID, paths.GetPath))
	if err == nil {
		log.Fatalf("this VPN already exists: %s", ID)
	}

	entry := entry.NewVpnEntry()

	entry.ID = ID
	entry.Cloud = cloud

	fmt.Println("[INFO] Initializing new VPN action ...")

	return &NewVpnAction{entry}
}

func (act *NewVpnAction) Prepare() {
	fmt.Println("[INFO] Preparing to deploy Wireguard ...")

	// Wireguard keys
	var err error
	act.entry.WgServerPrivateKey, act.entry.WgServerPublicKey, err = utils.GenerateWireguardKeyPair()
	if err != nil {
		log.Fatalf(err.Error())
	}

	act.entry.SshPrivateKeyPath = paths.GetSshKeyFilePath(act.entry.ID, paths.MkDirAllPath)
	act.entry.SshPublicKeyPath = paths.GetSshKeyFilePath(fmt.Sprintf("%s.pub", act.entry.ID), paths.MkDirAllPath)

	fmt.Println("[OK] Wireguard server keys successfully generated!")

	// SSH keys
	err = utils.GenerateSshKeyPair(act.entry.SshPrivateKeyPath, act.entry.SshPublicKeyPath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("[OK] SSH keys successfully generated & saved!")

	// Wireguard interface address
	act.entry.WgServerInterfaceAddress = "10.0.0.1/24"

	fmt.Println("[INFO] Wireguard interface address: ", act.entry.WgServerInterfaceAddress)

	// Wireguard interface listen port
	act.entry.WgServerInterfaceListenPort = "51820"

	fmt.Println("[INFO] Wireguard interface listen port: ", act.entry.WgServerInterfaceListenPort)
}

func (act *NewVpnAction) ApplyInfra() {
	fmt.Println("[INFO] Preparing terraform ...")

	// Create VPN terraform
	composer, executor, err := infra.From(act.entry.Cloud)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := composer.LoadTemplates(); err != nil {
		log.Fatalf(err.Error())
	}

	composer.Compose(act.entry)
	composer.Save()

	act.entry.WgServerPublicIP, err = executor.Apply(paths.GetTerraformDirPath(act.entry.ID, paths.GetPath))
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("[OK] Terraform finished!")
}

func (act *NewVpnAction) Save() error {
	path := paths.BuildVpnFilePath(act.entry.ID, paths.MkDirAllPath)

	err := utils.WriteBinaryFile(path, act.entry)
	if err != nil {
		return err
	}

	fmt.Println("[OK] VPN entry saved!")

	return nil
}

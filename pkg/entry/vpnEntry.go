package entry

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/tiptophelmet/mywireguard/pkg/cloud"
	"github.com/tiptophelmet/mywireguard/pkg/utils"
)

type VpnEntry struct {
	ID string

	Cloud cloud.Cloud

	WgServerPrivateKey          string `terraform:"wireguard_server_private_key"`
	WgServerInterfaceAddress    string `terraform:"wireguard_interface_address"`
	WgServerInterfaceListenPort string `terraform:"wireguard_interface_listen_port"`

	WgServerPublicKey string `wgclient:"wireguard_server_public_key"`

	SshPrivateKeyPath string `terraform:"digitalocean_ssh_private_key_path"`
	SshPublicKeyPath  string `terraform:"digitalocean_ssh_public_key_path"`

	WgServerPublicIP string `wgclient:"wireguard_server_public_ip"`
}

func NewVpnEntry() *VpnEntry {
	entry := &VpnEntry{}
	gob.Register(entry)
	return entry
}

func (vpn *VpnEntry) executeCommand(cmdStr string) string {
	sshDestination := fmt.Sprintf("root@%s", vpn.WgServerPublicIP)

	cmd := exec.Command("ssh", "-t", "-i", vpn.SshPrivateKeyPath, "-o", "StrictHostKeyChecking=no", sshDestination, cmdStr)

	var out bytes.Buffer
	var stdErr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		log.Fatalf(stdErr.String())
	}

	return out.String()
}

func (vpn *VpnEntry) getAllocatedPeerIPs() ([]string, error) {
	// Run the "wg" command to get the list of peers and their allowed IPs
	output := vpn.executeCommand("wg show wg0 | grep 'allowed ips' | awk -F': ' '{print $2}'")

	if strings.TrimSpace(output) == "" {
		return []string{}, nil
	}

	output = strings.ReplaceAll(output, "\r\n", "\n")

	// Parse the allowed ips and return them as a slice of strings
	return strings.Split(strings.TrimSpace(output), "\n"), nil
}

func (vpn *VpnEntry) CreateWireguardPeer(clientPublicKey string) (allowedIP string) {
	// Get the list of already allocated IPs
	allocatedIPs, err := vpn.getAllocatedPeerIPs()
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	if len(allocatedIPs) == 0 {
		fmt.Println("[INFO] No Wireguard peer allocated IPs so far")
		allowedIP = "10.0.0.2/32"
	} else {
		fmt.Println("[OK] Retrieved Wireguard peer allocated IPs")

		// Find the next available IP
		allowedIP, err = utils.AssumeWireguardPeerIP(vpn.WgServerInterfaceAddress, allocatedIPs)
		if err != nil {
			log.Fatalf(err.Error())
			return
		}
	}

	fmt.Println("[INFO] Assumed new Wireguard peer allowed IP")

	// Build the command to add the new peer
	setWgPeerCmd := fmt.Sprintf("wg set wg0 peer %s allowed-ips %s persistent-keepalive 25", strings.TrimSpace(clientPublicKey), allowedIP)

	// Run the command to add the new peer
	vpn.executeCommand(setWgPeerCmd)

	fmt.Println("[OK] Wireguard peer successfully created!")

	return
}

func (vpn *VpnEntry) DeleteWireguardPeer(clientPublicKey string) error {
	RemoveWgPeerCmd := fmt.Sprintf("wg set wg0 peer %s remove", clientPublicKey)

	// Run the command to remove the peer
	vpn.executeCommand(RemoveWgPeerCmd)

	return nil
}

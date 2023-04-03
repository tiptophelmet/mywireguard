package entry

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
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

func (vpn *VpnEntry) executeCommand(cmdStr string) (string, error) {
	sshDestination := fmt.Sprintf("root@%s", vpn.WgServerPublicIP)

	cmd := exec.Command("ssh", "-t", "-i", vpn.SshPrivateKeyPath, "-o", "StrictHostKeyChecking=no", sshDestination, cmdStr)

	var out bytes.Buffer
	var stdErr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		return "", errors.New(stdErr.String())
	}

	output := strings.ReplaceAll(out.String(), "\r\n", "\n")

	return output, nil
}

func (vpn *VpnEntry) getAllocatedPeerIPs() ([]string, error) {
	// Run the "wg" command to get the list of peers and their allowed IPs
	output, err := vpn.executeCommand("wg show wg0 | grep 'allowed ips' | awk -F': ' '{print $2}'")
	if err != nil {
		return []string{}, err
	}

	if strings.TrimSpace(output) == "" {
		return []string{}, nil
	}

	// Parse the allowed ips and return them as a slice of strings
	return strings.Split(strings.TrimSpace(output), "\n"), nil
}

func (vpn *VpnEntry) CreateWireguardPeer(clientPublicKey string) (string, error) {
	// Get the list of already allocated IPs
	allocatedIPs, err := vpn.getAllocatedPeerIPs()
	if err != nil {
		return "", err
	}

	var allowedIP string

	if len(allocatedIPs) == 0 {
		fmt.Println("[INFO] No Wireguard peer allocated IPs so far")
		allowedIP = "10.0.0.2/32"
	} else {
		fmt.Println("[OK] Retrieved Wireguard peer allocated IPs")

		// Find the next available IP
		allowedIP, err = utils.AssumeWireguardPeerIP(vpn.WgServerInterfaceAddress, allocatedIPs)
		if err != nil {
			return "", err
		}
	}

	fmt.Println("[INFO] Assumed new Wireguard peer allowed IP")

	// Build the command to add the new peer
	setWgPeerCmd := fmt.Sprintf("wg set wg0 peer %s allowed-ips %s persistent-keepalive 25", strings.TrimSpace(clientPublicKey), allowedIP)

	// Run the command to add the new peer
	_, err = vpn.executeCommand(setWgPeerCmd)
	if err != nil {
		return "", err
	}

	fmt.Println("[OK] Wireguard peer successfully created!")

	return allowedIP, nil
}

func (vpn *VpnEntry) DeleteWireguardPeer(clientPublicKey string) error {
	RemoveWgPeerCmd := fmt.Sprintf("wg set wg0 peer %s remove", clientPublicKey)

	// Run the command to remove the peer
	_, err := vpn.executeCommand(RemoveWgPeerCmd)
	if err != nil {
		return err
	}

	return nil
}

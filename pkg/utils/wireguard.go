package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
	"os/exec"
	"strings"
)

func GenerateWireguardKeyPair() (privateKey, publicKey string, err error) {
	genkey := exec.Command("wg", "genkey")
	genkeyOutput := &bytes.Buffer{}
	genkey.Stdout = genkeyOutput
	err = genkey.Run()
	if err != nil {
		return "", "", errors.New("failed to generate private key")
	}

	privateKey = strings.TrimSpace(genkeyOutput.String())

	pubkey := exec.Command("wg", "pubkey")
	pubkey.Stdin = genkeyOutput
	pubkeyOutput := &bytes.Buffer{}
	pubkey.Stdout = pubkeyOutput
	err = pubkey.Run()
	if err != nil {
		return "", "", errors.New("failed to generate public key")
	}

	publicKey = strings.TrimSpace(pubkeyOutput.String())

	return
}

func nextIP(ip *net.IP) {
	ipInt := binary.BigEndian.Uint32((*ip).To4())
	ipInt++
	binary.BigEndian.PutUint32((*ip).To4(), ipInt)
}

func setStartIP(ip *net.IP, ipnet *net.IPNet) net.IP {
	startIP := ip.Mask(ipnet.Mask)
	nextIP(&startIP)
	nextIP(&startIP)
	return startIP
}

func AssumeWireguardPeerIP(wgInterfaceAddress string, allocated []string) (string, error) {
	ip, ipnet, err := net.ParseCIDR(wgInterfaceAddress)
	if err != nil {
		return "", err
	}

	// Set the initial IP address to skip the network and gateway IPs
	ip = setStartIP(&ip, ipnet)

	for ; ipnet.Contains(ip); nextIP(&ip) {
		ipWithMask := net.IPNet{IP: ip, Mask: net.CIDRMask(32, 32)}
		if !IsStrInSlice(allocated, ipWithMask.String()) {
			return ipWithMask.String(), nil
		}
	}

	return "", errors.New("all peer IPs are occupied")
}

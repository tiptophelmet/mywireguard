package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func GenerateSshKeyPair(sshPrivateKeyPath, sshPublicKeyPath string) error {
	// Generate a new RSA private key with 2048 bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// Convert the private key to a PEM-encoded string
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	privateKeyFile, err := os.Create(sshPrivateKeyPath)
	if err != nil {
		return fmt.Errorf("error creating private key file: %v", err)
	}

	err = pem.Encode(privateKeyFile, privateKeyBlock)
	if err != nil {
		return fmt.Errorf("error encoding private key: %v", err)
	}
	privateKeyFile.Close()

	// Create a new public key from the private key
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}

	// Convert the public key to a string
	publicKeyBytes := ssh.MarshalAuthorizedKey(publicKey)

	err = os.WriteFile(sshPublicKeyPath, publicKeyBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

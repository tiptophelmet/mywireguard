package paths

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func BuildVpnClientFilePath(vpnID string, clientID string, mode PathBuildMode) string {
	filename := fmt.Sprintf("%s.mywg", clientID)
	path := filepath.Join(getStorageDir(), "vpn", vpnID, "clients", filename)

	if mode == MkDirAllPath {
		// Create the directory path if it does not already exist
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			log.Fatalf("error creating directory: %s\n", err)
		}
	}

	return path
}

func BuildVpnClientsDirPath(vpnID string, mode PathBuildMode) string {
	path := filepath.Join(getStorageDir(), "vpn", vpnID, "clients")

	if mode == MkDirAllPath {
		// Create the directory path if it does not already exist
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			log.Fatalf("error creating directory: %s\n", err)
		}
	}

	return path
}

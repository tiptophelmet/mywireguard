package paths

import (
	"log"
	"os"
	"path/filepath"
)

func BuildVpnFilePath(vpnID string, mode PathBuildMode) string {
	path := filepath.Join(getStorageDir(), "vpn", vpnID, "vpn.mywg")

	if mode == MkDirAllPath {
		// Create the directory path if it does not already exist
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			log.Fatalf("error creating directory: %s\n", err)
		}
	}

	return path
}

func BuildVpnDirPath(vpnID string, mode PathBuildMode) string {
	path := filepath.Join(getStorageDir(), "vpn", vpnID)

	if mode == MkDirAllPath {
		// Create the directory path if it does not already exist
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			log.Fatalf("error creating directory: %s\n", err)
		}
	}

	return path
}

func BuildVpnsDirPath(mode PathBuildMode) string {
	path := filepath.Join(getStorageDir(), "vpn")

	if mode == MkDirAllPath {
		// Create the directory path if it does not already exist
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			log.Fatalf("error creating directory: %s\n", err)
		}
	}

	return path
}

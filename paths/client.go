package paths

import (
	"fmt"
	"path/filepath"
)

func BuildVpnClientFilePath(vpnID string, clientID string, mode PathBuildMode) string {
	filename := fmt.Sprintf("%s.mywg", clientID)
	path := filepath.Join(getStorageDir(), "vpn", vpnID, "clients", filename)
	
	if mode == MkDirAllPath {
		mkDirAllPath(path, 0755)
	}

	return path
}

func BuildVpnClientsDirPath(vpnID string, mode PathBuildMode) string {
	path := filepath.Join(getStorageDir(), "vpn", vpnID, "clients")
	if mode == MkDirAllPath {
		mkDirAllPath(path, 0755)
	}

	return path
}

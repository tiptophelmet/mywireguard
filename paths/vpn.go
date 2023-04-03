package paths

import (
	"path/filepath"
)

func BuildVpnFilePath(vpnID string, mode PathBuildMode) string {
	path := filepath.Join(getStorageDir(), "vpn", vpnID, "vpn.mywg")
	if mode == MkDirAllPath {
		mkDirAllPath(path, 0755)
	}

	return path
}

func BuildVpnDirPath(vpnID string, mode PathBuildMode) string {
	path := filepath.Join(getStorageDir(), "vpn", vpnID)
	if mode == MkDirAllPath {
		mkDirAllPath(path, 0755)
	}

	return path
}

func BuildVpnsDirPath(mode PathBuildMode) string {
	path := filepath.Join(getStorageDir(), "vpn")
	if mode == MkDirAllPath {
		mkDirAllPath(path, 0755)
	}

	return path
}

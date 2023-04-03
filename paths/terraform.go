package paths

import (
	"path/filepath"
)

func GetTerraformDirPath(vpnID string, mode PathBuildMode) string {
	path := filepath.Join(getStorageDir(), "terraform", vpnID)
	if mode == MkDirAllPath {
		mkDirAllPath(path, 0755)
	}

	return path
}

package paths

import (
	"path/filepath"
)

func GetSshKeyFilePath(keyFilename string, mode PathBuildMode) string {
	path := filepath.ToSlash(filepath.Join(getStorageDir(), "ssh", keyFilename))
	if mode == MkDirAllPath {
		mkDirAllPath(path, 0755)
	}

	return path
}

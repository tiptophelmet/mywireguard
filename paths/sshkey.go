package paths

import (
	"log"
	"os"
	"path/filepath"
)

func GetSshKeyFilePath(keyFilename string, mode PathBuildMode) string {
	path := filepath.ToSlash(filepath.Join(getStorageDir(), "ssh", keyFilename))

	if mode == MkDirAllPath {
		// Create the directory path if it does not already exist
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			log.Fatalf("error creating directory: %s\n", err)
		}
	}

	return path
}

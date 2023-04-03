package paths

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

type PathBuildMode int

const (
	GetPath PathBuildMode = iota
	MkDirAllPath
)

func getStorageDir() string {
	homeDir, err := homedir.Dir()
	if err != nil {
		log.Fatalf("failed to find home directory: %v", err)
	}

	storageDir := filepath.Join(homeDir, ".mywireguard")
	if _, err := os.Stat(storageDir); os.IsNotExist(err) {
		err := os.Mkdir(storageDir, 0755)
		if err != nil {
			log.Fatalf("failed to create storage directory (%s): %v", storageDir, err)
		}
	}

	return storageDir
}

func mkDirAllPath(path string, perm fs.FileMode) {
	err := os.MkdirAll(filepath.Dir(path), perm)
	if err != nil {
		log.Fatalf("failed to create path: %s (path: %s)", err.Error(), path)
	}
}

package paths

import (
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
	// Create storage directory if it doesn't exist
	homeDir, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Error: cannot find home directory: %v\n", err)
	}

	storageDir := filepath.Join(homeDir, ".mywireguard")
	if _, err := os.Stat(storageDir); os.IsNotExist(err) {
		err := os.Mkdir(storageDir, 0755)
		if err != nil {
			log.Fatalf("Error: cannot create storage directory (%s): %v\n", storageDir, err)
		}
	}

	return storageDir
}

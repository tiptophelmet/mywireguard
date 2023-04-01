package cloud

import (
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Instance interface {
	GetImage() string
	GetName() string
	GetRegion() string
	GetSize() string
}

type Cloud interface {
	GetInstance() Instance
	GetApiToken() string
}

type CloudConfig struct {
	tomlData []byte
}

func NewCloudConfig() *CloudConfig {
	return &CloudConfig{}
}

func (cc *CloudConfig) ImportToml(tomlPath string) {
	// Read the TOML file
	fmt.Println("[INFO] Reading .toml ...")

	tomlData, err := os.ReadFile(tomlPath)
	if err != nil {
		log.Fatalf("Error reading TOML file: %s", err)
	}
	cc.tomlData = tomlData
}

func (cc *CloudConfig) InitCloud() (Cloud, error) {
	// Decode the TOML data into a map
	var tomlMap map[string]interface{}
	if _, err := toml.Decode(string(cc.tomlData), &tomlMap); err != nil {
		log.Fatalf("Error decoding .toml: %s", err)
	}

	if _, ok := tomlMap["provider"]; !ok {
		log.Fatalf(".toml does not have a provider field")
	}

	switch tomlMap["provider"] {
	case "digitalocean":
		var digitalOceanCloud DigitalOceanCloud
		gob.Register(&DigitalOceanCloud{})

		// Decode the TOML data directly into the DigitalOceanCloud struct
		if _, err := toml.Decode(string(cc.tomlData), &digitalOceanCloud); err != nil {
			log.Fatalf("Error decoding TOML: %s", err)
		}

		fmt.Println("[INFO] DigitalOcean cloud provider identified!")

		return &digitalOceanCloud, nil
	default:
		return nil, errors.New("no supported cloud provider found in the .toml")
	}
}

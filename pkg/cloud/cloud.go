package cloud

import (
	"encoding/gob"
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

func (cc *CloudConfig) ImportToml(tomlPath string) error {
	// Read the TOML file
	fmt.Println("[INFO] Reading .toml ...")

	tomlData, err := os.ReadFile(tomlPath)
	if err != nil {
		return fmt.Errorf("error reading TOML file: %s", err)
	}
	cc.tomlData = tomlData

	return nil
}

func (cc *CloudConfig) InitCloud() Cloud {
	// Decode the TOML data into a map
	var tomlMap map[string]interface{}
	if _, err := toml.Decode(string(cc.tomlData), &tomlMap); err != nil {
		log.Fatalf("error decoding .toml: %s", err)
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
			log.Fatalf("error decoding TOML: %s", err)
		}

		fmt.Println("[INFO] DigitalOcean cloud provider identified!")

		return &digitalOceanCloud
	default:
		log.Fatalf("no supported cloud provider found in the .toml")
	}

	return nil
}

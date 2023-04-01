package infra

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tiptophelmet/mywireguard/paths"
	"github.com/tiptophelmet/mywireguard/pkg/entry"
	"github.com/tiptophelmet/mywireguard/pkg/utils"
)

type InfraComposer struct {
	templates map[string]string
	values    map[string]string
	infraPath string
}

func InitInfraComposer() *InfraComposer {
	return &InfraComposer{
		templates: map[string]string{},
		values:    map[string]string{},
	}
}

func (ic *InfraComposer) addTemplate(rawPath string) {
	fileBytes, err := os.ReadFile(rawPath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	baseName := filepath.Base(rawPath)
	ic.templates[baseName] = string(fileBytes)
}

func (ic *InfraComposer) LoadTemplates() error {
	ic.addTemplate("pkg/infra/digitalocean/terraform/raw/main.tf")
	ic.addTemplate("pkg/infra/digitalocean/terraform/raw/wireguard_setup.sh")

	fmt.Println("[OK] Terraform templates loaded!")

	return nil
}

func (ic *InfraComposer) Compose(vpnEntry *entry.VpnEntry) {
	ic.infraPath = paths.GetTerraformDirPath(vpnEntry.ID, paths.MkDirAllPath)
	ic.values = utils.ExtractTagMap("terraform", vpnEntry)

	if len(ic.values) == 0 {
		log.Fatalf("infra values are absent")
	}

	for baseName, template := range ic.templates {
		ic.templates[baseName] = utils.StrCompose(template, ic.values)
	}

	fmt.Println("[OK] Terraform composed!")
}

func (ic *InfraComposer) Save() {
	for baseName, composed := range ic.templates {
		path := filepath.Join(ic.infraPath, baseName)

		// Create the directory path if it does not already exist
		dirPath := filepath.Dir(path)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			log.Fatalf("error creating directory: %s\n", err)
		}

		err := os.WriteFile(path, []byte(composed), 0644)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	fmt.Println("[OK] Terraform saved!")
}

package infra

import (
	"errors"
	"fmt"
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

func (ic *InfraComposer) addTemplate(rawPath string) error {
	fileBytes, err := os.ReadFile(rawPath)
	if err != nil {
		return err
	}

	baseName := filepath.Base(rawPath)
	ic.templates[baseName] = string(fileBytes)

	return nil
}

func (ic *InfraComposer) LoadTemplates() error {
	err := ic.addTemplate("pkg/infra/digitalocean/terraform/raw/main.tf")
	if err != nil {
		return err
	}

	err = ic.addTemplate("pkg/infra/digitalocean/terraform/raw/wireguard_setup.sh")
	if err != nil {
		return err
	}

	fmt.Println("[OK] Terraform templates loaded!")

	return nil
}

func (ic *InfraComposer) Compose(vpnEntry *entry.VpnEntry) error {
	ic.infraPath = paths.GetTerraformDirPath(vpnEntry.ID, paths.MkDirAllPath)

	var err error
	ic.values, err = utils.ExtractTagMap("terraform", vpnEntry)

	if err != nil {
		return err
	}

	if len(ic.values) == 0 {
		return errors.New("infra values are absent")
	}

	for baseName, template := range ic.templates {
		ic.templates[baseName] = utils.StrCompose(template, ic.values)
	}

	fmt.Println("[OK] Terraform composed!")
	return nil
}

func (ic *InfraComposer) Save() error {
	for baseName, composed := range ic.templates {
		path := filepath.Join(ic.infraPath, baseName)

		dirPath := filepath.Dir(path)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create dir: %s", err)
		}

		err := os.WriteFile(path, []byte(composed), 0644)
		if err != nil {
			return err
		}
	}

	fmt.Println("[OK] Terraform saved!")
	return nil
}

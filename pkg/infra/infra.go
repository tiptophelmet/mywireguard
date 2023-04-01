package infra

import (
	"log"
	"reflect"

	"github.com/tiptophelmet/mywireguard/pkg/cloud"
	"github.com/tiptophelmet/mywireguard/pkg/entry"
	digitalocean "github.com/tiptophelmet/mywireguard/pkg/infra/digitalocean"
)

type InfraExecutor interface {
	Apply(execPath string) (publicIP string, err error)
	Destroy(execPath string) error
}

type InfraComposer interface {
	LoadTemplates() error
	Compose(*entry.VpnEntry)
	Save()
}

func From(cl cloud.Cloud) (InfraComposer, InfraExecutor) {
	switch cl.(type) {
	case *cloud.DigitalOceanCloud:
		return digitalocean.InitInfraComposer(), digitalocean.InitInfraExecutor()
	}

	log.Fatalf("no supported infra found for cloud %s", reflect.TypeOf(cl).Name())
	return nil, nil
}

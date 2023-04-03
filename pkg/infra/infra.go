package infra

import (
	"fmt"
	"reflect"

	"github.com/tiptophelmet/mywireguard/pkg/cloud"
	"github.com/tiptophelmet/mywireguard/pkg/entry"
	digitalocean "github.com/tiptophelmet/mywireguard/pkg/infra/digitalocean"
)

type InfraExecutor interface {
	Apply(execPath string) (vpnPublicIP string, err error)
	Destroy(execPath string) error
}

type InfraComposer interface {
	LoadTemplates() error
	Compose(*entry.VpnEntry) error
	Save() error
}

func From(cl cloud.Cloud) (InfraComposer, InfraExecutor, error) {
	switch cl.(type) {
	case *cloud.DigitalOceanCloud:
		return digitalocean.InitInfraComposer(), digitalocean.InitInfraExecutor(), nil
	}

	return nil, nil, fmt.Errorf("no supported infra found for cloud %s", reflect.TypeOf(cl).Name())
}

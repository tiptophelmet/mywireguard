package entry

import "encoding/gob"

type ClientEntry struct {
	ID    string `terraform:"wireguard_client_id"`
	VPNID string

	WgClientPrivateKey string `wgclient:"wireguard_client_private_key"`
	WgClientPublicKey  string

	WgClientAllowedIP string `wgclient:"wireguard_client_allowed_ip"`
}

func NewClientEntry() *ClientEntry {
	entry := &ClientEntry{}
	gob.Register(entry)
	return entry
}

# __VPN Protocols Support__

## Supported:
- Wireguard
## Not supported:
- OpenVPN
- IPSEC
- IKEAv2

# __Cloud Support__

## Supported:
- DigitalOcean
## Not yet supported:
- AWS
- GCP

# __VPN clients limitations__
Currently, all connected clients have 1 shared static IP.
Dedicated IPs support is planned.

Please, keep in mind that option with 1 shared static IP is the cheapest infra cost option.
Dedicated IPs require much more network infra resources, leading to higher cloud costs.
# __Dependencies__
Dependencies are not yet included into the app! Please make sure they are installed on your local machine.
- Terraform (https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli)
- Wireguard (https://www.wireguard.com/install/)

# __CLI design__

##  __New VPN__

```
mywg new-vpn --toml=~/path/to/toml/team-vpn.toml
```
VPN .toml example:
```
provider = "digitalocean"
api_token = "dop_v1_...*redacted*..."

[droplet]
image = "ubuntu-20-04-x64"
name = "team-vpn"
region = "fra1"
size = "s-1vcpu-1gb"
```

## __New VPN client__


```
mywg new-client --vpnid=team-vpn --conf=/output/path/for/vpn-client-james.conf
```
Generated Wireguard VPN client .conf example:
```
[Interface]
PrivateKey = g7IcPqp...*redacted*...BR3w93a
Address = 10.0.0.2/32
DNS = 1.1.1.1

[Peer]
PublicKey = sDoskH7S...*redacted*...Ze+PX9e
Endpoint = <VPN PUBLIC IP>:51820
AllowedIPs = 0.0.0.0/0, ::/0
PersistentKeepalive = 25
```

## __List VPNs__

```
mywg list-vpn
```
## __List VPN clients__

```
mywg list-client --vpnid=team-vpn
```

## __Delete VPN client__

```
mywg del-client --clientid=vpn-client-james --vpnid=team-vpn 
```

## __Delete VPN__

```
mywg del-vpn --vpnid=team-vpn
```

# __Roadmap__
- Firewall for DigitalOcean-hosted VPN
- Better error handling
- Tests
- Docs
- AWS Support
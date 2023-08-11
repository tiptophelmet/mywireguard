# 🔒 __MyWireguard__


👩‍💻 Deploy a Wireguard VPN to your cloud provider in minutes and create VPN clients within seconds - 😲 right from your console/terminal.

With our app, you can easily set up a secure VPN connection between your devices and access resources on your cloud network as if you were physically present there. 🔒

💻 If you're not familiar with Wireguard, it's a modern VPN protocol that is designed to be fast, simple, and secure. Wireguard uses state-of-the-art cryptography to ensure the confidentiality, integrity, and authenticity of data transmitted over the VPN. Compared to other VPN protocols, such as OpenVPN and IPSec, Wireguard is faster, more lightweight, and easier to set up and manage. 🚀

💰 By using MyWireguard CLI app, you can save valuable time and effort that would otherwise be spent on setting up the VPN manually. Our app provides a simple and automated solution that is perfect for small teams looking for their own cloud-hosted low-cost and cheap VPN. 

🔒💰 With a secure Wireguard VPN protocol & this CLI app in pair - you get a strong balance between cost & security.

&nbsp;

# 📌 __Legend__

Most workers dealing with sensitive firewall-protected web resources (like AWS/GCP VPCs) need a static IP.

It's great, if you, as a worker, is provided with a company's VPN & a Static IP. Otherwise, surprisingly, you need to deal with it yourself.

You can find numerous VPN providers offering Static IP as well. 

🤔🌐 Such solutions will match your needs, but keep in mind these concerns: 
- You pay twice: for a VPN access & for a Static IP.
- You may be unsure of how your business/product-sensitive data flows through VPN providers
- You might want to consider hosting your own VPN to be 100% sure your business/product-sensitive data is treated properly

🙂🛡️ MyWireguard CLI app exactly targets such concerns & provides the following countermeasures:
- With Wireguard VPN deployed in a chosen cloud - you always have a Static IP, as all cloud compute instances always have an attached Static IP
- Your cloud. Your cloud-hosted VPN. Totally your VPN server. You control it 100%.
- MyWireguard CLI app is fully open-source & does not include any compiled binaries, executables (i.e. no spyware, no trackers, etc.) All required dependencies are installed separately from official websites. 
All 100% of the code is open to you.

&nbsp;

# ⭐ __Features__
- ☁️ Choose your cloud & region
- 🪄 Deploy Wireguard VPN with a single command
- 💁‍♀️ Provision VPN clients within seconds
- ⚡ Automatically generate VPN client connection .conf(s)
- ⛔ Immediately revoke VPN access from your clients
- 🔥 Destroy your VPN within minutes
- 🎮 Use features via single user-friendly commands
- 🆗 Track command execution statuses

&nbsp;

# 👩‍💻 __How to use__

## ➕ __New VPN__

```
mywg new-vpn --toml=~/path/to/toml/team-vpn.toml
```
VPN .toml example:
```
provider = "digitalocean"
api_token = "dop_v1_...*redacted*..."

[droplet]
image = "ubuntu-20-04-x64"
region = "fra1"
size = "s-1vcpu-1gb"
```

## ➕ __New VPN client__

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

## 📋 __List VPNs__

```
mywg list-vpn
```

## 📋 __List VPN clients__

```
mywg list-client --vpnid=team-vpn
```

## ❌ __Delete VPN client__

```
mywg del-client --clientid=vpn-client-james --vpnid=team-vpn 
```

## ❌ __Delete VPN__

```
mywg del-vpn --vpnid=team-vpn
```

&nbsp;

# ⚠️ __VPN clients limitations__
Currently, all connected clients have 1 shared static IP. Dedicated IPs support is planned.

Please, keep in mind that option with 1 shared static IP is the cheapest infra cost option.
Dedicated IPs require much more network infra resources, leading to higher cloud costs.

&nbsp;

# 📦 __Dependencies__
Dependencies are not yet included into the app! Please make sure they are installed on your local machine.
- Go (v1.20 or higher)
Download: https://go.dev/dl/
- Terraform (v1.4.2 or higher) 
Download: https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli
- Wireguard (wireguard-tools v1.0.20210914 or higher) 
Download: https://www.wireguard.com/install/

&nbsp;

# ☁️ __Cloud Support__

✅ Supported:
- DigitalOcean

❌ Not yet supported, 📅 but planned:
- AWS
- Microsoft Azure
- Google Cloud

&nbsp;

# 📊 __Cloud Suitability for Dedicated IPs Feature__

A quick research made to help you out with choosing a cloud, if you decide to focus on  Dedicated IPs feature.

Despite the different cost effectiveness of each cloud for Dedicated IPs feature, it is still planned to cover most of the clouds.

| ☁️ | IPs per instance | Cost effectiveness | Dedicated IPs Feature Status |
|--------|--------|:--------:|:--------:|
| AWS | 2+ dedicated IPs possible per instance, see quotas below | ✅ OK | 📅 planned |
| Microsoft Azure | 2+ dedicated IPs possible per instance, see quotas below | ✅🤔 OK, but pricier than AWS | 📅 planned |
| Google Cloud | 2+ dedicated IPs possible per instance, see quotas below | ✅🤔 OK, but pricier than AWS & low initial quotas | 📅 planned |
| DigitalOcean | 1 dedicated IP per instance | ⚠️ costly | ❓ under consideration |

## __Links__: 
- ### __AWS__
    Quotas for network interfaces per region:

    https://us-east-2.console.aws.amazon.com/servicequotas/home/services/vpc/quotas (search for "Network interfaces per Region")

    Quotas for network interfaces per instance:
    
    https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI 
    
    <u>1 network interface == 1 dedicated Elastic IP.</u> 

    __Example__: EC2 t3.medium instance is allowed to have 3 network interfaces at max. So 3 network interfaces == 3 dedicated IPs for t3.medium instance.
    
    Use https://calculator.aws/ to calculate your approx. costs

- ### __Microsoft Azure__

    Quotas for network interfaces per VM:

    https://learn.microsoft.com/en-us/azure/virtual-machines/sizes

    <u>1 network interface == 1 dedicated Static IP.</u>

    __Example__: Standard_B1ms VM is allowed to have 2 network interfaces at max. 
    
    https://learn.microsoft.com/en-us/azure/virtual-machines/sizes-b-series-burstable
    
    So 2 network interfaces == 2 dedicated IPs for Standard_B1ms VM.

    Use https://azure.microsoft.com/en-us/pricing/calculator/ to calculate your approx. costs

- ### __Google Cloud__

    Maximum possible number of network interfaces:
    https://cloud.google.com/vpc/docs/quota#per_instance (see "Maximum number of network interfaces" in quotas table)

    ☝️ Pay attention that "Maximum number of network interfaces" is the highest possible number, meaning that specific VMs will have their own network interfaces limits

    Quotas for network interfaces per project:

    https://console.cloud.google.com/iam-admin/quotas (then select your project & see "Networks" in quotas table)

    ⚠️ __Warnings:__ 
    - Google Cloud Static IPs cannot be attached to network interfaces, only 1 Static IP can be attached directly to your VM
    - Having Network Interfaces attached to your VM, you get Ephemeral Public IPs pointing to your VM, not Static IPs
    - Ephemeral IPs might be reset with each VM start/stop, VM reboot does not change Ephemeral IPs
    - <b><u>Initial networks per project quota is 5</u></b> meaning no more than 5 public IPs in total for all VMs, consider requesting an increase

    &nbsp;

    🎯 For an accurate number of network interfaces per instance, refer to:

    VM types:
    https://cloud.google.com/compute
    
    Network interfaces number depending on vCPU count:
    https://cloud.google.com/vpc/docs/create-use-multiple-interfaces#max-interfaces

&nbsp;

# 🗺️ __Roadmap__
- 👌 IMPROVE: CLI commands design
- 👌 IMPROVE: Better error handling
- 👌 IMPROVE: Better logging (logger, colored output, timer output for operations with long-term execution)
- 📖 DOC: Provide boilerplate .toml(s) for setting up a VPN
- 👌 IMPROVE: Back & forward slashes for command filepath options
- 👌 IMPROVE: Optional conf path for new-client command
- 📦 NEW: Feature to allow specific outbound TCP ports instead of all TCP ports
- 🤖 TEST: Tests
- 📖 DOC: Code docs
- 🚀 __RELEASE: <u>v0.1.0__</u>
- 📦 NEW: Force delete VPN
- 📦 NEW: Bulk generation of VPN client .conf(s)
- 🚀 __RELEASE: <u>v0.2.0__</u>
- 📦 NEW: AWS Support
- 🚀 __RELEASE: <u>v0.3.0__</u>
- 📦 NEW: Google Cloud Support
- 🚀 __RELEASE: <u>v0.4.0__</u>
- 📦 NEW: Microsoft Azure Support
- 🚀 __RELEASE: <u>v0.5.0__</u>
- 📦 NEW: Track VPN max capacity according to max network interfaces
- 📦 NEW: Track max instances quota per cloud
- 🚀 __RELEASE: <u>v0.6.0__</u>
- 📦 NEW: Dedicated IPs for VPN clients
- 🚀 __RELEASE: <u>v1.6.0__</u>
- 📦 NEW: Makefile

&nbsp;

# 📃 __License__
This app is licensed under the MIT License. See the LICENSE file for details.
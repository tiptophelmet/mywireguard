terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.26.0"
    }
  }
}

provider "digitalocean" {
  token = "{digital_ocean_api_token}"
}

resource "digitalocean_ssh_key" "mywireguard_ssh_key" {
  name       = "{vpn_identifier}-ssh-key"
  public_key = file("{digitalocean_ssh_public_key_path}")
}

resource "digitalocean_droplet" "mywireguard_vpn" {
  image  = "{digitalocean_droplet_image}"
  name   = "{vpn_identifier}-droplet"
  region = "{digitalocean_droplet_region}"
  size   = "{digitalocean_droplet_size}"
  ssh_keys = [
    digitalocean_ssh_key.mywireguard_ssh_key.id
  ]

  connection {
    host        = self.ipv4_address
    user        = "root"
    type        = "ssh"
    private_key = file("{digitalocean_ssh_private_key_path}")
  }

  provisioner "file" {
    source      = "wireguard_setup.sh"
    destination = "/root/wireguard_setup.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /root/wireguard_setup.sh",
      "/root/wireguard_setup.sh"
    ]
  }
}

resource "digitalocean_firewall" "mywireguard_vpn_firewall" {
  name = "{vpn_identifier}-firewall"
  droplet_ids = [
    digitalocean_droplet.mywireguard_vpn.id
  ]

  inbound_rule {
    protocol         = "tcp"
    port_range       = "22"
    source_addresses = ["0.0.0.0/0", "::/0"]
  }

  inbound_rule {
    protocol         = "udp"
    port_range       = "{wireguard_interface_listen_port}"
    source_addresses = ["0.0.0.0/0", "::/0"]
  }

  # Feature is planned: allow specific outbound TCP ports instead of all TCP ports
  outbound_rule {
    protocol              = "tcp"
    port_range            = "1-65535"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }

  outbound_rule {
    protocol              = "udp"
    port_range            = "1-65535"
    destination_addresses = ["0.0.0.0/0", "::/0"]
  }
}

output "public_ip" {
  value = digitalocean_droplet.mywireguard_vpn.ipv4_address
}

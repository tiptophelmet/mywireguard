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

resource "digitalocean_ssh_key" "wireguard_ssh_key" {
  name       = "wireguard_ssh_key"
  public_key = file("{digitalocean_ssh_public_key_path}")
}

resource "digitalocean_droplet" "mywireguard_vpn" {
  image  = "{digitalocean_droplet_image}"
  name   = "{digitalocean_droplet_name}"
  region = "{digitalocean_droplet_region}"
  size   = "{digitalocean_droplet_size}"
  ssh_keys = [
    digitalocean_ssh_key.wireguard_ssh_key.id
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

output "public_ip" {
  value = digitalocean_droplet.mywireguard_vpn.ipv4_address
}

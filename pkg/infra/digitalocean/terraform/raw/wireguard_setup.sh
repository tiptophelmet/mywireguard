#!/bin/bash

# Install WireGuard
add-apt-repository ppa:wireguard/wireguard -y
apt-get update
apt-get install -y wireguard

# Configure WireGuard
cat > /etc/wireguard/wg0.conf <<EOL
[Interface]
PrivateKey = {wireguard_server_private_key}
Address = {wireguard_interface_address}
ListenPort = {wireguard_interface_listen_port}
PostUp = iptables -A FORWARD -i wg0 -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
PostDown = iptables -D FORWARD -i wg0 -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE

EOL

# Enable IP forwarding
echo "net.ipv4.ip_forward=1" | tee -a /etc/sysctl.conf
sysctl -p

# Install UFW (if not already installed)
apt-get install -y ufw

# Configure UFW to allow incoming traffic on the WireGuard port
ufw allow {wireguard_interface_listen_port}/udp

# Enable and start WireGuard
systemctl enable wg-quick@wg0
systemctl start wg-quick@wg0 || { systemctl status wg-quick@wg0.service; journalctl -xe; }

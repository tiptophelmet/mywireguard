package cloud

type DigitalOceanDroplet struct {
	Image  string `toml:"image" terraform:"digitalocean_droplet_image"`
	Name   string `toml:"name" terraform:"digitalocean_droplet_name"`
	Region string `toml:"region" terraform:"digitalocean_droplet_region"`
	Size   string `toml:"size" terraform:"digitalocean_droplet_size"`
}

func (droplet *DigitalOceanDroplet) GetImage() string {
	return droplet.Image
}

func (droplet *DigitalOceanDroplet) GetName() string {
	return droplet.Name
}

func (droplet *DigitalOceanDroplet) GetRegion() string {
	return droplet.Region
}

func (droplet *DigitalOceanDroplet) GetSize() string {
	return droplet.Size
}

type DigitalOceanCloud struct {
	Provider string               `toml:"provider"`
	ApiToken string               `toml:"api_token" terraform:"digital_ocean_api_token"`
	Droplet  *DigitalOceanDroplet `toml:"droplet"`
}

func (cloud *DigitalOceanCloud) GetInstance() Instance {
	return cloud.Droplet
}

func (cloud *DigitalOceanCloud) GetApiToken() string {
	return cloud.ApiToken
}

package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getComputeVPNTunnelRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_compute_vpn_tunnel",
		CoreRFunc: NewComputeVPNTunnel,
	}
}

func NewComputeVPNTunnel(d *schema.ResourceData) schema.CoreResource {
	r := &google.ComputeVPNTunnel{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	return r
}

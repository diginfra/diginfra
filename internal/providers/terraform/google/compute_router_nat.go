package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getComputeRouterNATRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_compute_router_nat",
		CoreRFunc: NewComputeRouterNAT,
	}
}

func NewComputeRouterNAT(d *schema.ResourceData) schema.CoreResource {
	r := &google.ComputeRouterNAT{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	return r
}

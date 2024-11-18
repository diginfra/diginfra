package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getComputeMachineImageRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_compute_machine_image",
		CoreRFunc: newComputeMachineImage,
	}
}

func newComputeMachineImage(d *schema.ResourceData) schema.CoreResource {
	region := d.Get("region").String()

	r := &google.ComputeMachineImage{
		Address: d.Address,
		Region:  region,
	}
	return r
}

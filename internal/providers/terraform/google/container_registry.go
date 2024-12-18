package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getContainerRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:                "google_container_registry",
		CoreRFunc:           NewContainerRegistry,
		ReferenceAttributes: []string{},
	}
}
func NewContainerRegistry(d *schema.ResourceData) schema.CoreResource {
	r := &google.ContainerRegistry{
		Address:      d.Address,
		Region:       d.Get("region").String(),
		Location:     d.Get("location").String(),
		StorageClass: d.Get("storage_class").String(),
	}
	return r
}

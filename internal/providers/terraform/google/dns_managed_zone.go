package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getDNSManagedZoneRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_dns_managed_zone",
		CoreRFunc: NewDNSManagedZone,
	}
}

func NewDNSManagedZone(d *schema.ResourceData) schema.CoreResource {
	r := &google.DNSManagedZone{
		Address: d.Address,
	}

	return r
}

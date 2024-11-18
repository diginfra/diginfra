package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getDNSRecordSetRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_dns_record_set",
		CoreRFunc: NewDNSRecordSet,
	}
}

func NewDNSRecordSet(d *schema.ResourceData) schema.CoreResource {
	r := &google.DNSRecordSet{
		Address: d.Address,
	}

	return r
}

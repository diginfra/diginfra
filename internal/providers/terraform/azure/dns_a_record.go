package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getDNSARecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_dns_a_record",
		CoreRFunc: NewDNSARecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewDNSARecord(d *schema.ResourceData) schema.CoreResource {
	r := &azure.DNSARecord{Address: d.Address, Region: d.Region}
	return r
}

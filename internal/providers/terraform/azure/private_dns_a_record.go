package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getPrivateDNSARecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_private_dns_a_record",
		CoreRFunc: NewPrivateDNSARecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewPrivateDNSARecord(d *schema.ResourceData) schema.CoreResource {
	r := &azure.PrivateDNSARecord{Address: d.Address, Region: d.Region}
	return r
}

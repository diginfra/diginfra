package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getPrivateDNSAAAARecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_private_dns_aaaa_record",
		CoreRFunc: NewPrivateDNSAAAARecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewPrivateDNSAAAARecord(d *schema.ResourceData) schema.CoreResource {
	r := &azure.PrivateDNSAAAARecord{Address: d.Address, Region: d.Region}
	return r
}

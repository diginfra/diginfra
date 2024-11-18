package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getPrivateDNSPTRRecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_private_dns_ptr_record",
		CoreRFunc: NewPrivateDNSPTRRecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewPrivateDNSPTRRecord(d *schema.ResourceData) schema.CoreResource {
	r := &azure.PrivateDNSPTRRecord{Address: d.Address, Region: d.Region}
	return r
}

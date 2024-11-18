package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getDNSSrvRecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_dns_srv_record",
		CoreRFunc: NewDNSSrvRecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewDNSSrvRecord(d *schema.ResourceData) schema.CoreResource {
	r := &azure.DNSSrvRecord{Address: d.Address, Region: d.Region}
	return r
}

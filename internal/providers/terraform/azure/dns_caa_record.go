package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getDNSCAARecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_dns_caa_record",
		CoreRFunc: NewDNSCAARecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewDNSCAARecord(d *schema.ResourceData) schema.CoreResource {
	r := &azure.DNSCAARecord{Address: d.Address, Region: d.Region}
	return r
}

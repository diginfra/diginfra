package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getDNSNSRecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_dns_ns_record",
		CoreRFunc: NewDNSNSRecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewDNSNSRecord(d *schema.ResourceData) schema.CoreResource {
	r := &azure.DNSNSRecord{Address: d.Address, Region: d.Region}
	return r
}

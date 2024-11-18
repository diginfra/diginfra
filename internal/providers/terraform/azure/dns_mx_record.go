package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getDNSMXRecordRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_dns_mx_record",
		CoreRFunc: NewDNSMXRecord,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewDNSMXRecord(d *schema.ResourceData) schema.CoreResource {
	r := &azure.DNSMXRecord{Address: d.Address, Region: d.Region}
	return r
}

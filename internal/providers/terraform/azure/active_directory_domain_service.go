package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getActiveDirectoryDomainServiceRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_active_directory_domain_service",
		CoreRFunc: NewActiveDirectoryDomainService,
	}
}
func NewActiveDirectoryDomainService(d *schema.ResourceData) schema.CoreResource {
	r := &azure.ActiveDirectoryDomainService{Address: d.Address, Region: d.Region, SKU: d.Get("sku").String()}
	return r
}

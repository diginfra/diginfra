package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getBastionHostRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_bastion_host",
		CoreRFunc: NewBastionHost,
	}
}
func NewBastionHost(d *schema.ResourceData) schema.CoreResource {
	r := &azure.BastionHost{Address: d.Address, Region: d.Region}
	return r
}

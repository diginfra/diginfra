package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getDataFactoryRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_data_factory",
		CoreRFunc: newDataFactory,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newDataFactory(d *schema.ResourceData) schema.CoreResource {
	region := d.Region

	r := &azure.DataFactory{
		Address: d.Address,
		Region:  region,
	}
	return r
}

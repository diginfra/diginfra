package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getEventgridSystemTopicRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name: "azurerm_eventgrid_system_topic",
		CoreRFunc: func(d *schema.ResourceData) schema.CoreResource {
			return &azure.EventGridTopic{
				Address: d.Address,
				Region:  d.Region,
			}
		},
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getMonitorDataCollectionRuleRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_monitor_data_collection_rule",
		CoreRFunc: newMonitorDataCollectionRule,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newMonitorDataCollectionRule(d *schema.ResourceData) schema.CoreResource {
	region := d.Region
	return &azure.MonitorDataCollectionRule{
		Address: d.Address,
		Region:  region,
	}
}

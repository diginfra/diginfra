package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getAutomationDSCConfigurationRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_automation_dsc_configuration",
		CoreRFunc: NewAutomationDSCConfiguration,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewAutomationDSCConfiguration(d *schema.ResourceData) schema.CoreResource {
	r := &azure.AutomationDSCConfiguration{Address: d.Address, Region: d.Region}
	return r
}

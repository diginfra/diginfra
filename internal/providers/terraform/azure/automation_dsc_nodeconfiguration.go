package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getAutomationDSCNodeConfigurationRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_automation_dsc_nodeconfiguration",
		CoreRFunc: NewAutomationDSCNodeConfiguration,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewAutomationDSCNodeConfiguration(d *schema.ResourceData) schema.CoreResource {
	r := &azure.AutomationDSCNodeConfiguration{Address: d.Address, Region: d.Region}
	return r
}

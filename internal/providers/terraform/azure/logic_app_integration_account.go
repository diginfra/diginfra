package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getLogicAppIntegrationAccountRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_logic_app_integration_account",
		CoreRFunc: newLogicAppIntegrationAccount,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newLogicAppIntegrationAccount(d *schema.ResourceData) schema.CoreResource {
	region := d.Region

	return azure.NewLogicAppIntegrationAccount(d.Address, region, d.GetStringOrDefault("sku_name", "free"))
}

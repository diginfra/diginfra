package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getApplicationInsightsStandardWebTestRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_application_insights_standard_web_test",
		CoreRFunc: newApplicationInsightsStandardWebTest,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newApplicationInsightsStandardWebTest(d *schema.ResourceData) schema.CoreResource {
	region := d.Region
	return &azure.ApplicationInsightsStandardWebTest{
		Address:   d.Address,
		Region:    region,
		Enabled:   d.GetBoolOrDefault("enabled", true),
		Frequency: d.GetInt64OrDefault("frequency", 300),
	}
}

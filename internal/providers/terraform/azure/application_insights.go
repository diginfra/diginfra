package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getApplicationInsightsRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_application_insights",
		CoreRFunc: NewApplicationInsights,
	}
}
func NewApplicationInsights(d *schema.ResourceData) schema.CoreResource {
	r := &azure.ApplicationInsights{
		Address:         d.Address,
		Region:          d.Region,
		RetentionInDays: d.Get("retention_in_days").Int(),
	}
	return r
}

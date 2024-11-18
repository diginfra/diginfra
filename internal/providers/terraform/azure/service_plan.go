package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getServicePlanRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name: "azurerm_service_plan",
		CoreRFunc: func(d *schema.ResourceData) schema.CoreResource {
			return &azure.ServicePlan{
				Address:     d.Address,
				Region:      d.Region,
				SKUName:     d.Get("sku_name").String(),
				WorkerCount: d.GetInt64OrDefault("worker_count", 1),
				OSType:      d.Get("os_type").String(),
			}
		},
	}
}

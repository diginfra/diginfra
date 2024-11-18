package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getAutomationJobScheduleRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_automation_job_schedule",
		CoreRFunc: NewAutomationJobSchedule,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}
func NewAutomationJobSchedule(d *schema.ResourceData) schema.CoreResource {
	r := &azure.AutomationJobSchedule{Address: d.Address, Region: d.Region}
	return r
}

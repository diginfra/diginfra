package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getCloudwatchDashboardRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_cloudwatch_dashboard",
		CoreRFunc: NewCloudwatchDashboard,
	}
}
func NewCloudwatchDashboard(d *schema.ResourceData) schema.CoreResource {
	r := &aws.CloudwatchDashboard{
		Address: d.Address,
	}
	return r
}

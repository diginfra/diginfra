package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getCloudwatchEventBusItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_cloudwatch_event_bus",
		CoreRFunc: NewCloudwatchEventBus,
	}
}
func NewCloudwatchEventBus(d *schema.ResourceData) schema.CoreResource {
	r := &aws.CloudwatchEventBus{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}

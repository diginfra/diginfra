package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getKinesisAnalyticsApplicationRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_kinesis_analytics_application",
		CoreRFunc: NewKinesisAnalyticsApplication,
	}
}

func NewKinesisAnalyticsApplication(d *schema.ResourceData) schema.CoreResource {
	r := &aws.KinesisAnalyticsApplication{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}

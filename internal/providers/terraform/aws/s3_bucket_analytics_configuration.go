package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getS3BucketAnalyticsConfigurationRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_s3_bucket_analytics_configuration",
		CoreRFunc: NewS3BucketAnalyticsConfiguration,
	}
}

func NewS3BucketAnalyticsConfiguration(d *schema.ResourceData) schema.CoreResource {
	r := &aws.S3BucketAnalyticsConfiguration{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}

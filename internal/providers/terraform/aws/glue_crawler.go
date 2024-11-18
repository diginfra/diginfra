package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getGlueCrawlerRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_glue_crawler",
		CoreRFunc: newGlueCrawler,
	}
}

func newGlueCrawler(d *schema.ResourceData) schema.CoreResource {
	region := d.Get("region").String()
	r := &aws.GlueCrawler{
		Address: d.Address,
		Region:  region,
	}

	return r
}

package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getGlueCatalogDatabaseRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_glue_catalog_database",
		CoreRFunc: newGlueCatalogDatabase,
	}
}

func newGlueCatalogDatabase(d *schema.ResourceData) schema.CoreResource {
	region := d.Get("region").String()
	r := &aws.GlueCatalogDatabase{
		Address: d.Address,
		Region:  region,
	}

	return r
}

package aws

import (
	"github.com/diginfra/diginfra/internal/schema"
)

func getOpensearchDomainRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_opensearch_domain",
		CoreRFunc: newSearchDomain,
	}
}

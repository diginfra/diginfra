package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getConfigOrganizationCustomRuleItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_config_organization_custom_rule",
		CoreRFunc: NewConfigOrganizationCustomRule,
	}
}
func NewConfigOrganizationCustomRule(d *schema.ResourceData) schema.CoreResource {
	r := &aws.ConfigConfigRule{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}

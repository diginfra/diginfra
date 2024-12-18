package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getSSMParameterRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_ssm_parameter",
		CoreRFunc: NewSSMParameter,
	}
}

func NewSSMParameter(d *schema.ResourceData) schema.CoreResource {
	r := &aws.SSMParameter{
		Address: d.Address,
		Region:  d.Get("region").String(),
		Tier:    d.Get("tier").String(),
	}
	return r
}

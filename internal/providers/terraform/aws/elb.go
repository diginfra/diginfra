package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getELBRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_elb",
		CoreRFunc: NewELB,
	}
}
func NewELB(d *schema.ResourceData) schema.CoreResource {
	r := &aws.ELB{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}

package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getDXConnectionRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_dx_connection",
		CoreRFunc: NewDXConnection,
	}
}

func NewDXConnection(d *schema.ResourceData) schema.CoreResource {
	r := &aws.DXConnection{
		Address:   d.Address,
		Region:    d.Get("region").String(),
		Bandwidth: d.Get("bandwidth").String(),
		Location:  d.Get("location").String(),
	}
	return r
}

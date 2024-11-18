package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getGlobalAcceleratorRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_globalaccelerator_accelerator",
		CoreRFunc: newGlobalAccelerator,
	}
}

func newGlobalAccelerator(d *schema.ResourceData) schema.CoreResource {

	r := &aws.GlobalAccelerator{
		Address: d.Address,
	}

	return r
}

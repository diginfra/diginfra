package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getEC2ClientVPNNetworkAssociationRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_ec2_client_vpn_network_association",
		CoreRFunc: NewEC2ClientVPNNetworkAssociation,
	}
}
func NewEC2ClientVPNNetworkAssociation(d *schema.ResourceData) schema.CoreResource {
	r := &aws.EC2ClientVPNNetworkAssociation{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}

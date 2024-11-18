package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getNewEKSFargateProfileItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_eks_fargate_profile",
		CoreRFunc: NewEKSFargateProfile,
	}
}
func NewEKSFargateProfile(d *schema.ResourceData) schema.CoreResource {
	r := &aws.EKSFargateProfile{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}

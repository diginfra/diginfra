package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getEC2TrafficMirrorSessionRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_ec2_traffic_mirror_session",
		CoreRFunc: NewEC2TrafficMirrorSession,
	}
}
func NewEC2TrafficMirrorSession(d *schema.ResourceData) schema.CoreResource {
	r := &aws.EC2TrafficMirrorSession{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}

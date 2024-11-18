package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getAPIGatewayRestAPIRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_api_gateway_rest_api",
		CoreRFunc: NewAPIGatewayRestAPI,
	}
}
func NewAPIGatewayRestAPI(d *schema.ResourceData) schema.CoreResource {
	r := &aws.APIGatewayRestAPI{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}

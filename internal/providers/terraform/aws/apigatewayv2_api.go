package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getAPIGatewayV2APIRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_apigatewayv2_api",
		CoreRFunc: NewAPIGatewayV2API,
	}
}
func NewAPIGatewayV2API(d *schema.ResourceData) schema.CoreResource {
	r := &aws.APIGatewayV2API{
		Address:      d.Address,
		ProtocolType: d.Get("protocol_type").String(),
		Region:       d.Get("region").String(),
	}
	return r
}

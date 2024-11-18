package aws

import (
	"github.com/diginfra/diginfra/internal/resources"
	"github.com/diginfra/diginfra/internal/schema"
)

type EC2TransitGatewayPeeringAttachment struct {
	Address              string
	Region               string
	TransitGatewayRegion string
}

func (r *EC2TransitGatewayPeeringAttachment) CoreType() string {
	return "EC2TransitGatewayPeeringAttachment"
}

func (r *EC2TransitGatewayPeeringAttachment) UsageSchema() []*schema.UsageItem {
	return []*schema.UsageItem{}
}

func (r *EC2TransitGatewayPeeringAttachment) PopulateUsage(u *schema.UsageData) {
	resources.PopulateArgsWithUsage(r, u)
}

func (r *EC2TransitGatewayPeeringAttachment) BuildResource() *schema.Resource {
	region := r.Region
	if r.TransitGatewayRegion != "" {
		region = r.TransitGatewayRegion
	}

	return &schema.Resource{
		Name: r.Address,
		CostComponents: []*schema.CostComponent{
			transitGatewayAttachmentCostComponent(region, "TransitGatewayPeering"),
		}, UsageSchema: r.UsageSchema(),
	}
}

package aws

import (
	"github.com/diginfra/diginfra/internal/resources"
	"github.com/diginfra/diginfra/internal/schema"

	"github.com/shopspring/decimal"
)

type VPNConnection struct {
	Address                string
	Region                 string
	TransitGatewayID       string
	MonthlyDataProcessedGB *float64 `diginfra_usage:"monthly_data_processed_gb"`
}

func (r *VPNConnection) CoreType() string {
	return "VPNConnection"
}

func (r *VPNConnection) UsageSchema() []*schema.UsageItem {
	return []*schema.UsageItem{{Key: "monthly_data_processed_gb", ValueType: schema.Float64, DefaultValue: 0}}
}

func (r *VPNConnection) PopulateUsage(u *schema.UsageData) {
	resources.PopulateArgsWithUsage(r, u)
}

func (r *VPNConnection) BuildResource() *schema.Resource {
	region := r.Region

	var gbDataProcessed *decimal.Decimal

	costComponents := []*schema.CostComponent{
		{
			Name:           "VPN connection",
			Unit:           "hours",
			UnitMultiplier: decimal.NewFromInt(1),
			HourlyQuantity: decimalPtr(decimal.NewFromInt(1)),
			ProductFilter: &schema.ProductFilter{
				VendorName:    strPtr("aws"),
				Region:        strPtr(region),
				Service:       strPtr("AmazonVPC"),
				ProductFamily: strPtr("Cloud Connectivity"),
			},
		},
	}

	if r.TransitGatewayID != "" {
		costComponents = append(costComponents, transitGatewayAttachmentCostComponent(region, "TransitGatewayVPN"))

		if r.MonthlyDataProcessedGB != nil {
			gbDataProcessed = decimalPtr(decimal.NewFromFloat(*r.MonthlyDataProcessedGB))
		}

		costComponents = append(costComponents, transitGatewayDataProcessingCostComponent(region, "TransitGatewayVPN", gbDataProcessed))
	}

	return &schema.Resource{
		Name:           r.Address,
		CostComponents: costComponents,
		UsageSchema:    r.UsageSchema(),
	}
}

package azure

import (
	"github.com/diginfra/diginfra/internal/resources"
	"github.com/diginfra/diginfra/internal/schema"
)

type PrivateDNSSRVRecord struct {
	Address        string
	Region         string
	MonthlyQueries *int64 `diginfra_usage:"monthly_queries"`
}

func (r *PrivateDNSSRVRecord) CoreType() string {
	return "PrivateDNSSRVRecord"
}

func (r *PrivateDNSSRVRecord) UsageSchema() []*schema.UsageItem {
	return []*schema.UsageItem{{Key: "monthly_queries", ValueType: schema.Int64, DefaultValue: 0}}
}

func (r *PrivateDNSSRVRecord) PopulateUsage(u *schema.UsageData) {
	resources.PopulateArgsWithUsage(r, u)
}

func (r *PrivateDNSSRVRecord) BuildResource() *schema.Resource {
	return &schema.Resource{
		Name:           r.Address,
		CostComponents: dnsQueriesCostComponent(r.Region, r.MonthlyQueries),
		UsageSchema:    r.UsageSchema(),
	}
}

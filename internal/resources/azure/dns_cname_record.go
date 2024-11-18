package azure

import (
	"github.com/diginfra/diginfra/internal/resources"
	"github.com/diginfra/diginfra/internal/schema"
)

type DNSCNameRecord struct {
	Address        string
	Region         string
	MonthlyQueries *int64 `diginfra_usage:"monthly_queries"`
}

func (r *DNSCNameRecord) CoreType() string {
	return "DNSCNameRecord"
}

func (r *DNSCNameRecord) UsageSchema() []*schema.UsageItem {
	return []*schema.UsageItem{{Key: "monthly_queries", ValueType: schema.Int64, DefaultValue: 0}}
}

func (r *DNSCNameRecord) PopulateUsage(u *schema.UsageData) {
	resources.PopulateArgsWithUsage(r, u)
}

func (r *DNSCNameRecord) BuildResource() *schema.Resource {
	return &schema.Resource{
		Name:           r.Address,
		CostComponents: dnsQueriesCostComponent(r.Region, r.MonthlyQueries),
		UsageSchema:    r.UsageSchema(),
	}
}

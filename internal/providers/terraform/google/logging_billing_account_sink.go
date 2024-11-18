package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getLoggingBillingAccountSinkRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_logging_billing_account_sink",
		CoreRFunc: NewLoggingBillingAccountSink,
	}
}

func NewLoggingBillingAccountSink(d *schema.ResourceData) schema.CoreResource {
	r := &google.Logging{
		Address: d.Address,
	}

	return r
}

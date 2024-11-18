package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getLoggingBucketConfigRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_logging_project_bucket_config",
		CoreRFunc: NewLoggingProjectBucketConfig,
	}
}

func NewLoggingProjectBucketConfig(d *schema.ResourceData) schema.CoreResource {
	r := &google.Logging{
		Address: d.Address,
	}

	return r
}

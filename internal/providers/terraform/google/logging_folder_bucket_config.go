package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getLoggingFolderBucketConfigRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_logging_folder_bucket_config",
		CoreRFunc: NewLoggingFolderBucketConfig,
	}
}

func NewLoggingFolderBucketConfig(d *schema.ResourceData) schema.CoreResource {
	r := &google.Logging{
		Address: d.Address,
	}

	return r
}

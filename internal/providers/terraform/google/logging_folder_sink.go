package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getLoggingFolderSinkRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_logging_folder_sink",
		CoreRFunc: NewLoggingFolderSink,
	}
}

func NewLoggingFolderSink(d *schema.ResourceData) schema.CoreResource {
	r := &google.Logging{
		Address: d.Address,
	}

	return r
}

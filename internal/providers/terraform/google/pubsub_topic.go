package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getPubSubTopicRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_pubsub_topic",
		CoreRFunc: NewPubSubTopic,
	}
}

func NewPubSubTopic(d *schema.ResourceData) schema.CoreResource {
	r := &google.PubSubTopic{
		Address: d.Address,
	}

	return r
}

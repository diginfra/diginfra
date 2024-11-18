package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getPubSubSubscriptionRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_pubsub_subscription",
		CoreRFunc: NewPubSubSubscription,
	}
}

func NewPubSubSubscription(d *schema.ResourceData) schema.CoreResource {
	r := &google.PubSubSubscription{
		Address: d.Address,
	}

	return r
}

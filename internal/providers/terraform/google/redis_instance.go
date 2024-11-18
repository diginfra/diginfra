package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getRedisInstanceRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "google_redis_instance",
		CoreRFunc: NewRedisInstance,
	}
}

func NewRedisInstance(d *schema.ResourceData) schema.CoreResource {
	r := &google.RedisInstance{
		Address:      d.Address,
		Region:       d.Get("region").String(),
		MemorySizeGB: d.Get("memory_size_gb").Float(),
		Tier:         d.Get("tier").String(),
	}

	return r
}

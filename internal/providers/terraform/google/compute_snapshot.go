package google

import (
	"github.com/diginfra/diginfra/internal/resources/google"
	"github.com/diginfra/diginfra/internal/schema"
)

func getComputeSnapshotRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:                "google_compute_snapshot",
		CoreRFunc:           newComputeSnapshot,
		ReferenceAttributes: []string{"source_disk"},
	}
}

func newComputeSnapshot(d *schema.ResourceData) schema.CoreResource {
	region := d.Get("region").String()

	size := computeSnapshotDiskSize(d)

	r := &google.ComputeSnapshot{
		Address:  d.Address,
		Region:   region,
		DiskSize: size,
	}
	return r
}

package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getDocDBClusterSnapshotRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_docdb_cluster_snapshot",
		CoreRFunc: NewDocDBClusterSnapshot,
	}

}
func NewDocDBClusterSnapshot(d *schema.ResourceData) schema.CoreResource {
	r := &aws.DocDBClusterSnapshot{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}

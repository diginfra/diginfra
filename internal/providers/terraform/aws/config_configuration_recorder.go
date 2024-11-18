package aws

import (
	"github.com/diginfra/diginfra/internal/resources/aws"
	"github.com/diginfra/diginfra/internal/schema"
)

func getConfigurationRecorderItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "aws_config_configuration_recorder",
		CoreRFunc: NewConfigConfigurationRecorder,
	}
}
func NewConfigConfigurationRecorder(d *schema.ResourceData) schema.CoreResource {
	r := &aws.ConfigConfigurationRecorder{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}
	return r
}

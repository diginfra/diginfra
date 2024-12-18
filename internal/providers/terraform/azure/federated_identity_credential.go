package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getFederatedIdentityCredentialRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_federated_identity_credential",
		CoreRFunc: newFederatedIdentityCredential,
		ReferenceAttributes: []string{
			"resource_group_name",
		},
	}
}

func newFederatedIdentityCredential(d *schema.ResourceData) schema.CoreResource {
	region := d.Region
	return &azure.FederatedIdentityCredential{
		Address: d.Address,
		Region:  region,
	}
}

package azure

import (
	"github.com/diginfra/diginfra/internal/resources/azure"
	"github.com/diginfra/diginfra/internal/schema"
)

func getExpressRouteGatewayRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_express_route_gateway",
		CoreRFunc: newExpressRouteGateway,
	}
}

func newExpressRouteGateway(d *schema.ResourceData) schema.CoreResource {
	e := &azure.ExpressRouteGateway{
		Address:    d.Address,
		Region:     d.Get("region").String(),
		ScaleUnits: d.Get("scale_units").Int(),
	}

	return e
}

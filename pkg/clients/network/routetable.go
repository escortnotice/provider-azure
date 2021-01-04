package network

import (
	networkmgmt "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network"
	"github.com/crossplane/provider-azure/apis/network/v1alpha3"
	azure "github.com/crossplane/provider-azure/pkg/clients"
)

// UpdateRouteTableStatusFromAzure updates the status related to the external
// Azure RouteTable in the RouteTableStatus
func UpdateRouteTableStatusFromAzure(rt *v1alpha3.RouteTable, az networkmgmt.RouteTable) {
	rt.Status.State = azure.ToString(az.ProvisioningState)
	rt.Status.ID = azure.ToString(az.ID)
	rt.Status.Etag = azure.ToString(az.Etag)
	rt.Status.Type = azure.ToString(az.Type)
}

// NewRouteTableParameters returns an Azure VirtualNetwork object from a virtual network spec
func NewRouteTableParameters(v *v1alpha3.RouteTable) networkmgmt.RouteTable {
	return networkmgmt.RouteTable{
		RouteTablePropertiesFormat: &networkmgmt.RouteTablePropertiesFormat{
			Routes:                     setRoutes(v.Spec.Properties.Routes),
			Subnets:                    setSubnets(v.Spec.Properties.Subnets),
			DisableBgpRoutePropagation: azure.ToBoolPtr(v.Spec.Properties.DisableBgpRoutePropagation),
		},
		Name:     azure.ToStringPtr(v.Spec.Name),
		Location: azure.ToStringPtr(v.Spec.Location),
		Tags:     azure.ToStringPtrMap(v.Spec.Tags),
	}
}

func setSubnets(subnets []v1alpha3.Subnet) *[]networkmgmt.Subnet {
	return nil
}

func setRoutes(routes []v1alpha3.Route) *[]networkmgmt.Route {
	return nil
}

package network

import (
	"reflect"

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

// NewRouteTableParameters returns an Azure RouteTable object from a Route Table Spec
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
	var snets = new([]networkmgmt.Subnet)
	if subnets != nil {
		for _, rt := range subnets {
			var subnet = networkmgmt.Subnet{}
			subnet.ID = azure.ToStringPtr(rt.Status.ID)
			subnet.Name = azure.ToStringPtr(rt.Name)
			subnet.Etag = azure.ToStringPtr(rt.Status.Etag)
			var subnetProperties = &networkmgmt.SubnetPropertiesFormat{
				AddressPrefix:                     azure.ToStringPtr(rt.Spec.AddressPrefix),
				AddressPrefixes:                   nil,
				NetworkSecurityGroup:              nil,
				RouteTable:                        nil,
				NatGateway:                        nil,
				ServiceEndpoints:                  nil,
				ServiceEndpointPolicies:           nil,
				PrivateEndpoints:                  nil,
				IPConfigurations:                  nil,
				IPConfigurationProfiles:           nil,
				ResourceNavigationLinks:           nil,
				ServiceAssociationLinks:           nil,
				Delegations:                       nil,
				Purpose:                           nil,
				ProvisioningState:                 nil,
				PrivateEndpointNetworkPolicies:    nil,
				PrivateLinkServiceNetworkPolicies: nil,
			}
			subnet.SubnetPropertiesFormat = subnetProperties
			*snets = append(*snets, subnet)
		}
		return snets
	}
	return nil
}

func setRoutes(routes []v1alpha3.Route) *[]networkmgmt.Route {
	var rs = new([]networkmgmt.Route)
	if routes != nil {
		for _, r := range routes {
			var route = new(networkmgmt.Route)
			route.Name = azure.ToStringPtr(r.Name)
			route.Etag = azure.ToStringPtr(r.Etag)
			route.ID = azure.ToStringPtr(r.ID)
			var properties = &networkmgmt.RoutePropertiesFormat{
				AddressPrefix:     azure.ToStringPtr(r.Properties.AddressPrefix),
				NextHopType:       networkmgmt.RouteNextHopType(r.Properties.NextHopType),
				NextHopIPAddress:  azure.ToStringPtr(r.Properties.NextHopIPAddress),
				ProvisioningState: azure.ToStringPtr(r.Properties.ProvisioningState),
			}
			route.RoutePropertiesFormat = properties
			*rs = append(*rs, *route)
		}
		return rs
	}
	return nil
}

// RouteTableNeedsUpdate determines if a Route Table need to be updated
func RouteTableNeedsUpdate(kube *v1alpha3.RouteTable, az networkmgmt.RouteTable) bool {
	up := NewRouteTableParameters(kube)

	switch {
	case !reflect.DeepEqual(up.RouteTablePropertiesFormat.Subnets, az.RouteTablePropertiesFormat.Subnets):
		return true
	case !reflect.DeepEqual(up.RouteTablePropertiesFormat.Routes, az.RouteTablePropertiesFormat.Routes):
		return true
	case !reflect.DeepEqual(up.RouteTablePropertiesFormat.DisableBgpRoutePropagation, az.RouteTablePropertiesFormat.DisableBgpRoutePropagation):
		return true
	case !reflect.DeepEqual(up.Tags, az.Tags):
		return true
	}

	return false
}

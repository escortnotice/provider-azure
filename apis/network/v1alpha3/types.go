/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// AddressSpace contains an array of IP address ranges that can be used by
// subnets of the virtual network.
type AddressSpace struct {
	// AddressPrefixes - A list of address blocks reserved for this virtual
	// network in CIDR notation.
	AddressPrefixes []string `json:"addressPrefixes"`
}

// VirtualNetworkPropertiesFormat defines properties of a VirtualNetwork.
type VirtualNetworkPropertiesFormat struct {
	// AddressSpace - The AddressSpace that contains an array of IP address
	// ranges that can be used by subnets.
	// +optional
	AddressSpace AddressSpace `json:"addressSpace"`

	// EnableDDOSProtection - Indicates if DDoS protection is enabled for all
	// the protected resources in the virtual network. It requires a DDoS
	// protection plan associated with the resource.
	// +optional
	EnableDDOSProtection bool `json:"enableDdosProtection,omitempty"`

	// EnableVMProtection - Indicates if VM protection is enabled for all the
	// subnets in the virtual network.
	// +optional
	EnableVMProtection bool `json:"enableVmProtection,omitempty"`
}

// A VirtualNetworkSpec defines the desired state of a VirtualNetwork.
type VirtualNetworkSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`

	// ResourceGroupName - Name of the Virtual Network's resource group.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ResourceGroupNameRef - A reference to the the Virtual Network's resource
	// group.
	ResourceGroupNameRef *runtimev1alpha1.Reference `json:"resourceGroupNameRef,omitempty"`

	// ResourceGroupNameSelector - Select a reference to the the Virtual
	// Network's resource group.
	ResourceGroupNameSelector *runtimev1alpha1.Selector `json:"resourceGroupNameSelector,omitempty"`

	// VirtualNetworkPropertiesFormat - Properties of the virtual network.
	VirtualNetworkPropertiesFormat `json:"properties"`

	// Location - Resource location.
	Location string `json:"location"`

	// Tags - Resource tags.
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

// A VirtualNetworkStatus represents the observed state of a VirtualNetwork.
type VirtualNetworkStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`

	// State of this VirtualNetwork.
	State string `json:"state,omitempty"`

	// A Message providing detail about the state of this VirtualNetwork, if
	// any.
	Message string `json:"message,omitempty"`

	// ID of this VirtualNetwork.
	ID string `json:"id,omitempty"`

	// Etag - A unique read-only string that changes whenever the resource is
	// updated.
	Etag string `json:"etag,omitempty"`

	// ResourceGUID - The GUID of this VirtualNetwork.
	ResourceGUID string `json:"resourceGuid,omitempty"`

	// Type of this VirtualNetwork.
	Type string `json:"type,omitempty"`
}

// +kubebuilder:object:root=true

// A VirtualNetwork is a managed resource that represents an Azure Virtual
// Network.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualNetworkSpec   `json:"spec"`
	Status VirtualNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkList contains a list of VirtualNetwork items
type VirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetwork `json:"items"`
}

// ServiceEndpointPropertiesFormat defines properties of a service endpoint.
type ServiceEndpointPropertiesFormat struct {
	// Service - The type of the endpoint service.
	// +optional
	Service string `json:"service,omitempty"`

	// Locations - A list of locations.
	// +optional
	Locations []string `json:"locations,omitempty"`

	// ProvisioningState - The provisioning state of the resource.
	// +optional
	ProvisioningState string `json:"provisioningState,omitempty"`
}

// SubnetPropertiesFormat defines properties of a Subnet.
type SubnetPropertiesFormat struct {
	// AddressPrefix - The address prefix for the subnet.
	AddressPrefix string `json:"addressPrefix"`

	// ServiceEndpoints - An array of service endpoints.
	ServiceEndpoints []ServiceEndpointPropertiesFormat `json:"serviceEndpoints,omitempty"`
}

// A SubnetSpec defines the desired state of a Subnet.
type SubnetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`

	// VirtualNetworkName - Name of the Subnet's virtual network.
	VirtualNetworkName string `json:"virtualNetworkName,omitempty"`

	// VirtualNetworkNameRef references to a VirtualNetwork to retrieve its name
	VirtualNetworkNameRef *runtimev1alpha1.Reference `json:"virtualNetworkNameRef,omitempty"`

	// VirtualNetworkNameSelector selects a reference to a VirtualNetwork to
	// retrieve its name
	VirtualNetworkNameSelector *runtimev1alpha1.Selector `json:"virtualNetworkNameSelector,omitempty"`

	// ResourceGroupName - Name of the Subnet's resource group.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ResourceGroupNameRef - A reference to the the Subnets's resource group.
	ResourceGroupNameRef *runtimev1alpha1.Reference `json:"resourceGroupNameRef,omitempty"`

	// ResourceGroupNameSelector - Selects a reference to the the Subnets's
	// resource group.
	ResourceGroupNameSelector *runtimev1alpha1.Selector `json:"resourceGroupNameSelector,omitempty"`

	// SubnetPropertiesFormat - Properties of the subnet.
	SubnetPropertiesFormat `json:"properties"`
}

// A SubnetStatus represents the observed state of a Subnet.
type SubnetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`

	// State of this Subnet.
	State string `json:"state,omitempty"`

	// A Message providing detail about the state of this Subnet, if any.
	Message string `json:"message,omitempty"`

	// Etag - A unique string that changes whenever the resource is updated.
	Etag string `json:"etag,omitempty"`

	// ID of this Subnet.
	ID string `json:"id,omitempty"`

	// Purpose - A string identifying the intention of use for this subnet based
	// on delegations and other user-defined properties.
	Purpose string `json:"purpose,omitempty"`
}

// +kubebuilder:object:root=true

// A Subnet is a managed resource that represents an Azure Subnet.
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SubnetSpec   `json:"spec"`
	Status SubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetList contains a list of Subnet items
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

// A RouteTableSpec represents the observed state of a RouteTable.
type RouteTableSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`

	// ResourceGroupName - Name of the RouteTable's resource group.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ResourceGroupNameRef - A reference to the the RouteTable's resource group.
	ResourceGroupNameRef *runtimev1alpha1.Reference `json:"resourceGroupNameRef,omitempty"`

	// ResourceGroupNameSelector - Selects a reference to the the RouteTable's
	// resource group.
	ResourceGroupNameSelector *runtimev1alpha1.Selector `json:"resourceGroupNameSelector,omitempty"`

	// RouteTablePropertiesFormat - Properties of the route table.
	Properties RouteTablePropertiesFormat `json:"properties,omitempty"`
	// Name - Resource name.
	Name string `json:"name,omitempty"`

	// Location - Resource location.
	Location string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]string `json:"tags"`
}

// RouteTablePropertiesFormat route Table resource.
type RouteTablePropertiesFormat struct {
	// Routes - Collection of routes contained within a route table.
	Routes []Route `json:"routes,omitempty"`
	// Subnets - A collection of references to subnets.
	Subnets []Subnet `json:"subnets,omitempty"`
	// DisableBgpRoutePropagation - Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
	DisableBgpRoutePropagation bool `json:"disableBgpRoutePropagation,omitempty"`
	// ProvisioningState - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState string `json:"provisioningState,omitempty"`
}

// Route route resource.
type Route struct {
	// RoutePropertiesFormat - Properties of the route.
	Properties RoutePropertiesFormat `json:"properties,omitempty"`
	// Name - The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `json:"name,omitempty"`
	// Etag - A unique read-only string that changes whenever the resource is updated.
	Etag string `json:"etag,omitempty"`
	// ID - Resource ID.
	ID string `json:"id,omitempty"`
}

// RoutePropertiesFormat route resource.
type RoutePropertiesFormat struct {
	// AddressPrefix - The destination CIDR to which the route applies.
	AddressPrefix string `json:"addressPrefix,omitempty"`
	// NextHopType - The type of Azure hop the packet should be sent to. Possible values include: 'RouteNextHopTypeVirtualNetworkGateway', 'RouteNextHopTypeVnetLocal', 'RouteNextHopTypeInternet', 'RouteNextHopTypeVirtualAppliance', 'RouteNextHopTypeNone'
	NextHopType string `json:"nextHopType,omitempty"`
	// NextHopIPAddress - The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
	NextHopIPAddress string `json:"nextHopIpAddress,omitempty"`
	// ProvisioningState - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState string `json:"provisioningState,omitempty"`
}

// A RouteTableStatus represents the observed state of a RouteTable.
type RouteTableStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`

	// State of this RouteTable.
	State string `json:"state,omitempty"`

	// A Message providing detail about the state of this RouteTable, if any.
	Message string `json:"message,omitempty"`

	// Etag - A unique string that changes whenever the resource is updated.
	Etag string `json:"etag,omitempty"`

	// ID of this RouteTable.
	ID string `json:"id,omitempty"`

	// Type - Resource type.
	Type string `json:"type,omitempty"`

	// Purpose - A string identifying the intention of use for this subnet based
	// on delegations and other user-defined properties.
	Purpose string `json:"purpose,omitempty"`
}

// +kubebuilder:object:root=true

// A RouteTable is a managed resource that represents an Azure RouteTable.
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteTableSpec   `json:"spec"`
	Status RouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetList contains a list of Subnet items
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

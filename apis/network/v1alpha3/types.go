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



// +kubebuilder:object:root=true

// A PublicIPAddress is a managed resource that represents an Azure PublicIPAddress
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PublicIPAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   PublicIPAddressSpec   `json:"spec"`

	Status PublicIPAddressStatus `json:"status,omitempty"`
	
	// Etag - A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`
	// ID - Resource Identifier.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type *string `json:"type,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`

}
// PublicIPAddressSkuName enumerates the values for public ip address sku name.
type PublicIPAddressSkuName string

const (
	// PublicIPAddressSkuNameBasic ...
	PublicIPAddressSkuNameBasic PublicIPAddressSkuName = "Basic"
	// PublicIPAddressSkuNameStandard ...
	PublicIPAddressSkuNameStandard PublicIPAddressSkuName = "Standard"
)

// PublicIPAddressSpec contains PublicIPAddress specification
type PublicIPAddressSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`

	// ResourceGroupName - Name of the PublicIPAddress's resource group.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ResourceGroupNameRef - A reference to the the PublicIPAddress' resource
	// group.
	ResourceGroupNameRef *runtimev1alpha1.Reference `json:"resourceGroupNameRef,omitempty"`

	//PublicIPAddressPropertiesFormat - Properties of the application security group.
	Properties PublicIPAddressPropertiesFormat `json:"properties,omitempty"`
	// Etag - READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag string `json:"etag,omitempty"`
	// ID - Resource ID.
	ID string `json:"id,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name string `json:"name,omitempty"`
	// Type - READ-ONLY; Resource type.
	Type string `json:"type,omitempty"`
	// Location - Resource location.
	Location string `json:"location,omitempty"`
	// Tags - Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

//  PublicIPAddressPropertiesFormat Public IPAddress properties.
type PublicIPAddressPropertiesFormat struct {
	// PublicIPAllocationMethod - The public IP allocation method. Possible values are: 'Static' and 'Dynamic'. Possible values include: 'Static', 'Dynamic'
	PublicIPAllocationMethod IPAllocationMethod `json:"publicIPAllocationMethod"`
	IPAddressSkuName PublicIPAddressSkuName `json:"iPAddressSkuName,omitempty"`
	// PublicIPAddressVersion - The public IP address version. Possible values include: 'IPv4', 'IPv6'
	PublicIPAddressVersion IPVersion `json:"publicIPAddressVersion,omitempty"`
	// DNSSettings - The FQDN of the DNS record associated with the public IP address.
	DNSSettings *PublicIPAddressDNSSettings `json:"dnsSettings,omitempty"`
	IPAddress   *string                     `json:"ipAddress,omitempty"`
	// IdleTimeoutInMinutes - The idle timeout of the public IP address.
	IdleTimeoutInMinutes *int32 `json:"idleTimeoutInMinutes,omitempty"`
	// ResourceGUID - The resource GUID property of the public IP resource.
	//ResourceGUID *string `json:"resourceGuid,omitempty"`
	// ProvisioningState - The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// IPAllocationMethod enumerates the values for ip allocation method.
type IPAllocationMethod string

const (
	// Dynamic ...
	Dynamic IPAllocationMethod = "Dynamic"
	// Static ...
	Static IPAllocationMethod = "Static"
)

// PublicIPAddressDNSSettings contains FQDN of the DNS record associated with the public IP address
type PublicIPAddressDNSSettings struct {
	// DomainNameLabel - Gets or sets the Domain name label.The concatenation of the domain name label and the regionalized DNS zone make up the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel *string `json:"domainNameLabel,omitempty"`
	// Fqdn - Gets the FQDN, Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
	Fqdn *string `json:"fqdn,omitempty"`
	// ReverseFqdn - Gets or Sets the Reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn *string `json:"reverseFqdn,omitempty"`
}
// IPVersion enumerates the values for ip version.
type IPVersion string

const (
	// IPv4 ...
	IPv4 IPVersion = "IPv4"
	// IPv6 ...
	IPv6 IPVersion = "IPv6"
)


// PublicIPAddressStatus represents the observed state of a PublicIPAddress.
type PublicIPAddressStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`

	// State of this PublicIPAddress.
	State string `json:"state,omitempty"`

	// A Message providing detail about the state of this PublicIPAddress, if any.
	Message string `json:"message,omitempty"`

	// Etag - A unique string that changes whenever the resource is updated.
	Etag string `json:"etag,omitempty"`

	// ID of this PublicIPAddress.
	ID string `json:"id,omitempty"`

	// Purpose - A string identifying the intention of use for this subnet based
	// on delegations and other user-defined properties.
	Purpose string `json:"purpose,omitempty"`

	// Type of this PublicIPAddress.
	Type string `json:"type,omitempty"`

	// ResourceGUID - The GUID of this PublicIPAddress.
	ResourceGUID string `json:"resourceGuid,omitempty"`
}

// +kubebuilder:object:root=true

// PublicIPAddressList contains a list of PublicIPAddress items
type PublicIPAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicIPAddress `json:"items"`
}

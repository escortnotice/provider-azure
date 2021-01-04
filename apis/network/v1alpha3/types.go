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

	// Etag - A unique string that changes whenever the resource is
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

//Network Security Group structs
// SecurityRuleProtocol enumerates the values for security rule protocol.
type SecurityRuleProtocol string

// ApplicationSecurityGroupPropertiesFormat application security group properties.
/*type ApplicationSecurityGroupPropertiesFormat struct {
	// ResourceGUID - READ-ONLY; The resource GUID property of the application security group resource. It uniquely identifies a resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
	ResourceGUID string `json:"resourceGuid,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioning state of the application security group resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState string `json:"provisioningState,omitempty"`
}*/

// ApplicationSecurityGroup an application security group in a resource group used in Network Security Group.
/*type ApplicationSecurityGroup struct {
	// ApplicationSecurityGroupPropertiesFormat - Properties of the application security group.
	Properties ApplicationSecurityGroupPropertiesFormat `json:"properties,omitempty"`
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
}*/

// SecurityRuleAccess enumerates the values for security rule access.
type SecurityRuleAccess string

// SecurityRuleDirection enumerates the values for security rule direction.
type SecurityRuleDirection string

// SecurityRulePropertiesFormat security rule resource.
type SecurityRulePropertiesFormat struct {
	// Description - A description for this rule. Restricted to 140 chars.
	Description string `json:"description,omitempty"`
	// Protocol - Network protocol this rule applies to. Possible values include: 'SecurityRuleProtocolTCP', 'SecurityRuleProtocolUDP', 'SecurityRuleProtocolIcmp', 'SecurityRuleProtocolEsp', 'SecurityRuleProtocolAsterisk'
	Protocol SecurityRuleProtocol `json:"protocol,omitempty"`
	// SourcePortRange - The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	SourcePortRange string `json:"sourcePortRange,omitempty"`
	// DestinationPortRange - The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
	DestinationPortRange string `json:"destinationPortRange,omitempty"`
	// SourceAddressPrefix - The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from.
	SourceAddressPrefix string `json:"sourceAddressPrefix,omitempty"`
	// SourceAddressPrefixes - The CIDR or source IP ranges.
	SourceAddressPrefixes []string `json:"sourceAddressPrefixes,omitempty"`
	// SourceApplicationSecurityGroups - The application security group specified as source.
	SourceApplicationSecurityGroups []ApplicationSecurityGroup `json:"sourceApplicationSecurityGroups,omitempty"`
	// DestinationAddressPrefix - The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix string `json:"destinationAddressPrefix,omitempty"`
	// DestinationAddressPrefixes - The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes []string `json:"destinationAddressPrefixes,omitempty"`
	// DestinationApplicationSecurityGroups - The application security group specified as destination.
	DestinationApplicationSecurityGroups []ApplicationSecurityGroup `json:"destinationApplicationSecurityGroups,omitempty"`
	// SourcePortRanges - The source port ranges.
	SourcePortRanges []string `json:"sourcePortRanges,omitempty"`
	// DestinationPortRanges - The destination port ranges.
	DestinationPortRanges []string `json:"destinationPortRanges,omitempty"`
	// Access - The network traffic is allowed or denied. Possible values include: 'SecurityRuleAccessAllow', 'SecurityRuleAccessDeny'
	Access SecurityRuleAccess `json:"access,omitempty"`
	// Priority - The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority int32 `json:"priority,omitempty"`
	// Direction - The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values include: 'SecurityRuleDirectionInbound', 'SecurityRuleDirectionOutbound'
	Direction SecurityRuleDirection `json:"direction,omitempty"`
	// ProvisioningState - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState string `json:"provisioningState,omitempty"`
}

// +kubebuilder:object:root=true
// SecurityRule network security rule.
// +kubebuilder:printcolumn:name="RECLAIM-POLICY",type="string",JSONPath=".spec.reclaimPolicy"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// SecurityRulePropertiesFormat - Properties of the security rule.
	Properties SecurityRulePropertiesFormat `json:"properties,omitempty"`
	// Name - The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `json:"name,omitempty"`
	// Etag - A unique read-only string that changes whenever the resource is updated.
	Etag string `json:"etag,omitempty"`
	// ID - Resource ID.
	ID string `json:"id,omitempty"`
}

// A SecurityGroupSpec defines the desired state of a SecurityGroup.
type SecurityGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`

	// ResourceGroupName - Name of the SecurityGroup's resource group.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ResourceGroupNameRef - A reference to the the SecurityGroup's resource
	// group.
	ResourceGroupNameRef *runtimev1alpha1.Reference `json:"resourceGroupNameRef,omitempty"`

	// ResourceGroupNameSelector - Select a reference to the the Security
	// group's resource group.
	ResourceGroupNameSelector *runtimev1alpha1.Selector `json:"resourceGroupNameSelector,omitempty"`

	// Location - Resource location.
	Location string `json:"location"`

	//SecurityGroPropertiesFormat - Properties of security group
	SecurityGroupPropertiesFormat `json:"properties,omitempty"`

	// Tags - Resource tags.
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

// A SecurityGroupStatus represents the observed status of a SecurityGroup.
type SecurityGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`

	// State of this SecurityGroup.
	State string `json:"state,omitempty"`

	// A Message providing detail about the state of this SecurityGroup, if
	// any.
	Message string `json:"message,omitempty"`

	// ID of this SecurityGroup.
	ID string `json:"id,omitempty"`

	// Etag - A unique read-only string that changes whenever the resource is
	// updated.
	Etag string `json:"etag,omitempty"`

	// ResourceGUID - The GUID of this SecurityGroup.
	ResourceGUID string `json:"resourceGuid,omitempty"`

	// Type of this SecurityGroup.
	Type string `json:"type,omitempty"`
}

// SecurityGroupPropertiesFormat network Security Group resource.
type SecurityGroupPropertiesFormat struct {
	// SecurityRules - A collection of security rules of the network security group.
	SecurityRules *[]SecurityRule `json:"securityRules,omitempty"`
	// DefaultSecurityRules - The default security rules of network security group.
	DefaultSecurityRules *[]SecurityRule `json:"defaultSecurityRules,omitempty"`
	// NetworkInterfaces - READ-ONLY; A collection of references to network interfaces.
	//NetworkInterfaces *[]Interface `json:"networkInterfaces,omitempty"`
	// Subnets - READ-ONLY; A collection of references to subnets.
	//Subnets *[]Subnet `json:"subnets,omitempty"`
	// ResourceGUID - The resource GUID property of the network security group resource.
	ResourceGUID *string `json:"resourceGuid,omitempty"`
	// ProvisioningState - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// +kubebuilder:object:root=true
// A SecurityGroup is a managed resource that represents an Azure Security
// Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.location"
// +kubebuilder:printcolumn:name="RECLAIM-POLICY",type="string",JSONPath=".spec.reclaimPolicy"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityGroupSpec   `json:"spec"`
	Status SecurityGroupStatus `json:"status,omitempty"`
	///Properties SecurityGroupPropertiesFormat   `json:"properties,omitempty"`
}

// +kubebuilder:object:root=true
// SecurityGroupList contains a list of Security Groups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

//Azure Firewall Structs
// +kubebuilder:object:root=true
// A AzureFirewall is a managed resource that represents an Azure Firewall
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AzureFirewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureFirewallSpec   `json:"spec"`
	Status AzureFirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// AzureFirewallList contains a list of Security Groups
type AzureFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureFirewall `json:"items"`
}

// A AzureFirewallSpec defines the desired state of a AzureFirewall.
type AzureFirewallSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`

	// ResourceGroupName - Name of the Azure Firewall's resource group.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ResourceGroupNameRef - A reference to the the Azure Firewall's resource
	// group.
	ResourceGroupNameRef *runtimev1alpha1.Reference `json:"resourceGroupNameRef,omitempty"`

	// ResourceGroupNameSelector - Select a reference to the the Azure Firewall
	// resource group.
	ResourceGroupNameSelector *runtimev1alpha1.Selector `json:"resourceGroupNameSelector,omitempty"`

	// Location - Resource location.
	Location string `json:"location"`

	// AzureFirewallPropertiesFormat - Properties of AzureFirewall
	AzureFirewallPropertiesFormat `json:"properties,omitempty"`

	// Tags - Resource tags.
	// +optional
	Tags map[string]string `json:"tags,omitempty"`

	// Zones - A list of availability zones denoting where the resource needs to come from.
	Zones []string `json:"zones,omitempty"`

	// Etag - Gets a unique string that changes whenever the resource is updated.
	Etag string `json:"etag,omitempty"`

	// ID - Resource ID.
	ID string `json:"id,omitempty"`

	// Name - Resource name.
	Name string `json:"name,omitempty"`

	// Type - Resource type.
	Type string `json:"type,omitempty"`
}

// AzureFirewallPropertiesFormat properties of the Azure Firewall.
type AzureFirewallPropertiesFormat struct {
	// ApplicationRuleCollections - Collection of application rule collections used by Azure Firewall.
	ApplicationRuleCollections *[]AzureFirewallApplicationRuleCollection `json:"applicationRuleCollections,omitempty"`
	// NatRuleCollections - Collection of NAT rule collections used by Azure Firewall.
	NatRuleCollections *[]AzureFirewallNatRuleCollection `json:"natRuleCollections,omitempty"`
	// NetworkRuleCollections - Collection of network rule collections used by Azure Firewall.
	NetworkRuleCollections *[]AzureFirewallNetworkRuleCollection `json:"networkRuleCollections,omitempty"`
	// IPConfigurations - IP configuration of the Azure Firewall resource.
	IPConfigurations *[]AzureFirewallIPConfiguration `json:"ipConfigurations,omitempty"`
	// ProvisioningState - The provisioning state of the resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'
	ProvisioningState string `json:"provisioningState,omitempty"`
	// ThreatIntelMode - The operation mode for Threat Intelligence. Possible values include: 'AzureFirewallThreatIntelModeAlert', 'AzureFirewallThreatIntelModeDeny', 'AzureFirewallThreatIntelModeOff'
	ThreatIntelMode string `json:"threatIntelMode,omitempty"`
	// VirtualHub - The virtualHub to which the firewall belongs.
	VirtualHub *SubResource `json:"virtualHub,omitempty"`
	// FirewallPolicy - The firewallPolicy associated with this azure firewall.
	FirewallPolicy *SubResource `json:"firewallPolicy,omitempty"`
	// HubIPAddresses - IP addresses associated with AzureFirewall.
	HubIPAddresses *HubIPAddresses `json:"hubIpAddresses,omitempty"`
}

// AzureFirewallApplicationRuleCollection application rule collection resource.
type AzureFirewallApplicationRuleCollection struct {
	// AzureFirewallApplicationRuleCollectionPropertiesFormat - Properties of the azure firewall application rule collection.
	Properties AzureFirewallApplicationRuleCollectionPropertiesFormat `json:"properties,omitempty"`
	// Name - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `json:"name,omitempty"`
	// Etag - READ-ONLY; Gets a unique read-only string that changes whenever the resource is updated.
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
  
// AzureFirewallApplicationRuleCollectionPropertiesFormat properties of the application rule collection.
type AzureFirewallApplicationRuleCollectionPropertiesFormat struct {
	// Priority - Priority of the application rule collection resource.
	Priority int32 `json:"priority,omitempty"`
	// Action - The action type of a rule collection.
	Action string `json:"action,omitempty"`
	// Rules - Collection of rules used by a application rule collection.
	Rules []AzureFirewallApplicationRule `json:"rules,omitempty"`
	// ProvisioningState - The provisioning state of the resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'
	ProvisioningState string `json:"provisioningState,omitempty"`
}

// AzureFirewallApplicationRule properties of an application rule.
type AzureFirewallApplicationRule struct {
	// Name - Name of the application rule.
	Name string `json:"name,omitempty"`
	// Description - Description of the rule.
	Description string `json:"description,omitempty"`
	// SourceAddresses - List of source IP addresses for this rule.
	SourceAddresses []string `json:"sourceAddresses,omitempty"`
	// Protocols - Array of ApplicationRuleProtocols.
	Protocols []AzureFirewallApplicationRuleProtocol `json:"protocols,omitempty"`
	// TargetFqdns - List of FQDNs for this rule.
	TargetFqdns []string `json:"targetFqdns,omitempty"`
	// FqdnTags - List of FQDN Tags for this rule.
	FqdnTags []string `json:"fqdnTags,omitempty"`
}

// AzureFirewallApplicationRuleProtocol properties of the application rule protocol.
type AzureFirewallApplicationRuleProtocol struct {
	// ProtocolType - Protocol type. Possible values include: 'AzureFirewallApplicationRuleProtocolTypeHTTP', 'AzureFirewallApplicationRuleProtocolTypeHTTPS'
	ProtocolType string `json:"protocolType,omitempty"`
	// Port - Port number for the protocol, cannot be greater than 64000. This field is optional.
	Port int32 `json:"port,omitempty"`
}

// AzureFirewallIPConfiguration IP configuration of an Azure Firewall.
type AzureFirewallIPConfiguration struct {
	// AzureFirewallIPConfigurationPropertiesFormat - Properties of the azure firewall IP configuration.
	AzureFirewallIPConfigurationPropertiesFormat AzureFirewallIPConfigurationPropertiesFormat `json:"properties,omitempty"`
	// Name - Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`
	// Etag - A unique string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`
	// ID - Resource ID.
	ID *string `json:"id,omitempty"`
}

// AzureFirewallIPConfigurationPropertiesFormat properties of IP configuration of an Azure Firewall.
type AzureFirewallIPConfigurationPropertiesFormat struct {
	// PrivateIPAddress - The Firewall Internal Load Balancer IP to be used as the next hop in User Defined Routes.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`
	// Subnet - Reference of the subnet resource. This resource must be named 'AzureFirewallSubnet'.
	Subnet *SubResource `json:"subnet,omitempty"`
	// PublicIPAddress - Reference of the PublicIP resource. This field is a mandatory input if subnet is not null.
	PublicIPAddress *SubResource `json:"publicIPAddress,omitempty"`
	// ProvisioningState - The provisioning state of the resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

// SubResource reference to another subresource.
type SubResource struct {
	// ID - Resource ID.
	ID string `json:"id,omitempty"`
}

// HubIPAddresses IP addresses associated with azure firewall.
type HubIPAddresses struct {
	// PublicIPAddresses - List of Public IP addresses associated with azure firewall.
	PublicIPAddresses []AzureFirewallPublicIPAddress `json:"publicIPAddresses,omitempty"`
	// PrivateIPAddress - Private IP Address associated with azure firewall.
	PrivateIPAddress string `json:"privateIPAddress,omitempty"`
}

// AzureFirewallPublicIPAddress public IP Address associated with azure firewall.
type AzureFirewallPublicIPAddress struct {
	// Address - Public IP Address value.
	Address *string `json:"address,omitempty"`
}

// A AzureFirewallStatus represents the observed status of a AzureFirewall.
type AzureFirewallStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`

	// A Message providing detail about the state of this AzureFirewall, if
	// any.
	Message string `json:"message,omitempty"`

	// ID of this AzureFirewall.
	ID string `json:"id,omitempty"`

	// Etag - A unique string that changes whenever the resource is
	// updated.
	Etag string `json:"etag,omitempty"`

	// ResourceGUID - The GUID of this AzureFirewall.
	ResourceGUID string `json:"resourceGuid,omitempty"`

	// Type of this AzureFirewall.
	Type string `json:"type,omitempty"`

	// State of this VirtualNetwork.
	State string `json:"state,omitempty"`
}

//Rules Structs
// AzureFirewallNatRule properties of a NAT rule.
type AzureFirewallNatRule struct {
	// Name - Name of the NAT rule.
	Name string `json:"name,omitempty"`
	// Description - Description of the rule.
	Description string `json:"description,omitempty"`
	// SourceAddresses - List of source IP addresses for this rule.
	SourceAddresses []string `json:"sourceAddresses,omitempty"`
	// DestinationAddresses - List of destination IP addresses for this rule. Supports IP ranges, prefixes, and service tags.
	DestinationAddresses []string `json:"destinationAddresses,omitempty"`
	// DestinationPorts - List of destination ports.
	DestinationPorts []string `json:"destinationPorts,omitempty"`
	// Protocols - Array of AzureFirewallNetworkRuleProtocols applicable to this NAT rule.
	Protocols []string `json:"protocols,omitempty"`
	// TranslatedAddress - The translated address for this NAT rule.
	TranslatedAddress string `json:"translatedAddress,omitempty"`
	// TranslatedPort - The translated port for this NAT rule.
	TranslatedPort string `json:"translatedPort,omitempty"`
}

// AzureFirewallNatRuleCollectionProperties properties of the NAT rule collection.
type AzureFirewallNatRuleCollectionProperties struct {
	// Priority - Priority of the NAT rule collection resource.
	Priority int32 `json:"priority,omitempty"`
	// Action - The action type of a NAT rule collection.
	Action string `json:"action,omitempty"`
	// Rules - Collection of rules used by a NAT rule collection.
	Rules []AzureFirewallNatRule `json:"rules,omitempty"`
	// ProvisioningState - The provisioning state of the resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'
	ProvisioningState string `json:"provisioningState,omitempty"`
}

// AzureFirewallNatRuleCollection NAT rule collection resource.
type AzureFirewallNatRuleCollection struct {
	// AzureFirewallNatRuleCollectionProperties - Properties of the azure firewall NAT rule collection.
	Properties AzureFirewallNatRuleCollectionProperties `json:"properties,omitempty"`
	// Name - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `json:"name,omitempty"`
	// Etag - Gets a unique string that changes whenever the resource is updated.
	Etag string `json:"etag,omitempty"`
	// ID - Resource ID.
	ID string `json:"id,omitempty"`
}

// AzureFirewallNetworkRuleCollection network rule collection resource.
type AzureFirewallNetworkRuleCollection struct {
	// AzureFirewallNetworkRuleCollectionPropertiesFormat - Properties of the azure firewall network rule collection.
	Properties AzureFirewallNetworkRuleCollectionPropertiesFormat `json:"properties,omitempty"`
	// Name - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `json:"name,omitempty"`
	// Etag - Gets a unique string that changes whenever the resource is updated.
	Etag string `json:"etag,omitempty"`
	// ID - Resource ID.
	ID string `json:"id,omitempty"`
}

// AzureFirewallNetworkRuleCollectionPropertiesFormat properties of the network rule collection.
type AzureFirewallNetworkRuleCollectionPropertiesFormat struct {
	// Priority - Priority of the network rule collection resource.
	Priority int32 `json:"priority,omitempty"`
	// Action - The action type of a rule collection.
	Action string `json:"action,omitempty"`
	// Rules - Collection of rules used by a network rule collection.
	Rules []AzureFirewallNetworkRule `json:"rules,omitempty"`
	// ProvisioningState - The provisioning state of the resource. Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'
	ProvisioningState string `json:"provisioningState,omitempty"`
}

// AzureFirewallNetworkRule properties of the network rule.
type AzureFirewallNetworkRule struct {
	// Name - Name of the network rule.
	Name string `json:"name,omitempty"`
	// Description - Description of the rule.
	Description string `json:"description,omitempty"`
	// Protocols - Array of AzureFirewallNetworkRuleProtocols.
	Protocols []string `json:"protocols,omitempty"`
	// SourceAddresses - List of source IP addresses for this rule.
	SourceAddresses []string `json:"sourceAddresses,omitempty"`
	// DestinationAddresses - List of destination IP addresses.
	DestinationAddresses []string `json:"destinationAddresses,omitempty"`
	// DestinationPorts - List of destination ports.
	DestinationPorts []string `json:"destinationPorts,omitempty"`
}

// +kubebuilder:object:root=true
// A ApplicationSecurityGroup is a managed resource that represents an Azure ApplicationSecurityGroup
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="LOCATION",type="string",JSONPath=".spec.location"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationSecurityGroupSpec   `json:"spec"`
	Status ApplicationSecurityGroupStatus `json:"status,omitempty"`

}

// ApplicationSecurityGroupStatus represents the observed state of a ApplicationSecurityGroup.
type ApplicationSecurityGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	// State of this ApplicationSecurityGroup.
	State string `json:"state,omitempty"`
	// A Message providing detail about the state of this ApplicationSecurityGroup, if any.
	Message string `json:"message,omitempty"`

	// Etag - A unique string that changes whenever the resource is updated.
	Etag string `json:"etag,omitempty"`
  	// ID of this ApplicationSecurityGroup.
	ID string `json:"id,omitempty"`
  	// Purpose - A string identifying the intention of use for this subnet based
	// on delegations and other user-defined properties.
	Purpose string `json:"purpose,omitempty"`
  	// Type of this ApplicationSecurityGroup.
	Type string `json:"type,omitempty"`
}

// ApplicationSecurityGroup contains a list of ApplicationSecurityGroup items
type ApplicationSecurityGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`

	// ResourceGroupName - Name of the ApplicationSecurityGroup's resource group.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ResourceGroupNameRef - A reference to the the ApplicationSecurityGroup's resource
	// group.
	ResourceGroupNameRef *runtimev1alpha1.Reference `json:"resourceGroupNameRef,omitempty"`

	// ApplicationSecurityGroupPropertiesFormat - Properties of the application security group.
	Properties ApplicationSecurityGroupPropertiesFormat `json:"properties,omitempty"`
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

//  ApplicationSecurityGroupPropertiesFormat application security group properties.
type ApplicationSecurityGroupPropertiesFormat struct {
	// ResourceGUID - READ-ONLY; The resource GUID property of the application security group resource. It uniquely identifies a resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
	ResourceGUID *string `json:"resourceGuid,omitempty"`
	// ProvisioningState - READ-ONLY; The provisioning state of the application security group resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}
// +kubebuilder:object:root=true
  // ApplicationSecurityGroupList contains a list of ApplicationSecurityGroup items
type ApplicationSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationSecurityGroup `json:"items"`
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

type PublicIPAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   PublicIPAddressSpec   `json:"spec"`
	Status PublicIPAddressStatus `json:"status,omitempty"`
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

	// PublicIPAddressPropertiesFormat - Properties of the public ip address.
	Properties PublicIPAddressPropertiesFormat `json:"properties,omitempty"`
	// Name - READ-ONLY; Resource name.
	Name string `json:"name,omitempty"`
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
	// IdleTimeoutInMinutes - The idle timeout of the public IP address.
	IdleTimeoutInMinutes int `json:"idleTimeoutInMinutes,omitempty"`
	// ProvisioningState - The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState string `json:"provisioningState,omitempty"`
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

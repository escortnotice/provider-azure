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
	"reflect"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

// Package type metadata.
const (
	Group   = "network.azure.crossplane.io"
	Version = "v1alpha3"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

// VirtualNetwork type metadata.
var (
	VirtualNetworkKind             = reflect.TypeOf(VirtualNetwork{}).Name()
	VirtualNetworkGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualNetworkKind}.String()
	VirtualNetworkKindAPIVersion   = VirtualNetworkKind + "." + SchemeGroupVersion.String()
	VirtualNetworkGroupVersionKind = SchemeGroupVersion.WithKind(VirtualNetworkKind)
)

// Subnet type metadata.
var (
	SubnetKind             = reflect.TypeOf(Subnet{}).Name()
	SubnetGroupKind        = schema.GroupKind{Group: Group, Kind: SubnetKind}.String()
	SubnetKindAPIVersion   = SubnetKind + "." + SchemeGroupVersion.String()
	SubnetGroupVersionKind = SchemeGroupVersion.WithKind(SubnetKind)
)


// RouteTable type metadata.
var (
	RouteTableKind             = reflect.TypeOf(RouteTable{}).Name()
	RouteTableGroupKind        = schema.GroupKind{Group: Group, Kind: RouteTableKind}.String()
	RouteTableKindAPIVersion   = RouteTableKind + "." + SchemeGroupVersion.String()
	RouteTableGroupVersionKind = SchemeGroupVersion.WithKind(RouteTableKind)

// PublicIPAddress type metadata.
var (
	PublicIPAddressKind             = reflect.TypeOf(PublicIPAddress{}).Name()
	PublicIPAddressGroupKind        = schema.GroupKind{Group: Group, Kind: PublicIPAddressKind}.String()
	PublicIPAddressKindAPIVersion   = PublicIPAddressKind + "." + SchemeGroupVersion.String()
	PublicIPAddressGroupVersionKind = SchemeGroupVersion.WithKind(PublicIPAddressKind)
)

// SecurityGroup type Metadata.
var (
	SecurityGroupKind             = reflect.TypeOf(SecurityGroup{}).Name()
	SecurityGroupGroupKind        = schema.GroupKind{Group: Group, Kind: SecurityGroupKind}.String()
	SecurityGroupKindAPIVersion   = SecurityGroupKind + "." + SchemeGroupVersion.String()
	SecurityGroupGroupVersionKind = SchemeGroupVersion.WithKind(SecurityGroupKind)
)

// AzureFirewall type Metadata.
var (
	AzureFirewallKind             = reflect.TypeOf(AzureFirewall{}).Name()
	AzureFirewallGroupKind        = schema.GroupKind{Group: Group, Kind: AzureFirewallKind}.String()
	AzureFirewallKindAPIVersion   = AzureFirewallKind + "." + SchemeGroupVersion.String()
	AzureFirewallGroupVersionKind = SchemeGroupVersion.WithKind(AzureFirewallKind)
)
// ApplicationSecurityGroup type metadata.
var (
	ApplicationSecurityGroupKind             = reflect.TypeOf(ApplicationSecurityGroup{}).Name()
	ApplicationSecurityGroupGroupKind        = schema.GroupKind{Group: Group, Kind: ApplicationSecurityGroupKind}.String()
	ApplicationSecurityGroupKindAPIVersion   = ApplicationSecurityGroupKind + "." + SchemeGroupVersion.String()
	ApplicationSecurityGroupGroupVersionKind = SchemeGroupVersion.WithKind(ApplicationSecurityGroupKind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetwork{}, &VirtualNetworkList{})
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
	SchemeBuilder.Register(&SecurityRule{})
	SchemeBuilder.Register(&AzureFirewall{}, &AzureFirewallList{})
	SchemeBuilder.Register(&ApplicationSecurityGroup{}, &ApplicationSecurityGroupList{})
	SchemeBuilder.Register(&PublicIPAddress{}, &PublicIPAddressList{})
}

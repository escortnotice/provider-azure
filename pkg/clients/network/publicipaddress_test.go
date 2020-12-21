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

package network

import (
	"testing"

	networkmgmt "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network"
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/crossplane/provider-azure/apis/network/v1alpha3"
	azure "github.com/crossplane/provider-azure/pkg/clients"
)

var (
	uidIp      = types.UID("definitely-a-uuid")
	locationIp = "cool-location"
	tagsIp     = map[string]string{"one": "test", "two": "test"}

	idIp               = "a-very-cool-id"
	etagIp             = "a-very-cool-etag"
	resourcetypeIp     = "resource-type"
	domainNameLabel    = "cooldomain"
	fqdn               =  "cooldomain.fqdn.net"
	name               = "cool-ip"
)

func TestNewPublicIpAddressParameters(t *testing.T) {
	cases := []struct {
		name string
		r    *v1alpha3.PublicIPAddress
		want networkmgmt.PublicIPAddress
	}{
		{
			name: "SuccessfulFull",
			r: &v1alpha3.PublicIPAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:       name,
					UID:        uidIp,
					Finalizers: []string{},
				},
				Spec: v1alpha3.PublicIPAddressSpec{
					Properties: v1alpha3.PublicIPAddressPropertiesFormat{
						IPAddressSkuName : v1alpha3.PublicIPAddressSkuNameBasic,
						PublicIPAddressVersion: v1alpha3.IPv4,
						DNSSettings: &v1alpha3.PublicIPAddressDNSSettings{
							DomainNameLabel:azure.ToStringPtr(domainNameLabel),
							Fqdn: azure.ToStringPtr(fqdn),
						},
						IdleTimeoutInMinutes  : 4,
						PublicIPAllocationMethod   :
							v1alpha3.Static,
					},
					Location: locationIp,
					Tags:     tagsIp,
					Name:     name,
				},
			},
			want: networkmgmt.PublicIPAddress{
				Location: azure.ToStringPtr(locationIp),
				Tags:     azure.ToStringPtrMap(tagsIp),
				Name: azure.ToStringPtr(name),
				Sku: &networkmgmt.PublicIPAddressSku{
					Name: networkmgmt.PublicIPAddressSkuNameBasic,
				},
				PublicIPAddressPropertiesFormat : &networkmgmt.PublicIPAddressPropertiesFormat{
					PublicIPAddressVersion: networkmgmt.IPv4,
					PublicIPAllocationMethod: networkmgmt.Static,
					IdleTimeoutInMinutes: azure.ToInt32Ptr(4),
					DNSSettings: &networkmgmt.PublicIPAddressDNSSettings{
						DomainNameLabel: azure.ToStringPtr(domainNameLabel),
						Fqdn: azure.ToStringPtr(fqdn),
					},


				},
			},
		},
/*		{
			name: "SuccessfulPartial",
			r: nil,
			want: nil,
		},*/
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := NewPublicIpAddressParameters(tc.r)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("NewPublicIpAddressParameters(...): -want, +got\n%s", diff)
			}
		})
	}
}

func TestPublicIpAdressNeedsUpdate(t *testing.T) {
	cases := []struct {
		name string
		kube *v1alpha3.PublicIPAddress
		az   networkmgmt.PublicIPAddress
		want bool
	}{
		{
			name: "NeedsUpdateTags",
			kube: &v1alpha3.PublicIPAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:       name,
					UID:        uidIp,
					Finalizers: []string{},
				},
				Spec: v1alpha3.PublicIPAddressSpec{
					Properties: v1alpha3.PublicIPAddressPropertiesFormat{
						IPAddressSkuName : v1alpha3.PublicIPAddressSkuNameBasic,
						PublicIPAddressVersion: v1alpha3.IPv4,
						DNSSettings: &v1alpha3.PublicIPAddressDNSSettings{
							DomainNameLabel:azure.ToStringPtr(domainNameLabel),
							Fqdn: azure.ToStringPtr(fqdn),
						},
						IdleTimeoutInMinutes  : 4,
						PublicIPAllocationMethod   : v1alpha3.Static,
					},
					Location: locationIp,
					Tags:     map[string]string{"three": "test"},
					Name:     name,
				},
			},
			az : networkmgmt.PublicIPAddress{
				Location: azure.ToStringPtr(locationIp),
				Tags:     azure.ToStringPtrMap(tagsIp),
				Name: azure.ToStringPtr(name),
				Sku: &networkmgmt.PublicIPAddressSku{
					Name: networkmgmt.PublicIPAddressSkuNameBasic,
				},
				PublicIPAddressPropertiesFormat : &networkmgmt.PublicIPAddressPropertiesFormat{
					PublicIPAddressVersion: networkmgmt.IPv4,
					PublicIPAllocationMethod: networkmgmt.Static,
					IdleTimeoutInMinutes: azure.ToInt32Ptr(4),
					DNSSettings: &networkmgmt.PublicIPAddressDNSSettings{
						DomainNameLabel: azure.ToStringPtr(domainNameLabel),
						Fqdn: azure.ToStringPtr(fqdn),
					},


				},
			},
			want: true,
		},
		{
			name: "NeedsUpdateDNSName",
			kube: &v1alpha3.PublicIPAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:       name,
					UID:        uidIp,
					Finalizers: []string{},
				},
				Spec: v1alpha3.PublicIPAddressSpec{
					Properties: v1alpha3.PublicIPAddressPropertiesFormat{
						IPAddressSkuName : v1alpha3.PublicIPAddressSkuNameBasic,
						PublicIPAddressVersion: v1alpha3.IPv4,
						DNSSettings: &v1alpha3.PublicIPAddressDNSSettings{
							DomainNameLabel:azure.ToStringPtr("dns1"),
							Fqdn: azure.ToStringPtr(fqdn),
						},
						IdleTimeoutInMinutes  : 4,
						PublicIPAllocationMethod   : v1alpha3.Static,
					},
					Location: locationIp,
					Tags:     map[string]string(tagsIp),
					Name:     name,
				},
			},
			az : networkmgmt.PublicIPAddress{
				Location: azure.ToStringPtr(locationIp),
				Tags:     azure.ToStringPtrMap(tagsIp),
				Name: azure.ToStringPtr(name),
				Sku: &networkmgmt.PublicIPAddressSku{
					Name: networkmgmt.PublicIPAddressSkuNameBasic,
				},
				PublicIPAddressPropertiesFormat : &networkmgmt.PublicIPAddressPropertiesFormat{
					PublicIPAddressVersion: networkmgmt.IPv4,
					PublicIPAllocationMethod: networkmgmt.Static,
					IdleTimeoutInMinutes: azure.ToInt32Ptr(4),
					DNSSettings: &networkmgmt.PublicIPAddressDNSSettings{
						DomainNameLabel: azure.ToStringPtr(domainNameLabel),
						Fqdn: azure.ToStringPtr(fqdn),
					},


				},
			},
			want: true,
		},
		{
			name: "NeedsUpdateTimeout",
			kube: &v1alpha3.PublicIPAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:       name,
					UID:        uidIp,
					Finalizers: []string{},
				},
				Spec: v1alpha3.PublicIPAddressSpec{
					Properties: v1alpha3.PublicIPAddressPropertiesFormat{
						IPAddressSkuName : v1alpha3.PublicIPAddressSkuNameBasic,
						PublicIPAddressVersion: v1alpha3.IPv4,
						DNSSettings: &v1alpha3.PublicIPAddressDNSSettings{
							DomainNameLabel:azure.ToStringPtr(domainNameLabel),
							Fqdn: azure.ToStringPtr(fqdn),
						},
						IdleTimeoutInMinutes  : 8,
						PublicIPAllocationMethod   : v1alpha3.Static,
					},
					Location: locationIp,
					Tags:     map[string]string(tagsIp),
					Name:     name,
				},
			},
			az : networkmgmt.PublicIPAddress{
				Location: azure.ToStringPtr(locationIp),
				Tags:     azure.ToStringPtrMap(tagsIp),
				Name: azure.ToStringPtr(name),
				Sku: &networkmgmt.PublicIPAddressSku{
					Name: networkmgmt.PublicIPAddressSkuNameBasic,
				},
				PublicIPAddressPropertiesFormat : &networkmgmt.PublicIPAddressPropertiesFormat{
					PublicIPAddressVersion: networkmgmt.IPv4,
					PublicIPAllocationMethod: networkmgmt.Static,
					IdleTimeoutInMinutes: azure.ToInt32Ptr(4),
					DNSSettings: &networkmgmt.PublicIPAddressDNSSettings{
						DomainNameLabel: azure.ToStringPtr(domainNameLabel),
						Fqdn: azure.ToStringPtr(fqdn),
					},


				},
			},
			want: true,
		},
		{
			name: "NeedsUpdateAllocationMethod",
			kube: &v1alpha3.PublicIPAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:       name,
					UID:        uidIp,
					Finalizers: []string{},
				},
				Spec: v1alpha3.PublicIPAddressSpec{
					Properties: v1alpha3.PublicIPAddressPropertiesFormat{
						IPAddressSkuName : v1alpha3.PublicIPAddressSkuNameBasic,
						PublicIPAddressVersion: v1alpha3.IPv6,
						DNSSettings: &v1alpha3.PublicIPAddressDNSSettings{
							DomainNameLabel:azure.ToStringPtr(domainNameLabel),
							Fqdn: azure.ToStringPtr(fqdn),
						},
						IdleTimeoutInMinutes  : 4,
						PublicIPAllocationMethod   : v1alpha3.Static,
					},
					Location: locationIp,
					Tags:     map[string]string{"three": "test"},
					Name:     name,
				},
			},
			az : networkmgmt.PublicIPAddress{
				Location: azure.ToStringPtr(locationIp),
				Tags:     azure.ToStringPtrMap(tagsIp),
				Name: azure.ToStringPtr(name),
				Sku: &networkmgmt.PublicIPAddressSku{
					Name: networkmgmt.PublicIPAddressSkuNameBasic,
				},
				PublicIPAddressPropertiesFormat : &networkmgmt.PublicIPAddressPropertiesFormat{
					PublicIPAddressVersion: networkmgmt.IPv4,
					PublicIPAllocationMethod: networkmgmt.Static,
					IdleTimeoutInMinutes: azure.ToInt32Ptr(4),
					DNSSettings: &networkmgmt.PublicIPAddressDNSSettings{
						DomainNameLabel: azure.ToStringPtr(domainNameLabel),
						Fqdn: azure.ToStringPtr(fqdn),
					},


				},
			},
			want: true,
		},
		{
			name: "NoUpdate",
			kube: &v1alpha3.PublicIPAddress{
				ObjectMeta: metav1.ObjectMeta{
					Name:       name,
					UID:        uidIp,
					Finalizers: []string{},
				},
				Spec: v1alpha3.PublicIPAddressSpec{
					Properties: v1alpha3.PublicIPAddressPropertiesFormat{
						IPAddressSkuName : v1alpha3.PublicIPAddressSkuNameBasic,
						PublicIPAddressVersion: v1alpha3.IPv4,
						DNSSettings: &v1alpha3.PublicIPAddressDNSSettings{
							DomainNameLabel:azure.ToStringPtr(domainNameLabel),
							Fqdn: azure.ToStringPtr(fqdn),
						},
						IdleTimeoutInMinutes  : 4,
						PublicIPAllocationMethod   : v1alpha3.Static,
					},
					Location: locationIp,
					Tags:     tagsIp,
					Name:     name,
				},
			},
			az: networkmgmt.PublicIPAddress{
				Location: azure.ToStringPtr(locationIp),
				Tags:     azure.ToStringPtrMap(tagsIp),
				Name: azure.ToStringPtr(name),
				Sku: &networkmgmt.PublicIPAddressSku{
					Name: networkmgmt.PublicIPAddressSkuNameBasic,
				},
				PublicIPAddressPropertiesFormat : &networkmgmt.PublicIPAddressPropertiesFormat{
					PublicIPAddressVersion: networkmgmt.IPv4,
					PublicIPAllocationMethod: networkmgmt.Static,
					IdleTimeoutInMinutes: azure.ToInt32Ptr(4),
					DNSSettings: &networkmgmt.PublicIPAddressDNSSettings{
						DomainNameLabel: azure.ToStringPtr(domainNameLabel),
						Fqdn: azure.ToStringPtr(fqdn),
					},


				},
			},
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := PublicIpAdressNeedsUpdate(tc.kube, tc.az)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("PublicIpAdressNeedsUpdate(...): -want, +got\n%s", diff)
			}
		})
	}
}

func TestUpdatePublicIpAddressStatusFromAzure(t *testing.T) {
	mockCondition := runtimev1alpha1.Condition{Message: "mockMessage"}
	resourceStatus := runtimev1alpha1.ResourceStatus{
		ConditionedStatus: runtimev1alpha1.ConditionedStatus{
			Conditions: []runtimev1alpha1.Condition{mockCondition},
		},
	}

	cases := []struct {
		name string
		r    networkmgmt.PublicIPAddress
		want v1alpha3.PublicIPAddressStatus
	}{
		{
			name: "SuccessfulFull",
			r: networkmgmt.PublicIPAddress{
				Location: azure.ToStringPtr(locationIp),
				Etag:     azure.ToStringPtr(etagIp),
				ID:       azure.ToStringPtr(idIp),
				Type:     azure.ToStringPtr(resourcetypeIp),
				Tags:     azure.ToStringPtrMap(nil),
				PublicIPAddressPropertiesFormat : &networkmgmt.PublicIPAddressPropertiesFormat{
					PublicIPAddressVersion: networkmgmt.IPv4,
					PublicIPAllocationMethod: networkmgmt.Static,
					IdleTimeoutInMinutes: azure.ToInt32Ptr(4),
					DNSSettings: &networkmgmt.PublicIPAddressDNSSettings{
						DomainNameLabel: azure.ToStringPtr(domainNameLabel),
						Fqdn: azure.ToStringPtr(fqdn),
					},
					ProvisioningState: azure.ToStringPtr("Succeeded"),
					ResourceGUID:      azure.ToStringPtr(string(uidIp)),

				},
			},
			want: v1alpha3.PublicIPAddressStatus{
					State:        string(networkmgmt.Succeeded),
					ID:           idIp,
					Etag:         etagIp,
					Type:         resourcetypeIp,
					ResourceGUID: string(uidIp),
				},
			},
		{
			name: "SuccessfulPartial",
			r: networkmgmt.PublicIPAddress{
				Location: azure.ToStringPtr(locationIp),
				Type:     azure.ToStringPtr(resourcetypeIp),
				Tags:     azure.ToStringPtrMap(nil),
				PublicIPAddressPropertiesFormat : &networkmgmt.PublicIPAddressPropertiesFormat{
					PublicIPAddressVersion: networkmgmt.IPv4,
					PublicIPAllocationMethod: networkmgmt.Static,
					IdleTimeoutInMinutes: azure.ToInt32Ptr(4),
					ProvisioningState: azure.ToStringPtr("Succeeded"),
					ResourceGUID:      azure.ToStringPtr(string(uidIp)),

				},
			},
			want: v1alpha3.PublicIPAddressStatus{
				State:        string(networkmgmt.Succeeded),
				ResourceGUID: string(uidIp),
				Type:         resourcetypeIp,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			v := &v1alpha3.PublicIPAddress{
				Status: v1alpha3.PublicIPAddressStatus{
					ResourceStatus: resourceStatus,
				},
			}

			UpdatePublicIpAddressStatusFromAzure(v, tc.r)

			// make sure that internal resource status hasn't changed
			if diff := cmp.Diff(mockCondition, v.Status.ResourceStatus.Conditions[0]); diff != "" {
				t.Errorf("UpdatePublicIpAddressStatusFromAzure(...): -want, +got\n%s", diff)
			}

			// make sure that other resource parameters are updated
			tc.want.ResourceStatus = resourceStatus
			if diff := cmp.Diff(tc.want, v.Status); diff != "" {
				t.Errorf("UpdatePublicIpAddressStatusFromAzure(...): -want, +got\n%s", diff)
			}
		})
	}
}


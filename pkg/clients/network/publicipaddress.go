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
	"reflect"

	networkmgmt "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network"
	"github.com/crossplane/provider-azure/apis/network/v1alpha3"
	azure "github.com/crossplane/provider-azure/pkg/clients"
)

// UpdatePublicIpAddressStatusFromAzure updates the status related to the external
// Azure Public IP in the PublicIpAddress Status
func UpdatePublicIpAddressStatusFromAzure(v *v1alpha3.PublicIPAddress, az networkmgmt.PublicIPAddress) {
	v.Status.State = azure.ToString(az.ProvisioningState)
	v.Status.ID = azure.ToString(az.ID)
	v.Status.Etag = azure.ToString(az.Etag)
	v.Status.ResourceGUID = azure.ToString(az.ResourceGUID)
	v.Status.Type = azure.ToString(az.Type)
}

// NewPublicIpAddressParameters returns an Azure public ip address object from a Public Ip Address Spec
func NewPublicIpAddressParameters(pub *v1alpha3.PublicIPAddress) networkmgmt.PublicIPAddress {
	return networkmgmt.PublicIPAddress{
		Name:     azure.ToStringPtr(pub.Spec.Name),
		Type:     azure.ToStringPtr(pub.Spec.Type),
		Location: azure.ToStringPtr(pub.Spec.Location),
		Tags: azure.ToStringPtrMap(pub.Spec.Tags),

		PublicIPAddressPropertiesFormat: &networkmgmt.PublicIPAddressPropertiesFormat{
			ProvisioningState:        azure.ToStringPtr(pub.Spec.Properties.ProvisioningState),
			IdleTimeoutInMinutes:     azure.ToInt32Ptr(pub.Spec.Properties.IdleTimeoutInMinutes),
			PublicIPAddressVersion:   networkmgmt.IPVersion(pub.Spec.Properties.PublicIPAddressVersion),
			PublicIPAllocationMethod: networkmgmt.IPAllocationMethod(pub.Spec.Properties.PublicIPAllocationMethod),
			DNSSettings:              setDNS(pub.Spec.Properties.DNSSettings),
		},
		Sku: setSKU(pub.Spec.Properties.IPAddressSkuName),
	}
}

func setSKU(skuName v1alpha3.PublicIPAddressSkuName) *networkmgmt.PublicIPAddressSku {
	if skuName=="" {
		return &networkmgmt.PublicIPAddressSku{
			Name: networkmgmt.PublicIPAddressSkuName("Basic"),
		}
	}
			return &networkmgmt.PublicIPAddressSku{
				Name:  networkmgmt.PublicIPAddressSkuName(skuName),

		}
	}

func setDNS(settings *v1alpha3.PublicIPAddressDNSSettings) *networkmgmt.PublicIPAddressDNSSettings {
	if nil!=settings{
		return &networkmgmt.PublicIPAddressDNSSettings{
			DomainNameLabel: settings.DomainNameLabel,
			Fqdn: settings.Fqdn,
		}
	}
	return nil
}


// PublicIpAdressNeedsUpdate determines if a PublicIpAdress need to be updated
func PublicIpAdressNeedsUpdate(sg *v1alpha3.PublicIPAddress, az networkmgmt.PublicIPAddress) bool {

	if !reflect.DeepEqual(azure.ToStringPtr(sg.Name),az.Name){
		return true
	}
	if !reflect.DeepEqual(azure.ToStringPtr(sg.Spec.Location) ,az.Location){
		return true
	}
	if !reflect.DeepEqual(azure.ToStringPtrMap(sg.Spec.Tags),az.Tags){
		return true
	}

	return false
}

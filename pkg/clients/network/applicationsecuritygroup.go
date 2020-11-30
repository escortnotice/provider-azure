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

// UpdateApplicationSecurityGroupStatusFromAzure updates the status related to the external
// Azure ApplicationSecurityGroup in the ApplicationSecurityGroupStatus
func UpdateApplicationSecurityGroupStatusFromAzure(asg *v1alpha3.ApplicationSecurityGroup, az networkmgmt.ApplicationSecurityGroup) {
	asg.Status.State = azure.ToString(az.ProvisioningState)
	asg.Status.ID = azure.ToString(az.ID)
	asg.Status.Etag = azure.ToString(az.Etag)
	asg.Spec.Properties.ResourceGUID = az.ResourceGUID
	asg.Status.Type = azure.ToString(az.Type)
}

// NewSubnetParameters returns an Azure ApplicationSecurityGroup object from a ApplicationSecurityGroup spec
func NewApplicationSecurityGroupParameters(asg *v1alpha3.ApplicationSecurityGroup) networkmgmt.ApplicationSecurityGroup {
	return networkmgmt.ApplicationSecurityGroup{
		ApplicationSecurityGroupPropertiesFormat: &networkmgmt.ApplicationSecurityGroupPropertiesFormat{
			ResourceGUID:      asg.Spec.Properties.ResourceGUID,
			ProvisioningState: asg.Spec.Properties.ProvisioningState,
		},
		Etag:                                     azure.ToStringPtr(asg.Spec.Etag),
		ID:                                       azure.ToStringPtr(asg.Spec.ID),
		Name:                                     azure.ToStringPtr(asg.Spec.Name),
		Type:                                     azure.ToStringPtr(asg.Spec.Type),
		Location:                                 azure.ToStringPtr(asg.Spec.Location),
		Tags:                                     azure.ToStringPtrMap(asg.Spec.Tags),
	}
}

// ApplicationSecurityGroupNeedsUpdate determines if a ApplicationSecurityGroup need to be updated
func ApplicationSecurityGroupNeedsUpdate(kube *v1alpha3.ApplicationSecurityGroup, az networkmgmt.ApplicationSecurityGroup) bool {
	asg := NewApplicationSecurityGroupParameters(kube)

	switch {

	case !reflect.DeepEqual(asg.Tags, az.Tags):
		return true
	}
	return false
}
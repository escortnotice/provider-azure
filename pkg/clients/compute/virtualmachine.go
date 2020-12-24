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

package compute

//import (
//	computemgmt "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
//	networkmgmt "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network"
//
//	//networkmgmt "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network"
//
//	"github.com/crossplane/provider-azure/apis/compute/v1alpha3"
//	"reflect"
//
//	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute/computeapi"
//	azure "github.com/crossplane/provider-azure/pkg/clients"
//)
//
//// A GroupsClient handles CRUD operations for Azure Virtual Machine resources.
//type GroupsClient computeapi.VirtualMachinesClientAPI
//
//// UpdateSecurityGroupStatusFromAzure updates the status related to the external
//// Azure Security Group in the SecurityGroupStatus
//func UpdateVirtualMachineStatusFromAzure(v *v1alpha3.VirtualMachine, az computemgmt.VirtualMachine) {
//	v.Status.State = azure.ToString(az.ProvisioningState)
//	v.Status.ID = azure.ToString(az.ID)
//	//v.Status.Etag = azure.ToString(az.Etag)
//	//v.Status.ResourceGUID = azure.ToString(az.ResourceGUID)
//	v.Status.Type = azure.ToString(az.Type)
//}
//
//// NewSecurityGroupParameters returns an Azure SecurityGroup object from a Security Group Spec
//func NewVirtualMachineParameters(v *v1alpha3.VirtualMachine) computemgmt.VirtualMachine {
//	return computemgmt.VirtualMachine{
//		Location: azure.ToStringPtr(v.Spec.Location),
//		Tags:     azure.ToStringPtrMap(v.Spec.Tags),
//		VirtualMachineProperties: &computemgmt.VirtualMachineProperties{
//			// Default spec changes will be added if needed here
//			SecurityRules: SetSecurityRulesToSecurityGroup(v.Spec.SecurityGroupPropertiesFormat.SecurityRules),
//		},
//	}
//}

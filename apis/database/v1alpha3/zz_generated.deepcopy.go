// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha3

import (
	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CosmosDBAccount) DeepCopyInto(out *CosmosDBAccount) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CosmosDBAccount.
func (in *CosmosDBAccount) DeepCopy() *CosmosDBAccount {
	if in == nil {
		return nil
	}
	out := new(CosmosDBAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CosmosDBAccount) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CosmosDBAccountConsistencyPolicy) DeepCopyInto(out *CosmosDBAccountConsistencyPolicy) {
	*out = *in
	if in.MaxStalenessPrefix != nil {
		in, out := &in.MaxStalenessPrefix, &out.MaxStalenessPrefix
		*out = new(int64)
		**out = **in
	}
	if in.MaxIntervalInSeconds != nil {
		in, out := &in.MaxIntervalInSeconds, &out.MaxIntervalInSeconds
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CosmosDBAccountConsistencyPolicy.
func (in *CosmosDBAccountConsistencyPolicy) DeepCopy() *CosmosDBAccountConsistencyPolicy {
	if in == nil {
		return nil
	}
	out := new(CosmosDBAccountConsistencyPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CosmosDBAccountList) DeepCopyInto(out *CosmosDBAccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CosmosDBAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CosmosDBAccountList.
func (in *CosmosDBAccountList) DeepCopy() *CosmosDBAccountList {
	if in == nil {
		return nil
	}
	out := new(CosmosDBAccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CosmosDBAccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CosmosDBAccountLocation) DeepCopyInto(out *CosmosDBAccountLocation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CosmosDBAccountLocation.
func (in *CosmosDBAccountLocation) DeepCopy() *CosmosDBAccountLocation {
	if in == nil {
		return nil
	}
	out := new(CosmosDBAccountLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CosmosDBAccountObservation) DeepCopyInto(out *CosmosDBAccountObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CosmosDBAccountObservation.
func (in *CosmosDBAccountObservation) DeepCopy() *CosmosDBAccountObservation {
	if in == nil {
		return nil
	}
	out := new(CosmosDBAccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CosmosDBAccountParameters) DeepCopyInto(out *CosmosDBAccountParameters) {
	*out = *in
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	in.Properties.DeepCopyInto(&out.Properties)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CosmosDBAccountParameters.
func (in *CosmosDBAccountParameters) DeepCopy() *CosmosDBAccountParameters {
	if in == nil {
		return nil
	}
	out := new(CosmosDBAccountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CosmosDBAccountProperties) DeepCopyInto(out *CosmosDBAccountProperties) {
	*out = *in
	if in.ConsistencyPolicy != nil {
		in, out := &in.ConsistencyPolicy, &out.ConsistencyPolicy
		*out = new(CosmosDBAccountConsistencyPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]CosmosDBAccountLocation, len(*in))
		copy(*out, *in)
	}
	if in.IPRangeFilter != nil {
		in, out := &in.IPRangeFilter, &out.IPRangeFilter
		*out = new(string)
		**out = **in
	}
	if in.EnableAutomaticFailover != nil {
		in, out := &in.EnableAutomaticFailover, &out.EnableAutomaticFailover
		*out = new(bool)
		**out = **in
	}
	if in.EnableMultipleWriteLocations != nil {
		in, out := &in.EnableMultipleWriteLocations, &out.EnableMultipleWriteLocations
		*out = new(bool)
		**out = **in
	}
	if in.EnableCassandraConnector != nil {
		in, out := &in.EnableCassandraConnector, &out.EnableCassandraConnector
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CosmosDBAccountProperties.
func (in *CosmosDBAccountProperties) DeepCopy() *CosmosDBAccountProperties {
	if in == nil {
		return nil
	}
	out := new(CosmosDBAccountProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CosmosDBAccountSpec) DeepCopyInto(out *CosmosDBAccountSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CosmosDBAccountSpec.
func (in *CosmosDBAccountSpec) DeepCopy() *CosmosDBAccountSpec {
	if in == nil {
		return nil
	}
	out := new(CosmosDBAccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CosmosDBAccountStatus) DeepCopyInto(out *CosmosDBAccountStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	if in.AtProvider != nil {
		in, out := &in.AtProvider, &out.AtProvider
		*out = new(CosmosDBAccountObservation)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CosmosDBAccountStatus.
func (in *CosmosDBAccountStatus) DeepCopy() *CosmosDBAccountStatus {
	if in == nil {
		return nil
	}
	out := new(CosmosDBAccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleObservation) DeepCopyInto(out *FirewallRuleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleObservation.
func (in *FirewallRuleObservation) DeepCopy() *FirewallRuleObservation {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleParameters) DeepCopyInto(out *FirewallRuleParameters) {
	*out = *in
	if in.ServerNameRef != nil {
		in, out := &in.ServerNameRef, &out.ServerNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ServerNameSelector != nil {
		in, out := &in.ServerNameSelector, &out.ServerNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	out.FirewallRuleProperties = in.FirewallRuleProperties
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleParameters.
func (in *FirewallRuleParameters) DeepCopy() *FirewallRuleParameters {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleProperties) DeepCopyInto(out *FirewallRuleProperties) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleProperties.
func (in *FirewallRuleProperties) DeepCopy() *FirewallRuleProperties {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleSpec) DeepCopyInto(out *FirewallRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleSpec.
func (in *FirewallRuleSpec) DeepCopy() *FirewallRuleSpec {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirewallRuleStatus) DeepCopyInto(out *FirewallRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirewallRuleStatus.
func (in *FirewallRuleStatus) DeepCopy() *FirewallRuleStatus {
	if in == nil {
		return nil
	}
	out := new(FirewallRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServerFirewallRule) DeepCopyInto(out *MySQLServerFirewallRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServerFirewallRule.
func (in *MySQLServerFirewallRule) DeepCopy() *MySQLServerFirewallRule {
	if in == nil {
		return nil
	}
	out := new(MySQLServerFirewallRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLServerFirewallRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServerFirewallRuleList) DeepCopyInto(out *MySQLServerFirewallRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLServerFirewallRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServerFirewallRuleList.
func (in *MySQLServerFirewallRuleList) DeepCopy() *MySQLServerFirewallRuleList {
	if in == nil {
		return nil
	}
	out := new(MySQLServerFirewallRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLServerFirewallRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServerVirtualNetworkRule) DeepCopyInto(out *MySQLServerVirtualNetworkRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServerVirtualNetworkRule.
func (in *MySQLServerVirtualNetworkRule) DeepCopy() *MySQLServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(MySQLServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLServerVirtualNetworkRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServerVirtualNetworkRuleList) DeepCopyInto(out *MySQLServerVirtualNetworkRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLServerVirtualNetworkRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServerVirtualNetworkRuleList.
func (in *MySQLServerVirtualNetworkRuleList) DeepCopy() *MySQLServerVirtualNetworkRuleList {
	if in == nil {
		return nil
	}
	out := new(MySQLServerVirtualNetworkRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLServerVirtualNetworkRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLVirtualNetworkRuleSpec) DeepCopyInto(out *MySQLVirtualNetworkRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.ServerNameRef != nil {
		in, out := &in.ServerNameRef, &out.ServerNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ServerNameSelector != nil {
		in, out := &in.ServerNameSelector, &out.ServerNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	in.VirtualNetworkRuleProperties.DeepCopyInto(&out.VirtualNetworkRuleProperties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLVirtualNetworkRuleSpec.
func (in *MySQLVirtualNetworkRuleSpec) DeepCopy() *MySQLVirtualNetworkRuleSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLVirtualNetworkRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServerFirewallRule) DeepCopyInto(out *PostgreSQLServerFirewallRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServerFirewallRule.
func (in *PostgreSQLServerFirewallRule) DeepCopy() *PostgreSQLServerFirewallRule {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServerFirewallRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLServerFirewallRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServerFirewallRuleList) DeepCopyInto(out *PostgreSQLServerFirewallRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgreSQLServerFirewallRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServerFirewallRuleList.
func (in *PostgreSQLServerFirewallRuleList) DeepCopy() *PostgreSQLServerFirewallRuleList {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServerFirewallRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLServerFirewallRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServerVirtualNetworkRule) DeepCopyInto(out *PostgreSQLServerVirtualNetworkRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServerVirtualNetworkRule.
func (in *PostgreSQLServerVirtualNetworkRule) DeepCopy() *PostgreSQLServerVirtualNetworkRule {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServerVirtualNetworkRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLServerVirtualNetworkRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServerVirtualNetworkRuleList) DeepCopyInto(out *PostgreSQLServerVirtualNetworkRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgreSQLServerVirtualNetworkRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServerVirtualNetworkRuleList.
func (in *PostgreSQLServerVirtualNetworkRuleList) DeepCopy() *PostgreSQLServerVirtualNetworkRuleList {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServerVirtualNetworkRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLServerVirtualNetworkRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLVirtualNetworkRuleSpec) DeepCopyInto(out *PostgreSQLVirtualNetworkRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	if in.ServerNameRef != nil {
		in, out := &in.ServerNameRef, &out.ServerNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ServerNameSelector != nil {
		in, out := &in.ServerNameSelector, &out.ServerNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
	in.VirtualNetworkRuleProperties.DeepCopyInto(&out.VirtualNetworkRuleProperties)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLVirtualNetworkRuleSpec.
func (in *PostgreSQLVirtualNetworkRuleSpec) DeepCopy() *PostgreSQLVirtualNetworkRuleSpec {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLVirtualNetworkRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleProperties) DeepCopyInto(out *VirtualNetworkRuleProperties) {
	*out = *in
	if in.VirtualNetworkSubnetIDRef != nil {
		in, out := &in.VirtualNetworkSubnetIDRef, &out.VirtualNetworkSubnetIDRef
		*out = new(v1alpha1.Reference)
		**out = **in
	}
	if in.VirtualNetworkSubnetIDSelector != nil {
		in, out := &in.VirtualNetworkSubnetIDSelector, &out.VirtualNetworkSubnetIDSelector
		*out = new(v1alpha1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleProperties.
func (in *VirtualNetworkRuleProperties) DeepCopy() *VirtualNetworkRuleProperties {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNetworkRuleStatus) DeepCopyInto(out *VirtualNetworkRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNetworkRuleStatus.
func (in *VirtualNetworkRuleStatus) DeepCopy() *VirtualNetworkRuleStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualNetworkRuleStatus)
	in.DeepCopyInto(out)
	return out
}

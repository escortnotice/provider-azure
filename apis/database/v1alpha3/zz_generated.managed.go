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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha3

import (
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

// GetBindingPhase of this CosmosDBAccount.
func (mg *CosmosDBAccount) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this CosmosDBAccount.
func (mg *CosmosDBAccount) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetClassReference of this CosmosDBAccount.
func (mg *CosmosDBAccount) GetClassReference() *corev1.ObjectReference {
	return mg.Spec.ClassReference
}

// GetCondition of this CosmosDBAccount.
func (mg *CosmosDBAccount) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetProviderReference of this CosmosDBAccount.
func (mg *CosmosDBAccount) GetProviderReference() *corev1.ObjectReference {
	return mg.Spec.ProviderReference
}

// GetReclaimPolicy of this CosmosDBAccount.
func (mg *CosmosDBAccount) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this CosmosDBAccount.
func (mg *CosmosDBAccount) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this CosmosDBAccount.
func (mg *CosmosDBAccount) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this CosmosDBAccount.
func (mg *CosmosDBAccount) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetClassReference of this CosmosDBAccount.
func (mg *CosmosDBAccount) SetClassReference(r *corev1.ObjectReference) {
	mg.Spec.ClassReference = r
}

// SetConditions of this CosmosDBAccount.
func (mg *CosmosDBAccount) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetProviderReference of this CosmosDBAccount.
func (mg *CosmosDBAccount) SetProviderReference(r *corev1.ObjectReference) {
	mg.Spec.ProviderReference = r
}

// SetReclaimPolicy of this CosmosDBAccount.
func (mg *CosmosDBAccount) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this CosmosDBAccount.
func (mg *CosmosDBAccount) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetClassReference of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) GetClassReference() *corev1.ObjectReference {
	return mg.Spec.ClassReference
}

// GetCondition of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetProviderReference of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) GetProviderReference() *corev1.ObjectReference {
	return mg.Spec.ProviderReference
}

// GetReclaimPolicy of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetClassReference of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) SetClassReference(r *corev1.ObjectReference) {
	mg.Spec.ClassReference = r
}

// SetConditions of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetProviderReference of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) SetProviderReference(r *corev1.ObjectReference) {
	mg.Spec.ProviderReference = r
}

// SetReclaimPolicy of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this MySQLServerFirewallRule.
func (mg *MySQLServerFirewallRule) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetClassReference of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) GetClassReference() *corev1.ObjectReference {
	return mg.Spec.ClassReference
}

// GetCondition of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetProviderReference of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) GetProviderReference() *corev1.ObjectReference {
	return mg.Spec.ProviderReference
}

// GetReclaimPolicy of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetClassReference of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) SetClassReference(r *corev1.ObjectReference) {
	mg.Spec.ClassReference = r
}

// SetConditions of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetProviderReference of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) SetProviderReference(r *corev1.ObjectReference) {
	mg.Spec.ProviderReference = r
}

// SetReclaimPolicy of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this MySQLServerVirtualNetworkRule.
func (mg *MySQLServerVirtualNetworkRule) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetClassReference of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) GetClassReference() *corev1.ObjectReference {
	return mg.Spec.ClassReference
}

// GetCondition of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetProviderReference of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) GetProviderReference() *corev1.ObjectReference {
	return mg.Spec.ProviderReference
}

// GetReclaimPolicy of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetClassReference of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) SetClassReference(r *corev1.ObjectReference) {
	mg.Spec.ClassReference = r
}

// SetConditions of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetProviderReference of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) SetProviderReference(r *corev1.ObjectReference) {
	mg.Spec.ProviderReference = r
}

// SetReclaimPolicy of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this PostgreSQLServerFirewallRule.
func (mg *PostgreSQLServerFirewallRule) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return mg.Status.GetBindingPhase()
}

// GetClaimReference of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) GetClaimReference() *corev1.ObjectReference {
	return mg.Spec.ClaimReference
}

// GetClassReference of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) GetClassReference() *corev1.ObjectReference {
	return mg.Spec.ClassReference
}

// GetCondition of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetProviderReference of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) GetProviderReference() *corev1.ObjectReference {
	return mg.Spec.ProviderReference
}

// GetReclaimPolicy of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return mg.Spec.ReclaimPolicy
}

// GetWriteConnectionSecretToReference of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) GetWriteConnectionSecretToReference() *runtimev1alpha1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	mg.Status.SetBindingPhase(p)
}

// SetClaimReference of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) SetClaimReference(r *corev1.ObjectReference) {
	mg.Spec.ClaimReference = r
}

// SetClassReference of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) SetClassReference(r *corev1.ObjectReference) {
	mg.Spec.ClassReference = r
}

// SetConditions of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) SetConditions(c ...runtimev1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetProviderReference of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) SetProviderReference(r *corev1.ObjectReference) {
	mg.Spec.ProviderReference = r
}

// SetReclaimPolicy of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	mg.Spec.ReclaimPolicy = r
}

// SetWriteConnectionSecretToReference of this PostgreSQLServerVirtualNetworkRule.
func (mg *PostgreSQLServerVirtualNetworkRule) SetWriteConnectionSecretToReference(r *runtimev1alpha1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

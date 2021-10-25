/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this KustoAttachedDatabaseConfiguration.
func (mg *KustoAttachedDatabaseConfiguration) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this KustoAttachedDatabaseConfiguration.
func (mg *KustoAttachedDatabaseConfiguration) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this KustoAttachedDatabaseConfiguration.
func (mg *KustoAttachedDatabaseConfiguration) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this KustoAttachedDatabaseConfiguration.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *KustoAttachedDatabaseConfiguration) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this KustoAttachedDatabaseConfiguration.
func (mg *KustoAttachedDatabaseConfiguration) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this KustoAttachedDatabaseConfiguration.
func (mg *KustoAttachedDatabaseConfiguration) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this KustoAttachedDatabaseConfiguration.
func (mg *KustoAttachedDatabaseConfiguration) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this KustoAttachedDatabaseConfiguration.
func (mg *KustoAttachedDatabaseConfiguration) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this KustoAttachedDatabaseConfiguration.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *KustoAttachedDatabaseConfiguration) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this KustoAttachedDatabaseConfiguration.
func (mg *KustoAttachedDatabaseConfiguration) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this KustoCluster.
func (mg *KustoCluster) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this KustoCluster.
func (mg *KustoCluster) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this KustoCluster.
func (mg *KustoCluster) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this KustoCluster.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *KustoCluster) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this KustoCluster.
func (mg *KustoCluster) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this KustoCluster.
func (mg *KustoCluster) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this KustoCluster.
func (mg *KustoCluster) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this KustoCluster.
func (mg *KustoCluster) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this KustoCluster.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *KustoCluster) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this KustoCluster.
func (mg *KustoCluster) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this KustoClusterCustomerManagedKey.
func (mg *KustoClusterCustomerManagedKey) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this KustoClusterCustomerManagedKey.
func (mg *KustoClusterCustomerManagedKey) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this KustoClusterCustomerManagedKey.
func (mg *KustoClusterCustomerManagedKey) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this KustoClusterCustomerManagedKey.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *KustoClusterCustomerManagedKey) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this KustoClusterCustomerManagedKey.
func (mg *KustoClusterCustomerManagedKey) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this KustoClusterCustomerManagedKey.
func (mg *KustoClusterCustomerManagedKey) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this KustoClusterCustomerManagedKey.
func (mg *KustoClusterCustomerManagedKey) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this KustoClusterCustomerManagedKey.
func (mg *KustoClusterCustomerManagedKey) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this KustoClusterCustomerManagedKey.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *KustoClusterCustomerManagedKey) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this KustoClusterCustomerManagedKey.
func (mg *KustoClusterCustomerManagedKey) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this KustoClusterPrincipalAssignment.
func (mg *KustoClusterPrincipalAssignment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this KustoClusterPrincipalAssignment.
func (mg *KustoClusterPrincipalAssignment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this KustoClusterPrincipalAssignment.
func (mg *KustoClusterPrincipalAssignment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this KustoClusterPrincipalAssignment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *KustoClusterPrincipalAssignment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this KustoClusterPrincipalAssignment.
func (mg *KustoClusterPrincipalAssignment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this KustoClusterPrincipalAssignment.
func (mg *KustoClusterPrincipalAssignment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this KustoClusterPrincipalAssignment.
func (mg *KustoClusterPrincipalAssignment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this KustoClusterPrincipalAssignment.
func (mg *KustoClusterPrincipalAssignment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this KustoClusterPrincipalAssignment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *KustoClusterPrincipalAssignment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this KustoClusterPrincipalAssignment.
func (mg *KustoClusterPrincipalAssignment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this KustoDatabase.
func (mg *KustoDatabase) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this KustoDatabase.
func (mg *KustoDatabase) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this KustoDatabase.
func (mg *KustoDatabase) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this KustoDatabase.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *KustoDatabase) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this KustoDatabase.
func (mg *KustoDatabase) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this KustoDatabase.
func (mg *KustoDatabase) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this KustoDatabase.
func (mg *KustoDatabase) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this KustoDatabase.
func (mg *KustoDatabase) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this KustoDatabase.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *KustoDatabase) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this KustoDatabase.
func (mg *KustoDatabase) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this KustoDatabasePrincipal.
func (mg *KustoDatabasePrincipal) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this KustoDatabasePrincipal.
func (mg *KustoDatabasePrincipal) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this KustoDatabasePrincipal.
func (mg *KustoDatabasePrincipal) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this KustoDatabasePrincipal.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *KustoDatabasePrincipal) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this KustoDatabasePrincipal.
func (mg *KustoDatabasePrincipal) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this KustoDatabasePrincipal.
func (mg *KustoDatabasePrincipal) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this KustoDatabasePrincipal.
func (mg *KustoDatabasePrincipal) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this KustoDatabasePrincipal.
func (mg *KustoDatabasePrincipal) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this KustoDatabasePrincipal.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *KustoDatabasePrincipal) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this KustoDatabasePrincipal.
func (mg *KustoDatabasePrincipal) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this KustoDatabasePrincipalAssignment.
func (mg *KustoDatabasePrincipalAssignment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this KustoDatabasePrincipalAssignment.
func (mg *KustoDatabasePrincipalAssignment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this KustoDatabasePrincipalAssignment.
func (mg *KustoDatabasePrincipalAssignment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this KustoDatabasePrincipalAssignment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *KustoDatabasePrincipalAssignment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this KustoDatabasePrincipalAssignment.
func (mg *KustoDatabasePrincipalAssignment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this KustoDatabasePrincipalAssignment.
func (mg *KustoDatabasePrincipalAssignment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this KustoDatabasePrincipalAssignment.
func (mg *KustoDatabasePrincipalAssignment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this KustoDatabasePrincipalAssignment.
func (mg *KustoDatabasePrincipalAssignment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this KustoDatabasePrincipalAssignment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *KustoDatabasePrincipalAssignment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this KustoDatabasePrincipalAssignment.
func (mg *KustoDatabasePrincipalAssignment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this KustoEventgridDataConnection.
func (mg *KustoEventgridDataConnection) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this KustoEventgridDataConnection.
func (mg *KustoEventgridDataConnection) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this KustoEventgridDataConnection.
func (mg *KustoEventgridDataConnection) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this KustoEventgridDataConnection.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *KustoEventgridDataConnection) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this KustoEventgridDataConnection.
func (mg *KustoEventgridDataConnection) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this KustoEventgridDataConnection.
func (mg *KustoEventgridDataConnection) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this KustoEventgridDataConnection.
func (mg *KustoEventgridDataConnection) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this KustoEventgridDataConnection.
func (mg *KustoEventgridDataConnection) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this KustoEventgridDataConnection.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *KustoEventgridDataConnection) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this KustoEventgridDataConnection.
func (mg *KustoEventgridDataConnection) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this KustoEventhubDataConnection.
func (mg *KustoEventhubDataConnection) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this KustoEventhubDataConnection.
func (mg *KustoEventhubDataConnection) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this KustoEventhubDataConnection.
func (mg *KustoEventhubDataConnection) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this KustoEventhubDataConnection.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *KustoEventhubDataConnection) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this KustoEventhubDataConnection.
func (mg *KustoEventhubDataConnection) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this KustoEventhubDataConnection.
func (mg *KustoEventhubDataConnection) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this KustoEventhubDataConnection.
func (mg *KustoEventhubDataConnection) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this KustoEventhubDataConnection.
func (mg *KustoEventhubDataConnection) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this KustoEventhubDataConnection.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *KustoEventhubDataConnection) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this KustoEventhubDataConnection.
func (mg *KustoEventhubDataConnection) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this KustoIothubDataConnection.
func (mg *KustoIothubDataConnection) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this KustoIothubDataConnection.
func (mg *KustoIothubDataConnection) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this KustoIothubDataConnection.
func (mg *KustoIothubDataConnection) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this KustoIothubDataConnection.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *KustoIothubDataConnection) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this KustoIothubDataConnection.
func (mg *KustoIothubDataConnection) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this KustoIothubDataConnection.
func (mg *KustoIothubDataConnection) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this KustoIothubDataConnection.
func (mg *KustoIothubDataConnection) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this KustoIothubDataConnection.
func (mg *KustoIothubDataConnection) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this KustoIothubDataConnection.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *KustoIothubDataConnection) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this KustoIothubDataConnection.
func (mg *KustoIothubDataConnection) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

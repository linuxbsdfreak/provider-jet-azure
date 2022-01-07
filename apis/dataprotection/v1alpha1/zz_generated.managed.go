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

// GetCondition of this BackupInstanceBlobStorage.
func (mg *BackupInstanceBlobStorage) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BackupInstanceBlobStorage.
func (mg *BackupInstanceBlobStorage) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BackupInstanceBlobStorage.
func (mg *BackupInstanceBlobStorage) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BackupInstanceBlobStorage.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BackupInstanceBlobStorage) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BackupInstanceBlobStorage.
func (mg *BackupInstanceBlobStorage) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BackupInstanceBlobStorage.
func (mg *BackupInstanceBlobStorage) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BackupInstanceBlobStorage.
func (mg *BackupInstanceBlobStorage) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BackupInstanceBlobStorage.
func (mg *BackupInstanceBlobStorage) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BackupInstanceBlobStorage.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BackupInstanceBlobStorage) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BackupInstanceBlobStorage.
func (mg *BackupInstanceBlobStorage) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BackupInstanceDisk.
func (mg *BackupInstanceDisk) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BackupInstanceDisk.
func (mg *BackupInstanceDisk) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BackupInstanceDisk.
func (mg *BackupInstanceDisk) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BackupInstanceDisk.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BackupInstanceDisk) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BackupInstanceDisk.
func (mg *BackupInstanceDisk) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BackupInstanceDisk.
func (mg *BackupInstanceDisk) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BackupInstanceDisk.
func (mg *BackupInstanceDisk) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BackupInstanceDisk.
func (mg *BackupInstanceDisk) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BackupInstanceDisk.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BackupInstanceDisk) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BackupInstanceDisk.
func (mg *BackupInstanceDisk) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BackupInstancePostgreSQL.
func (mg *BackupInstancePostgreSQL) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BackupInstancePostgreSQL.
func (mg *BackupInstancePostgreSQL) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BackupInstancePostgreSQL.
func (mg *BackupInstancePostgreSQL) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BackupInstancePostgreSQL.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BackupInstancePostgreSQL) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BackupInstancePostgreSQL.
func (mg *BackupInstancePostgreSQL) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BackupInstancePostgreSQL.
func (mg *BackupInstancePostgreSQL) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BackupInstancePostgreSQL.
func (mg *BackupInstancePostgreSQL) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BackupInstancePostgreSQL.
func (mg *BackupInstancePostgreSQL) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BackupInstancePostgreSQL.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BackupInstancePostgreSQL) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BackupInstancePostgreSQL.
func (mg *BackupInstancePostgreSQL) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BackupPolicyBlobStorage.
func (mg *BackupPolicyBlobStorage) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BackupPolicyBlobStorage.
func (mg *BackupPolicyBlobStorage) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BackupPolicyBlobStorage.
func (mg *BackupPolicyBlobStorage) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BackupPolicyBlobStorage.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BackupPolicyBlobStorage) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BackupPolicyBlobStorage.
func (mg *BackupPolicyBlobStorage) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BackupPolicyBlobStorage.
func (mg *BackupPolicyBlobStorage) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BackupPolicyBlobStorage.
func (mg *BackupPolicyBlobStorage) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BackupPolicyBlobStorage.
func (mg *BackupPolicyBlobStorage) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BackupPolicyBlobStorage.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BackupPolicyBlobStorage) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BackupPolicyBlobStorage.
func (mg *BackupPolicyBlobStorage) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BackupPolicyDisk.
func (mg *BackupPolicyDisk) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BackupPolicyDisk.
func (mg *BackupPolicyDisk) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BackupPolicyDisk.
func (mg *BackupPolicyDisk) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BackupPolicyDisk.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BackupPolicyDisk) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BackupPolicyDisk.
func (mg *BackupPolicyDisk) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BackupPolicyDisk.
func (mg *BackupPolicyDisk) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BackupPolicyDisk.
func (mg *BackupPolicyDisk) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BackupPolicyDisk.
func (mg *BackupPolicyDisk) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BackupPolicyDisk.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BackupPolicyDisk) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BackupPolicyDisk.
func (mg *BackupPolicyDisk) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BackupPolicyPostgreSQL.
func (mg *BackupPolicyPostgreSQL) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BackupPolicyPostgreSQL.
func (mg *BackupPolicyPostgreSQL) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BackupPolicyPostgreSQL.
func (mg *BackupPolicyPostgreSQL) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BackupPolicyPostgreSQL.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BackupPolicyPostgreSQL) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BackupPolicyPostgreSQL.
func (mg *BackupPolicyPostgreSQL) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BackupPolicyPostgreSQL.
func (mg *BackupPolicyPostgreSQL) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BackupPolicyPostgreSQL.
func (mg *BackupPolicyPostgreSQL) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BackupPolicyPostgreSQL.
func (mg *BackupPolicyPostgreSQL) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BackupPolicyPostgreSQL.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BackupPolicyPostgreSQL) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BackupPolicyPostgreSQL.
func (mg *BackupPolicyPostgreSQL) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this BackupVault.
func (mg *BackupVault) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this BackupVault.
func (mg *BackupVault) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this BackupVault.
func (mg *BackupVault) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this BackupVault.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *BackupVault) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this BackupVault.
func (mg *BackupVault) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this BackupVault.
func (mg *BackupVault) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this BackupVault.
func (mg *BackupVault) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this BackupVault.
func (mg *BackupVault) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this BackupVault.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *BackupVault) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this BackupVault.
func (mg *BackupVault) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

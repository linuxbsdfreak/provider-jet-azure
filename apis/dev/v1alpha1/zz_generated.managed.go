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

// GetCondition of this DevTestGlobalVmShutdownSchedule.
func (mg *DevTestGlobalVmShutdownSchedule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this DevTestGlobalVmShutdownSchedule.
func (mg *DevTestGlobalVmShutdownSchedule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this DevTestGlobalVmShutdownSchedule.
func (mg *DevTestGlobalVmShutdownSchedule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this DevTestGlobalVmShutdownSchedule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *DevTestGlobalVmShutdownSchedule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this DevTestGlobalVmShutdownSchedule.
func (mg *DevTestGlobalVmShutdownSchedule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this DevTestGlobalVmShutdownSchedule.
func (mg *DevTestGlobalVmShutdownSchedule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this DevTestGlobalVmShutdownSchedule.
func (mg *DevTestGlobalVmShutdownSchedule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this DevTestGlobalVmShutdownSchedule.
func (mg *DevTestGlobalVmShutdownSchedule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this DevTestGlobalVmShutdownSchedule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *DevTestGlobalVmShutdownSchedule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this DevTestGlobalVmShutdownSchedule.
func (mg *DevTestGlobalVmShutdownSchedule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this DevTestLab.
func (mg *DevTestLab) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this DevTestLab.
func (mg *DevTestLab) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this DevTestLab.
func (mg *DevTestLab) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this DevTestLab.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *DevTestLab) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this DevTestLab.
func (mg *DevTestLab) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this DevTestLab.
func (mg *DevTestLab) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this DevTestLab.
func (mg *DevTestLab) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this DevTestLab.
func (mg *DevTestLab) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this DevTestLab.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *DevTestLab) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this DevTestLab.
func (mg *DevTestLab) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this DevTestLinuxVirtualMachine.
func (mg *DevTestLinuxVirtualMachine) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this DevTestLinuxVirtualMachine.
func (mg *DevTestLinuxVirtualMachine) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this DevTestLinuxVirtualMachine.
func (mg *DevTestLinuxVirtualMachine) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this DevTestLinuxVirtualMachine.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *DevTestLinuxVirtualMachine) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this DevTestLinuxVirtualMachine.
func (mg *DevTestLinuxVirtualMachine) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this DevTestLinuxVirtualMachine.
func (mg *DevTestLinuxVirtualMachine) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this DevTestLinuxVirtualMachine.
func (mg *DevTestLinuxVirtualMachine) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this DevTestLinuxVirtualMachine.
func (mg *DevTestLinuxVirtualMachine) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this DevTestLinuxVirtualMachine.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *DevTestLinuxVirtualMachine) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this DevTestLinuxVirtualMachine.
func (mg *DevTestLinuxVirtualMachine) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this DevTestPolicy.
func (mg *DevTestPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this DevTestPolicy.
func (mg *DevTestPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this DevTestPolicy.
func (mg *DevTestPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this DevTestPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *DevTestPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this DevTestPolicy.
func (mg *DevTestPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this DevTestPolicy.
func (mg *DevTestPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this DevTestPolicy.
func (mg *DevTestPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this DevTestPolicy.
func (mg *DevTestPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this DevTestPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *DevTestPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this DevTestPolicy.
func (mg *DevTestPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this DevTestSchedule.
func (mg *DevTestSchedule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this DevTestSchedule.
func (mg *DevTestSchedule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this DevTestSchedule.
func (mg *DevTestSchedule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this DevTestSchedule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *DevTestSchedule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this DevTestSchedule.
func (mg *DevTestSchedule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this DevTestSchedule.
func (mg *DevTestSchedule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this DevTestSchedule.
func (mg *DevTestSchedule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this DevTestSchedule.
func (mg *DevTestSchedule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this DevTestSchedule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *DevTestSchedule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this DevTestSchedule.
func (mg *DevTestSchedule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this DevTestVirtualNetwork.
func (mg *DevTestVirtualNetwork) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this DevTestVirtualNetwork.
func (mg *DevTestVirtualNetwork) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this DevTestVirtualNetwork.
func (mg *DevTestVirtualNetwork) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this DevTestVirtualNetwork.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *DevTestVirtualNetwork) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this DevTestVirtualNetwork.
func (mg *DevTestVirtualNetwork) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this DevTestVirtualNetwork.
func (mg *DevTestVirtualNetwork) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this DevTestVirtualNetwork.
func (mg *DevTestVirtualNetwork) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this DevTestVirtualNetwork.
func (mg *DevTestVirtualNetwork) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this DevTestVirtualNetwork.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *DevTestVirtualNetwork) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this DevTestVirtualNetwork.
func (mg *DevTestVirtualNetwork) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this DevTestWindowsVirtualMachine.
func (mg *DevTestWindowsVirtualMachine) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this DevTestWindowsVirtualMachine.
func (mg *DevTestWindowsVirtualMachine) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this DevTestWindowsVirtualMachine.
func (mg *DevTestWindowsVirtualMachine) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this DevTestWindowsVirtualMachine.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *DevTestWindowsVirtualMachine) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this DevTestWindowsVirtualMachine.
func (mg *DevTestWindowsVirtualMachine) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this DevTestWindowsVirtualMachine.
func (mg *DevTestWindowsVirtualMachine) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this DevTestWindowsVirtualMachine.
func (mg *DevTestWindowsVirtualMachine) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this DevTestWindowsVirtualMachine.
func (mg *DevTestWindowsVirtualMachine) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this DevTestWindowsVirtualMachine.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *DevTestWindowsVirtualMachine) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this DevTestWindowsVirtualMachine.
func (mg *DevTestWindowsVirtualMachine) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

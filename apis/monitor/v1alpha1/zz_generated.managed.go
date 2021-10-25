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

// GetCondition of this MonitorAadDiagnosticSetting.
func (mg *MonitorAadDiagnosticSetting) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorAadDiagnosticSetting.
func (mg *MonitorAadDiagnosticSetting) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MonitorAadDiagnosticSetting.
func (mg *MonitorAadDiagnosticSetting) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorAadDiagnosticSetting.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorAadDiagnosticSetting) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MonitorAadDiagnosticSetting.
func (mg *MonitorAadDiagnosticSetting) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorAadDiagnosticSetting.
func (mg *MonitorAadDiagnosticSetting) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorAadDiagnosticSetting.
func (mg *MonitorAadDiagnosticSetting) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MonitorAadDiagnosticSetting.
func (mg *MonitorAadDiagnosticSetting) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorAadDiagnosticSetting.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorAadDiagnosticSetting) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MonitorAadDiagnosticSetting.
func (mg *MonitorAadDiagnosticSetting) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorActionGroup.
func (mg *MonitorActionGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorActionGroup.
func (mg *MonitorActionGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MonitorActionGroup.
func (mg *MonitorActionGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorActionGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorActionGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MonitorActionGroup.
func (mg *MonitorActionGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorActionGroup.
func (mg *MonitorActionGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorActionGroup.
func (mg *MonitorActionGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MonitorActionGroup.
func (mg *MonitorActionGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorActionGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorActionGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MonitorActionGroup.
func (mg *MonitorActionGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorActionRuleActionGroup.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorActionRuleActionGroup) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorActionRuleActionGroup.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorActionRuleActionGroup) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MonitorActionRuleActionGroup.
func (mg *MonitorActionRuleActionGroup) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorActionRuleSuppression.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorActionRuleSuppression) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorActionRuleSuppression.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorActionRuleSuppression) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MonitorActionRuleSuppression.
func (mg *MonitorActionRuleSuppression) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorActivityLogAlert.
func (mg *MonitorActivityLogAlert) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorActivityLogAlert.
func (mg *MonitorActivityLogAlert) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MonitorActivityLogAlert.
func (mg *MonitorActivityLogAlert) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorActivityLogAlert.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorActivityLogAlert) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MonitorActivityLogAlert.
func (mg *MonitorActivityLogAlert) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorActivityLogAlert.
func (mg *MonitorActivityLogAlert) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorActivityLogAlert.
func (mg *MonitorActivityLogAlert) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MonitorActivityLogAlert.
func (mg *MonitorActivityLogAlert) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorActivityLogAlert.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorActivityLogAlert) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MonitorActivityLogAlert.
func (mg *MonitorActivityLogAlert) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorAutoscaleSetting.
func (mg *MonitorAutoscaleSetting) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorAutoscaleSetting.
func (mg *MonitorAutoscaleSetting) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MonitorAutoscaleSetting.
func (mg *MonitorAutoscaleSetting) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorAutoscaleSetting.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorAutoscaleSetting) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MonitorAutoscaleSetting.
func (mg *MonitorAutoscaleSetting) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorAutoscaleSetting.
func (mg *MonitorAutoscaleSetting) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorAutoscaleSetting.
func (mg *MonitorAutoscaleSetting) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MonitorAutoscaleSetting.
func (mg *MonitorAutoscaleSetting) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorAutoscaleSetting.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorAutoscaleSetting) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MonitorAutoscaleSetting.
func (mg *MonitorAutoscaleSetting) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorDiagnosticSetting.
func (mg *MonitorDiagnosticSetting) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorDiagnosticSetting.
func (mg *MonitorDiagnosticSetting) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MonitorDiagnosticSetting.
func (mg *MonitorDiagnosticSetting) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorDiagnosticSetting.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorDiagnosticSetting) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MonitorDiagnosticSetting.
func (mg *MonitorDiagnosticSetting) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorDiagnosticSetting.
func (mg *MonitorDiagnosticSetting) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorDiagnosticSetting.
func (mg *MonitorDiagnosticSetting) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MonitorDiagnosticSetting.
func (mg *MonitorDiagnosticSetting) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorDiagnosticSetting.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorDiagnosticSetting) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MonitorDiagnosticSetting.
func (mg *MonitorDiagnosticSetting) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorMetricAlert.
func (mg *MonitorMetricAlert) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorMetricAlert.
func (mg *MonitorMetricAlert) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MonitorMetricAlert.
func (mg *MonitorMetricAlert) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorMetricAlert.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorMetricAlert) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MonitorMetricAlert.
func (mg *MonitorMetricAlert) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorMetricAlert.
func (mg *MonitorMetricAlert) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorMetricAlert.
func (mg *MonitorMetricAlert) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MonitorMetricAlert.
func (mg *MonitorMetricAlert) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorMetricAlert.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorMetricAlert) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MonitorMetricAlert.
func (mg *MonitorMetricAlert) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorScheduledQueryRulesAlert.
func (mg *MonitorScheduledQueryRulesAlert) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorScheduledQueryRulesAlert.
func (mg *MonitorScheduledQueryRulesAlert) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MonitorScheduledQueryRulesAlert.
func (mg *MonitorScheduledQueryRulesAlert) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorScheduledQueryRulesAlert.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorScheduledQueryRulesAlert) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MonitorScheduledQueryRulesAlert.
func (mg *MonitorScheduledQueryRulesAlert) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorScheduledQueryRulesAlert.
func (mg *MonitorScheduledQueryRulesAlert) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorScheduledQueryRulesAlert.
func (mg *MonitorScheduledQueryRulesAlert) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MonitorScheduledQueryRulesAlert.
func (mg *MonitorScheduledQueryRulesAlert) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorScheduledQueryRulesAlert.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorScheduledQueryRulesAlert) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MonitorScheduledQueryRulesAlert.
func (mg *MonitorScheduledQueryRulesAlert) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorScheduledQueryRulesLog.
func (mg *MonitorScheduledQueryRulesLog) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorScheduledQueryRulesLog.
func (mg *MonitorScheduledQueryRulesLog) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MonitorScheduledQueryRulesLog.
func (mg *MonitorScheduledQueryRulesLog) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorScheduledQueryRulesLog.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorScheduledQueryRulesLog) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MonitorScheduledQueryRulesLog.
func (mg *MonitorScheduledQueryRulesLog) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorScheduledQueryRulesLog.
func (mg *MonitorScheduledQueryRulesLog) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorScheduledQueryRulesLog.
func (mg *MonitorScheduledQueryRulesLog) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MonitorScheduledQueryRulesLog.
func (mg *MonitorScheduledQueryRulesLog) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorScheduledQueryRulesLog.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorScheduledQueryRulesLog) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MonitorScheduledQueryRulesLog.
func (mg *MonitorScheduledQueryRulesLog) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MonitorSmartDetectorAlertRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MonitorSmartDetectorAlertRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MonitorSmartDetectorAlertRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MonitorSmartDetectorAlertRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this MonitorSmartDetectorAlertRule.
func (mg *MonitorSmartDetectorAlertRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

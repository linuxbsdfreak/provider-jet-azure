/*
Copyright 2022 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LogObservation struct {
}

type LogParameters struct {

	// +kubebuilder:validation:Required
	Category *string `json:"category" tf:"category,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy" tf:"retention_policy,omitempty"`
}

type MonitorAADDiagnosticSettingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MonitorAADDiagnosticSettingParameters struct {

	// +kubebuilder:validation:Optional
	EventHubAuthorizationRuleID *string `json:"eventhubAuthorizationRuleId,omitempty" tf:"eventhub_authorization_rule_id,omitempty"`

	// +kubebuilder:validation:Optional
	EventHubName *string `json:"eventhubName,omitempty" tf:"eventhub_name,omitempty"`

	// +kubebuilder:validation:Required
	Log []LogParameters `json:"log" tf:"log,omitempty"`

	// +kubebuilder:validation:Optional
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId,omitempty" tf:"log_analytics_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type RetentionPolicyObservation struct {
}

type RetentionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

// MonitorAADDiagnosticSettingSpec defines the desired state of MonitorAADDiagnosticSetting
type MonitorAADDiagnosticSettingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorAADDiagnosticSettingParameters `json:"forProvider"`
}

// MonitorAADDiagnosticSettingStatus defines the observed state of MonitorAADDiagnosticSetting.
type MonitorAADDiagnosticSettingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorAADDiagnosticSettingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorAADDiagnosticSetting is the Schema for the MonitorAADDiagnosticSettings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type MonitorAADDiagnosticSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorAADDiagnosticSettingSpec   `json:"spec"`
	Status            MonitorAADDiagnosticSettingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorAADDiagnosticSettingList contains a list of MonitorAADDiagnosticSettings
type MonitorAADDiagnosticSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorAADDiagnosticSetting `json:"items"`
}

// Repository type metadata.
var (
	MonitorAADDiagnosticSetting_Kind             = "MonitorAADDiagnosticSetting"
	MonitorAADDiagnosticSetting_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorAADDiagnosticSetting_Kind}.String()
	MonitorAADDiagnosticSetting_KindAPIVersion   = MonitorAADDiagnosticSetting_Kind + "." + CRDGroupVersion.String()
	MonitorAADDiagnosticSetting_GroupVersionKind = CRDGroupVersion.WithKind(MonitorAADDiagnosticSetting_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorAADDiagnosticSetting{}, &MonitorAADDiagnosticSettingList{})
}

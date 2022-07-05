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

type BackupPolicyDiskObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackupPolicyDiskParameters struct {

	// +kubebuilder:validation:Required
	BackupRepeatingTimeIntervals []*string `json:"backupRepeatingTimeIntervals" tf:"backup_repeating_time_intervals,omitempty"`

	// +kubebuilder:validation:Required
	DefaultRetentionDuration *string `json:"defaultRetentionDuration" tf:"default_retention_duration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionRule []RetentionRuleParameters `json:"retentionRule,omitempty" tf:"retention_rule,omitempty"`

	// +kubebuilder:validation:Required
	VaultID *string `json:"vaultId" tf:"vault_id,omitempty"`
}

type CriteriaObservation struct {
}

type CriteriaParameters struct {

	// +kubebuilder:validation:Optional
	AbsoluteCriteria *string `json:"absoluteCriteria,omitempty" tf:"absolute_criteria,omitempty"`
}

type RetentionRuleObservation struct {
}

type RetentionRuleParameters struct {

	// +kubebuilder:validation:Required
	Criteria []CriteriaParameters `json:"criteria" tf:"criteria,omitempty"`

	// +kubebuilder:validation:Required
	Duration *string `json:"duration" tf:"duration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`
}

// BackupPolicyDiskSpec defines the desired state of BackupPolicyDisk
type BackupPolicyDiskSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupPolicyDiskParameters `json:"forProvider"`
}

// BackupPolicyDiskStatus defines the observed state of BackupPolicyDisk.
type BackupPolicyDiskStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupPolicyDiskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyDisk is the Schema for the BackupPolicyDisks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type BackupPolicyDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupPolicyDiskSpec   `json:"spec"`
	Status            BackupPolicyDiskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyDiskList contains a list of BackupPolicyDisks
type BackupPolicyDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPolicyDisk `json:"items"`
}

// Repository type metadata.
var (
	BackupPolicyDisk_Kind             = "BackupPolicyDisk"
	BackupPolicyDisk_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupPolicyDisk_Kind}.String()
	BackupPolicyDisk_KindAPIVersion   = BackupPolicyDisk_Kind + "." + CRDGroupVersion.String()
	BackupPolicyDisk_GroupVersionKind = CRDGroupVersion.WithKind(BackupPolicyDisk_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupPolicyDisk{}, &BackupPolicyDiskList{})
}

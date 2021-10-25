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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BackupPolicyVmBackupObservation struct {
}

type BackupPolicyVmBackupParameters struct {

	// +kubebuilder:validation:Required
	Frequency *string `json:"frequency" tf:"frequency,omitempty"`

	// +kubebuilder:validation:Required
	Time *string `json:"time" tf:"time,omitempty"`

	// +kubebuilder:validation:Optional
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type BackupPolicyVmObservation struct {
}

type BackupPolicyVmParameters struct {

	// +kubebuilder:validation:Required
	Backup []BackupPolicyVmBackupParameters `json:"backup" tf:"backup,omitempty"`

	// +kubebuilder:validation:Optional
	InstantRestoreRetentionDays *int64 `json:"instantRestoreRetentionDays,omitempty" tf:"instant_restore_retention_days,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryVaultName *string `json:"recoveryVaultName" tf:"recovery_vault_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionDaily []BackupPolicyVmRetentionDailyParameters `json:"retentionDaily,omitempty" tf:"retention_daily,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionMonthly []BackupPolicyVmRetentionMonthlyParameters `json:"retentionMonthly,omitempty" tf:"retention_monthly,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionWeekly []BackupPolicyVmRetentionWeeklyParameters `json:"retentionWeekly,omitempty" tf:"retention_weekly,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionYearly []BackupPolicyVmRetentionYearlyParameters `json:"retentionYearly,omitempty" tf:"retention_yearly,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type BackupPolicyVmRetentionDailyObservation struct {
}

type BackupPolicyVmRetentionDailyParameters struct {

	// +kubebuilder:validation:Required
	Count *int64 `json:"count" tf:"count,omitempty"`
}

type BackupPolicyVmRetentionMonthlyObservation struct {
}

type BackupPolicyVmRetentionMonthlyParameters struct {

	// +kubebuilder:validation:Required
	Count *int64 `json:"count" tf:"count,omitempty"`

	// +kubebuilder:validation:Required
	Weekdays []*string `json:"weekdays" tf:"weekdays,omitempty"`

	// +kubebuilder:validation:Required
	Weeks []*string `json:"weeks" tf:"weeks,omitempty"`
}

type BackupPolicyVmRetentionWeeklyObservation struct {
}

type BackupPolicyVmRetentionWeeklyParameters struct {

	// +kubebuilder:validation:Required
	Count *int64 `json:"count" tf:"count,omitempty"`

	// +kubebuilder:validation:Required
	Weekdays []*string `json:"weekdays" tf:"weekdays,omitempty"`
}

type BackupPolicyVmRetentionYearlyObservation struct {
}

type BackupPolicyVmRetentionYearlyParameters struct {

	// +kubebuilder:validation:Required
	Count *int64 `json:"count" tf:"count,omitempty"`

	// +kubebuilder:validation:Required
	Months []*string `json:"months" tf:"months,omitempty"`

	// +kubebuilder:validation:Required
	Weekdays []*string `json:"weekdays" tf:"weekdays,omitempty"`

	// +kubebuilder:validation:Required
	Weeks []*string `json:"weeks" tf:"weeks,omitempty"`
}

// BackupPolicyVmSpec defines the desired state of BackupPolicyVm
type BackupPolicyVmSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupPolicyVmParameters `json:"forProvider"`
}

// BackupPolicyVmStatus defines the observed state of BackupPolicyVm.
type BackupPolicyVmStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupPolicyVmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyVm is the Schema for the BackupPolicyVms API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupPolicyVm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupPolicyVmSpec   `json:"spec"`
	Status            BackupPolicyVmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyVmList contains a list of BackupPolicyVms
type BackupPolicyVmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPolicyVm `json:"items"`
}

// Repository type metadata.
var (
	BackupPolicyVmKind             = "BackupPolicyVm"
	BackupPolicyVmGroupKind        = schema.GroupKind{Group: Group, Kind: BackupPolicyVmKind}.String()
	BackupPolicyVmKindAPIVersion   = BackupPolicyVmKind + "." + GroupVersion.String()
	BackupPolicyVmGroupVersionKind = GroupVersion.WithKind(BackupPolicyVmKind)
)

func init() {
	SchemeBuilder.Register(&BackupPolicyVm{}, &BackupPolicyVmList{})
}

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

type WorkspaceExtendedAuditingPolicyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type WorkspaceExtendedAuditingPolicyParameters struct {

	// +kubebuilder:validation:Optional
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionInDays *int64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`

	// +kubebuilder:validation:Required
	SynapseWorkspaceID *string `json:"synapseWorkspaceId" tf:"synapse_workspace_id,omitempty"`
}

// WorkspaceExtendedAuditingPolicySpec defines the desired state of WorkspaceExtendedAuditingPolicy
type WorkspaceExtendedAuditingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceExtendedAuditingPolicyParameters `json:"forProvider"`
}

// WorkspaceExtendedAuditingPolicyStatus defines the observed state of WorkspaceExtendedAuditingPolicy.
type WorkspaceExtendedAuditingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceExtendedAuditingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceExtendedAuditingPolicy is the Schema for the WorkspaceExtendedAuditingPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type WorkspaceExtendedAuditingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceExtendedAuditingPolicySpec   `json:"spec"`
	Status            WorkspaceExtendedAuditingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceExtendedAuditingPolicyList contains a list of WorkspaceExtendedAuditingPolicys
type WorkspaceExtendedAuditingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspaceExtendedAuditingPolicy `json:"items"`
}

// Repository type metadata.
var (
	WorkspaceExtendedAuditingPolicy_Kind             = "WorkspaceExtendedAuditingPolicy"
	WorkspaceExtendedAuditingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkspaceExtendedAuditingPolicy_Kind}.String()
	WorkspaceExtendedAuditingPolicy_KindAPIVersion   = WorkspaceExtendedAuditingPolicy_Kind + "." + CRDGroupVersion.String()
	WorkspaceExtendedAuditingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(WorkspaceExtendedAuditingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkspaceExtendedAuditingPolicy{}, &WorkspaceExtendedAuditingPolicyList{})
}

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

type PolicyAssignmentIdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type PolicyAssignmentIdentityParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PolicyAssignmentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PolicyAssignmentParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	EnforcementMode *bool `json:"enforcementMode,omitempty" tf:"enforcement_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []PolicyAssignmentIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NotScopes []*string `json:"notScopes,omitempty" tf:"not_scopes,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters *string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	PolicyDefinitionID *string `json:"policyDefinitionId" tf:"policy_definition_id,omitempty"`

	// +kubebuilder:validation:Required
	Scope *string `json:"scope" tf:"scope,omitempty"`
}

// PolicyAssignmentSpec defines the desired state of PolicyAssignment
type PolicyAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyAssignmentParameters `json:"forProvider"`
}

// PolicyAssignmentStatus defines the observed state of PolicyAssignment.
type PolicyAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyAssignment is the Schema for the PolicyAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type PolicyAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyAssignmentSpec   `json:"spec"`
	Status            PolicyAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyAssignmentList contains a list of PolicyAssignments
type PolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyAssignment `json:"items"`
}

// Repository type metadata.
var (
	PolicyAssignment_Kind             = "PolicyAssignment"
	PolicyAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PolicyAssignment_Kind}.String()
	PolicyAssignment_KindAPIVersion   = PolicyAssignment_Kind + "." + CRDGroupVersion.String()
	PolicyAssignment_GroupVersionKind = CRDGroupVersion.WithKind(PolicyAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&PolicyAssignment{}, &PolicyAssignmentList{})
}

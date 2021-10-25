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

type PolicyRemediationObservation struct {
}

type PolicyRemediationParameters struct {

	// +kubebuilder:validation:Optional
	LocationFilters []*string `json:"locationFilters,omitempty" tf:"location_filters,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PolicyAssignmentID *string `json:"policyAssignmentId" tf:"policy_assignment_id,omitempty"`

	// +kubebuilder:validation:Optional
	PolicyDefinitionReferenceID *string `json:"policyDefinitionReferenceId,omitempty" tf:"policy_definition_reference_id,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceDiscoveryMode *string `json:"resourceDiscoveryMode,omitempty" tf:"resource_discovery_mode,omitempty"`

	// +kubebuilder:validation:Required
	Scope *string `json:"scope" tf:"scope,omitempty"`
}

// PolicyRemediationSpec defines the desired state of PolicyRemediation
type PolicyRemediationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyRemediationParameters `json:"forProvider"`
}

// PolicyRemediationStatus defines the observed state of PolicyRemediation.
type PolicyRemediationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyRemediationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyRemediation is the Schema for the PolicyRemediations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PolicyRemediation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyRemediationSpec   `json:"spec"`
	Status            PolicyRemediationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyRemediationList contains a list of PolicyRemediations
type PolicyRemediationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyRemediation `json:"items"`
}

// Repository type metadata.
var (
	PolicyRemediationKind             = "PolicyRemediation"
	PolicyRemediationGroupKind        = schema.GroupKind{Group: Group, Kind: PolicyRemediationKind}.String()
	PolicyRemediationKindAPIVersion   = PolicyRemediationKind + "." + GroupVersion.String()
	PolicyRemediationGroupVersionKind = GroupVersion.WithKind(PolicyRemediationKind)
)

func init() {
	SchemeBuilder.Register(&PolicyRemediation{}, &PolicyRemediationList{})
}

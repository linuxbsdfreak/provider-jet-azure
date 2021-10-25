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

type AuthorizationObservation struct {
}

type AuthorizationParameters struct {

	// +kubebuilder:validation:Optional
	DelegatedRoleDefinitionIds []*string `json:"delegatedRoleDefinitionIds,omitempty" tf:"delegated_role_definition_ids,omitempty"`

	// +kubebuilder:validation:Optional
	PrincipalDisplayName *string `json:"principalDisplayName,omitempty" tf:"principal_display_name,omitempty"`

	// +kubebuilder:validation:Required
	PrincipalID *string `json:"principalId" tf:"principal_id,omitempty"`

	// +kubebuilder:validation:Required
	RoleDefinitionID *string `json:"roleDefinitionId" tf:"role_definition_id,omitempty"`
}

type LighthouseDefinitionObservation struct {
}

type LighthouseDefinitionParameters struct {

	// +kubebuilder:validation:Required
	Authorization []AuthorizationParameters `json:"authorization" tf:"authorization,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	LighthouseDefinitionID *string `json:"lighthouseDefinitionId,omitempty" tf:"lighthouse_definition_id,omitempty"`

	// +kubebuilder:validation:Required
	ManagingTenantID *string `json:"managingTenantId" tf:"managing_tenant_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Plan []PlanParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// +kubebuilder:validation:Required
	Scope *string `json:"scope" tf:"scope,omitempty"`
}

type PlanObservation struct {
}

type PlanParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Product *string `json:"product" tf:"product,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

// LighthouseDefinitionSpec defines the desired state of LighthouseDefinition
type LighthouseDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LighthouseDefinitionParameters `json:"forProvider"`
}

// LighthouseDefinitionStatus defines the observed state of LighthouseDefinition.
type LighthouseDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LighthouseDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LighthouseDefinition is the Schema for the LighthouseDefinitions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LighthouseDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LighthouseDefinitionSpec   `json:"spec"`
	Status            LighthouseDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LighthouseDefinitionList contains a list of LighthouseDefinitions
type LighthouseDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LighthouseDefinition `json:"items"`
}

// Repository type metadata.
var (
	LighthouseDefinitionKind             = "LighthouseDefinition"
	LighthouseDefinitionGroupKind        = schema.GroupKind{Group: Group, Kind: LighthouseDefinitionKind}.String()
	LighthouseDefinitionKindAPIVersion   = LighthouseDefinitionKind + "." + GroupVersion.String()
	LighthouseDefinitionGroupVersionKind = GroupVersion.WithKind(LighthouseDefinitionKind)
)

func init() {
	SchemeBuilder.Register(&LighthouseDefinition{}, &LighthouseDefinitionList{})
}

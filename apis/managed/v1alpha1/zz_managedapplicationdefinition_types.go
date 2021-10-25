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

	// +kubebuilder:validation:Required
	RoleDefinitionID *string `json:"roleDefinitionId" tf:"role_definition_id,omitempty"`

	// +kubebuilder:validation:Required
	ServicePrincipalID *string `json:"servicePrincipalId" tf:"service_principal_id,omitempty"`
}

type ManagedApplicationDefinitionObservation struct {
}

type ManagedApplicationDefinitionParameters struct {

	// +kubebuilder:validation:Optional
	Authorization []AuthorizationParameters `json:"authorization,omitempty" tf:"authorization,omitempty"`

	// +kubebuilder:validation:Optional
	CreateUIDefinition *string `json:"createUiDefinition,omitempty" tf:"create_ui_definition,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	LockLevel *string `json:"lockLevel" tf:"lock_level,omitempty"`

	// +kubebuilder:validation:Optional
	MainTemplate *string `json:"mainTemplate,omitempty" tf:"main_template,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PackageEnabled *bool `json:"packageEnabled,omitempty" tf:"package_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	PackageFileURI *string `json:"packageFileUri,omitempty" tf:"package_file_uri,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ManagedApplicationDefinitionSpec defines the desired state of ManagedApplicationDefinition
type ManagedApplicationDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedApplicationDefinitionParameters `json:"forProvider"`
}

// ManagedApplicationDefinitionStatus defines the observed state of ManagedApplicationDefinition.
type ManagedApplicationDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedApplicationDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedApplicationDefinition is the Schema for the ManagedApplicationDefinitions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagedApplicationDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedApplicationDefinitionSpec   `json:"spec"`
	Status            ManagedApplicationDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedApplicationDefinitionList contains a list of ManagedApplicationDefinitions
type ManagedApplicationDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedApplicationDefinition `json:"items"`
}

// Repository type metadata.
var (
	ManagedApplicationDefinitionKind             = "ManagedApplicationDefinition"
	ManagedApplicationDefinitionGroupKind        = schema.GroupKind{Group: Group, Kind: ManagedApplicationDefinitionKind}.String()
	ManagedApplicationDefinitionKindAPIVersion   = ManagedApplicationDefinitionKind + "." + GroupVersion.String()
	ManagedApplicationDefinitionGroupVersionKind = GroupVersion.WithKind(ManagedApplicationDefinitionKind)
)

func init() {
	SchemeBuilder.Register(&ManagedApplicationDefinition{}, &ManagedApplicationDefinitionList{})
}

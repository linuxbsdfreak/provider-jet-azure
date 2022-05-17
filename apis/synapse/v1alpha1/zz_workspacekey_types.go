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

type WorkspaceKeyObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type WorkspaceKeyParameters struct {

	// +kubebuilder:validation:Required
	Active *bool `json:"active" tf:"active,omitempty"`

	// +kubebuilder:validation:Required
	CusomterManagedKeyName *string `json:"cusomterManagedKeyName" tf:"cusomter_managed_key_name,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerManagedKeyVersionlessID *string `json:"customerManagedKeyVersionlessId,omitempty" tf:"customer_managed_key_versionless_id,omitempty"`

	// +kubebuilder:validation:Required
	SynapseWorkspaceID *string `json:"synapseWorkspaceId" tf:"synapse_workspace_id,omitempty"`
}

// WorkspaceKeySpec defines the desired state of WorkspaceKey
type WorkspaceKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WorkspaceKeyParameters `json:"forProvider"`
}

// WorkspaceKeyStatus defines the observed state of WorkspaceKey.
type WorkspaceKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WorkspaceKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceKey is the Schema for the WorkspaceKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type WorkspaceKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceKeySpec   `json:"spec"`
	Status            WorkspaceKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorkspaceKeyList contains a list of WorkspaceKeys
type WorkspaceKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspaceKey `json:"items"`
}

// Repository type metadata.
var (
	WorkspaceKey_Kind             = "WorkspaceKey"
	WorkspaceKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WorkspaceKey_Kind}.String()
	WorkspaceKey_KindAPIVersion   = WorkspaceKey_Kind + "." + CRDGroupVersion.String()
	WorkspaceKey_GroupVersionKind = CRDGroupVersion.WithKind(WorkspaceKey_Kind)
)

func init() {
	SchemeBuilder.Register(&WorkspaceKey{}, &WorkspaceKeyList{})
}

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

type AceObservation struct {
}

type AceParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Permissions *string `json:"permissions" tf:"permissions,omitempty"`

	// +kubebuilder:validation:Optional
	Scope *string `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type DataLakeGen2FileSystemObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataLakeGen2FileSystemParameters struct {

	// +kubebuilder:validation:Optional
	Ace []AceParameters `json:"ace,omitempty" tf:"ace,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

// DataLakeGen2FileSystemSpec defines the desired state of DataLakeGen2FileSystem
type DataLakeGen2FileSystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataLakeGen2FileSystemParameters `json:"forProvider"`
}

// DataLakeGen2FileSystemStatus defines the observed state of DataLakeGen2FileSystem.
type DataLakeGen2FileSystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataLakeGen2FileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataLakeGen2FileSystem is the Schema for the DataLakeGen2FileSystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DataLakeGen2FileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataLakeGen2FileSystemSpec   `json:"spec"`
	Status            DataLakeGen2FileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataLakeGen2FileSystemList contains a list of DataLakeGen2FileSystems
type DataLakeGen2FileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataLakeGen2FileSystem `json:"items"`
}

// Repository type metadata.
var (
	DataLakeGen2FileSystem_Kind             = "DataLakeGen2FileSystem"
	DataLakeGen2FileSystem_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataLakeGen2FileSystem_Kind}.String()
	DataLakeGen2FileSystem_KindAPIVersion   = DataLakeGen2FileSystem_Kind + "." + CRDGroupVersion.String()
	DataLakeGen2FileSystem_GroupVersionKind = CRDGroupVersion.WithKind(DataLakeGen2FileSystem_Kind)
)

func init() {
	SchemeBuilder.Register(&DataLakeGen2FileSystem{}, &DataLakeGen2FileSystemList{})
}

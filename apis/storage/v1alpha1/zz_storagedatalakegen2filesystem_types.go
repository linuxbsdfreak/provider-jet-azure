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

type StorageDataLakeGen2FilesystemObservation struct {
}

type StorageDataLakeGen2FilesystemParameters struct {

	// +kubebuilder:validation:Optional
	Ace []AceParameters `json:"ace,omitempty" tf:"ace,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

// StorageDataLakeGen2FilesystemSpec defines the desired state of StorageDataLakeGen2Filesystem
type StorageDataLakeGen2FilesystemSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageDataLakeGen2FilesystemParameters `json:"forProvider"`
}

// StorageDataLakeGen2FilesystemStatus defines the observed state of StorageDataLakeGen2Filesystem.
type StorageDataLakeGen2FilesystemStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageDataLakeGen2FilesystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StorageDataLakeGen2Filesystem is the Schema for the StorageDataLakeGen2Filesystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StorageDataLakeGen2Filesystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageDataLakeGen2FilesystemSpec   `json:"spec"`
	Status            StorageDataLakeGen2FilesystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageDataLakeGen2FilesystemList contains a list of StorageDataLakeGen2Filesystems
type StorageDataLakeGen2FilesystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageDataLakeGen2Filesystem `json:"items"`
}

// Repository type metadata.
var (
	StorageDataLakeGen2FilesystemKind             = "StorageDataLakeGen2Filesystem"
	StorageDataLakeGen2FilesystemGroupKind        = schema.GroupKind{Group: Group, Kind: StorageDataLakeGen2FilesystemKind}.String()
	StorageDataLakeGen2FilesystemKindAPIVersion   = StorageDataLakeGen2FilesystemKind + "." + GroupVersion.String()
	StorageDataLakeGen2FilesystemGroupVersionKind = GroupVersion.WithKind(StorageDataLakeGen2FilesystemKind)
)

func init() {
	SchemeBuilder.Register(&StorageDataLakeGen2Filesystem{}, &StorageDataLakeGen2FilesystemList{})
}

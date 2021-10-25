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

type ApiManagementApiReleaseObservation struct {
}

type ApiManagementApiReleaseParameters struct {

	// +kubebuilder:validation:Required
	APIID *string `json:"apiId" tf:"api_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`
}

// ApiManagementApiReleaseSpec defines the desired state of ApiManagementApiRelease
type ApiManagementApiReleaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApiManagementApiReleaseParameters `json:"forProvider"`
}

// ApiManagementApiReleaseStatus defines the observed state of ApiManagementApiRelease.
type ApiManagementApiReleaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApiManagementApiReleaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementApiRelease is the Schema for the ApiManagementApiReleases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementApiRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementApiReleaseSpec   `json:"spec"`
	Status            ApiManagementApiReleaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementApiReleaseList contains a list of ApiManagementApiReleases
type ApiManagementApiReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementApiRelease `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementApiReleaseKind             = "ApiManagementApiRelease"
	ApiManagementApiReleaseGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementApiReleaseKind}.String()
	ApiManagementApiReleaseKindAPIVersion   = ApiManagementApiReleaseKind + "." + GroupVersion.String()
	ApiManagementApiReleaseGroupVersionKind = GroupVersion.WithKind(ApiManagementApiReleaseKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementApiRelease{}, &ApiManagementApiReleaseList{})
}

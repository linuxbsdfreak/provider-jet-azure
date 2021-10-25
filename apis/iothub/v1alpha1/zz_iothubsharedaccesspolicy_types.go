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

type IothubSharedAccessPolicyObservation struct {
}

type IothubSharedAccessPolicyParameters struct {

	// +kubebuilder:validation:Optional
	DeviceConnect *bool `json:"deviceConnect,omitempty" tf:"device_connect,omitempty"`

	// +kubebuilder:validation:Required
	IothubName *string `json:"iothubName" tf:"iothub_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryRead *bool `json:"registryRead,omitempty" tf:"registry_read,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryWrite *bool `json:"registryWrite,omitempty" tf:"registry_write,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceConnect *bool `json:"serviceConnect,omitempty" tf:"service_connect,omitempty"`
}

// IothubSharedAccessPolicySpec defines the desired state of IothubSharedAccessPolicy
type IothubSharedAccessPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IothubSharedAccessPolicyParameters `json:"forProvider"`
}

// IothubSharedAccessPolicyStatus defines the observed state of IothubSharedAccessPolicy.
type IothubSharedAccessPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IothubSharedAccessPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IothubSharedAccessPolicy is the Schema for the IothubSharedAccessPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IothubSharedAccessPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubSharedAccessPolicySpec   `json:"spec"`
	Status            IothubSharedAccessPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IothubSharedAccessPolicyList contains a list of IothubSharedAccessPolicys
type IothubSharedAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IothubSharedAccessPolicy `json:"items"`
}

// Repository type metadata.
var (
	IothubSharedAccessPolicyKind             = "IothubSharedAccessPolicy"
	IothubSharedAccessPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: IothubSharedAccessPolicyKind}.String()
	IothubSharedAccessPolicyKindAPIVersion   = IothubSharedAccessPolicyKind + "." + GroupVersion.String()
	IothubSharedAccessPolicyGroupVersionKind = GroupVersion.WithKind(IothubSharedAccessPolicyKind)
)

func init() {
	SchemeBuilder.Register(&IothubSharedAccessPolicy{}, &IothubSharedAccessPolicyList{})
}

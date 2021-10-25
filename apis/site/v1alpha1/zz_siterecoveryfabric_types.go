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

type SiteRecoveryFabricObservation struct {
}

type SiteRecoveryFabricParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryVaultName *string `json:"recoveryVaultName" tf:"recovery_vault_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// SiteRecoveryFabricSpec defines the desired state of SiteRecoveryFabric
type SiteRecoveryFabricSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteRecoveryFabricParameters `json:"forProvider"`
}

// SiteRecoveryFabricStatus defines the observed state of SiteRecoveryFabric.
type SiteRecoveryFabricStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteRecoveryFabricObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryFabric is the Schema for the SiteRecoveryFabrics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SiteRecoveryFabric struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SiteRecoveryFabricSpec   `json:"spec"`
	Status            SiteRecoveryFabricStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteRecoveryFabricList contains a list of SiteRecoveryFabrics
type SiteRecoveryFabricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteRecoveryFabric `json:"items"`
}

// Repository type metadata.
var (
	SiteRecoveryFabricKind             = "SiteRecoveryFabric"
	SiteRecoveryFabricGroupKind        = schema.GroupKind{Group: Group, Kind: SiteRecoveryFabricKind}.String()
	SiteRecoveryFabricKindAPIVersion   = SiteRecoveryFabricKind + "." + GroupVersion.String()
	SiteRecoveryFabricGroupVersionKind = GroupVersion.WithKind(SiteRecoveryFabricKind)
)

func init() {
	SchemeBuilder.Register(&SiteRecoveryFabric{}, &SiteRecoveryFabricList{})
}

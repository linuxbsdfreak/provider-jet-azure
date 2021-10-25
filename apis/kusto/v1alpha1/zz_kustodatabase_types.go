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

type KustoDatabaseObservation struct {
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`
}

type KustoDatabaseParameters struct {

	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Optional
	HotCachePeriod *string `json:"hotCachePeriod,omitempty" tf:"hot_cache_period,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	SoftDeletePeriod *string `json:"softDeletePeriod,omitempty" tf:"soft_delete_period,omitempty"`
}

// KustoDatabaseSpec defines the desired state of KustoDatabase
type KustoDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KustoDatabaseParameters `json:"forProvider"`
}

// KustoDatabaseStatus defines the observed state of KustoDatabase.
type KustoDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KustoDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KustoDatabase is the Schema for the KustoDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type KustoDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KustoDatabaseSpec   `json:"spec"`
	Status            KustoDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KustoDatabaseList contains a list of KustoDatabases
type KustoDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KustoDatabase `json:"items"`
}

// Repository type metadata.
var (
	KustoDatabaseKind             = "KustoDatabase"
	KustoDatabaseGroupKind        = schema.GroupKind{Group: Group, Kind: KustoDatabaseKind}.String()
	KustoDatabaseKindAPIVersion   = KustoDatabaseKind + "." + GroupVersion.String()
	KustoDatabaseGroupVersionKind = GroupVersion.WithKind(KustoDatabaseKind)
)

func init() {
	SchemeBuilder.Register(&KustoDatabase{}, &KustoDatabaseList{})
}

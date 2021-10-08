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

type CosmosdbSqlFunctionObservation struct {
}

type CosmosdbSqlFunctionParameters struct {

	// +kubebuilder:validation:Required
	Body *string `json:"body" tf:"body"`

	// +crossplane:generate:reference:type=CosmosdbSqlContainer
	// +kubebuilder:validation:Optional
	ContainerID *string `json:"containerId,omitempty" tf:"container_id"`

	// +kubebuilder:validation:Optional
	ContainerIDRef *v1.Reference `json:"containerIDRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ContainerIDSelector *v1.Selector `json:"containerIDSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name"`
}

// CosmosdbSqlFunctionSpec defines the desired state of CosmosdbSqlFunction
type CosmosdbSqlFunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CosmosdbSqlFunctionParameters `json:"forProvider"`
}

// CosmosdbSqlFunctionStatus defines the observed state of CosmosdbSqlFunction.
type CosmosdbSqlFunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CosmosdbSqlFunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CosmosdbSqlFunction is the Schema for the CosmosdbSqlFunctions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CosmosdbSqlFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CosmosdbSqlFunctionSpec   `json:"spec"`
	Status            CosmosdbSqlFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CosmosdbSqlFunctionList contains a list of CosmosdbSqlFunctions
type CosmosdbSqlFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CosmosdbSqlFunction `json:"items"`
}

// Repository type metadata.
var (
	CosmosdbSqlFunctionKind             = "CosmosdbSqlFunction"
	CosmosdbSqlFunctionGroupKind        = schema.GroupKind{Group: Group, Kind: CosmosdbSqlFunctionKind}.String()
	CosmosdbSqlFunctionKindAPIVersion   = CosmosdbSqlFunctionKind + "." + GroupVersion.String()
	CosmosdbSqlFunctionGroupVersionKind = GroupVersion.WithKind(CosmosdbSqlFunctionKind)
)

func init() {
	SchemeBuilder.Register(&CosmosdbSqlFunction{}, &CosmosdbSqlFunctionList{})
}
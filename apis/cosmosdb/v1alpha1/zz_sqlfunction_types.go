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

type SqlFunctionObservation struct {
}

type SqlFunctionParameters struct {

	// +kubebuilder:validation:Required
	Body *string `json:"body" tf:"body,omitempty"`

	// +crossplane:generate:reference:type=SqlContainer
	// +kubebuilder:validation:Optional
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// +kubebuilder:validation:Optional
	ContainerIDRef *v1.Reference `json:"containerIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ContainerIDSelector *v1.Selector `json:"containerIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// SqlFunctionSpec defines the desired state of SqlFunction
type SqlFunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SqlFunctionParameters `json:"forProvider"`
}

// SqlFunctionStatus defines the observed state of SqlFunction.
type SqlFunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SqlFunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SqlFunction is the Schema for the SqlFunctions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type SqlFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlFunctionSpec   `json:"spec"`
	Status            SqlFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlFunctionList contains a list of SqlFunctions
type SqlFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlFunction `json:"items"`
}

// Repository type metadata.
var (
	SqlFunction_Kind             = "SqlFunction"
	SqlFunction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SqlFunction_Kind}.String()
	SqlFunction_KindAPIVersion   = SqlFunction_Kind + "." + CRDGroupVersion.String()
	SqlFunction_GroupVersionKind = CRDGroupVersion.WithKind(SqlFunction_Kind)
)

func init() {
	SchemeBuilder.Register(&SqlFunction{}, &SqlFunctionList{})
}

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

type IothubFallbackRouteObservation struct {
}

type IothubFallbackRouteParameters struct {

	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	EndpointNames []*string `json:"endpointNames" tf:"endpoint_names,omitempty"`

	// +kubebuilder:validation:Required
	IothubName *string `json:"iothubName" tf:"iothub_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// IothubFallbackRouteSpec defines the desired state of IothubFallbackRoute
type IothubFallbackRouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IothubFallbackRouteParameters `json:"forProvider"`
}

// IothubFallbackRouteStatus defines the observed state of IothubFallbackRoute.
type IothubFallbackRouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IothubFallbackRouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IothubFallbackRoute is the Schema for the IothubFallbackRoutes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IothubFallbackRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubFallbackRouteSpec   `json:"spec"`
	Status            IothubFallbackRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IothubFallbackRouteList contains a list of IothubFallbackRoutes
type IothubFallbackRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IothubFallbackRoute `json:"items"`
}

// Repository type metadata.
var (
	IothubFallbackRouteKind             = "IothubFallbackRoute"
	IothubFallbackRouteGroupKind        = schema.GroupKind{Group: Group, Kind: IothubFallbackRouteKind}.String()
	IothubFallbackRouteKindAPIVersion   = IothubFallbackRouteKind + "." + GroupVersion.String()
	IothubFallbackRouteGroupVersionKind = GroupVersion.WithKind(IothubFallbackRouteKind)
)

func init() {
	SchemeBuilder.Register(&IothubFallbackRoute{}, &IothubFallbackRouteList{})
}

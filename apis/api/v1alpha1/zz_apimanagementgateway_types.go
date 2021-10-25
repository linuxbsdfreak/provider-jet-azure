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

type ApiManagementGatewayObservation struct {
}

type ApiManagementGatewayParameters struct {

	// +kubebuilder:validation:Required
	APIManagementID *string `json:"apiManagementId" tf:"api_management_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	LocationData []LocationDataParameters `json:"locationData" tf:"location_data,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type LocationDataObservation struct {
}

type LocationDataParameters struct {

	// +kubebuilder:validation:Optional
	City *string `json:"city,omitempty" tf:"city,omitempty"`

	// +kubebuilder:validation:Optional
	District *string `json:"district,omitempty" tf:"district,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// ApiManagementGatewaySpec defines the desired state of ApiManagementGateway
type ApiManagementGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApiManagementGatewayParameters `json:"forProvider"`
}

// ApiManagementGatewayStatus defines the observed state of ApiManagementGateway.
type ApiManagementGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApiManagementGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementGateway is the Schema for the ApiManagementGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementGatewaySpec   `json:"spec"`
	Status            ApiManagementGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementGatewayList contains a list of ApiManagementGateways
type ApiManagementGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementGateway `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementGatewayKind             = "ApiManagementGateway"
	ApiManagementGatewayGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementGatewayKind}.String()
	ApiManagementGatewayKindAPIVersion   = ApiManagementGatewayKind + "." + GroupVersion.String()
	ApiManagementGatewayGroupVersionKind = GroupVersion.WithKind(ApiManagementGatewayKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementGateway{}, &ApiManagementGatewayList{})
}

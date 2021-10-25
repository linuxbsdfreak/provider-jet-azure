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

type BgpSettingsObservation struct {
}

type BgpSettingsParameters struct {

	// +kubebuilder:validation:Required
	Asn *int64 `json:"asn" tf:"asn,omitempty"`

	// +kubebuilder:validation:Required
	BgpPeeringAddress *string `json:"bgpPeeringAddress" tf:"bgp_peering_address,omitempty"`

	// +kubebuilder:validation:Optional
	PeerWeight *int64 `json:"peerWeight,omitempty" tf:"peer_weight,omitempty"`
}

type LocalNetworkGatewayObservation struct {
}

type LocalNetworkGatewayParameters struct {

	// +kubebuilder:validation:Optional
	AddressSpace []*string `json:"addressSpace,omitempty" tf:"address_space,omitempty"`

	// +kubebuilder:validation:Optional
	BgpSettings []BgpSettingsParameters `json:"bgpSettings,omitempty" tf:"bgp_settings,omitempty"`

	// +kubebuilder:validation:Optional
	GatewayAddress *string `json:"gatewayAddress,omitempty" tf:"gateway_address,omitempty"`

	// +kubebuilder:validation:Optional
	GatewayFqdn *string `json:"gatewayFqdn,omitempty" tf:"gateway_fqdn,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LocalNetworkGatewaySpec defines the desired state of LocalNetworkGateway
type LocalNetworkGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocalNetworkGatewayParameters `json:"forProvider"`
}

// LocalNetworkGatewayStatus defines the observed state of LocalNetworkGateway.
type LocalNetworkGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocalNetworkGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LocalNetworkGateway is the Schema for the LocalNetworkGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LocalNetworkGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocalNetworkGatewaySpec   `json:"spec"`
	Status            LocalNetworkGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocalNetworkGatewayList contains a list of LocalNetworkGateways
type LocalNetworkGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocalNetworkGateway `json:"items"`
}

// Repository type metadata.
var (
	LocalNetworkGatewayKind             = "LocalNetworkGateway"
	LocalNetworkGatewayGroupKind        = schema.GroupKind{Group: Group, Kind: LocalNetworkGatewayKind}.String()
	LocalNetworkGatewayKindAPIVersion   = LocalNetworkGatewayKind + "." + GroupVersion.String()
	LocalNetworkGatewayGroupVersionKind = GroupVersion.WithKind(LocalNetworkGatewayKind)
)

func init() {
	SchemeBuilder.Register(&LocalNetworkGateway{}, &LocalNetworkGatewayList{})
}

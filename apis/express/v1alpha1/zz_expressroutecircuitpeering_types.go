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

type ExpressRouteCircuitPeeringObservation struct {
	AzureAsn *int64 `json:"azureAsn,omitempty" tf:"azure_asn,omitempty"`

	PrimaryAzurePort *string `json:"primaryAzurePort,omitempty" tf:"primary_azure_port,omitempty"`

	SecondaryAzurePort *string `json:"secondaryAzurePort,omitempty" tf:"secondary_azure_port,omitempty"`
}

type ExpressRouteCircuitPeeringParameters struct {

	// +kubebuilder:validation:Required
	ExpressRouteCircuitName *string `json:"expressRouteCircuitName" tf:"express_route_circuit_name,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6 []IPv6Parameters `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// +kubebuilder:validation:Optional
	MicrosoftPeeringConfig []MicrosoftPeeringConfigParameters `json:"microsoftPeeringConfig,omitempty" tf:"microsoft_peering_config,omitempty"`

	// +kubebuilder:validation:Optional
	PeerAsn *int64 `json:"peerAsn,omitempty" tf:"peer_asn,omitempty"`

	// +kubebuilder:validation:Required
	PeeringType *string `json:"peeringType" tf:"peering_type,omitempty"`

	// +kubebuilder:validation:Required
	PrimaryPeerAddressPrefix *string `json:"primaryPeerAddressPrefix" tf:"primary_peer_address_prefix,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	RouteFilterID *string `json:"routeFilterId,omitempty" tf:"route_filter_id,omitempty"`

	// +kubebuilder:validation:Required
	SecondaryPeerAddressPrefix *string `json:"secondaryPeerAddressPrefix" tf:"secondary_peer_address_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	SharedKeySecretRef v1.SecretKeySelector `json:"sharedKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	VlanID *int64 `json:"vlanId" tf:"vlan_id,omitempty"`
}

type IPv6Observation struct {
}

type IPv6Parameters struct {

	// +kubebuilder:validation:Required
	MicrosoftPeering []MicrosoftPeeringParameters `json:"microsoftPeering" tf:"microsoft_peering,omitempty"`

	// +kubebuilder:validation:Required
	PrimaryPeerAddressPrefix *string `json:"primaryPeerAddressPrefix" tf:"primary_peer_address_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	RouteFilterID *string `json:"routeFilterId,omitempty" tf:"route_filter_id,omitempty"`

	// +kubebuilder:validation:Required
	SecondaryPeerAddressPrefix *string `json:"secondaryPeerAddressPrefix" tf:"secondary_peer_address_prefix,omitempty"`
}

type MicrosoftPeeringConfigObservation struct {
}

type MicrosoftPeeringConfigParameters struct {

	// +kubebuilder:validation:Required
	AdvertisedPublicPrefixes []*string `json:"advertisedPublicPrefixes" tf:"advertised_public_prefixes,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerAsn *int64 `json:"customerAsn,omitempty" tf:"customer_asn,omitempty"`

	// +kubebuilder:validation:Optional
	RoutingRegistryName *string `json:"routingRegistryName,omitempty" tf:"routing_registry_name,omitempty"`
}

type MicrosoftPeeringObservation struct {
}

type MicrosoftPeeringParameters struct {

	// +kubebuilder:validation:Optional
	AdvertisedPublicPrefixes []*string `json:"advertisedPublicPrefixes,omitempty" tf:"advertised_public_prefixes,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerAsn *int64 `json:"customerAsn,omitempty" tf:"customer_asn,omitempty"`

	// +kubebuilder:validation:Optional
	RoutingRegistryName *string `json:"routingRegistryName,omitempty" tf:"routing_registry_name,omitempty"`
}

// ExpressRouteCircuitPeeringSpec defines the desired state of ExpressRouteCircuitPeering
type ExpressRouteCircuitPeeringSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRouteCircuitPeeringParameters `json:"forProvider"`
}

// ExpressRouteCircuitPeeringStatus defines the observed state of ExpressRouteCircuitPeering.
type ExpressRouteCircuitPeeringStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRouteCircuitPeeringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitPeering is the Schema for the ExpressRouteCircuitPeerings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRouteCircuitPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitPeeringSpec   `json:"spec"`
	Status            ExpressRouteCircuitPeeringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitPeeringList contains a list of ExpressRouteCircuitPeerings
type ExpressRouteCircuitPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRouteCircuitPeering `json:"items"`
}

// Repository type metadata.
var (
	ExpressRouteCircuitPeeringKind             = "ExpressRouteCircuitPeering"
	ExpressRouteCircuitPeeringGroupKind        = schema.GroupKind{Group: Group, Kind: ExpressRouteCircuitPeeringKind}.String()
	ExpressRouteCircuitPeeringKindAPIVersion   = ExpressRouteCircuitPeeringKind + "." + GroupVersion.String()
	ExpressRouteCircuitPeeringGroupVersionKind = GroupVersion.WithKind(ExpressRouteCircuitPeeringKind)
)

func init() {
	SchemeBuilder.Register(&ExpressRouteCircuitPeering{}, &ExpressRouteCircuitPeeringList{})
}

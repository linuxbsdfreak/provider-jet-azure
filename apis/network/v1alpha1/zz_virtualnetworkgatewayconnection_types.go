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

type IpsecPolicyObservation struct {
}

type IpsecPolicyParameters struct {

	// +kubebuilder:validation:Required
	DhGroup *string `json:"dhGroup" tf:"dh_group,omitempty"`

	// +kubebuilder:validation:Required
	IkeEncryption *string `json:"ikeEncryption" tf:"ike_encryption,omitempty"`

	// +kubebuilder:validation:Required
	IkeIntegrity *string `json:"ikeIntegrity" tf:"ike_integrity,omitempty"`

	// +kubebuilder:validation:Required
	IpsecEncryption *string `json:"ipsecEncryption" tf:"ipsec_encryption,omitempty"`

	// +kubebuilder:validation:Required
	IpsecIntegrity *string `json:"ipsecIntegrity" tf:"ipsec_integrity,omitempty"`

	// +kubebuilder:validation:Required
	PfsGroup *string `json:"pfsGroup" tf:"pfs_group,omitempty"`

	// +kubebuilder:validation:Optional
	SaDatasize *int64 `json:"saDatasize,omitempty" tf:"sa_datasize,omitempty"`

	// +kubebuilder:validation:Optional
	SaLifetime *int64 `json:"saLifetime,omitempty" tf:"sa_lifetime,omitempty"`
}

type TrafficSelectorPolicyObservation struct {
}

type TrafficSelectorPolicyParameters struct {

	// +kubebuilder:validation:Required
	LocalAddressCidrs []*string `json:"localAddressCidrs" tf:"local_address_cidrs,omitempty"`

	// +kubebuilder:validation:Required
	RemoteAddressCidrs []*string `json:"remoteAddressCidrs" tf:"remote_address_cidrs,omitempty"`
}

type VirtualNetworkGatewayConnectionObservation struct {
}

type VirtualNetworkGatewayConnectionParameters struct {

	// +kubebuilder:validation:Optional
	AuthorizationKeySecretRef *v1.SecretKeySelector `json:"authorizationKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ConnectionProtocol *string `json:"connectionProtocol,omitempty" tf:"connection_protocol,omitempty"`

	// +kubebuilder:validation:Optional
	DpdTimeoutSeconds *int64 `json:"dpdTimeoutSeconds,omitempty" tf:"dpd_timeout_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	EnableBgp *bool `json:"enableBgp,omitempty" tf:"enable_bgp,omitempty"`

	// +kubebuilder:validation:Optional
	ExpressRouteCircuitID *string `json:"expressRouteCircuitId,omitempty" tf:"express_route_circuit_id,omitempty"`

	// +kubebuilder:validation:Optional
	ExpressRouteGatewayBypass *bool `json:"expressRouteGatewayBypass,omitempty" tf:"express_route_gateway_bypass,omitempty"`

	// +kubebuilder:validation:Optional
	IpsecPolicy []IpsecPolicyParameters `json:"ipsecPolicy,omitempty" tf:"ipsec_policy,omitempty"`

	// +kubebuilder:validation:Optional
	LocalAzureIPAddressEnabled *bool `json:"localAzureIpAddressEnabled,omitempty" tf:"local_azure_ip_address_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LocalNetworkGatewayID *string `json:"localNetworkGatewayId,omitempty" tf:"local_network_gateway_id,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +crossplane:generate:reference:type=VirtualNetworkGateway
	// +kubebuilder:validation:Optional
	PeerVirtualNetworkGatewayID *string `json:"peerVirtualNetworkGatewayId,omitempty" tf:"peer_virtual_network_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	PeerVirtualNetworkGatewayIDRef *v1.Reference `json:"peerVirtualNetworkGatewayIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	PeerVirtualNetworkGatewayIDSelector *v1.Selector `json:"peerVirtualNetworkGatewayIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/azure/v1alpha1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoutingWeight *int64 `json:"routingWeight,omitempty" tf:"routing_weight,omitempty"`

	// +kubebuilder:validation:Optional
	SharedKeySecretRef *v1.SecretKeySelector `json:"sharedKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TrafficSelectorPolicy []TrafficSelectorPolicyParameters `json:"trafficSelectorPolicy,omitempty" tf:"traffic_selector_policy,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	UsePolicyBasedTrafficSelectors *bool `json:"usePolicyBasedTrafficSelectors,omitempty" tf:"use_policy_based_traffic_selectors,omitempty"`

	// +crossplane:generate:reference:type=VirtualNetworkGateway
	// +kubebuilder:validation:Optional
	VirtualNetworkGatewayID *string `json:"virtualNetworkGatewayId,omitempty" tf:"virtual_network_gateway_id,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkGatewayIDRef *v1.Reference `json:"virtualNetworkGatewayIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VirtualNetworkGatewayIDSelector *v1.Selector `json:"virtualNetworkGatewayIdSelector,omitempty" tf:"-"`
}

// VirtualNetworkGatewayConnectionSpec defines the desired state of VirtualNetworkGatewayConnection
type VirtualNetworkGatewayConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNetworkGatewayConnectionParameters `json:"forProvider"`
}

// VirtualNetworkGatewayConnectionStatus defines the observed state of VirtualNetworkGatewayConnection.
type VirtualNetworkGatewayConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNetworkGatewayConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkGatewayConnection is the Schema for the VirtualNetworkGatewayConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type VirtualNetworkGatewayConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkGatewayConnectionSpec   `json:"spec"`
	Status            VirtualNetworkGatewayConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkGatewayConnectionList contains a list of VirtualNetworkGatewayConnections
type VirtualNetworkGatewayConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkGatewayConnection `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetworkGatewayConnection_Kind             = "VirtualNetworkGatewayConnection"
	VirtualNetworkGatewayConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualNetworkGatewayConnection_Kind}.String()
	VirtualNetworkGatewayConnection_KindAPIVersion   = VirtualNetworkGatewayConnection_Kind + "." + CRDGroupVersion.String()
	VirtualNetworkGatewayConnection_GroupVersionKind = CRDGroupVersion.WithKind(VirtualNetworkGatewayConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetworkGatewayConnection{}, &VirtualNetworkGatewayConnectionList{})
}

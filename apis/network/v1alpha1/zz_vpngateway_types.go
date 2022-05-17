/*
Copyright 2022 The Crossplane Authors.

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

type Instance0BGPPeeringAddressObservation struct {
	DefaultIps []*string `json:"defaultIps,omitempty" tf:"default_ips,omitempty"`

	IPConfigurationID *string `json:"ipConfigurationId,omitempty" tf:"ip_configuration_id,omitempty"`

	TunnelIps []*string `json:"tunnelIps,omitempty" tf:"tunnel_ips,omitempty"`
}

type Instance0BGPPeeringAddressParameters struct {

	// +kubebuilder:validation:Required
	CustomIps []*string `json:"customIps" tf:"custom_ips,omitempty"`
}

type Instance1BGPPeeringAddressObservation struct {
	DefaultIps []*string `json:"defaultIps,omitempty" tf:"default_ips,omitempty"`

	IPConfigurationID *string `json:"ipConfigurationId,omitempty" tf:"ip_configuration_id,omitempty"`

	TunnelIps []*string `json:"tunnelIps,omitempty" tf:"tunnel_ips,omitempty"`
}

type Instance1BGPPeeringAddressParameters struct {

	// +kubebuilder:validation:Required
	CustomIps []*string `json:"customIps" tf:"custom_ips,omitempty"`
}

type VPNGatewayBGPSettingsObservation struct {
	BGPPeeringAddress *string `json:"bgpPeeringAddress,omitempty" tf:"bgp_peering_address,omitempty"`
}

type VPNGatewayBGPSettingsParameters struct {

	// +kubebuilder:validation:Required
	Asn *float64 `json:"asn" tf:"asn,omitempty"`

	// +kubebuilder:validation:Optional
	Instance0BGPPeeringAddress []Instance0BGPPeeringAddressParameters `json:"instance0BgpPeeringAddress,omitempty" tf:"instance_0_bgp_peering_address,omitempty"`

	// +kubebuilder:validation:Optional
	Instance1BGPPeeringAddress []Instance1BGPPeeringAddressParameters `json:"instance1BgpPeeringAddress,omitempty" tf:"instance_1_bgp_peering_address,omitempty"`

	// +kubebuilder:validation:Required
	PeerWeight *float64 `json:"peerWeight" tf:"peer_weight,omitempty"`
}

type VPNGatewayObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VPNGatewayParameters struct {

	// +kubebuilder:validation:Optional
	BGPSettings []VPNGatewayBGPSettingsParameters `json:"bgpSettings,omitempty" tf:"bgp_settings,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ScaleUnit *float64 `json:"scaleUnit,omitempty" tf:"scale_unit,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	VirtualHubID *string `json:"virtualHubId" tf:"virtual_hub_id,omitempty"`
}

// VPNGatewaySpec defines the desired state of VPNGateway
type VPNGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPNGatewayParameters `json:"forProvider"`
}

// VPNGatewayStatus defines the observed state of VPNGateway.
type VPNGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPNGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGateway is the Schema for the VPNGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VPNGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPNGatewaySpec   `json:"spec"`
	Status            VPNGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPNGatewayList contains a list of VPNGateways
type VPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPNGateway `json:"items"`
}

// Repository type metadata.
var (
	VPNGateway_Kind             = "VPNGateway"
	VPNGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPNGateway_Kind}.String()
	VPNGateway_KindAPIVersion   = VPNGateway_Kind + "." + CRDGroupVersion.String()
	VPNGateway_GroupVersionKind = CRDGroupVersion.WithKind(VPNGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&VPNGateway{}, &VPNGatewayList{})
}

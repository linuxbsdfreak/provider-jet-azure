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

type ExpressRoutePortIdentityObservation struct {
}

type ExpressRoutePortIdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ExpressRoutePortObservation struct {
	Ethertype *string `json:"ethertype,omitempty" tf:"ethertype,omitempty"`

	GUID *string `json:"guid,omitempty" tf:"guid,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Link1 []Link1Observation `json:"link1,omitempty" tf:"link1,omitempty"`

	Link2 []Link2Observation `json:"link2,omitempty" tf:"link2,omitempty"`

	Mtu *string `json:"mtu,omitempty" tf:"mtu,omitempty"`
}

type ExpressRoutePortParameters struct {

	// +kubebuilder:validation:Required
	BandwidthInGbps *float64 `json:"bandwidthInGbps" tf:"bandwidth_in_gbps,omitempty"`

	// +kubebuilder:validation:Required
	Encapsulation *string `json:"encapsulation" tf:"encapsulation,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []ExpressRoutePortIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Optional
	Link1 []Link1Parameters `json:"link1,omitempty" tf:"link1,omitempty"`

	// +kubebuilder:validation:Optional
	Link2 []Link2Parameters `json:"link2,omitempty" tf:"link2,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PeeringLocation *string `json:"peeringLocation" tf:"peering_location,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type Link1Observation struct {
	ConnectorType *string `json:"connectorType,omitempty" tf:"connector_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InterfaceName *string `json:"interfaceName,omitempty" tf:"interface_name,omitempty"`

	PatchPanelID *string `json:"patchPanelId,omitempty" tf:"patch_panel_id,omitempty"`

	RackID *string `json:"rackId,omitempty" tf:"rack_id,omitempty"`

	RouterName *string `json:"routerName,omitempty" tf:"router_name,omitempty"`
}

type Link1Parameters struct {

	// +kubebuilder:validation:Optional
	AdminEnabled *bool `json:"adminEnabled,omitempty" tf:"admin_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	MacsecCakKeyvaultSecretID *string `json:"macsecCakKeyvaultSecretId,omitempty" tf:"macsec_cak_keyvault_secret_id,omitempty"`

	// +kubebuilder:validation:Optional
	MacsecCipher *string `json:"macsecCipher,omitempty" tf:"macsec_cipher,omitempty"`

	// +kubebuilder:validation:Optional
	MacsecCknKeyvaultSecretID *string `json:"macsecCknKeyvaultSecretId,omitempty" tf:"macsec_ckn_keyvault_secret_id,omitempty"`
}

type Link2Observation struct {
	ConnectorType *string `json:"connectorType,omitempty" tf:"connector_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	InterfaceName *string `json:"interfaceName,omitempty" tf:"interface_name,omitempty"`

	PatchPanelID *string `json:"patchPanelId,omitempty" tf:"patch_panel_id,omitempty"`

	RackID *string `json:"rackId,omitempty" tf:"rack_id,omitempty"`

	RouterName *string `json:"routerName,omitempty" tf:"router_name,omitempty"`
}

type Link2Parameters struct {

	// +kubebuilder:validation:Optional
	AdminEnabled *bool `json:"adminEnabled,omitempty" tf:"admin_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	MacsecCakKeyvaultSecretID *string `json:"macsecCakKeyvaultSecretId,omitempty" tf:"macsec_cak_keyvault_secret_id,omitempty"`

	// +kubebuilder:validation:Optional
	MacsecCipher *string `json:"macsecCipher,omitempty" tf:"macsec_cipher,omitempty"`

	// +kubebuilder:validation:Optional
	MacsecCknKeyvaultSecretID *string `json:"macsecCknKeyvaultSecretId,omitempty" tf:"macsec_ckn_keyvault_secret_id,omitempty"`
}

// ExpressRoutePortSpec defines the desired state of ExpressRoutePort
type ExpressRoutePortSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRoutePortParameters `json:"forProvider"`
}

// ExpressRoutePortStatus defines the observed state of ExpressRoutePort.
type ExpressRoutePortStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRoutePortObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRoutePort is the Schema for the ExpressRoutePorts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ExpressRoutePort struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRoutePortSpec   `json:"spec"`
	Status            ExpressRoutePortStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRoutePortList contains a list of ExpressRoutePorts
type ExpressRoutePortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRoutePort `json:"items"`
}

// Repository type metadata.
var (
	ExpressRoutePort_Kind             = "ExpressRoutePort"
	ExpressRoutePort_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExpressRoutePort_Kind}.String()
	ExpressRoutePort_KindAPIVersion   = ExpressRoutePort_Kind + "." + CRDGroupVersion.String()
	ExpressRoutePort_GroupVersionKind = CRDGroupVersion.WithKind(ExpressRoutePort_Kind)
)

func init() {
	SchemeBuilder.Register(&ExpressRoutePort{}, &ExpressRoutePortList{})
}

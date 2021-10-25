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

type EventhubNamespaceObservation struct {
}

type EventhubNamespaceParameters struct {

	// +kubebuilder:validation:Optional
	AutoInflateEnabled *bool `json:"autoInflateEnabled,omitempty" tf:"auto_inflate_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Capacity *int64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Optional
	DedicatedClusterID *string `json:"dedicatedClusterId,omitempty" tf:"dedicated_cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MaximumThroughputUnits *int64 `json:"maximumThroughputUnits,omitempty" tf:"maximum_throughput_units,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkRulesets []NetworkRulesetsParameters `json:"networkRulesets,omitempty" tf:"network_rulesets,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type IPRuleObservation struct {
}

type IPRuleParameters struct {

	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	IPMask *string `json:"ipMask" tf:"ip_mask,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkRulesetsObservation struct {
}

type NetworkRulesetsParameters struct {

	// +kubebuilder:validation:Required
	DefaultAction *string `json:"defaultAction" tf:"default_action,omitempty"`

	// +kubebuilder:validation:Optional
	IPRule []IPRuleParameters `json:"ipRule,omitempty" tf:"ip_rule,omitempty"`

	// +kubebuilder:validation:Optional
	TrustedServiceAccessEnabled *bool `json:"trustedServiceAccessEnabled,omitempty" tf:"trusted_service_access_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkRule []VirtualNetworkRuleParameters `json:"virtualNetworkRule,omitempty" tf:"virtual_network_rule,omitempty"`
}

type VirtualNetworkRuleObservation struct {
}

type VirtualNetworkRuleParameters struct {

	// +kubebuilder:validation:Optional
	IgnoreMissingVirtualNetworkServiceEndpoint *bool `json:"ignoreMissingVirtualNetworkServiceEndpoint,omitempty" tf:"ignore_missing_virtual_network_service_endpoint,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
}

// EventhubNamespaceSpec defines the desired state of EventhubNamespace
type EventhubNamespaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventhubNamespaceParameters `json:"forProvider"`
}

// EventhubNamespaceStatus defines the observed state of EventhubNamespace.
type EventhubNamespaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventhubNamespaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubNamespace is the Schema for the EventhubNamespaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EventhubNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubNamespaceSpec   `json:"spec"`
	Status            EventhubNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubNamespaceList contains a list of EventhubNamespaces
type EventhubNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventhubNamespace `json:"items"`
}

// Repository type metadata.
var (
	EventhubNamespaceKind             = "EventhubNamespace"
	EventhubNamespaceGroupKind        = schema.GroupKind{Group: Group, Kind: EventhubNamespaceKind}.String()
	EventhubNamespaceKindAPIVersion   = EventhubNamespaceKind + "." + GroupVersion.String()
	EventhubNamespaceGroupVersionKind = GroupVersion.WithKind(EventhubNamespaceKind)
)

func init() {
	SchemeBuilder.Register(&EventhubNamespace{}, &EventhubNamespaceList{})
}

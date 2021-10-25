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

type LbOutboundRuleFrontendIPConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LbOutboundRuleFrontendIPConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type LbOutboundRuleObservation struct {
}

type LbOutboundRuleParameters struct {

	// +kubebuilder:validation:Optional
	AllocatedOutboundPorts *int64 `json:"allocatedOutboundPorts,omitempty" tf:"allocated_outbound_ports,omitempty"`

	// +kubebuilder:validation:Required
	BackendAddressPoolID *string `json:"backendAddressPoolId" tf:"backend_address_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	EnableTCPReset *bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`

	// +kubebuilder:validation:Optional
	FrontendIPConfiguration []LbOutboundRuleFrontendIPConfigurationParameters `json:"frontendIpConfiguration,omitempty" tf:"frontend_ip_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *int64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// +kubebuilder:validation:Required
	LoadbalancerID *string `json:"loadbalancerId" tf:"loadbalancer_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// LbOutboundRuleSpec defines the desired state of LbOutboundRule
type LbOutboundRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LbOutboundRuleParameters `json:"forProvider"`
}

// LbOutboundRuleStatus defines the observed state of LbOutboundRule.
type LbOutboundRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LbOutboundRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LbOutboundRule is the Schema for the LbOutboundRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LbOutboundRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbOutboundRuleSpec   `json:"spec"`
	Status            LbOutboundRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbOutboundRuleList contains a list of LbOutboundRules
type LbOutboundRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbOutboundRule `json:"items"`
}

// Repository type metadata.
var (
	LbOutboundRuleKind             = "LbOutboundRule"
	LbOutboundRuleGroupKind        = schema.GroupKind{Group: Group, Kind: LbOutboundRuleKind}.String()
	LbOutboundRuleKindAPIVersion   = LbOutboundRuleKind + "." + GroupVersion.String()
	LbOutboundRuleGroupVersionKind = GroupVersion.WithKind(LbOutboundRuleKind)
)

func init() {
	SchemeBuilder.Register(&LbOutboundRule{}, &LbOutboundRuleList{})
}

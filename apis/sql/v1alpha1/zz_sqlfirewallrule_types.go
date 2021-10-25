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

type SqlFirewallRuleObservation struct {
}

type SqlFirewallRuleParameters struct {

	// +kubebuilder:validation:Required
	EndIPAddress *string `json:"endIpAddress" tf:"end_ip_address,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	ServerName *string `json:"serverName" tf:"server_name,omitempty"`

	// +kubebuilder:validation:Required
	StartIPAddress *string `json:"startIpAddress" tf:"start_ip_address,omitempty"`
}

// SqlFirewallRuleSpec defines the desired state of SqlFirewallRule
type SqlFirewallRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SqlFirewallRuleParameters `json:"forProvider"`
}

// SqlFirewallRuleStatus defines the observed state of SqlFirewallRule.
type SqlFirewallRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SqlFirewallRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SqlFirewallRule is the Schema for the SqlFirewallRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SqlFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlFirewallRuleSpec   `json:"spec"`
	Status            SqlFirewallRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqlFirewallRuleList contains a list of SqlFirewallRules
type SqlFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlFirewallRule `json:"items"`
}

// Repository type metadata.
var (
	SqlFirewallRuleKind             = "SqlFirewallRule"
	SqlFirewallRuleGroupKind        = schema.GroupKind{Group: Group, Kind: SqlFirewallRuleKind}.String()
	SqlFirewallRuleKindAPIVersion   = SqlFirewallRuleKind + "." + GroupVersion.String()
	SqlFirewallRuleGroupVersionKind = GroupVersion.WithKind(SqlFirewallRuleKind)
)

func init() {
	SchemeBuilder.Register(&SqlFirewallRule{}, &SqlFirewallRuleList{})
}

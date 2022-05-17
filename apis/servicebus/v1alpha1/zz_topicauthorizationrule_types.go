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

type TopicAuthorizationRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TopicAuthorizationRuleParameters struct {

	// +kubebuilder:validation:Optional
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// +kubebuilder:validation:Optional
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NamespaceName *string `json:"namespaceName" tf:"namespace_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`

	// +kubebuilder:validation:Required
	TopicName *string `json:"topicName" tf:"topic_name,omitempty"`
}

// TopicAuthorizationRuleSpec defines the desired state of TopicAuthorizationRule
type TopicAuthorizationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicAuthorizationRuleParameters `json:"forProvider"`
}

// TopicAuthorizationRuleStatus defines the observed state of TopicAuthorizationRule.
type TopicAuthorizationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicAuthorizationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TopicAuthorizationRule is the Schema for the TopicAuthorizationRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type TopicAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicAuthorizationRuleSpec   `json:"spec"`
	Status            TopicAuthorizationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicAuthorizationRuleList contains a list of TopicAuthorizationRules
type TopicAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopicAuthorizationRule `json:"items"`
}

// Repository type metadata.
var (
	TopicAuthorizationRule_Kind             = "TopicAuthorizationRule"
	TopicAuthorizationRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TopicAuthorizationRule_Kind}.String()
	TopicAuthorizationRule_KindAPIVersion   = TopicAuthorizationRule_Kind + "." + CRDGroupVersion.String()
	TopicAuthorizationRule_GroupVersionKind = CRDGroupVersion.WithKind(TopicAuthorizationRule_Kind)
)

func init() {
	SchemeBuilder.Register(&TopicAuthorizationRule{}, &TopicAuthorizationRuleList{})
}

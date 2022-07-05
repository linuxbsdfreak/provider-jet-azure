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

type QueueAuthorizationRuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type QueueAuthorizationRuleParameters struct {

	// +kubebuilder:validation:Optional
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// +kubebuilder:validation:Optional
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NamespaceName *string `json:"namespaceName" tf:"namespace_name,omitempty"`

	// +kubebuilder:validation:Required
	QueueName *string `json:"queueName" tf:"queue_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
}

// QueueAuthorizationRuleSpec defines the desired state of QueueAuthorizationRule
type QueueAuthorizationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueAuthorizationRuleParameters `json:"forProvider"`
}

// QueueAuthorizationRuleStatus defines the observed state of QueueAuthorizationRule.
type QueueAuthorizationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueAuthorizationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QueueAuthorizationRule is the Schema for the QueueAuthorizationRules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type QueueAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueueAuthorizationRuleSpec   `json:"spec"`
	Status            QueueAuthorizationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueAuthorizationRuleList contains a list of QueueAuthorizationRules
type QueueAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QueueAuthorizationRule `json:"items"`
}

// Repository type metadata.
var (
	QueueAuthorizationRule_Kind             = "QueueAuthorizationRule"
	QueueAuthorizationRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QueueAuthorizationRule_Kind}.String()
	QueueAuthorizationRule_KindAPIVersion   = QueueAuthorizationRule_Kind + "." + CRDGroupVersion.String()
	QueueAuthorizationRule_GroupVersionKind = CRDGroupVersion.WithKind(QueueAuthorizationRule_Kind)
)

func init() {
	SchemeBuilder.Register(&QueueAuthorizationRule{}, &QueueAuthorizationRuleList{})
}

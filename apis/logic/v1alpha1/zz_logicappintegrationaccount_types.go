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

type LogicAppIntegrationAccountObservation struct {
}

type LogicAppIntegrationAccountParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LogicAppIntegrationAccountSpec defines the desired state of LogicAppIntegrationAccount
type LogicAppIntegrationAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogicAppIntegrationAccountParameters `json:"forProvider"`
}

// LogicAppIntegrationAccountStatus defines the observed state of LogicAppIntegrationAccount.
type LogicAppIntegrationAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogicAppIntegrationAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogicAppIntegrationAccount is the Schema for the LogicAppIntegrationAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogicAppIntegrationAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppIntegrationAccountSpec   `json:"spec"`
	Status            LogicAppIntegrationAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogicAppIntegrationAccountList contains a list of LogicAppIntegrationAccounts
type LogicAppIntegrationAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogicAppIntegrationAccount `json:"items"`
}

// Repository type metadata.
var (
	LogicAppIntegrationAccountKind             = "LogicAppIntegrationAccount"
	LogicAppIntegrationAccountGroupKind        = schema.GroupKind{Group: Group, Kind: LogicAppIntegrationAccountKind}.String()
	LogicAppIntegrationAccountKindAPIVersion   = LogicAppIntegrationAccountKind + "." + GroupVersion.String()
	LogicAppIntegrationAccountGroupVersionKind = GroupVersion.WithKind(LogicAppIntegrationAccountKind)
)

func init() {
	SchemeBuilder.Register(&LogicAppIntegrationAccount{}, &LogicAppIntegrationAccountList{})
}

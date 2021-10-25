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

type LogicAppIntegrationAccountSessionObservation struct {
}

type LogicAppIntegrationAccountSessionParameters struct {

	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// +kubebuilder:validation:Required
	IntegrationAccountName *string `json:"integrationAccountName" tf:"integration_account_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// LogicAppIntegrationAccountSessionSpec defines the desired state of LogicAppIntegrationAccountSession
type LogicAppIntegrationAccountSessionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogicAppIntegrationAccountSessionParameters `json:"forProvider"`
}

// LogicAppIntegrationAccountSessionStatus defines the observed state of LogicAppIntegrationAccountSession.
type LogicAppIntegrationAccountSessionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogicAppIntegrationAccountSessionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogicAppIntegrationAccountSession is the Schema for the LogicAppIntegrationAccountSessions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogicAppIntegrationAccountSession struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppIntegrationAccountSessionSpec   `json:"spec"`
	Status            LogicAppIntegrationAccountSessionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogicAppIntegrationAccountSessionList contains a list of LogicAppIntegrationAccountSessions
type LogicAppIntegrationAccountSessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogicAppIntegrationAccountSession `json:"items"`
}

// Repository type metadata.
var (
	LogicAppIntegrationAccountSessionKind             = "LogicAppIntegrationAccountSession"
	LogicAppIntegrationAccountSessionGroupKind        = schema.GroupKind{Group: Group, Kind: LogicAppIntegrationAccountSessionKind}.String()
	LogicAppIntegrationAccountSessionKindAPIVersion   = LogicAppIntegrationAccountSessionKind + "." + GroupVersion.String()
	LogicAppIntegrationAccountSessionGroupVersionKind = GroupVersion.WithKind(LogicAppIntegrationAccountSessionKind)
)

func init() {
	SchemeBuilder.Register(&LogicAppIntegrationAccountSession{}, &LogicAppIntegrationAccountSessionList{})
}

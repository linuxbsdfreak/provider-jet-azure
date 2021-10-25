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

type DataFactoryIntegrationRuntimeSelfHostedObservation struct {
	AuthKey1 *string `json:"authKey1,omitempty" tf:"auth_key_1,omitempty"`

	AuthKey2 *string `json:"authKey2,omitempty" tf:"auth_key_2,omitempty"`
}

type DataFactoryIntegrationRuntimeSelfHostedParameters struct {

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RbacAuthorization []RbacAuthorizationParameters `json:"rbacAuthorization,omitempty" tf:"rbac_authorization,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

type RbacAuthorizationObservation struct {
}

type RbacAuthorizationParameters struct {

	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`
}

// DataFactoryIntegrationRuntimeSelfHostedSpec defines the desired state of DataFactoryIntegrationRuntimeSelfHosted
type DataFactoryIntegrationRuntimeSelfHostedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataFactoryIntegrationRuntimeSelfHostedParameters `json:"forProvider"`
}

// DataFactoryIntegrationRuntimeSelfHostedStatus defines the observed state of DataFactoryIntegrationRuntimeSelfHosted.
type DataFactoryIntegrationRuntimeSelfHostedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataFactoryIntegrationRuntimeSelfHostedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryIntegrationRuntimeSelfHosted is the Schema for the DataFactoryIntegrationRuntimeSelfHosteds API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFactoryIntegrationRuntimeSelfHosted struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryIntegrationRuntimeSelfHostedSpec   `json:"spec"`
	Status            DataFactoryIntegrationRuntimeSelfHostedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryIntegrationRuntimeSelfHostedList contains a list of DataFactoryIntegrationRuntimeSelfHosteds
type DataFactoryIntegrationRuntimeSelfHostedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFactoryIntegrationRuntimeSelfHosted `json:"items"`
}

// Repository type metadata.
var (
	DataFactoryIntegrationRuntimeSelfHostedKind             = "DataFactoryIntegrationRuntimeSelfHosted"
	DataFactoryIntegrationRuntimeSelfHostedGroupKind        = schema.GroupKind{Group: Group, Kind: DataFactoryIntegrationRuntimeSelfHostedKind}.String()
	DataFactoryIntegrationRuntimeSelfHostedKindAPIVersion   = DataFactoryIntegrationRuntimeSelfHostedKind + "." + GroupVersion.String()
	DataFactoryIntegrationRuntimeSelfHostedGroupVersionKind = GroupVersion.WithKind(DataFactoryIntegrationRuntimeSelfHostedKind)
)

func init() {
	SchemeBuilder.Register(&DataFactoryIntegrationRuntimeSelfHosted{}, &DataFactoryIntegrationRuntimeSelfHostedList{})
}

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

type DataFactoryLinkedServiceSynapseKeyVaultPasswordObservation struct {
}

type DataFactoryLinkedServiceSynapseKeyVaultPasswordParameters struct {

	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type DataFactoryLinkedServiceSynapseObservation struct {
}

type DataFactoryLinkedServiceSynapseParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	ConnectionString *string `json:"connectionString" tf:"connection_string,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultPassword []DataFactoryLinkedServiceSynapseKeyVaultPasswordParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// DataFactoryLinkedServiceSynapseSpec defines the desired state of DataFactoryLinkedServiceSynapse
type DataFactoryLinkedServiceSynapseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataFactoryLinkedServiceSynapseParameters `json:"forProvider"`
}

// DataFactoryLinkedServiceSynapseStatus defines the observed state of DataFactoryLinkedServiceSynapse.
type DataFactoryLinkedServiceSynapseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataFactoryLinkedServiceSynapseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryLinkedServiceSynapse is the Schema for the DataFactoryLinkedServiceSynapses API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFactoryLinkedServiceSynapse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryLinkedServiceSynapseSpec   `json:"spec"`
	Status            DataFactoryLinkedServiceSynapseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryLinkedServiceSynapseList contains a list of DataFactoryLinkedServiceSynapses
type DataFactoryLinkedServiceSynapseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFactoryLinkedServiceSynapse `json:"items"`
}

// Repository type metadata.
var (
	DataFactoryLinkedServiceSynapseKind             = "DataFactoryLinkedServiceSynapse"
	DataFactoryLinkedServiceSynapseGroupKind        = schema.GroupKind{Group: Group, Kind: DataFactoryLinkedServiceSynapseKind}.String()
	DataFactoryLinkedServiceSynapseKindAPIVersion   = DataFactoryLinkedServiceSynapseKind + "." + GroupVersion.String()
	DataFactoryLinkedServiceSynapseGroupVersionKind = GroupVersion.WithKind(DataFactoryLinkedServiceSynapseKind)
)

func init() {
	SchemeBuilder.Register(&DataFactoryLinkedServiceSynapse{}, &DataFactoryLinkedServiceSynapseList{})
}

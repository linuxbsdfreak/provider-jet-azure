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

type DataFactoryLinkedServiceWebObservation struct {
}

type DataFactoryLinkedServiceWebParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	AuthenticationType *string `json:"authenticationType" tf:"authentication_type,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// DataFactoryLinkedServiceWebSpec defines the desired state of DataFactoryLinkedServiceWeb
type DataFactoryLinkedServiceWebSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataFactoryLinkedServiceWebParameters `json:"forProvider"`
}

// DataFactoryLinkedServiceWebStatus defines the observed state of DataFactoryLinkedServiceWeb.
type DataFactoryLinkedServiceWebStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataFactoryLinkedServiceWebObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryLinkedServiceWeb is the Schema for the DataFactoryLinkedServiceWebs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFactoryLinkedServiceWeb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryLinkedServiceWebSpec   `json:"spec"`
	Status            DataFactoryLinkedServiceWebStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryLinkedServiceWebList contains a list of DataFactoryLinkedServiceWebs
type DataFactoryLinkedServiceWebList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFactoryLinkedServiceWeb `json:"items"`
}

// Repository type metadata.
var (
	DataFactoryLinkedServiceWebKind             = "DataFactoryLinkedServiceWeb"
	DataFactoryLinkedServiceWebGroupKind        = schema.GroupKind{Group: Group, Kind: DataFactoryLinkedServiceWebKind}.String()
	DataFactoryLinkedServiceWebKindAPIVersion   = DataFactoryLinkedServiceWebKind + "." + GroupVersion.String()
	DataFactoryLinkedServiceWebGroupVersionKind = GroupVersion.WithKind(DataFactoryLinkedServiceWebKind)
)

func init() {
	SchemeBuilder.Register(&DataFactoryLinkedServiceWeb{}, &DataFactoryLinkedServiceWebList{})
}

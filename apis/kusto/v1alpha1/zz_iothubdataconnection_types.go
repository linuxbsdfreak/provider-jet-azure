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

type IothubDataConnectionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IothubDataConnectionParameters struct {

	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// +kubebuilder:validation:Required
	ConsumerGroup *string `json:"consumerGroup" tf:"consumer_group,omitempty"`

	// +kubebuilder:validation:Optional
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	EventSystemProperties []*string `json:"eventSystemProperties,omitempty" tf:"event_system_properties,omitempty"`

	// +kubebuilder:validation:Required
	IothubID *string `json:"iothubId" tf:"iothub_id,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MappingRuleName *string `json:"mappingRuleName,omitempty" tf:"mapping_rule_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SharedAccessPolicyName *string `json:"sharedAccessPolicyName" tf:"shared_access_policy_name,omitempty"`

	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

// IothubDataConnectionSpec defines the desired state of IothubDataConnection
type IothubDataConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IothubDataConnectionParameters `json:"forProvider"`
}

// IothubDataConnectionStatus defines the observed state of IothubDataConnection.
type IothubDataConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IothubDataConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IothubDataConnection is the Schema for the IothubDataConnections API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type IothubDataConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubDataConnectionSpec   `json:"spec"`
	Status            IothubDataConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IothubDataConnectionList contains a list of IothubDataConnections
type IothubDataConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IothubDataConnection `json:"items"`
}

// Repository type metadata.
var (
	IothubDataConnection_Kind             = "IothubDataConnection"
	IothubDataConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IothubDataConnection_Kind}.String()
	IothubDataConnection_KindAPIVersion   = IothubDataConnection_Kind + "." + CRDGroupVersion.String()
	IothubDataConnection_GroupVersionKind = CRDGroupVersion.WithKind(IothubDataConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&IothubDataConnection{}, &IothubDataConnectionList{})
}
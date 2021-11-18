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

type AutoscaleSettingsObservation struct {
}

type AutoscaleSettingsParameters struct {

	// +kubebuilder:validation:Optional
	MaxThroughput *int64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type CassandraKeyspaceObservation struct {
}

type CassandraKeyspaceParameters struct {

	// +crossplane:generate:reference:type=Account
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AutoscaleSettings []AutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/azure/v1alpha1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Throughput *int64 `json:"throughput,omitempty" tf:"throughput,omitempty"`
}

// CassandraKeyspaceSpec defines the desired state of CassandraKeyspace
type CassandraKeyspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CassandraKeyspaceParameters `json:"forProvider"`
}

// CassandraKeyspaceStatus defines the observed state of CassandraKeyspace.
type CassandraKeyspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CassandraKeyspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraKeyspace is the Schema for the CassandraKeyspaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type CassandraKeyspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CassandraKeyspaceSpec   `json:"spec"`
	Status            CassandraKeyspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CassandraKeyspaceList contains a list of CassandraKeyspaces
type CassandraKeyspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CassandraKeyspace `json:"items"`
}

// Repository type metadata.
var (
	CassandraKeyspace_Kind             = "CassandraKeyspace"
	CassandraKeyspace_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CassandraKeyspace_Kind}.String()
	CassandraKeyspace_KindAPIVersion   = CassandraKeyspace_Kind + "." + CRDGroupVersion.String()
	CassandraKeyspace_GroupVersionKind = CRDGroupVersion.WithKind(CassandraKeyspace_Kind)
)

func init() {
	SchemeBuilder.Register(&CassandraKeyspace{}, &CassandraKeyspaceList{})
}

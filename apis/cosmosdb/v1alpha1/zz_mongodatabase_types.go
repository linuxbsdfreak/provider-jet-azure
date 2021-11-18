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

type MongoDatabaseAutoscaleSettingsObservation struct {
}

type MongoDatabaseAutoscaleSettingsParameters struct {

	// +kubebuilder:validation:Optional
	MaxThroughput *int64 `json:"maxThroughput,omitempty" tf:"max_throughput,omitempty"`
}

type MongoDatabaseObservation struct {
}

type MongoDatabaseParameters struct {

	// +crossplane:generate:reference:type=Account
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AutoscaleSettings []MongoDatabaseAutoscaleSettingsParameters `json:"autoscaleSettings,omitempty" tf:"autoscale_settings,omitempty"`

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

// MongoDatabaseSpec defines the desired state of MongoDatabase
type MongoDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MongoDatabaseParameters `json:"forProvider"`
}

// MongoDatabaseStatus defines the observed state of MongoDatabase.
type MongoDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MongoDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MongoDatabase is the Schema for the MongoDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type MongoDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MongoDatabaseSpec   `json:"spec"`
	Status            MongoDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MongoDatabaseList contains a list of MongoDatabases
type MongoDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongoDatabase `json:"items"`
}

// Repository type metadata.
var (
	MongoDatabase_Kind             = "MongoDatabase"
	MongoDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MongoDatabase_Kind}.String()
	MongoDatabase_KindAPIVersion   = MongoDatabase_Kind + "." + CRDGroupVersion.String()
	MongoDatabase_GroupVersionKind = CRDGroupVersion.WithKind(MongoDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&MongoDatabase{}, &MongoDatabaseList{})
}

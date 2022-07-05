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

type PrivateDNSTXTRecordObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateDNSTXTRecordParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Record []PrivateDNSTXTRecordRecordParameters `json:"record" tf:"record,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	TTL *float64 `json:"ttl" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	ZoneName *string `json:"zoneName" tf:"zone_name,omitempty"`
}

type PrivateDNSTXTRecordRecordObservation struct {
}

type PrivateDNSTXTRecordRecordParameters struct {

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// PrivateDNSTXTRecordSpec defines the desired state of PrivateDNSTXTRecord
type PrivateDNSTXTRecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSTXTRecordParameters `json:"forProvider"`
}

// PrivateDNSTXTRecordStatus defines the observed state of PrivateDNSTXTRecord.
type PrivateDNSTXTRecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSTXTRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSTXTRecord is the Schema for the PrivateDNSTXTRecords API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type PrivateDNSTXTRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDNSTXTRecordSpec   `json:"spec"`
	Status            PrivateDNSTXTRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSTXTRecordList contains a list of PrivateDNSTXTRecords
type PrivateDNSTXTRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSTXTRecord `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSTXTRecord_Kind             = "PrivateDNSTXTRecord"
	PrivateDNSTXTRecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSTXTRecord_Kind}.String()
	PrivateDNSTXTRecord_KindAPIVersion   = PrivateDNSTXTRecord_Kind + "." + CRDGroupVersion.String()
	PrivateDNSTXTRecord_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSTXTRecord_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSTXTRecord{}, &PrivateDNSTXTRecordList{})
}

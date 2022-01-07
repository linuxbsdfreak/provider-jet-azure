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

type OutputBlobObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type OutputBlobParameters struct {

	// +kubebuilder:validation:Optional
	BatchMaxWaitTime *string `json:"batchMaxWaitTime,omitempty" tf:"batch_max_wait_time,omitempty"`

	// +kubebuilder:validation:Optional
	BatchMinRows *float64 `json:"batchMinRows,omitempty" tf:"batch_min_rows,omitempty"`

	// +kubebuilder:validation:Required
	DateFormat *string `json:"dateFormat" tf:"date_format,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PathPattern *string `json:"pathPattern" tf:"path_pattern,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Serialization []SerializationParameters `json:"serialization" tf:"serialization,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountKeySecretRef v1.SecretKeySelector `json:"storageAccountKeySecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	StorageAccountName *string `json:"storageAccountName" tf:"storage_account_name,omitempty"`

	// +kubebuilder:validation:Required
	StorageContainerName *string `json:"storageContainerName" tf:"storage_container_name,omitempty"`

	// +kubebuilder:validation:Required
	StreamAnalyticsJobName *string `json:"streamAnalyticsJobName" tf:"stream_analytics_job_name,omitempty"`

	// +kubebuilder:validation:Required
	TimeFormat *string `json:"timeFormat" tf:"time_format,omitempty"`
}

type SerializationObservation struct {
}

type SerializationParameters struct {

	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// +kubebuilder:validation:Optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// OutputBlobSpec defines the desired state of OutputBlob
type OutputBlobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OutputBlobParameters `json:"forProvider"`
}

// OutputBlobStatus defines the observed state of OutputBlob.
type OutputBlobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OutputBlobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OutputBlob is the Schema for the OutputBlobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type OutputBlob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OutputBlobSpec   `json:"spec"`
	Status            OutputBlobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutputBlobList contains a list of OutputBlobs
type OutputBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutputBlob `json:"items"`
}

// Repository type metadata.
var (
	OutputBlob_Kind             = "OutputBlob"
	OutputBlob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OutputBlob_Kind}.String()
	OutputBlob_KindAPIVersion   = OutputBlob_Kind + "." + CRDGroupVersion.String()
	OutputBlob_GroupVersionKind = CRDGroupVersion.WithKind(OutputBlob_Kind)
)

func init() {
	SchemeBuilder.Register(&OutputBlob{}, &OutputBlobList{})
}

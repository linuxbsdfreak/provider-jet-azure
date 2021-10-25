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

type StreamAnalyticsStreamInputBlobObservation struct {
}

type StreamAnalyticsStreamInputBlobParameters struct {

	// +kubebuilder:validation:Required
	DateFormat *string `json:"dateFormat" tf:"date_format,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PathPattern *string `json:"pathPattern" tf:"path_pattern,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	Serialization []StreamAnalyticsStreamInputBlobSerializationParameters `json:"serialization" tf:"serialization,omitempty"`

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

type StreamAnalyticsStreamInputBlobSerializationObservation struct {
}

type StreamAnalyticsStreamInputBlobSerializationParameters struct {

	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// +kubebuilder:validation:Optional
	FieldDelimiter *string `json:"fieldDelimiter,omitempty" tf:"field_delimiter,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// StreamAnalyticsStreamInputBlobSpec defines the desired state of StreamAnalyticsStreamInputBlob
type StreamAnalyticsStreamInputBlobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StreamAnalyticsStreamInputBlobParameters `json:"forProvider"`
}

// StreamAnalyticsStreamInputBlobStatus defines the observed state of StreamAnalyticsStreamInputBlob.
type StreamAnalyticsStreamInputBlobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StreamAnalyticsStreamInputBlobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StreamAnalyticsStreamInputBlob is the Schema for the StreamAnalyticsStreamInputBlobs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StreamAnalyticsStreamInputBlob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StreamAnalyticsStreamInputBlobSpec   `json:"spec"`
	Status            StreamAnalyticsStreamInputBlobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StreamAnalyticsStreamInputBlobList contains a list of StreamAnalyticsStreamInputBlobs
type StreamAnalyticsStreamInputBlobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StreamAnalyticsStreamInputBlob `json:"items"`
}

// Repository type metadata.
var (
	StreamAnalyticsStreamInputBlobKind             = "StreamAnalyticsStreamInputBlob"
	StreamAnalyticsStreamInputBlobGroupKind        = schema.GroupKind{Group: Group, Kind: StreamAnalyticsStreamInputBlobKind}.String()
	StreamAnalyticsStreamInputBlobKindAPIVersion   = StreamAnalyticsStreamInputBlobKind + "." + GroupVersion.String()
	StreamAnalyticsStreamInputBlobGroupVersionKind = GroupVersion.WithKind(StreamAnalyticsStreamInputBlobKind)
)

func init() {
	SchemeBuilder.Register(&StreamAnalyticsStreamInputBlob{}, &StreamAnalyticsStreamInputBlobList{})
}

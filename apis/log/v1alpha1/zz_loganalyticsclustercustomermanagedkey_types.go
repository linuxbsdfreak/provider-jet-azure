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

type LogAnalyticsClusterCustomerManagedKeyObservation struct {
}

type LogAnalyticsClusterCustomerManagedKeyParameters struct {

	// +kubebuilder:validation:Required
	KeyVaultKeyID *string `json:"keyVaultKeyId" tf:"key_vault_key_id,omitempty"`

	// +kubebuilder:validation:Required
	LogAnalyticsClusterID *string `json:"logAnalyticsClusterId" tf:"log_analytics_cluster_id,omitempty"`
}

// LogAnalyticsClusterCustomerManagedKeySpec defines the desired state of LogAnalyticsClusterCustomerManagedKey
type LogAnalyticsClusterCustomerManagedKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsClusterCustomerManagedKeyParameters `json:"forProvider"`
}

// LogAnalyticsClusterCustomerManagedKeyStatus defines the observed state of LogAnalyticsClusterCustomerManagedKey.
type LogAnalyticsClusterCustomerManagedKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsClusterCustomerManagedKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsClusterCustomerManagedKey is the Schema for the LogAnalyticsClusterCustomerManagedKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogAnalyticsClusterCustomerManagedKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsClusterCustomerManagedKeySpec   `json:"spec"`
	Status            LogAnalyticsClusterCustomerManagedKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsClusterCustomerManagedKeyList contains a list of LogAnalyticsClusterCustomerManagedKeys
type LogAnalyticsClusterCustomerManagedKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsClusterCustomerManagedKey `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsClusterCustomerManagedKeyKind             = "LogAnalyticsClusterCustomerManagedKey"
	LogAnalyticsClusterCustomerManagedKeyGroupKind        = schema.GroupKind{Group: Group, Kind: LogAnalyticsClusterCustomerManagedKeyKind}.String()
	LogAnalyticsClusterCustomerManagedKeyKindAPIVersion   = LogAnalyticsClusterCustomerManagedKeyKind + "." + GroupVersion.String()
	LogAnalyticsClusterCustomerManagedKeyGroupVersionKind = GroupVersion.WithKind(LogAnalyticsClusterCustomerManagedKeyKind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsClusterCustomerManagedKey{}, &LogAnalyticsClusterCustomerManagedKeyList{})
}

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

type LogAnalyticsWorkspaceObservation struct {
	PortalURL *string `json:"portalUrl,omitempty" tf:"portal_url,omitempty"`

	WorkspaceID *string `json:"workspaceId,omitempty" tf:"workspace_id,omitempty"`
}

type LogAnalyticsWorkspaceParameters struct {

	// +kubebuilder:validation:Optional
	DailyQuotaGb *float64 `json:"dailyQuotaGb,omitempty" tf:"daily_quota_gb,omitempty"`

	// +kubebuilder:validation:Optional
	InternetIngestionEnabled *bool `json:"internetIngestionEnabled,omitempty" tf:"internet_ingestion_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	InternetQueryEnabled *bool `json:"internetQueryEnabled,omitempty" tf:"internet_query_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ReservationCapcityInGbPerDay *int64 `json:"reservationCapcityInGbPerDay,omitempty" tf:"reservation_capcity_in_gb_per_day,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionInDays *int64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LogAnalyticsWorkspaceSpec defines the desired state of LogAnalyticsWorkspace
type LogAnalyticsWorkspaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsWorkspaceParameters `json:"forProvider"`
}

// LogAnalyticsWorkspaceStatus defines the observed state of LogAnalyticsWorkspace.
type LogAnalyticsWorkspaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsWorkspaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsWorkspace is the Schema for the LogAnalyticsWorkspaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogAnalyticsWorkspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsWorkspaceSpec   `json:"spec"`
	Status            LogAnalyticsWorkspaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsWorkspaceList contains a list of LogAnalyticsWorkspaces
type LogAnalyticsWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsWorkspace `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsWorkspaceKind             = "LogAnalyticsWorkspace"
	LogAnalyticsWorkspaceGroupKind        = schema.GroupKind{Group: Group, Kind: LogAnalyticsWorkspaceKind}.String()
	LogAnalyticsWorkspaceKindAPIVersion   = LogAnalyticsWorkspaceKind + "." + GroupVersion.String()
	LogAnalyticsWorkspaceGroupVersionKind = GroupVersion.WithKind(LogAnalyticsWorkspaceKind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsWorkspace{}, &LogAnalyticsWorkspaceList{})
}

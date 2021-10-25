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

type SentinelAlertRuleFusionObservation struct {
}

type SentinelAlertRuleFusionParameters struct {

	// +kubebuilder:validation:Required
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid" tf:"alert_rule_template_guid,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// SentinelAlertRuleFusionSpec defines the desired state of SentinelAlertRuleFusion
type SentinelAlertRuleFusionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SentinelAlertRuleFusionParameters `json:"forProvider"`
}

// SentinelAlertRuleFusionStatus defines the observed state of SentinelAlertRuleFusion.
type SentinelAlertRuleFusionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SentinelAlertRuleFusionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleFusion is the Schema for the SentinelAlertRuleFusions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelAlertRuleFusion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SentinelAlertRuleFusionSpec   `json:"spec"`
	Status            SentinelAlertRuleFusionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleFusionList contains a list of SentinelAlertRuleFusions
type SentinelAlertRuleFusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelAlertRuleFusion `json:"items"`
}

// Repository type metadata.
var (
	SentinelAlertRuleFusionKind             = "SentinelAlertRuleFusion"
	SentinelAlertRuleFusionGroupKind        = schema.GroupKind{Group: Group, Kind: SentinelAlertRuleFusionKind}.String()
	SentinelAlertRuleFusionKindAPIVersion   = SentinelAlertRuleFusionKind + "." + GroupVersion.String()
	SentinelAlertRuleFusionGroupVersionKind = GroupVersion.WithKind(SentinelAlertRuleFusionKind)
)

func init() {
	SchemeBuilder.Register(&SentinelAlertRuleFusion{}, &SentinelAlertRuleFusionList{})
}

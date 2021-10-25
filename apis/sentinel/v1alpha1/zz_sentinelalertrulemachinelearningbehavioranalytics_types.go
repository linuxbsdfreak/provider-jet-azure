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

type SentinelAlertRuleMachineLearningBehaviorAnalyticsObservation struct {
}

type SentinelAlertRuleMachineLearningBehaviorAnalyticsParameters struct {

	// +kubebuilder:validation:Required
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid" tf:"alert_rule_template_guid,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// SentinelAlertRuleMachineLearningBehaviorAnalyticsSpec defines the desired state of SentinelAlertRuleMachineLearningBehaviorAnalytics
type SentinelAlertRuleMachineLearningBehaviorAnalyticsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SentinelAlertRuleMachineLearningBehaviorAnalyticsParameters `json:"forProvider"`
}

// SentinelAlertRuleMachineLearningBehaviorAnalyticsStatus defines the observed state of SentinelAlertRuleMachineLearningBehaviorAnalytics.
type SentinelAlertRuleMachineLearningBehaviorAnalyticsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SentinelAlertRuleMachineLearningBehaviorAnalyticsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleMachineLearningBehaviorAnalytics is the Schema for the SentinelAlertRuleMachineLearningBehaviorAnalyticss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelAlertRuleMachineLearningBehaviorAnalytics struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SentinelAlertRuleMachineLearningBehaviorAnalyticsSpec   `json:"spec"`
	Status            SentinelAlertRuleMachineLearningBehaviorAnalyticsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleMachineLearningBehaviorAnalyticsList contains a list of SentinelAlertRuleMachineLearningBehaviorAnalyticss
type SentinelAlertRuleMachineLearningBehaviorAnalyticsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelAlertRuleMachineLearningBehaviorAnalytics `json:"items"`
}

// Repository type metadata.
var (
	SentinelAlertRuleMachineLearningBehaviorAnalyticsKind             = "SentinelAlertRuleMachineLearningBehaviorAnalytics"
	SentinelAlertRuleMachineLearningBehaviorAnalyticsGroupKind        = schema.GroupKind{Group: Group, Kind: SentinelAlertRuleMachineLearningBehaviorAnalyticsKind}.String()
	SentinelAlertRuleMachineLearningBehaviorAnalyticsKindAPIVersion   = SentinelAlertRuleMachineLearningBehaviorAnalyticsKind + "." + GroupVersion.String()
	SentinelAlertRuleMachineLearningBehaviorAnalyticsGroupVersionKind = GroupVersion.WithKind(SentinelAlertRuleMachineLearningBehaviorAnalyticsKind)
)

func init() {
	SchemeBuilder.Register(&SentinelAlertRuleMachineLearningBehaviorAnalytics{}, &SentinelAlertRuleMachineLearningBehaviorAnalyticsList{})
}

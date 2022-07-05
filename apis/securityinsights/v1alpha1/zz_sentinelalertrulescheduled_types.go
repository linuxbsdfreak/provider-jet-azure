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

type EventGroupingObservation struct {
}

type EventGroupingParameters struct {

	// +kubebuilder:validation:Required
	AggregationMethod *string `json:"aggregationMethod" tf:"aggregation_method,omitempty"`
}

type GroupingObservation struct {
}

type GroupingParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EntityMatchingMethod *string `json:"entityMatchingMethod,omitempty" tf:"entity_matching_method,omitempty"`

	// +kubebuilder:validation:Optional
	GroupBy []*string `json:"groupBy,omitempty" tf:"group_by,omitempty"`

	// +kubebuilder:validation:Optional
	LookbackDuration *string `json:"lookbackDuration,omitempty" tf:"lookback_duration,omitempty"`

	// +kubebuilder:validation:Optional
	ReopenClosedIncidents *bool `json:"reopenClosedIncidents,omitempty" tf:"reopen_closed_incidents,omitempty"`
}

type IncidentConfigurationObservation struct {
}

type IncidentConfigurationParameters struct {

	// +kubebuilder:validation:Required
	CreateIncident *bool `json:"createIncident" tf:"create_incident,omitempty"`

	// +kubebuilder:validation:Required
	Grouping []GroupingParameters `json:"grouping" tf:"grouping,omitempty"`
}

type SentinelAlertRuleScheduledObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SentinelAlertRuleScheduledParameters struct {

	// +kubebuilder:validation:Optional
	AlertRuleTemplateGUID *string `json:"alertRuleTemplateGuid,omitempty" tf:"alert_rule_template_guid,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EventGrouping []EventGroupingParameters `json:"eventGrouping,omitempty" tf:"event_grouping,omitempty"`

	// +kubebuilder:validation:Optional
	IncidentConfiguration []IncidentConfigurationParameters `json:"incidentConfiguration,omitempty" tf:"incident_configuration,omitempty"`

	// +kubebuilder:validation:Required
	LogAnalyticsWorkspaceID *string `json:"logAnalyticsWorkspaceId" tf:"log_analytics_workspace_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Query *string `json:"query" tf:"query,omitempty"`

	// +kubebuilder:validation:Optional
	QueryFrequency *string `json:"queryFrequency,omitempty" tf:"query_frequency,omitempty"`

	// +kubebuilder:validation:Optional
	QueryPeriod *string `json:"queryPeriod,omitempty" tf:"query_period,omitempty"`

	// +kubebuilder:validation:Required
	Severity *string `json:"severity" tf:"severity,omitempty"`

	// +kubebuilder:validation:Optional
	SuppressionDuration *string `json:"suppressionDuration,omitempty" tf:"suppression_duration,omitempty"`

	// +kubebuilder:validation:Optional
	SuppressionEnabled *bool `json:"suppressionEnabled,omitempty" tf:"suppression_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Tactics []*string `json:"tactics,omitempty" tf:"tactics,omitempty"`

	// +kubebuilder:validation:Optional
	TriggerOperator *string `json:"triggerOperator,omitempty" tf:"trigger_operator,omitempty"`

	// +kubebuilder:validation:Optional
	TriggerThreshold *float64 `json:"triggerThreshold,omitempty" tf:"trigger_threshold,omitempty"`
}

// SentinelAlertRuleScheduledSpec defines the desired state of SentinelAlertRuleScheduled
type SentinelAlertRuleScheduledSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SentinelAlertRuleScheduledParameters `json:"forProvider"`
}

// SentinelAlertRuleScheduledStatus defines the observed state of SentinelAlertRuleScheduled.
type SentinelAlertRuleScheduledStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SentinelAlertRuleScheduledObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleScheduled is the Schema for the SentinelAlertRuleScheduleds API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SentinelAlertRuleScheduled struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SentinelAlertRuleScheduledSpec   `json:"spec"`
	Status            SentinelAlertRuleScheduledStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelAlertRuleScheduledList contains a list of SentinelAlertRuleScheduleds
type SentinelAlertRuleScheduledList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelAlertRuleScheduled `json:"items"`
}

// Repository type metadata.
var (
	SentinelAlertRuleScheduled_Kind             = "SentinelAlertRuleScheduled"
	SentinelAlertRuleScheduled_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SentinelAlertRuleScheduled_Kind}.String()
	SentinelAlertRuleScheduled_KindAPIVersion   = SentinelAlertRuleScheduled_Kind + "." + CRDGroupVersion.String()
	SentinelAlertRuleScheduled_GroupVersionKind = CRDGroupVersion.WithKind(SentinelAlertRuleScheduled_Kind)
)

func init() {
	SchemeBuilder.Register(&SentinelAlertRuleScheduled{}, &SentinelAlertRuleScheduledList{})
}

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

type ConditionAlertContextObservation struct {
}

type ConditionAlertContextParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionAlertRuleIDObservation struct {
}

type ConditionAlertRuleIDParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionDescriptionObservation struct {
}

type ConditionDescriptionParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionMonitorObservation struct {
}

type ConditionMonitorParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionMonitorServiceObservation struct {
}

type ConditionMonitorServiceParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionSeverityObservation struct {
}

type ConditionSeverityParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionTargetResourceTypeObservation struct {
}

type ConditionTargetResourceTypeParameters struct {

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorActionRuleSuppressionConditionObservation struct {
}

type MonitorActionRuleSuppressionConditionParameters struct {

	// +kubebuilder:validation:Optional
	AlertContext []ConditionAlertContextParameters `json:"alertContext,omitempty" tf:"alert_context,omitempty"`

	// +kubebuilder:validation:Optional
	AlertRuleID []ConditionAlertRuleIDParameters `json:"alertRuleId,omitempty" tf:"alert_rule_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description []ConditionDescriptionParameters `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Monitor []ConditionMonitorParameters `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// +kubebuilder:validation:Optional
	MonitorService []ConditionMonitorServiceParameters `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`

	// +kubebuilder:validation:Optional
	Severity []ConditionSeverityParameters `json:"severity,omitempty" tf:"severity,omitempty"`

	// +kubebuilder:validation:Optional
	TargetResourceType []ConditionTargetResourceTypeParameters `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type MonitorActionRuleSuppressionObservation struct {
}

type MonitorActionRuleSuppressionParameters struct {

	// +kubebuilder:validation:Optional
	Condition []MonitorActionRuleSuppressionConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Scope []MonitorActionRuleSuppressionScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// +kubebuilder:validation:Required
	Suppression []SuppressionParameters `json:"suppression" tf:"suppression,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorActionRuleSuppressionScopeObservation struct {
}

type MonitorActionRuleSuppressionScopeParameters struct {

	// +kubebuilder:validation:Required
	ResourceIds []*string `json:"resourceIds" tf:"resource_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Required
	EndDateUtc *string `json:"endDateUtc" tf:"end_date_utc,omitempty"`

	// +kubebuilder:validation:Optional
	RecurrenceMonthly []*int64 `json:"recurrenceMonthly,omitempty" tf:"recurrence_monthly,omitempty"`

	// +kubebuilder:validation:Optional
	RecurrenceWeekly []*string `json:"recurrenceWeekly,omitempty" tf:"recurrence_weekly,omitempty"`

	// +kubebuilder:validation:Required
	StartDateUtc *string `json:"startDateUtc" tf:"start_date_utc,omitempty"`
}

type SuppressionObservation struct {
}

type SuppressionParameters struct {

	// +kubebuilder:validation:Required
	RecurrenceType *string `json:"recurrenceType" tf:"recurrence_type,omitempty"`

	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

// MonitorActionRuleSuppressionSpec defines the desired state of MonitorActionRuleSuppression
type MonitorActionRuleSuppressionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorActionRuleSuppressionParameters `json:"forProvider"`
}

// MonitorActionRuleSuppressionStatus defines the observed state of MonitorActionRuleSuppression.
type MonitorActionRuleSuppressionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorActionRuleSuppressionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorActionRuleSuppression is the Schema for the MonitorActionRuleSuppressions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorActionRuleSuppression struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorActionRuleSuppressionSpec   `json:"spec"`
	Status            MonitorActionRuleSuppressionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorActionRuleSuppressionList contains a list of MonitorActionRuleSuppressions
type MonitorActionRuleSuppressionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorActionRuleSuppression `json:"items"`
}

// Repository type metadata.
var (
	MonitorActionRuleSuppressionKind             = "MonitorActionRuleSuppression"
	MonitorActionRuleSuppressionGroupKind        = schema.GroupKind{Group: Group, Kind: MonitorActionRuleSuppressionKind}.String()
	MonitorActionRuleSuppressionKindAPIVersion   = MonitorActionRuleSuppressionKind + "." + GroupVersion.String()
	MonitorActionRuleSuppressionGroupVersionKind = GroupVersion.WithKind(MonitorActionRuleSuppressionKind)
)

func init() {
	SchemeBuilder.Register(&MonitorActionRuleSuppression{}, &MonitorActionRuleSuppressionList{})
}

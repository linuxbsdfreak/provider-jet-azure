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

type AppTriggerRecurrenceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AppTriggerRecurrenceParameters struct {

	// +kubebuilder:validation:Required
	Frequency *string `json:"frequency" tf:"frequency,omitempty"`

	// +kubebuilder:validation:Required
	Interval *float64 `json:"interval" tf:"interval,omitempty"`

	// +kubebuilder:validation:Required
	LogicAppID *string `json:"logicAppId" tf:"logic_app_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Schedule []AppTriggerRecurrenceScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type AppTriggerRecurrenceScheduleObservation struct {
}

type AppTriggerRecurrenceScheduleParameters struct {

	// +kubebuilder:validation:Optional
	AtTheseHours []*float64 `json:"atTheseHours,omitempty" tf:"at_these_hours,omitempty"`

	// +kubebuilder:validation:Optional
	AtTheseMinutes []*float64 `json:"atTheseMinutes,omitempty" tf:"at_these_minutes,omitempty"`

	// +kubebuilder:validation:Optional
	OnTheseDays []*string `json:"onTheseDays,omitempty" tf:"on_these_days,omitempty"`
}

// AppTriggerRecurrenceSpec defines the desired state of AppTriggerRecurrence
type AppTriggerRecurrenceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppTriggerRecurrenceParameters `json:"forProvider"`
}

// AppTriggerRecurrenceStatus defines the observed state of AppTriggerRecurrence.
type AppTriggerRecurrenceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppTriggerRecurrenceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppTriggerRecurrence is the Schema for the AppTriggerRecurrences API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AppTriggerRecurrence struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppTriggerRecurrenceSpec   `json:"spec"`
	Status            AppTriggerRecurrenceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppTriggerRecurrenceList contains a list of AppTriggerRecurrences
type AppTriggerRecurrenceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppTriggerRecurrence `json:"items"`
}

// Repository type metadata.
var (
	AppTriggerRecurrence_Kind             = "AppTriggerRecurrence"
	AppTriggerRecurrence_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppTriggerRecurrence_Kind}.String()
	AppTriggerRecurrence_KindAPIVersion   = AppTriggerRecurrence_Kind + "." + CRDGroupVersion.String()
	AppTriggerRecurrence_GroupVersionKind = CRDGroupVersion.WithKind(AppTriggerRecurrence_Kind)
)

func init() {
	SchemeBuilder.Register(&AppTriggerRecurrence{}, &AppTriggerRecurrenceList{})
}

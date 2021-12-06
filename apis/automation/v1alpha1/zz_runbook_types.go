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

type PublishContentLinkHashObservation struct {
}

type PublishContentLinkHashParameters struct {

	// +kubebuilder:validation:Required
	Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type PublishContentLinkObservation struct {
}

type PublishContentLinkParameters struct {

	// +kubebuilder:validation:Optional
	Hash []PublishContentLinkHashParameters `json:"hash,omitempty" tf:"hash,omitempty"`

	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type RunbookJobScheduleObservation struct {
	JobScheduleID *string `json:"jobScheduleId,omitempty" tf:"job_schedule_id,omitempty"`
}

type RunbookJobScheduleParameters struct {

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	RunOn *string `json:"runOn,omitempty" tf:"run_on,omitempty"`

	// +kubebuilder:validation:Required
	ScheduleName *string `json:"scheduleName" tf:"schedule_name,omitempty"`
}

type RunbookObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RunbookParameters struct {

	// +kubebuilder:validation:Required
	AutomationAccountName *string `json:"automationAccountName" tf:"automation_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	JobSchedule []RunbookJobScheduleParameters `json:"jobSchedule,omitempty" tf:"job_schedule,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	LogProgress *bool `json:"logProgress" tf:"log_progress,omitempty"`

	// +kubebuilder:validation:Required
	LogVerbose *bool `json:"logVerbose" tf:"log_verbose,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PublishContentLink []PublishContentLinkParameters `json:"publishContentLink,omitempty" tf:"publish_content_link,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	RunbookType *string `json:"runbookType" tf:"runbook_type,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RunbookSpec defines the desired state of Runbook
type RunbookSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RunbookParameters `json:"forProvider"`
}

// RunbookStatus defines the observed state of Runbook.
type RunbookStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RunbookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Runbook is the Schema for the Runbooks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Runbook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RunbookSpec   `json:"spec"`
	Status            RunbookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RunbookList contains a list of Runbooks
type RunbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Runbook `json:"items"`
}

// Repository type metadata.
var (
	Runbook_Kind             = "Runbook"
	Runbook_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Runbook_Kind}.String()
	Runbook_KindAPIVersion   = Runbook_Kind + "." + CRDGroupVersion.String()
	Runbook_GroupVersionKind = CRDGroupVersion.WithKind(Runbook_Kind)
)

func init() {
	SchemeBuilder.Register(&Runbook{}, &RunbookList{})
}
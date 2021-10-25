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

type LogAnalyticsSolutionObservation struct {
}

type LogAnalyticsSolutionParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Plan []PlanParameters `json:"plan" tf:"plan,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SolutionName *string `json:"solutionName" tf:"solution_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	WorkspaceName *string `json:"workspaceName" tf:"workspace_name,omitempty"`

	// +kubebuilder:validation:Required
	WorkspaceResourceID *string `json:"workspaceResourceId" tf:"workspace_resource_id,omitempty"`
}

type PlanObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PlanParameters struct {

	// +kubebuilder:validation:Required
	Product *string `json:"product" tf:"product,omitempty"`

	// +kubebuilder:validation:Optional
	PromotionCode *string `json:"promotionCode,omitempty" tf:"promotion_code,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`
}

// LogAnalyticsSolutionSpec defines the desired state of LogAnalyticsSolution
type LogAnalyticsSolutionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogAnalyticsSolutionParameters `json:"forProvider"`
}

// LogAnalyticsSolutionStatus defines the observed state of LogAnalyticsSolution.
type LogAnalyticsSolutionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogAnalyticsSolutionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsSolution is the Schema for the LogAnalyticsSolutions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogAnalyticsSolution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsSolutionSpec   `json:"spec"`
	Status            LogAnalyticsSolutionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogAnalyticsSolutionList contains a list of LogAnalyticsSolutions
type LogAnalyticsSolutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogAnalyticsSolution `json:"items"`
}

// Repository type metadata.
var (
	LogAnalyticsSolutionKind             = "LogAnalyticsSolution"
	LogAnalyticsSolutionGroupKind        = schema.GroupKind{Group: Group, Kind: LogAnalyticsSolutionKind}.String()
	LogAnalyticsSolutionKindAPIVersion   = LogAnalyticsSolutionKind + "." + GroupVersion.String()
	LogAnalyticsSolutionGroupVersionKind = GroupVersion.WithKind(LogAnalyticsSolutionKind)
)

func init() {
	SchemeBuilder.Register(&LogAnalyticsSolution{}, &LogAnalyticsSolutionList{})
}

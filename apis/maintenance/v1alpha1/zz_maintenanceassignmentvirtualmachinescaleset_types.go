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

type MaintenanceAssignmentVirtualMachineScaleSetObservation struct {
}

type MaintenanceAssignmentVirtualMachineScaleSetParameters struct {

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	MaintenanceConfigurationID *string `json:"maintenanceConfigurationId" tf:"maintenance_configuration_id,omitempty"`

	// +kubebuilder:validation:Required
	VirtualMachineScaleSetID *string `json:"virtualMachineScaleSetId" tf:"virtual_machine_scale_set_id,omitempty"`
}

// MaintenanceAssignmentVirtualMachineScaleSetSpec defines the desired state of MaintenanceAssignmentVirtualMachineScaleSet
type MaintenanceAssignmentVirtualMachineScaleSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MaintenanceAssignmentVirtualMachineScaleSetParameters `json:"forProvider"`
}

// MaintenanceAssignmentVirtualMachineScaleSetStatus defines the observed state of MaintenanceAssignmentVirtualMachineScaleSet.
type MaintenanceAssignmentVirtualMachineScaleSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MaintenanceAssignmentVirtualMachineScaleSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceAssignmentVirtualMachineScaleSet is the Schema for the MaintenanceAssignmentVirtualMachineScaleSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MaintenanceAssignmentVirtualMachineScaleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MaintenanceAssignmentVirtualMachineScaleSetSpec   `json:"spec"`
	Status            MaintenanceAssignmentVirtualMachineScaleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceAssignmentVirtualMachineScaleSetList contains a list of MaintenanceAssignmentVirtualMachineScaleSets
type MaintenanceAssignmentVirtualMachineScaleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MaintenanceAssignmentVirtualMachineScaleSet `json:"items"`
}

// Repository type metadata.
var (
	MaintenanceAssignmentVirtualMachineScaleSetKind             = "MaintenanceAssignmentVirtualMachineScaleSet"
	MaintenanceAssignmentVirtualMachineScaleSetGroupKind        = schema.GroupKind{Group: Group, Kind: MaintenanceAssignmentVirtualMachineScaleSetKind}.String()
	MaintenanceAssignmentVirtualMachineScaleSetKindAPIVersion   = MaintenanceAssignmentVirtualMachineScaleSetKind + "." + GroupVersion.String()
	MaintenanceAssignmentVirtualMachineScaleSetGroupVersionKind = GroupVersion.WithKind(MaintenanceAssignmentVirtualMachineScaleSetKind)
)

func init() {
	SchemeBuilder.Register(&MaintenanceAssignmentVirtualMachineScaleSet{}, &MaintenanceAssignmentVirtualMachineScaleSetList{})
}

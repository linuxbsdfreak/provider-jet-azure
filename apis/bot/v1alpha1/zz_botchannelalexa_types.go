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

type BotChannelAlexaObservation struct {
}

type BotChannelAlexaParameters struct {

	// +kubebuilder:validation:Required
	BotName *string `json:"botName" tf:"bot_name,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SkillID *string `json:"skillId" tf:"skill_id,omitempty"`
}

// BotChannelAlexaSpec defines the desired state of BotChannelAlexa
type BotChannelAlexaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotChannelAlexaParameters `json:"forProvider"`
}

// BotChannelAlexaStatus defines the observed state of BotChannelAlexa.
type BotChannelAlexaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotChannelAlexaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelAlexa is the Schema for the BotChannelAlexas API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BotChannelAlexa struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BotChannelAlexaSpec   `json:"spec"`
	Status            BotChannelAlexaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotChannelAlexaList contains a list of BotChannelAlexas
type BotChannelAlexaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotChannelAlexa `json:"items"`
}

// Repository type metadata.
var (
	BotChannelAlexaKind             = "BotChannelAlexa"
	BotChannelAlexaGroupKind        = schema.GroupKind{Group: Group, Kind: BotChannelAlexaKind}.String()
	BotChannelAlexaKindAPIVersion   = BotChannelAlexaKind + "." + GroupVersion.String()
	BotChannelAlexaGroupVersionKind = GroupVersion.WithKind(BotChannelAlexaKind)
)

func init() {
	SchemeBuilder.Register(&BotChannelAlexa{}, &BotChannelAlexaList{})
}

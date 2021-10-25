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

type ApiManagementSubscriptionObservation struct {
}

type ApiManagementSubscriptionParameters struct {

	// +kubebuilder:validation:Optional
	APIID *string `json:"apiId,omitempty" tf:"api_id,omitempty"`

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Optional
	AllowTracing *bool `json:"allowTracing,omitempty" tf:"allow_tracing,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Optional
	PrimaryKeySecretRef v1.SecretKeySelector `json:"primaryKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProductID *string `json:"productId,omitempty" tf:"product_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	SecondaryKeySecretRef v1.SecretKeySelector `json:"secondaryKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// +kubebuilder:validation:Optional
	SubscriptionID *string `json:"subscriptionId,omitempty" tf:"subscription_id,omitempty"`

	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// ApiManagementSubscriptionSpec defines the desired state of ApiManagementSubscription
type ApiManagementSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApiManagementSubscriptionParameters `json:"forProvider"`
}

// ApiManagementSubscriptionStatus defines the observed state of ApiManagementSubscription.
type ApiManagementSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApiManagementSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementSubscription is the Schema for the ApiManagementSubscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementSubscriptionSpec   `json:"spec"`
	Status            ApiManagementSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementSubscriptionList contains a list of ApiManagementSubscriptions
type ApiManagementSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementSubscription `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementSubscriptionKind             = "ApiManagementSubscription"
	ApiManagementSubscriptionGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementSubscriptionKind}.String()
	ApiManagementSubscriptionKindAPIVersion   = ApiManagementSubscriptionKind + "." + GroupVersion.String()
	ApiManagementSubscriptionGroupVersionKind = GroupVersion.WithKind(ApiManagementSubscriptionKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementSubscription{}, &ApiManagementSubscriptionList{})
}

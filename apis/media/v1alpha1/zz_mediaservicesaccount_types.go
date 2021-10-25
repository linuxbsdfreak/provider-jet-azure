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

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type KeyDeliveryAccessControlObservation struct {
}

type KeyDeliveryAccessControlParameters struct {

	// +kubebuilder:validation:Optional
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// +kubebuilder:validation:Optional
	IPAllowList []*string `json:"ipAllowList,omitempty" tf:"ip_allow_list,omitempty"`
}

type MediaServicesAccountObservation struct {
}

type MediaServicesAccountParameters struct {

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Optional
	KeyDeliveryAccessControl []KeyDeliveryAccessControlParameters `json:"keyDeliveryAccessControl,omitempty" tf:"key_delivery_access_control,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccount []StorageAccountParameters `json:"storageAccount" tf:"storage_account,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAuthenticationType *string `json:"storageAuthenticationType,omitempty" tf:"storage_authentication_type,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StorageAccountObservation struct {
}

type StorageAccountParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	IsPrimary *bool `json:"isPrimary,omitempty" tf:"is_primary,omitempty"`
}

// MediaServicesAccountSpec defines the desired state of MediaServicesAccount
type MediaServicesAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MediaServicesAccountParameters `json:"forProvider"`
}

// MediaServicesAccountStatus defines the observed state of MediaServicesAccount.
type MediaServicesAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MediaServicesAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MediaServicesAccount is the Schema for the MediaServicesAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MediaServicesAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MediaServicesAccountSpec   `json:"spec"`
	Status            MediaServicesAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MediaServicesAccountList contains a list of MediaServicesAccounts
type MediaServicesAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MediaServicesAccount `json:"items"`
}

// Repository type metadata.
var (
	MediaServicesAccountKind             = "MediaServicesAccount"
	MediaServicesAccountGroupKind        = schema.GroupKind{Group: Group, Kind: MediaServicesAccountKind}.String()
	MediaServicesAccountKindAPIVersion   = MediaServicesAccountKind + "." + GroupVersion.String()
	MediaServicesAccountGroupVersionKind = GroupVersion.WithKind(MediaServicesAccountKind)
)

func init() {
	SchemeBuilder.Register(&MediaServicesAccount{}, &MediaServicesAccountList{})
}

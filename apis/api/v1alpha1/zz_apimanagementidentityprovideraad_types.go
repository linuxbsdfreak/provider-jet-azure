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

type ApiManagementIdentityProviderAadObservation struct {
}

type ApiManagementIdentityProviderAadParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Required
	AllowedTenants []*string `json:"allowedTenants" tf:"allowed_tenants,omitempty"`

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	SigninTenant *string `json:"signinTenant,omitempty" tf:"signin_tenant,omitempty"`
}

// ApiManagementIdentityProviderAadSpec defines the desired state of ApiManagementIdentityProviderAad
type ApiManagementIdentityProviderAadSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApiManagementIdentityProviderAadParameters `json:"forProvider"`
}

// ApiManagementIdentityProviderAadStatus defines the observed state of ApiManagementIdentityProviderAad.
type ApiManagementIdentityProviderAadStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApiManagementIdentityProviderAadObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementIdentityProviderAad is the Schema for the ApiManagementIdentityProviderAads API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementIdentityProviderAad struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementIdentityProviderAadSpec   `json:"spec"`
	Status            ApiManagementIdentityProviderAadStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementIdentityProviderAadList contains a list of ApiManagementIdentityProviderAads
type ApiManagementIdentityProviderAadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementIdentityProviderAad `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementIdentityProviderAadKind             = "ApiManagementIdentityProviderAad"
	ApiManagementIdentityProviderAadGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementIdentityProviderAadKind}.String()
	ApiManagementIdentityProviderAadKindAPIVersion   = ApiManagementIdentityProviderAadKind + "." + GroupVersion.String()
	ApiManagementIdentityProviderAadGroupVersionKind = GroupVersion.WithKind(ApiManagementIdentityProviderAadKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementIdentityProviderAad{}, &ApiManagementIdentityProviderAadList{})
}

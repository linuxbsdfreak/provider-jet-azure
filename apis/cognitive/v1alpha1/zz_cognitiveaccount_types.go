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

type CognitiveAccountObservation struct {
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
}

type CognitiveAccountParameters struct {

	// +kubebuilder:validation:Optional
	CustomSubdomainName *string `json:"customSubdomainName,omitempty" tf:"custom_subdomain_name,omitempty"`

	// +kubebuilder:validation:Optional
	Fqdns []*string `json:"fqdns,omitempty" tf:"fqdns,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Kind *string `json:"kind" tf:"kind,omitempty"`

	// +kubebuilder:validation:Optional
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MetricsAdvisorAadClientID *string `json:"metricsAdvisorAadClientId,omitempty" tf:"metrics_advisor_aad_client_id,omitempty"`

	// +kubebuilder:validation:Optional
	MetricsAdvisorAadTenantID *string `json:"metricsAdvisorAadTenantId,omitempty" tf:"metrics_advisor_aad_tenant_id,omitempty"`

	// +kubebuilder:validation:Optional
	MetricsAdvisorSuperUserName *string `json:"metricsAdvisorSuperUserName,omitempty" tf:"metrics_advisor_super_user_name,omitempty"`

	// +kubebuilder:validation:Optional
	MetricsAdvisorWebsiteName *string `json:"metricsAdvisorWebsiteName,omitempty" tf:"metrics_advisor_website_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkAcls []NetworkAclsParameters `json:"networkAcls,omitempty" tf:"network_acls,omitempty"`

	// +kubebuilder:validation:Optional
	OutboundNetworkAccessRestrited *bool `json:"outboundNetworkAccessRestrited,omitempty" tf:"outbound_network_access_restrited,omitempty"`

	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	QnaRuntimeEndpoint *string `json:"qnaRuntimeEndpoint,omitempty" tf:"qna_runtime_endpoint,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Storage []StorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkAclsObservation struct {
}

type NetworkAclsParameters struct {

	// +kubebuilder:validation:Required
	DefaultAction *string `json:"defaultAction" tf:"default_action,omitempty"`

	// +kubebuilder:validation:Optional
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkRules []VirtualNetworkRulesParameters `json:"virtualNetworkRules,omitempty" tf:"virtual_network_rules,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIds []*string `json:"virtualNetworkSubnetIds,omitempty" tf:"virtual_network_subnet_ids,omitempty"`
}

type StorageObservation struct {
}

type StorageParameters struct {

	// +kubebuilder:validation:Optional
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

type VirtualNetworkRulesObservation struct {
}

type VirtualNetworkRulesParameters struct {

	// +kubebuilder:validation:Optional
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
}

// CognitiveAccountSpec defines the desired state of CognitiveAccount
type CognitiveAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CognitiveAccountParameters `json:"forProvider"`
}

// CognitiveAccountStatus defines the observed state of CognitiveAccount.
type CognitiveAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CognitiveAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CognitiveAccount is the Schema for the CognitiveAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CognitiveAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitiveAccountSpec   `json:"spec"`
	Status            CognitiveAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CognitiveAccountList contains a list of CognitiveAccounts
type CognitiveAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CognitiveAccount `json:"items"`
}

// Repository type metadata.
var (
	CognitiveAccountKind             = "CognitiveAccount"
	CognitiveAccountGroupKind        = schema.GroupKind{Group: Group, Kind: CognitiveAccountKind}.String()
	CognitiveAccountKindAPIVersion   = CognitiveAccountKind + "." + GroupVersion.String()
	CognitiveAccountGroupVersionKind = GroupVersion.WithKind(CognitiveAccountKind)
)

func init() {
	SchemeBuilder.Register(&CognitiveAccount{}, &CognitiveAccountList{})
}

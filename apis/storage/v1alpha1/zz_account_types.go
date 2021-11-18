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

type AccountObservation struct {
	PrimaryBlobEndpoint *string `json:"primaryBlobEndpoint,omitempty" tf:"primary_blob_endpoint,omitempty"`

	PrimaryBlobHost *string `json:"primaryBlobHost,omitempty" tf:"primary_blob_host,omitempty"`

	PrimaryDfsEndpoint *string `json:"primaryDfsEndpoint,omitempty" tf:"primary_dfs_endpoint,omitempty"`

	PrimaryDfsHost *string `json:"primaryDfsHost,omitempty" tf:"primary_dfs_host,omitempty"`

	PrimaryFileEndpoint *string `json:"primaryFileEndpoint,omitempty" tf:"primary_file_endpoint,omitempty"`

	PrimaryFileHost *string `json:"primaryFileHost,omitempty" tf:"primary_file_host,omitempty"`

	PrimaryLocation *string `json:"primaryLocation,omitempty" tf:"primary_location,omitempty"`

	PrimaryQueueEndpoint *string `json:"primaryQueueEndpoint,omitempty" tf:"primary_queue_endpoint,omitempty"`

	PrimaryQueueHost *string `json:"primaryQueueHost,omitempty" tf:"primary_queue_host,omitempty"`

	PrimaryTableEndpoint *string `json:"primaryTableEndpoint,omitempty" tf:"primary_table_endpoint,omitempty"`

	PrimaryTableHost *string `json:"primaryTableHost,omitempty" tf:"primary_table_host,omitempty"`

	PrimaryWebEndpoint *string `json:"primaryWebEndpoint,omitempty" tf:"primary_web_endpoint,omitempty"`

	PrimaryWebHost *string `json:"primaryWebHost,omitempty" tf:"primary_web_host,omitempty"`

	SecondaryBlobEndpoint *string `json:"secondaryBlobEndpoint,omitempty" tf:"secondary_blob_endpoint,omitempty"`

	SecondaryBlobHost *string `json:"secondaryBlobHost,omitempty" tf:"secondary_blob_host,omitempty"`

	SecondaryDfsEndpoint *string `json:"secondaryDfsEndpoint,omitempty" tf:"secondary_dfs_endpoint,omitempty"`

	SecondaryDfsHost *string `json:"secondaryDfsHost,omitempty" tf:"secondary_dfs_host,omitempty"`

	SecondaryFileEndpoint *string `json:"secondaryFileEndpoint,omitempty" tf:"secondary_file_endpoint,omitempty"`

	SecondaryFileHost *string `json:"secondaryFileHost,omitempty" tf:"secondary_file_host,omitempty"`

	SecondaryLocation *string `json:"secondaryLocation,omitempty" tf:"secondary_location,omitempty"`

	SecondaryQueueEndpoint *string `json:"secondaryQueueEndpoint,omitempty" tf:"secondary_queue_endpoint,omitempty"`

	SecondaryQueueHost *string `json:"secondaryQueueHost,omitempty" tf:"secondary_queue_host,omitempty"`

	SecondaryTableEndpoint *string `json:"secondaryTableEndpoint,omitempty" tf:"secondary_table_endpoint,omitempty"`

	SecondaryTableHost *string `json:"secondaryTableHost,omitempty" tf:"secondary_table_host,omitempty"`

	SecondaryWebEndpoint *string `json:"secondaryWebEndpoint,omitempty" tf:"secondary_web_endpoint,omitempty"`

	SecondaryWebHost *string `json:"secondaryWebHost,omitempty" tf:"secondary_web_host,omitempty"`
}

type AccountParameters struct {

	// +kubebuilder:validation:Optional
	AccessTier *string `json:"accessTier,omitempty" tf:"access_tier,omitempty"`

	// +kubebuilder:validation:Optional
	AccountKind *string `json:"accountKind,omitempty" tf:"account_kind,omitempty"`

	// +kubebuilder:validation:Required
	AccountReplicationType *string `json:"accountReplicationType" tf:"account_replication_type,omitempty"`

	// +kubebuilder:validation:Required
	AccountTier *string `json:"accountTier" tf:"account_tier,omitempty"`

	// +kubebuilder:validation:Optional
	AllowBlobPublicAccess *bool `json:"allowBlobPublicAccess,omitempty" tf:"allow_blob_public_access,omitempty"`

	// +kubebuilder:validation:Optional
	AzureFilesAuthentication []AzureFilesAuthenticationParameters `json:"azureFilesAuthentication,omitempty" tf:"azure_files_authentication,omitempty"`

	// +kubebuilder:validation:Optional
	BlobProperties []BlobPropertiesParameters `json:"blobProperties,omitempty" tf:"blob_properties,omitempty"`

	// +kubebuilder:validation:Optional
	CustomDomain []CustomDomainParameters `json:"customDomain,omitempty" tf:"custom_domain,omitempty"`

	// +kubebuilder:validation:Optional
	EnableHTTPSTrafficOnly *bool `json:"enableHttpsTrafficOnly,omitempty" tf:"enable_https_traffic_only,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Optional
	IsHnsEnabled *bool `json:"isHnsEnabled,omitempty" tf:"is_hns_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	LargeFileShareEnabled *bool `json:"largeFileShareEnabled,omitempty" tf:"large_file_share_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MinTLSVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkRules []NetworkRulesParameters `json:"networkRules,omitempty" tf:"network_rules,omitempty"`

	// +kubebuilder:validation:Optional
	Nfsv3Enabled *bool `json:"nfsv3Enabled,omitempty" tf:"nfsv3_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	QueueProperties []QueuePropertiesParameters `json:"queueProperties,omitempty" tf:"queue_properties,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-azure/apis/azure/v1alpha1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Routing []RoutingParameters `json:"routing,omitempty" tf:"routing,omitempty"`

	// +kubebuilder:validation:Optional
	ShareProperties []SharePropertiesParameters `json:"shareProperties,omitempty" tf:"share_properties,omitempty"`

	// +kubebuilder:validation:Optional
	SharedAccessKeyEnabled *bool `json:"sharedAccessKeyEnabled,omitempty" tf:"shared_access_key_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	StaticWebsite []StaticWebsiteParameters `json:"staticWebsite,omitempty" tf:"static_website,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ActiveDirectoryObservation struct {
}

type ActiveDirectoryParameters struct {

	// +kubebuilder:validation:Required
	DomainGUID *string `json:"domainGuid" tf:"domain_guid,omitempty"`

	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Required
	DomainSid *string `json:"domainSid" tf:"domain_sid,omitempty"`

	// +kubebuilder:validation:Required
	ForestName *string `json:"forestName" tf:"forest_name,omitempty"`

	// +kubebuilder:validation:Required
	NetbiosDomainName *string `json:"netbiosDomainName" tf:"netbios_domain_name,omitempty"`

	// +kubebuilder:validation:Required
	StorageSid *string `json:"storageSid" tf:"storage_sid,omitempty"`
}

type AzureFilesAuthenticationObservation struct {
}

type AzureFilesAuthenticationParameters struct {

	// +kubebuilder:validation:Optional
	ActiveDirectory []ActiveDirectoryParameters `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`

	// +kubebuilder:validation:Required
	DirectoryType *string `json:"directoryType" tf:"directory_type,omitempty"`
}

type BlobPropertiesObservation struct {
}

type BlobPropertiesParameters struct {

	// +kubebuilder:validation:Optional
	ChangeFeedEnabled *bool `json:"changeFeedEnabled,omitempty" tf:"change_feed_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ContainerDeleteRetentionPolicy []ContainerDeleteRetentionPolicyParameters `json:"containerDeleteRetentionPolicy,omitempty" tf:"container_delete_retention_policy,omitempty"`

	// +kubebuilder:validation:Optional
	CorsRule []CorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultServiceVersion *string `json:"defaultServiceVersion,omitempty" tf:"default_service_version,omitempty"`

	// +kubebuilder:validation:Optional
	DeleteRetentionPolicy []DeleteRetentionPolicyParameters `json:"deleteRetentionPolicy,omitempty" tf:"delete_retention_policy,omitempty"`

	// +kubebuilder:validation:Optional
	LastAccessTimeEnabled *bool `json:"lastAccessTimeEnabled,omitempty" tf:"last_access_time_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	VersioningEnabled *bool `json:"versioningEnabled,omitempty" tf:"versioning_enabled,omitempty"`
}

type ContainerDeleteRetentionPolicyObservation struct {
}

type ContainerDeleteRetentionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`
}

type CorsRuleObservation struct {
}

type CorsRuleParameters struct {

	// +kubebuilder:validation:Required
	AllowedHeaders []*string `json:"allowedHeaders" tf:"allowed_headers,omitempty"`

	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// +kubebuilder:validation:Required
	ExposedHeaders []*string `json:"exposedHeaders" tf:"exposed_headers,omitempty"`

	// +kubebuilder:validation:Required
	MaxAgeInSeconds *int64 `json:"maxAgeInSeconds" tf:"max_age_in_seconds,omitempty"`
}

type CustomDomainObservation struct {
}

type CustomDomainParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	UseSubdomain *bool `json:"useSubdomain,omitempty" tf:"use_subdomain,omitempty"`
}

type DeleteRetentionPolicyObservation struct {
}

type DeleteRetentionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`
}

type HourMetricsObservation struct {
}

type HourMetricsParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeApis *bool `json:"includeApis,omitempty" tf:"include_apis,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionPolicyDays *int64 `json:"retentionPolicyDays,omitempty" tf:"retention_policy_days,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LoggingObservation struct {
}

type LoggingParameters struct {

	// +kubebuilder:validation:Required
	Delete *bool `json:"delete" tf:"delete,omitempty"`

	// +kubebuilder:validation:Required
	Read *bool `json:"read" tf:"read,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionPolicyDays *int64 `json:"retentionPolicyDays,omitempty" tf:"retention_policy_days,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`

	// +kubebuilder:validation:Required
	Write *bool `json:"write" tf:"write,omitempty"`
}

type MinuteMetricsObservation struct {
}

type MinuteMetricsParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeApis *bool `json:"includeApis,omitempty" tf:"include_apis,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionPolicyDays *int64 `json:"retentionPolicyDays,omitempty" tf:"retention_policy_days,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type NetworkRulesObservation struct {
}

type NetworkRulesParameters struct {

	// +kubebuilder:validation:Optional
	Bypass []*string `json:"bypass,omitempty" tf:"bypass,omitempty"`

	// +kubebuilder:validation:Required
	DefaultAction *string `json:"defaultAction" tf:"default_action,omitempty"`

	// +kubebuilder:validation:Optional
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateLinkAccess []PrivateLinkAccessParameters `json:"privateLinkAccess,omitempty" tf:"private_link_access,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIds []*string `json:"virtualNetworkSubnetIds,omitempty" tf:"virtual_network_subnet_ids,omitempty"`
}

type PrivateLinkAccessObservation struct {
}

type PrivateLinkAccessParameters struct {

	// +kubebuilder:validation:Required
	EndpointResourceID *string `json:"endpointResourceId" tf:"endpoint_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointTenantID *string `json:"endpointTenantId,omitempty" tf:"endpoint_tenant_id,omitempty"`
}

type QueuePropertiesCorsRuleObservation struct {
}

type QueuePropertiesCorsRuleParameters struct {

	// +kubebuilder:validation:Required
	AllowedHeaders []*string `json:"allowedHeaders" tf:"allowed_headers,omitempty"`

	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// +kubebuilder:validation:Required
	ExposedHeaders []*string `json:"exposedHeaders" tf:"exposed_headers,omitempty"`

	// +kubebuilder:validation:Required
	MaxAgeInSeconds *int64 `json:"maxAgeInSeconds" tf:"max_age_in_seconds,omitempty"`
}

type QueuePropertiesObservation struct {
}

type QueuePropertiesParameters struct {

	// +kubebuilder:validation:Optional
	CorsRule []QueuePropertiesCorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// +kubebuilder:validation:Optional
	HourMetrics []HourMetricsParameters `json:"hourMetrics,omitempty" tf:"hour_metrics,omitempty"`

	// +kubebuilder:validation:Optional
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// +kubebuilder:validation:Optional
	MinuteMetrics []MinuteMetricsParameters `json:"minuteMetrics,omitempty" tf:"minute_metrics,omitempty"`
}

type RetentionPolicyObservation struct {
}

type RetentionPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Days *int64 `json:"days,omitempty" tf:"days,omitempty"`
}

type RoutingObservation struct {
}

type RoutingParameters struct {

	// +kubebuilder:validation:Optional
	Choice *string `json:"choice,omitempty" tf:"choice,omitempty"`

	// +kubebuilder:validation:Optional
	PublishInternetEndpoints *bool `json:"publishInternetEndpoints,omitempty" tf:"publish_internet_endpoints,omitempty"`

	// +kubebuilder:validation:Optional
	PublishMicrosoftEndpoints *bool `json:"publishMicrosoftEndpoints,omitempty" tf:"publish_microsoft_endpoints,omitempty"`
}

type SharePropertiesCorsRuleObservation struct {
}

type SharePropertiesCorsRuleParameters struct {

	// +kubebuilder:validation:Required
	AllowedHeaders []*string `json:"allowedHeaders" tf:"allowed_headers,omitempty"`

	// +kubebuilder:validation:Required
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// +kubebuilder:validation:Required
	ExposedHeaders []*string `json:"exposedHeaders" tf:"exposed_headers,omitempty"`

	// +kubebuilder:validation:Required
	MaxAgeInSeconds *int64 `json:"maxAgeInSeconds" tf:"max_age_in_seconds,omitempty"`
}

type SharePropertiesObservation struct {
}

type SharePropertiesParameters struct {

	// +kubebuilder:validation:Optional
	CorsRule []SharePropertiesCorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// +kubebuilder:validation:Optional
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Smb []SmbParameters `json:"smb,omitempty" tf:"smb,omitempty"`
}

type SmbObservation struct {
}

type SmbParameters struct {

	// +kubebuilder:validation:Optional
	AuthenticationTypes []*string `json:"authenticationTypes,omitempty" tf:"authentication_types,omitempty"`

	// +kubebuilder:validation:Optional
	ChannelEncryptionType []*string `json:"channelEncryptionType,omitempty" tf:"channel_encryption_type,omitempty"`

	// +kubebuilder:validation:Optional
	KerberosTicketEncryptionType []*string `json:"kerberosTicketEncryptionType,omitempty" tf:"kerberos_ticket_encryption_type,omitempty"`

	// +kubebuilder:validation:Optional
	Versions []*string `json:"versions,omitempty" tf:"versions,omitempty"`
}

type StaticWebsiteObservation struct {
}

type StaticWebsiteParameters struct {

	// +kubebuilder:validation:Optional
	Error404Document *string `json:"error404Document,omitempty" tf:"error_404_document,omitempty"`

	// +kubebuilder:validation:Optional
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`
}

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountParameters `json:"forProvider"`
}

// AccountStatus defines the observed state of Account.
type AccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Account is the Schema for the Accounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfazure}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountSpec   `json:"spec"`
	Status            AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountList contains a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

// Repository type metadata.
var (
	Account_Kind             = "Account"
	Account_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Account_Kind}.String()
	Account_KindAPIVersion   = Account_Kind + "." + CRDGroupVersion.String()
	Account_GroupVersionKind = CRDGroupVersion.WithKind(Account_Kind)
)

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}

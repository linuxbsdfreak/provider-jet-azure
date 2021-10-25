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

type AadAuthObservation struct {
}

type AadAuthParameters struct {

	// +kubebuilder:validation:Optional
	IdentifierURI *string `json:"identifierUri,omitempty" tf:"identifier_uri,omitempty"`

	// +kubebuilder:validation:Required
	ObjectID *string `json:"objectId" tf:"object_id,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ArmRoleReceiverObservation struct {
}

type ArmRoleReceiverParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	RoleID *string `json:"roleId" tf:"role_id,omitempty"`

	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type AutomationRunbookReceiverObservation struct {
}

type AutomationRunbookReceiverParameters struct {

	// +kubebuilder:validation:Required
	AutomationAccountID *string `json:"automationAccountId" tf:"automation_account_id,omitempty"`

	// +kubebuilder:validation:Required
	IsGlobalRunbook *bool `json:"isGlobalRunbook" tf:"is_global_runbook,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	RunbookName *string `json:"runbookName" tf:"runbook_name,omitempty"`

	// +kubebuilder:validation:Required
	ServiceURI *string `json:"serviceUri" tf:"service_uri,omitempty"`

	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`

	// +kubebuilder:validation:Required
	WebhookResourceID *string `json:"webhookResourceId" tf:"webhook_resource_id,omitempty"`
}

type AzureAppPushReceiverObservation struct {
}

type AzureAppPushReceiverParameters struct {

	// +kubebuilder:validation:Required
	EmailAddress *string `json:"emailAddress" tf:"email_address,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type AzureFunctionReceiverObservation struct {
}

type AzureFunctionReceiverParameters struct {

	// +kubebuilder:validation:Required
	FunctionAppResourceID *string `json:"functionAppResourceId" tf:"function_app_resource_id,omitempty"`

	// +kubebuilder:validation:Required
	FunctionName *string `json:"functionName" tf:"function_name,omitempty"`

	// +kubebuilder:validation:Required
	HTTPTriggerURL *string `json:"httpTriggerUrl" tf:"http_trigger_url,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type EmailReceiverObservation struct {
}

type EmailReceiverParameters struct {

	// +kubebuilder:validation:Required
	EmailAddress *string `json:"emailAddress" tf:"email_address,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type ItsmReceiverObservation struct {
}

type ItsmReceiverParameters struct {

	// +kubebuilder:validation:Required
	ConnectionID *string `json:"connectionId" tf:"connection_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	TicketConfiguration *string `json:"ticketConfiguration" tf:"ticket_configuration,omitempty"`

	// +kubebuilder:validation:Required
	WorkspaceID *string `json:"workspaceId" tf:"workspace_id,omitempty"`
}

type LogicAppReceiverObservation struct {
}

type LogicAppReceiverParameters struct {

	// +kubebuilder:validation:Required
	CallbackURL *string `json:"callbackUrl" tf:"callback_url,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceID *string `json:"resourceId" tf:"resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

type MonitorActionGroupObservation struct {
}

type MonitorActionGroupParameters struct {

	// +kubebuilder:validation:Optional
	ArmRoleReceiver []ArmRoleReceiverParameters `json:"armRoleReceiver,omitempty" tf:"arm_role_receiver,omitempty"`

	// +kubebuilder:validation:Optional
	AutomationRunbookReceiver []AutomationRunbookReceiverParameters `json:"automationRunbookReceiver,omitempty" tf:"automation_runbook_receiver,omitempty"`

	// +kubebuilder:validation:Optional
	AzureAppPushReceiver []AzureAppPushReceiverParameters `json:"azureAppPushReceiver,omitempty" tf:"azure_app_push_receiver,omitempty"`

	// +kubebuilder:validation:Optional
	AzureFunctionReceiver []AzureFunctionReceiverParameters `json:"azureFunctionReceiver,omitempty" tf:"azure_function_receiver,omitempty"`

	// +kubebuilder:validation:Optional
	EmailReceiver []EmailReceiverParameters `json:"emailReceiver,omitempty" tf:"email_receiver,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ItsmReceiver []ItsmReceiverParameters `json:"itsmReceiver,omitempty" tf:"itsm_receiver,omitempty"`

	// +kubebuilder:validation:Optional
	LogicAppReceiver []LogicAppReceiverParameters `json:"logicAppReceiver,omitempty" tf:"logic_app_receiver,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	ShortName *string `json:"shortName" tf:"short_name,omitempty"`

	// +kubebuilder:validation:Optional
	SmsReceiver []SmsReceiverParameters `json:"smsReceiver,omitempty" tf:"sms_receiver,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VoiceReceiver []VoiceReceiverParameters `json:"voiceReceiver,omitempty" tf:"voice_receiver,omitempty"`

	// +kubebuilder:validation:Optional
	WebhookReceiver []WebhookReceiverParameters `json:"webhookReceiver,omitempty" tf:"webhook_receiver,omitempty"`
}

type SmsReceiverObservation struct {
}

type SmsReceiverParameters struct {

	// +kubebuilder:validation:Required
	CountryCode *string `json:"countryCode" tf:"country_code,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PhoneNumber *string `json:"phoneNumber" tf:"phone_number,omitempty"`
}

type VoiceReceiverObservation struct {
}

type VoiceReceiverParameters struct {

	// +kubebuilder:validation:Required
	CountryCode *string `json:"countryCode" tf:"country_code,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PhoneNumber *string `json:"phoneNumber" tf:"phone_number,omitempty"`
}

type WebhookReceiverObservation struct {
}

type WebhookReceiverParameters struct {

	// +kubebuilder:validation:Optional
	AadAuth []AadAuthParameters `json:"aadAuth,omitempty" tf:"aad_auth,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ServiceURI *string `json:"serviceUri" tf:"service_uri,omitempty"`

	// +kubebuilder:validation:Optional
	UseCommonAlertSchema *bool `json:"useCommonAlertSchema,omitempty" tf:"use_common_alert_schema,omitempty"`
}

// MonitorActionGroupSpec defines the desired state of MonitorActionGroup
type MonitorActionGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorActionGroupParameters `json:"forProvider"`
}

// MonitorActionGroupStatus defines the observed state of MonitorActionGroup.
type MonitorActionGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorActionGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorActionGroup is the Schema for the MonitorActionGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorActionGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorActionGroupSpec   `json:"spec"`
	Status            MonitorActionGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorActionGroupList contains a list of MonitorActionGroups
type MonitorActionGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorActionGroup `json:"items"`
}

// Repository type metadata.
var (
	MonitorActionGroupKind             = "MonitorActionGroup"
	MonitorActionGroupGroupKind        = schema.GroupKind{Group: Group, Kind: MonitorActionGroupKind}.String()
	MonitorActionGroupKindAPIVersion   = MonitorActionGroupKind + "." + GroupVersion.String()
	MonitorActionGroupGroupVersionKind = GroupVersion.WithKind(MonitorActionGroupKind)
)

func init() {
	SchemeBuilder.Register(&MonitorActionGroup{}, &MonitorActionGroupList{})
}

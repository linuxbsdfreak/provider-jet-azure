/*
Copyright 2022 The Crossplane Authors.

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
	"github.com/pkg/errors"

	"github.com/crossplane/terrajet/pkg/resource"
	"github.com/crossplane/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Namespace
func (mg *Namespace) GetTerraformResourceType() string {
	return "azurerm_servicebus_namespace"
}

// GetConnectionDetailsMapping for this Namespace
func (tr *Namespace) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"default_primary_connection_string": "status.atProvider.defaultPrimaryConnectionString", "default_primary_key": "status.atProvider.defaultPrimaryKey", "default_secondary_connection_string": "status.atProvider.defaultSecondaryConnectionString", "default_secondary_key": "status.atProvider.defaultSecondaryKey"}
}

// GetObservation of this Namespace
func (tr *Namespace) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Namespace
func (tr *Namespace) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Namespace
func (tr *Namespace) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Namespace
func (tr *Namespace) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Namespace
func (tr *Namespace) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Namespace using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Namespace) LateInitialize(attrs []byte) (bool, error) {
	params := &NamespaceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Namespace) GetTerraformSchemaVersion() int {
	return 1
}

// GetTerraformResourceType returns Terraform resource type for this NamespaceAuthorizationRule
func (mg *NamespaceAuthorizationRule) GetTerraformResourceType() string {
	return "azurerm_servicebus_namespace_authorization_rule"
}

// GetConnectionDetailsMapping for this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"primary_connection_string": "status.atProvider.primaryConnectionString", "primary_connection_string_alias": "status.atProvider.primaryConnectionStringAlias", "primary_key": "status.atProvider.primaryKey", "secondary_connection_string": "status.atProvider.secondaryConnectionString", "secondary_connection_string_alias": "status.atProvider.secondaryConnectionStringAlias", "secondary_key": "status.atProvider.secondaryKey"}
}

// GetObservation of this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NamespaceAuthorizationRule
func (tr *NamespaceAuthorizationRule) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this NamespaceAuthorizationRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NamespaceAuthorizationRule) LateInitialize(attrs []byte) (bool, error) {
	params := &NamespaceAuthorizationRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NamespaceAuthorizationRule) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this NamespaceDisasterRecoveryConfig
func (mg *NamespaceDisasterRecoveryConfig) GetTerraformResourceType() string {
	return "azurerm_servicebus_namespace_disaster_recovery_config"
}

// GetConnectionDetailsMapping for this NamespaceDisasterRecoveryConfig
func (tr *NamespaceDisasterRecoveryConfig) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"default_primary_key": "status.atProvider.defaultPrimaryKey", "default_secondary_key": "status.atProvider.defaultSecondaryKey", "primary_connection_string_alias": "status.atProvider.primaryConnectionStringAlias", "secondary_connection_string_alias": "status.atProvider.secondaryConnectionStringAlias"}
}

// GetObservation of this NamespaceDisasterRecoveryConfig
func (tr *NamespaceDisasterRecoveryConfig) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NamespaceDisasterRecoveryConfig
func (tr *NamespaceDisasterRecoveryConfig) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NamespaceDisasterRecoveryConfig
func (tr *NamespaceDisasterRecoveryConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NamespaceDisasterRecoveryConfig
func (tr *NamespaceDisasterRecoveryConfig) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NamespaceDisasterRecoveryConfig
func (tr *NamespaceDisasterRecoveryConfig) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this NamespaceDisasterRecoveryConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NamespaceDisasterRecoveryConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &NamespaceDisasterRecoveryConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NamespaceDisasterRecoveryConfig) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this NamespaceNetworkRuleSet
func (mg *NamespaceNetworkRuleSet) GetTerraformResourceType() string {
	return "azurerm_servicebus_namespace_network_rule_set"
}

// GetConnectionDetailsMapping for this NamespaceNetworkRuleSet
func (tr *NamespaceNetworkRuleSet) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this NamespaceNetworkRuleSet
func (tr *NamespaceNetworkRuleSet) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this NamespaceNetworkRuleSet
func (tr *NamespaceNetworkRuleSet) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this NamespaceNetworkRuleSet
func (tr *NamespaceNetworkRuleSet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this NamespaceNetworkRuleSet
func (tr *NamespaceNetworkRuleSet) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this NamespaceNetworkRuleSet
func (tr *NamespaceNetworkRuleSet) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this NamespaceNetworkRuleSet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *NamespaceNetworkRuleSet) LateInitialize(attrs []byte) (bool, error) {
	params := &NamespaceNetworkRuleSetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *NamespaceNetworkRuleSet) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Queue
func (mg *Queue) GetTerraformResourceType() string {
	return "azurerm_servicebus_queue"
}

// GetConnectionDetailsMapping for this Queue
func (tr *Queue) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Queue
func (tr *Queue) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Queue
func (tr *Queue) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Queue
func (tr *Queue) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Queue
func (tr *Queue) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Queue
func (tr *Queue) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Queue using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Queue) LateInitialize(attrs []byte) (bool, error) {
	params := &QueueParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Queue) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this QueueAuthorizationRule
func (mg *QueueAuthorizationRule) GetTerraformResourceType() string {
	return "azurerm_servicebus_queue_authorization_rule"
}

// GetConnectionDetailsMapping for this QueueAuthorizationRule
func (tr *QueueAuthorizationRule) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"primary_connection_string": "status.atProvider.primaryConnectionString", "primary_connection_string_alias": "status.atProvider.primaryConnectionStringAlias", "primary_key": "status.atProvider.primaryKey", "secondary_connection_string": "status.atProvider.secondaryConnectionString", "secondary_connection_string_alias": "status.atProvider.secondaryConnectionStringAlias", "secondary_key": "status.atProvider.secondaryKey"}
}

// GetObservation of this QueueAuthorizationRule
func (tr *QueueAuthorizationRule) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this QueueAuthorizationRule
func (tr *QueueAuthorizationRule) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this QueueAuthorizationRule
func (tr *QueueAuthorizationRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this QueueAuthorizationRule
func (tr *QueueAuthorizationRule) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this QueueAuthorizationRule
func (tr *QueueAuthorizationRule) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this QueueAuthorizationRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *QueueAuthorizationRule) LateInitialize(attrs []byte) (bool, error) {
	params := &QueueAuthorizationRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *QueueAuthorizationRule) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Subscription
func (mg *Subscription) GetTerraformResourceType() string {
	return "azurerm_servicebus_subscription"
}

// GetConnectionDetailsMapping for this Subscription
func (tr *Subscription) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Subscription
func (tr *Subscription) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Subscription
func (tr *Subscription) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Subscription
func (tr *Subscription) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Subscription
func (tr *Subscription) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Subscription
func (tr *Subscription) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Subscription using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Subscription) LateInitialize(attrs []byte) (bool, error) {
	params := &SubscriptionParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Subscription) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SubscriptionRule
func (mg *SubscriptionRule) GetTerraformResourceType() string {
	return "azurerm_servicebus_subscription_rule"
}

// GetConnectionDetailsMapping for this SubscriptionRule
func (tr *SubscriptionRule) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SubscriptionRule
func (tr *SubscriptionRule) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SubscriptionRule
func (tr *SubscriptionRule) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SubscriptionRule
func (tr *SubscriptionRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SubscriptionRule
func (tr *SubscriptionRule) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SubscriptionRule
func (tr *SubscriptionRule) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SubscriptionRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SubscriptionRule) LateInitialize(attrs []byte) (bool, error) {
	params := &SubscriptionRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SubscriptionRule) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Topic
func (mg *Topic) GetTerraformResourceType() string {
	return "azurerm_servicebus_topic"
}

// GetConnectionDetailsMapping for this Topic
func (tr *Topic) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Topic
func (tr *Topic) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Topic
func (tr *Topic) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Topic
func (tr *Topic) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Topic
func (tr *Topic) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Topic
func (tr *Topic) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Topic using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Topic) LateInitialize(attrs []byte) (bool, error) {
	params := &TopicParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Topic) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this TopicAuthorizationRule
func (mg *TopicAuthorizationRule) GetTerraformResourceType() string {
	return "azurerm_servicebus_topic_authorization_rule"
}

// GetConnectionDetailsMapping for this TopicAuthorizationRule
func (tr *TopicAuthorizationRule) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"primary_connection_string": "status.atProvider.primaryConnectionString", "primary_connection_string_alias": "status.atProvider.primaryConnectionStringAlias", "primary_key": "status.atProvider.primaryKey", "secondary_connection_string": "status.atProvider.secondaryConnectionString", "secondary_connection_string_alias": "status.atProvider.secondaryConnectionStringAlias", "secondary_key": "status.atProvider.secondaryKey"}
}

// GetObservation of this TopicAuthorizationRule
func (tr *TopicAuthorizationRule) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this TopicAuthorizationRule
func (tr *TopicAuthorizationRule) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this TopicAuthorizationRule
func (tr *TopicAuthorizationRule) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this TopicAuthorizationRule
func (tr *TopicAuthorizationRule) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this TopicAuthorizationRule
func (tr *TopicAuthorizationRule) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this TopicAuthorizationRule using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *TopicAuthorizationRule) LateInitialize(attrs []byte) (bool, error) {
	params := &TopicAuthorizationRuleParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *TopicAuthorizationRule) GetTerraformSchemaVersion() int {
	return 0
}
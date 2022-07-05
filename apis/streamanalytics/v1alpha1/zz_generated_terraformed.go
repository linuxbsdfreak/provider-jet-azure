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

// GetTerraformResourceType returns Terraform resource type for this FunctionJavascriptUDF
func (mg *FunctionJavascriptUDF) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_function_javascript_udf"
}

// GetConnectionDetailsMapping for this FunctionJavascriptUDF
func (tr *FunctionJavascriptUDF) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this FunctionJavascriptUDF
func (tr *FunctionJavascriptUDF) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this FunctionJavascriptUDF
func (tr *FunctionJavascriptUDF) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this FunctionJavascriptUDF
func (tr *FunctionJavascriptUDF) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this FunctionJavascriptUDF
func (tr *FunctionJavascriptUDF) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this FunctionJavascriptUDF
func (tr *FunctionJavascriptUDF) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this FunctionJavascriptUDF using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *FunctionJavascriptUDF) LateInitialize(attrs []byte) (bool, error) {
	params := &FunctionJavascriptUDFParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *FunctionJavascriptUDF) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Job
func (mg *Job) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_job"
}

// GetConnectionDetailsMapping for this Job
func (tr *Job) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Job
func (tr *Job) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Job
func (tr *Job) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Job
func (tr *Job) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Job
func (tr *Job) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Job
func (tr *Job) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Job using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Job) LateInitialize(attrs []byte) (bool, error) {
	params := &JobParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Job) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this OutputBlob
func (mg *OutputBlob) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_output_blob"
}

// GetConnectionDetailsMapping for this OutputBlob
func (tr *OutputBlob) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"storage_account_key": "spec.forProvider.storageAccountKeySecretRef"}
}

// GetObservation of this OutputBlob
func (tr *OutputBlob) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this OutputBlob
func (tr *OutputBlob) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this OutputBlob
func (tr *OutputBlob) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this OutputBlob
func (tr *OutputBlob) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this OutputBlob
func (tr *OutputBlob) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this OutputBlob using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *OutputBlob) LateInitialize(attrs []byte) (bool, error) {
	params := &OutputBlobParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *OutputBlob) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this OutputEventHub
func (mg *OutputEventHub) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_output_eventhub"
}

// GetConnectionDetailsMapping for this OutputEventHub
func (tr *OutputEventHub) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"shared_access_policy_key": "spec.forProvider.sharedAccessPolicyKeySecretRef"}
}

// GetObservation of this OutputEventHub
func (tr *OutputEventHub) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this OutputEventHub
func (tr *OutputEventHub) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this OutputEventHub
func (tr *OutputEventHub) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this OutputEventHub
func (tr *OutputEventHub) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this OutputEventHub
func (tr *OutputEventHub) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this OutputEventHub using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *OutputEventHub) LateInitialize(attrs []byte) (bool, error) {
	params := &OutputEventHubParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *OutputEventHub) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this OutputMSSQL
func (mg *OutputMSSQL) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_output_mssql"
}

// GetConnectionDetailsMapping for this OutputMSSQL
func (tr *OutputMSSQL) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"password": "spec.forProvider.passwordSecretRef"}
}

// GetObservation of this OutputMSSQL
func (tr *OutputMSSQL) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this OutputMSSQL
func (tr *OutputMSSQL) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this OutputMSSQL
func (tr *OutputMSSQL) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this OutputMSSQL
func (tr *OutputMSSQL) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this OutputMSSQL
func (tr *OutputMSSQL) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this OutputMSSQL using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *OutputMSSQL) LateInitialize(attrs []byte) (bool, error) {
	params := &OutputMSSQLParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *OutputMSSQL) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this OutputServiceBusQueue
func (mg *OutputServiceBusQueue) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_output_servicebus_queue"
}

// GetConnectionDetailsMapping for this OutputServiceBusQueue
func (tr *OutputServiceBusQueue) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"shared_access_policy_key": "spec.forProvider.sharedAccessPolicyKeySecretRef"}
}

// GetObservation of this OutputServiceBusQueue
func (tr *OutputServiceBusQueue) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this OutputServiceBusQueue
func (tr *OutputServiceBusQueue) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this OutputServiceBusQueue
func (tr *OutputServiceBusQueue) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this OutputServiceBusQueue
func (tr *OutputServiceBusQueue) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this OutputServiceBusQueue
func (tr *OutputServiceBusQueue) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this OutputServiceBusQueue using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *OutputServiceBusQueue) LateInitialize(attrs []byte) (bool, error) {
	params := &OutputServiceBusQueueParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *OutputServiceBusQueue) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this OutputServiceBusTopic
func (mg *OutputServiceBusTopic) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_output_servicebus_topic"
}

// GetConnectionDetailsMapping for this OutputServiceBusTopic
func (tr *OutputServiceBusTopic) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"shared_access_policy_key": "spec.forProvider.sharedAccessPolicyKeySecretRef"}
}

// GetObservation of this OutputServiceBusTopic
func (tr *OutputServiceBusTopic) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this OutputServiceBusTopic
func (tr *OutputServiceBusTopic) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this OutputServiceBusTopic
func (tr *OutputServiceBusTopic) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this OutputServiceBusTopic
func (tr *OutputServiceBusTopic) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this OutputServiceBusTopic
func (tr *OutputServiceBusTopic) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this OutputServiceBusTopic using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *OutputServiceBusTopic) LateInitialize(attrs []byte) (bool, error) {
	params := &OutputServiceBusTopicParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *OutputServiceBusTopic) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ReferenceInputBlob
func (mg *ReferenceInputBlob) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_reference_input_blob"
}

// GetConnectionDetailsMapping for this ReferenceInputBlob
func (tr *ReferenceInputBlob) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"storage_account_key": "spec.forProvider.storageAccountKeySecretRef"}
}

// GetObservation of this ReferenceInputBlob
func (tr *ReferenceInputBlob) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ReferenceInputBlob
func (tr *ReferenceInputBlob) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ReferenceInputBlob
func (tr *ReferenceInputBlob) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ReferenceInputBlob
func (tr *ReferenceInputBlob) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ReferenceInputBlob
func (tr *ReferenceInputBlob) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ReferenceInputBlob using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ReferenceInputBlob) LateInitialize(attrs []byte) (bool, error) {
	params := &ReferenceInputBlobParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ReferenceInputBlob) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this StreamInputBlob
func (mg *StreamInputBlob) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_stream_input_blob"
}

// GetConnectionDetailsMapping for this StreamInputBlob
func (tr *StreamInputBlob) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"storage_account_key": "spec.forProvider.storageAccountKeySecretRef"}
}

// GetObservation of this StreamInputBlob
func (tr *StreamInputBlob) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this StreamInputBlob
func (tr *StreamInputBlob) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this StreamInputBlob
func (tr *StreamInputBlob) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this StreamInputBlob
func (tr *StreamInputBlob) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this StreamInputBlob
func (tr *StreamInputBlob) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this StreamInputBlob using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *StreamInputBlob) LateInitialize(attrs []byte) (bool, error) {
	params := &StreamInputBlobParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *StreamInputBlob) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this StreamInputEventHub
func (mg *StreamInputEventHub) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_stream_input_eventhub"
}

// GetConnectionDetailsMapping for this StreamInputEventHub
func (tr *StreamInputEventHub) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"shared_access_policy_key": "spec.forProvider.sharedAccessPolicyKeySecretRef"}
}

// GetObservation of this StreamInputEventHub
func (tr *StreamInputEventHub) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this StreamInputEventHub
func (tr *StreamInputEventHub) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this StreamInputEventHub
func (tr *StreamInputEventHub) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this StreamInputEventHub
func (tr *StreamInputEventHub) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this StreamInputEventHub
func (tr *StreamInputEventHub) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this StreamInputEventHub using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *StreamInputEventHub) LateInitialize(attrs []byte) (bool, error) {
	params := &StreamInputEventHubParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *StreamInputEventHub) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this StreamInputIOTHub
func (mg *StreamInputIOTHub) GetTerraformResourceType() string {
	return "azurerm_stream_analytics_stream_input_iothub"
}

// GetConnectionDetailsMapping for this StreamInputIOTHub
func (tr *StreamInputIOTHub) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"shared_access_policy_key": "spec.forProvider.sharedAccessPolicyKeySecretRef"}
}

// GetObservation of this StreamInputIOTHub
func (tr *StreamInputIOTHub) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this StreamInputIOTHub
func (tr *StreamInputIOTHub) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this StreamInputIOTHub
func (tr *StreamInputIOTHub) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this StreamInputIOTHub
func (tr *StreamInputIOTHub) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this StreamInputIOTHub
func (tr *StreamInputIOTHub) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this StreamInputIOTHub using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *StreamInputIOTHub) LateInitialize(attrs []byte) (bool, error) {
	params := &StreamInputIOTHubParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *StreamInputIOTHub) GetTerraformSchemaVersion() int {
	return 0
}

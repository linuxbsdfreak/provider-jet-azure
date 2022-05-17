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

// GetTerraformResourceType returns Terraform resource type for this BackupContainerStorageAccount
func (mg *BackupContainerStorageAccount) GetTerraformResourceType() string {
	return "azurerm_backup_container_storage_account"
}

// GetConnectionDetailsMapping for this BackupContainerStorageAccount
func (tr *BackupContainerStorageAccount) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BackupContainerStorageAccount
func (tr *BackupContainerStorageAccount) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BackupContainerStorageAccount
func (tr *BackupContainerStorageAccount) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BackupContainerStorageAccount
func (tr *BackupContainerStorageAccount) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BackupContainerStorageAccount
func (tr *BackupContainerStorageAccount) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BackupContainerStorageAccount
func (tr *BackupContainerStorageAccount) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BackupContainerStorageAccount using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BackupContainerStorageAccount) LateInitialize(attrs []byte) (bool, error) {
	params := &BackupContainerStorageAccountParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BackupContainerStorageAccount) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BackupPolicyFileShare
func (mg *BackupPolicyFileShare) GetTerraformResourceType() string {
	return "azurerm_backup_policy_file_share"
}

// GetConnectionDetailsMapping for this BackupPolicyFileShare
func (tr *BackupPolicyFileShare) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BackupPolicyFileShare
func (tr *BackupPolicyFileShare) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BackupPolicyFileShare
func (tr *BackupPolicyFileShare) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BackupPolicyFileShare
func (tr *BackupPolicyFileShare) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BackupPolicyFileShare
func (tr *BackupPolicyFileShare) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BackupPolicyFileShare
func (tr *BackupPolicyFileShare) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BackupPolicyFileShare using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BackupPolicyFileShare) LateInitialize(attrs []byte) (bool, error) {
	params := &BackupPolicyFileShareParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BackupPolicyFileShare) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BackupPolicyVM
func (mg *BackupPolicyVM) GetTerraformResourceType() string {
	return "azurerm_backup_policy_vm"
}

// GetConnectionDetailsMapping for this BackupPolicyVM
func (tr *BackupPolicyVM) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BackupPolicyVM
func (tr *BackupPolicyVM) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BackupPolicyVM
func (tr *BackupPolicyVM) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BackupPolicyVM
func (tr *BackupPolicyVM) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BackupPolicyVM
func (tr *BackupPolicyVM) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BackupPolicyVM
func (tr *BackupPolicyVM) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BackupPolicyVM using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BackupPolicyVM) LateInitialize(attrs []byte) (bool, error) {
	params := &BackupPolicyVMParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BackupPolicyVM) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BackupProtectedFileShare
func (mg *BackupProtectedFileShare) GetTerraformResourceType() string {
	return "azurerm_backup_protected_file_share"
}

// GetConnectionDetailsMapping for this BackupProtectedFileShare
func (tr *BackupProtectedFileShare) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BackupProtectedFileShare
func (tr *BackupProtectedFileShare) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BackupProtectedFileShare
func (tr *BackupProtectedFileShare) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BackupProtectedFileShare
func (tr *BackupProtectedFileShare) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BackupProtectedFileShare
func (tr *BackupProtectedFileShare) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BackupProtectedFileShare
func (tr *BackupProtectedFileShare) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BackupProtectedFileShare using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BackupProtectedFileShare) LateInitialize(attrs []byte) (bool, error) {
	params := &BackupProtectedFileShareParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BackupProtectedFileShare) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BackupProtectedVM
func (mg *BackupProtectedVM) GetTerraformResourceType() string {
	return "azurerm_backup_protected_vm"
}

// GetConnectionDetailsMapping for this BackupProtectedVM
func (tr *BackupProtectedVM) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BackupProtectedVM
func (tr *BackupProtectedVM) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BackupProtectedVM
func (tr *BackupProtectedVM) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BackupProtectedVM
func (tr *BackupProtectedVM) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BackupProtectedVM
func (tr *BackupProtectedVM) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BackupProtectedVM
func (tr *BackupProtectedVM) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BackupProtectedVM using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BackupProtectedVM) LateInitialize(attrs []byte) (bool, error) {
	params := &BackupProtectedVMParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BackupProtectedVM) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Vault
func (mg *Vault) GetTerraformResourceType() string {
	return "azurerm_recovery_services_vault"
}

// GetConnectionDetailsMapping for this Vault
func (tr *Vault) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Vault
func (tr *Vault) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Vault
func (tr *Vault) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Vault
func (tr *Vault) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Vault
func (tr *Vault) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Vault
func (tr *Vault) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Vault using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Vault) LateInitialize(attrs []byte) (bool, error) {
	params := &VaultParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Vault) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SiteRecoveryFabric
func (mg *SiteRecoveryFabric) GetTerraformResourceType() string {
	return "azurerm_site_recovery_fabric"
}

// GetConnectionDetailsMapping for this SiteRecoveryFabric
func (tr *SiteRecoveryFabric) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SiteRecoveryFabric
func (tr *SiteRecoveryFabric) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SiteRecoveryFabric
func (tr *SiteRecoveryFabric) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SiteRecoveryFabric
func (tr *SiteRecoveryFabric) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SiteRecoveryFabric
func (tr *SiteRecoveryFabric) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SiteRecoveryFabric
func (tr *SiteRecoveryFabric) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SiteRecoveryFabric using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SiteRecoveryFabric) LateInitialize(attrs []byte) (bool, error) {
	params := &SiteRecoveryFabricParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SiteRecoveryFabric) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SiteRecoveryNetworkMapping
func (mg *SiteRecoveryNetworkMapping) GetTerraformResourceType() string {
	return "azurerm_site_recovery_network_mapping"
}

// GetConnectionDetailsMapping for this SiteRecoveryNetworkMapping
func (tr *SiteRecoveryNetworkMapping) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SiteRecoveryNetworkMapping
func (tr *SiteRecoveryNetworkMapping) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SiteRecoveryNetworkMapping
func (tr *SiteRecoveryNetworkMapping) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SiteRecoveryNetworkMapping
func (tr *SiteRecoveryNetworkMapping) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SiteRecoveryNetworkMapping
func (tr *SiteRecoveryNetworkMapping) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SiteRecoveryNetworkMapping
func (tr *SiteRecoveryNetworkMapping) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SiteRecoveryNetworkMapping using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SiteRecoveryNetworkMapping) LateInitialize(attrs []byte) (bool, error) {
	params := &SiteRecoveryNetworkMappingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SiteRecoveryNetworkMapping) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SiteRecoveryProtectionContainer
func (mg *SiteRecoveryProtectionContainer) GetTerraformResourceType() string {
	return "azurerm_site_recovery_protection_container"
}

// GetConnectionDetailsMapping for this SiteRecoveryProtectionContainer
func (tr *SiteRecoveryProtectionContainer) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SiteRecoveryProtectionContainer
func (tr *SiteRecoveryProtectionContainer) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SiteRecoveryProtectionContainer
func (tr *SiteRecoveryProtectionContainer) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SiteRecoveryProtectionContainer
func (tr *SiteRecoveryProtectionContainer) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SiteRecoveryProtectionContainer
func (tr *SiteRecoveryProtectionContainer) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SiteRecoveryProtectionContainer
func (tr *SiteRecoveryProtectionContainer) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SiteRecoveryProtectionContainer using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SiteRecoveryProtectionContainer) LateInitialize(attrs []byte) (bool, error) {
	params := &SiteRecoveryProtectionContainerParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SiteRecoveryProtectionContainer) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SiteRecoveryProtectionContainerMapping
func (mg *SiteRecoveryProtectionContainerMapping) GetTerraformResourceType() string {
	return "azurerm_site_recovery_protection_container_mapping"
}

// GetConnectionDetailsMapping for this SiteRecoveryProtectionContainerMapping
func (tr *SiteRecoveryProtectionContainerMapping) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SiteRecoveryProtectionContainerMapping
func (tr *SiteRecoveryProtectionContainerMapping) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SiteRecoveryProtectionContainerMapping
func (tr *SiteRecoveryProtectionContainerMapping) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SiteRecoveryProtectionContainerMapping
func (tr *SiteRecoveryProtectionContainerMapping) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SiteRecoveryProtectionContainerMapping
func (tr *SiteRecoveryProtectionContainerMapping) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SiteRecoveryProtectionContainerMapping
func (tr *SiteRecoveryProtectionContainerMapping) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SiteRecoveryProtectionContainerMapping using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SiteRecoveryProtectionContainerMapping) LateInitialize(attrs []byte) (bool, error) {
	params := &SiteRecoveryProtectionContainerMappingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SiteRecoveryProtectionContainerMapping) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SiteRecoveryReplicatedVM
func (mg *SiteRecoveryReplicatedVM) GetTerraformResourceType() string {
	return "azurerm_site_recovery_replicated_vm"
}

// GetConnectionDetailsMapping for this SiteRecoveryReplicatedVM
func (tr *SiteRecoveryReplicatedVM) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SiteRecoveryReplicatedVM
func (tr *SiteRecoveryReplicatedVM) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SiteRecoveryReplicatedVM
func (tr *SiteRecoveryReplicatedVM) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SiteRecoveryReplicatedVM
func (tr *SiteRecoveryReplicatedVM) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SiteRecoveryReplicatedVM
func (tr *SiteRecoveryReplicatedVM) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SiteRecoveryReplicatedVM
func (tr *SiteRecoveryReplicatedVM) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SiteRecoveryReplicatedVM using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SiteRecoveryReplicatedVM) LateInitialize(attrs []byte) (bool, error) {
	params := &SiteRecoveryReplicatedVMParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SiteRecoveryReplicatedVM) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SiteRecoveryReplicationPolicy
func (mg *SiteRecoveryReplicationPolicy) GetTerraformResourceType() string {
	return "azurerm_site_recovery_replication_policy"
}

// GetConnectionDetailsMapping for this SiteRecoveryReplicationPolicy
func (tr *SiteRecoveryReplicationPolicy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SiteRecoveryReplicationPolicy
func (tr *SiteRecoveryReplicationPolicy) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SiteRecoveryReplicationPolicy
func (tr *SiteRecoveryReplicationPolicy) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SiteRecoveryReplicationPolicy
func (tr *SiteRecoveryReplicationPolicy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SiteRecoveryReplicationPolicy
func (tr *SiteRecoveryReplicationPolicy) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SiteRecoveryReplicationPolicy
func (tr *SiteRecoveryReplicationPolicy) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SiteRecoveryReplicationPolicy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SiteRecoveryReplicationPolicy) LateInitialize(attrs []byte) (bool, error) {
	params := &SiteRecoveryReplicationPolicyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SiteRecoveryReplicationPolicy) GetTerraformSchemaVersion() int {
	return 0
}

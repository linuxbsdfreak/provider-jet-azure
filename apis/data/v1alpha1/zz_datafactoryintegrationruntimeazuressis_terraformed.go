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
	"github.com/pkg/errors"

	"github.com/crossplane-contrib/terrajet/pkg/resource"
	"github.com/crossplane-contrib/terrajet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this DataFactoryIntegrationRuntimeAzureSsis
func (mg *DataFactoryIntegrationRuntimeAzureSsis) GetTerraformResourceType() string {
	return "azurerm_data_factory_integration_runtime_azure_ssis"
}

// GetTerraformResourceIDField returns Terraform identifier field for this DataFactoryIntegrationRuntimeAzureSsis
func (tr *DataFactoryIntegrationRuntimeAzureSsis) GetTerraformResourceIDField() string {
	return "id"
}

// GetConnectionDetailsMapping for this DataFactoryIntegrationRuntimeAzureSsis
func (tr *DataFactoryIntegrationRuntimeAzureSsis) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"catalog_info[*].administrator_password": "spec.forProvider.catalogInfo[*].administratorPasswordSecretRef", "custom_setup_script[*].sas_token": "spec.forProvider.customSetupScript[*].sasTokenSecretRef", "express_custom_setup[*].command_key[*].password": "spec.forProvider.expressCustomSetup[*].commandKey[*].passwordSecretRef", "express_custom_setup[*].component[*].license": "spec.forProvider.expressCustomSetup[*].component[*].licenseSecretRef"}
}

// GetObservation of this DataFactoryIntegrationRuntimeAzureSsis
func (tr *DataFactoryIntegrationRuntimeAzureSsis) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this DataFactoryIntegrationRuntimeAzureSsis
func (tr *DataFactoryIntegrationRuntimeAzureSsis) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetParameters of this DataFactoryIntegrationRuntimeAzureSsis
func (tr *DataFactoryIntegrationRuntimeAzureSsis) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this DataFactoryIntegrationRuntimeAzureSsis
func (tr *DataFactoryIntegrationRuntimeAzureSsis) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this DataFactoryIntegrationRuntimeAzureSsis using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *DataFactoryIntegrationRuntimeAzureSsis) LateInitialize(attrs []byte) (bool, error) {
	params := &DataFactoryIntegrationRuntimeAzureSsisParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *DataFactoryIntegrationRuntimeAzureSsis) GetTerraformSchemaVersion() int {
	return 0
}

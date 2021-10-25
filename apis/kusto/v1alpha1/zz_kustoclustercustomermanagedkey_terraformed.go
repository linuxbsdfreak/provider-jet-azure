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

// GetTerraformResourceType returns Terraform resource type for this KustoClusterCustomerManagedKey
func (mg *KustoClusterCustomerManagedKey) GetTerraformResourceType() string {
	return "azurerm_kusto_cluster_customer_managed_key"
}

// GetTerraformResourceIDField returns Terraform identifier field for this KustoClusterCustomerManagedKey
func (tr *KustoClusterCustomerManagedKey) GetTerraformResourceIDField() string {
	return "id"
}

// GetConnectionDetailsMapping for this KustoClusterCustomerManagedKey
func (tr *KustoClusterCustomerManagedKey) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this KustoClusterCustomerManagedKey
func (tr *KustoClusterCustomerManagedKey) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this KustoClusterCustomerManagedKey
func (tr *KustoClusterCustomerManagedKey) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetParameters of this KustoClusterCustomerManagedKey
func (tr *KustoClusterCustomerManagedKey) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this KustoClusterCustomerManagedKey
func (tr *KustoClusterCustomerManagedKey) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this KustoClusterCustomerManagedKey using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *KustoClusterCustomerManagedKey) LateInitialize(attrs []byte) (bool, error) {
	params := &KustoClusterCustomerManagedKeyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *KustoClusterCustomerManagedKey) GetTerraformSchemaVersion() int {
	return 0
}

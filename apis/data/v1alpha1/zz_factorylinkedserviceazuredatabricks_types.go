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

type FactoryLinkedServiceAzureDatabricksKeyVaultPasswordObservation struct {
}

type FactoryLinkedServiceAzureDatabricksKeyVaultPasswordParameters struct {

	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// +kubebuilder:validation:Required
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type FactoryLinkedServiceAzureDatabricksObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FactoryLinkedServiceAzureDatabricksParameters struct {

	// +kubebuilder:validation:Optional
	AccessTokenSecretRef *v1.SecretKeySelector `json:"accessTokenSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	AdbDomain *string `json:"adbDomain" tf:"adb_domain,omitempty"`

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ExistingClusterID *string `json:"existingClusterId,omitempty" tf:"existing_cluster_id,omitempty"`

	// +kubebuilder:validation:Optional
	InstancePool []InstancePoolParameters `json:"instancePool,omitempty" tf:"instance_pool,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultPassword []FactoryLinkedServiceAzureDatabricksKeyVaultPasswordParameters `json:"keyVaultPassword,omitempty" tf:"key_vault_password,omitempty"`

	// +kubebuilder:validation:Optional
	MsiWorkSpaceResourceID *string `json:"msiWorkSpaceResourceId,omitempty" tf:"msi_work_space_resource_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NewClusterConfig []NewClusterConfigParameters `json:"newClusterConfig,omitempty" tf:"new_cluster_config,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

type InstancePoolObservation struct {
}

type InstancePoolParameters struct {

	// +kubebuilder:validation:Required
	ClusterVersion *string `json:"clusterVersion" tf:"cluster_version,omitempty"`

	// +kubebuilder:validation:Required
	InstancePoolID *string `json:"instancePoolId" tf:"instance_pool_id,omitempty"`

	// +kubebuilder:validation:Optional
	MaxNumberOfWorkers *int64 `json:"maxNumberOfWorkers,omitempty" tf:"max_number_of_workers,omitempty"`

	// +kubebuilder:validation:Optional
	MinNumberOfWorkers *int64 `json:"minNumberOfWorkers,omitempty" tf:"min_number_of_workers,omitempty"`
}

type NewClusterConfigObservation struct {
}

type NewClusterConfigParameters struct {

	// +kubebuilder:validation:Required
	ClusterVersion *string `json:"clusterVersion" tf:"cluster_version,omitempty"`

	// +kubebuilder:validation:Optional
	CustomTags map[string]*string `json:"customTags,omitempty" tf:"custom_tags,omitempty"`

	// +kubebuilder:validation:Optional
	DriverNodeType *string `json:"driverNodeType,omitempty" tf:"driver_node_type,omitempty"`

	// +kubebuilder:validation:Optional
	InitScripts []*string `json:"initScripts,omitempty" tf:"init_scripts,omitempty"`

	// +kubebuilder:validation:Optional
	LogDestination *string `json:"logDestination,omitempty" tf:"log_destination,omitempty"`

	// +kubebuilder:validation:Optional
	MaxNumberOfWorkers *int64 `json:"maxNumberOfWorkers,omitempty" tf:"max_number_of_workers,omitempty"`

	// +kubebuilder:validation:Optional
	MinNumberOfWorkers *int64 `json:"minNumberOfWorkers,omitempty" tf:"min_number_of_workers,omitempty"`

	// +kubebuilder:validation:Required
	NodeType *string `json:"nodeType" tf:"node_type,omitempty"`

	// +kubebuilder:validation:Optional
	SparkConfig map[string]*string `json:"sparkConfig,omitempty" tf:"spark_config,omitempty"`

	// +kubebuilder:validation:Optional
	SparkEnvironmentVariables map[string]*string `json:"sparkEnvironmentVariables,omitempty" tf:"spark_environment_variables,omitempty"`
}

// FactoryLinkedServiceAzureDatabricksSpec defines the desired state of FactoryLinkedServiceAzureDatabricks
type FactoryLinkedServiceAzureDatabricksSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FactoryLinkedServiceAzureDatabricksParameters `json:"forProvider"`
}

// FactoryLinkedServiceAzureDatabricksStatus defines the observed state of FactoryLinkedServiceAzureDatabricks.
type FactoryLinkedServiceAzureDatabricksStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FactoryLinkedServiceAzureDatabricksObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceAzureDatabricks is the Schema for the FactoryLinkedServiceAzureDatabrickss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FactoryLinkedServiceAzureDatabricks struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryLinkedServiceAzureDatabricksSpec   `json:"spec"`
	Status            FactoryLinkedServiceAzureDatabricksStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceAzureDatabricksList contains a list of FactoryLinkedServiceAzureDatabrickss
type FactoryLinkedServiceAzureDatabricksList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FactoryLinkedServiceAzureDatabricks `json:"items"`
}

// Repository type metadata.
var (
	FactoryLinkedServiceAzureDatabricks_Kind             = "FactoryLinkedServiceAzureDatabricks"
	FactoryLinkedServiceAzureDatabricks_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FactoryLinkedServiceAzureDatabricks_Kind}.String()
	FactoryLinkedServiceAzureDatabricks_KindAPIVersion   = FactoryLinkedServiceAzureDatabricks_Kind + "." + CRDGroupVersion.String()
	FactoryLinkedServiceAzureDatabricks_GroupVersionKind = CRDGroupVersion.WithKind(FactoryLinkedServiceAzureDatabricks_Kind)
)

func init() {
	SchemeBuilder.Register(&FactoryLinkedServiceAzureDatabricks{}, &FactoryLinkedServiceAzureDatabricksList{})
}
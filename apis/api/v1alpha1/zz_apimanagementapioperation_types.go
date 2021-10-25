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

type ApiManagementApiOperationObservation struct {
}

type ApiManagementApiOperationParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Required
	APIName *string `json:"apiName" tf:"api_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Required
	Method *string `json:"method" tf:"method,omitempty"`

	// +kubebuilder:validation:Required
	OperationID *string `json:"operationId" tf:"operation_id,omitempty"`

	// +kubebuilder:validation:Optional
	Request []RequestParameters `json:"request,omitempty" tf:"request,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Response []ResponseParameters `json:"response,omitempty" tf:"response,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateParameter []TemplateParameterParameters `json:"templateParameter,omitempty" tf:"template_parameter,omitempty"`

	// +kubebuilder:validation:Required
	URLTemplate *string `json:"urlTemplate" tf:"url_template,omitempty"`
}

type FormParameterObservation struct {
}

type FormParameterParameters struct {

	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Required *bool `json:"required" tf:"required,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type HeaderObservation struct {
}

type HeaderParameters struct {

	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Required *bool `json:"required" tf:"required,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type QueryParameterObservation struct {
}

type QueryParameterParameters struct {

	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Required *bool `json:"required" tf:"required,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type RepresentationFormParameterObservation struct {
}

type RepresentationFormParameterParameters struct {

	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Required *bool `json:"required" tf:"required,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type RepresentationObservation struct {
}

type RepresentationParameters struct {

	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// +kubebuilder:validation:Optional
	FormParameter []FormParameterParameters `json:"formParameter,omitempty" tf:"form_parameter,omitempty"`

	// +kubebuilder:validation:Optional
	Sample *string `json:"sample,omitempty" tf:"sample,omitempty"`

	// +kubebuilder:validation:Optional
	SchemaID *string `json:"schemaId,omitempty" tf:"schema_id,omitempty"`

	// +kubebuilder:validation:Optional
	TypeName *string `json:"typeName,omitempty" tf:"type_name,omitempty"`
}

type RequestObservation struct {
}

type RequestParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Header []HeaderParameters `json:"header,omitempty" tf:"header,omitempty"`

	// +kubebuilder:validation:Optional
	QueryParameter []QueryParameterParameters `json:"queryParameter,omitempty" tf:"query_parameter,omitempty"`

	// +kubebuilder:validation:Optional
	Representation []RepresentationParameters `json:"representation,omitempty" tf:"representation,omitempty"`
}

type ResponseHeaderObservation struct {
}

type ResponseHeaderParameters struct {

	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Required *bool `json:"required" tf:"required,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

type ResponseObservation struct {
}

type ResponseParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Header []ResponseHeaderParameters `json:"header,omitempty" tf:"header,omitempty"`

	// +kubebuilder:validation:Optional
	Representation []ResponseRepresentationParameters `json:"representation,omitempty" tf:"representation,omitempty"`

	// +kubebuilder:validation:Required
	StatusCode *int64 `json:"statusCode" tf:"status_code,omitempty"`
}

type ResponseRepresentationObservation struct {
}

type ResponseRepresentationParameters struct {

	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`

	// +kubebuilder:validation:Optional
	FormParameter []RepresentationFormParameterParameters `json:"formParameter,omitempty" tf:"form_parameter,omitempty"`

	// +kubebuilder:validation:Optional
	Sample *string `json:"sample,omitempty" tf:"sample,omitempty"`

	// +kubebuilder:validation:Optional
	SchemaID *string `json:"schemaId,omitempty" tf:"schema_id,omitempty"`

	// +kubebuilder:validation:Optional
	TypeName *string `json:"typeName,omitempty" tf:"type_name,omitempty"`
}

type TemplateParameterObservation struct {
}

type TemplateParameterParameters struct {

	// +kubebuilder:validation:Optional
	DefaultValue *string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Required *bool `json:"required" tf:"required,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values,omitempty"`
}

// ApiManagementApiOperationSpec defines the desired state of ApiManagementApiOperation
type ApiManagementApiOperationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApiManagementApiOperationParameters `json:"forProvider"`
}

// ApiManagementApiOperationStatus defines the observed state of ApiManagementApiOperation.
type ApiManagementApiOperationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApiManagementApiOperationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementApiOperation is the Schema for the ApiManagementApiOperations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementApiOperation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementApiOperationSpec   `json:"spec"`
	Status            ApiManagementApiOperationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementApiOperationList contains a list of ApiManagementApiOperations
type ApiManagementApiOperationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementApiOperation `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementApiOperationKind             = "ApiManagementApiOperation"
	ApiManagementApiOperationGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementApiOperationKind}.String()
	ApiManagementApiOperationKindAPIVersion   = ApiManagementApiOperationKind + "." + GroupVersion.String()
	ApiManagementApiOperationGroupVersionKind = GroupVersion.WithKind(ApiManagementApiOperationKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementApiOperation{}, &ApiManagementApiOperationList{})
}

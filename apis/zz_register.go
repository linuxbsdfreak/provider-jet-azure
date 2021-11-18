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

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/crossplane-contrib/provider-tf-azure/apis/authorization/v1alpha1"
	v1alpha1azure "github.com/crossplane-contrib/provider-tf-azure/apis/azure/v1alpha1"
	v1alpha1containerservice "github.com/crossplane-contrib/provider-tf-azure/apis/containerservice/v1alpha1"
	v1alpha1cosmosdb "github.com/crossplane-contrib/provider-tf-azure/apis/cosmosdb/v1alpha1"
	v1alpha1network "github.com/crossplane-contrib/provider-tf-azure/apis/network/v1alpha1"
	v1alpha1postgresql "github.com/crossplane-contrib/provider-tf-azure/apis/postgresql/v1alpha1"
	v1alpha1resources "github.com/crossplane-contrib/provider-tf-azure/apis/resources/v1alpha1"
	v1alpha1sql "github.com/crossplane-contrib/provider-tf-azure/apis/sql/v1alpha1"
	v1alpha1storage "github.com/crossplane-contrib/provider-tf-azure/apis/storage/v1alpha1"
	v1alpha1apis "github.com/crossplane-contrib/provider-tf-azure/apis/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1azure.SchemeBuilder.AddToScheme,
		v1alpha1containerservice.SchemeBuilder.AddToScheme,
		v1alpha1cosmosdb.SchemeBuilder.AddToScheme,
		v1alpha1network.SchemeBuilder.AddToScheme,
		v1alpha1postgresql.SchemeBuilder.AddToScheme,
		v1alpha1resources.SchemeBuilder.AddToScheme,
		v1alpha1sql.SchemeBuilder.AddToScheme,
		v1alpha1storage.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}

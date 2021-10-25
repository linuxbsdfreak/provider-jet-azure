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

type AppServiceEnvironmentV3ClusterSettingObservation struct {
}

type AppServiceEnvironmentV3ClusterSettingParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type AppServiceEnvironmentV3Observation struct {
	DNSSuffix *string `json:"dnsSuffix,omitempty" tf:"dns_suffix,omitempty"`

	ExternalInboundIPAddresses []*string `json:"externalInboundIpAddresses,omitempty" tf:"external_inbound_ip_addresses,omitempty"`

	IPSslAddressCount *int64 `json:"ipSslAddressCount,omitempty" tf:"ip_ssl_address_count,omitempty"`

	InboundNetworkDependencies []InboundNetworkDependenciesObservation `json:"inboundNetworkDependencies,omitempty" tf:"inbound_network_dependencies,omitempty"`

	InternalInboundIPAddresses []*string `json:"internalInboundIpAddresses,omitempty" tf:"internal_inbound_ip_addresses,omitempty"`

	LinuxOutboundIPAddresses []*string `json:"linuxOutboundIpAddresses,omitempty" tf:"linux_outbound_ip_addresses,omitempty"`

	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	PricingTier *string `json:"pricingTier,omitempty" tf:"pricing_tier,omitempty"`

	WindowsOutboundIPAddresses []*string `json:"windowsOutboundIpAddresses,omitempty" tf:"windows_outbound_ip_addresses,omitempty"`
}

type AppServiceEnvironmentV3Parameters struct {

	// +kubebuilder:validation:Optional
	AllowNewPrivateEndpointConnections *bool `json:"allowNewPrivateEndpointConnections,omitempty" tf:"allow_new_private_endpoint_connections,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterSetting []AppServiceEnvironmentV3ClusterSettingParameters `json:"clusterSetting,omitempty" tf:"cluster_setting,omitempty"`

	// +kubebuilder:validation:Optional
	DedicatedHostCount *int64 `json:"dedicatedHostCount,omitempty" tf:"dedicated_host_count,omitempty"`

	// +kubebuilder:validation:Optional
	InternalLoadBalancingMode *string `json:"internalLoadBalancingMode,omitempty" tf:"internal_load_balancing_mode,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type InboundNetworkDependenciesObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	IPAddresses []*string `json:"ipAddresses,omitempty" tf:"ip_addresses,omitempty"`

	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`
}

type InboundNetworkDependenciesParameters struct {
}

// AppServiceEnvironmentV3Spec defines the desired state of AppServiceEnvironmentV3
type AppServiceEnvironmentV3Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppServiceEnvironmentV3Parameters `json:"forProvider"`
}

// AppServiceEnvironmentV3Status defines the observed state of AppServiceEnvironmentV3.
type AppServiceEnvironmentV3Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppServiceEnvironmentV3Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppServiceEnvironmentV3 is the Schema for the AppServiceEnvironmentV3s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AppServiceEnvironmentV3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceEnvironmentV3Spec   `json:"spec"`
	Status            AppServiceEnvironmentV3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppServiceEnvironmentV3List contains a list of AppServiceEnvironmentV3s
type AppServiceEnvironmentV3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppServiceEnvironmentV3 `json:"items"`
}

// Repository type metadata.
var (
	AppServiceEnvironmentV3Kind             = "AppServiceEnvironmentV3"
	AppServiceEnvironmentV3GroupKind        = schema.GroupKind{Group: Group, Kind: AppServiceEnvironmentV3Kind}.String()
	AppServiceEnvironmentV3KindAPIVersion   = AppServiceEnvironmentV3Kind + "." + GroupVersion.String()
	AppServiceEnvironmentV3GroupVersionKind = GroupVersion.WithKind(AppServiceEnvironmentV3Kind)
)

func init() {
	SchemeBuilder.Register(&AppServiceEnvironmentV3{}, &AppServiceEnvironmentV3List{})
}

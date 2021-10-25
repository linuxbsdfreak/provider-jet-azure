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

type DestinationObservation struct {
}

type DestinationParameters struct {

	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type EndpointObservation struct {
}

type EndpointParameters struct {

	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// +kubebuilder:validation:Optional
	CoverageLevel *string `json:"coverageLevel,omitempty" tf:"coverage_level,omitempty"`

	// +kubebuilder:validation:Optional
	ExcludedIPAddresses []*string `json:"excludedIpAddresses,omitempty" tf:"excluded_ip_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Optional
	IncludedIPAddresses []*string `json:"includedIpAddresses,omitempty" tf:"included_ip_addresses,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// +kubebuilder:validation:Optional
	TargetResourceType *string `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// +kubebuilder:validation:Optional
	Item []ItemParameters `json:"item,omitempty" tf:"item,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HTTPConfigurationObservation struct {
}

type HTTPConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PreferHTTPS *bool `json:"preferHttps,omitempty" tf:"prefer_https,omitempty"`

	// +kubebuilder:validation:Optional
	RequestHeader []RequestHeaderParameters `json:"requestHeader,omitempty" tf:"request_header,omitempty"`

	// +kubebuilder:validation:Optional
	ValidStatusCodeRanges []*string `json:"validStatusCodeRanges,omitempty" tf:"valid_status_code_ranges,omitempty"`
}

type IcmpConfigurationObservation struct {
}

type IcmpConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	TraceRouteEnabled *bool `json:"traceRouteEnabled,omitempty" tf:"trace_route_enabled,omitempty"`
}

type ItemObservation struct {
}

type ItemParameters struct {

	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type NetworkConnectionMonitorObservation struct {
}

type NetworkConnectionMonitorParameters struct {

	// +kubebuilder:validation:Optional
	AutoStart *bool `json:"autoStart,omitempty" tf:"auto_start,omitempty"`

	// +kubebuilder:validation:Optional
	Destination []DestinationParameters `json:"destination,omitempty" tf:"destination,omitempty"`

	// +kubebuilder:validation:Required
	Endpoint []EndpointParameters `json:"endpoint" tf:"endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NetworkWatcherID *string `json:"networkWatcherId" tf:"network_watcher_id,omitempty"`

	// +kubebuilder:validation:Optional
	Notes *string `json:"notes,omitempty" tf:"notes,omitempty"`

	// +kubebuilder:validation:Optional
	OutputWorkspaceResourceIds []*string `json:"outputWorkspaceResourceIds,omitempty" tf:"output_workspace_resource_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Source []SourceParameters `json:"source,omitempty" tf:"source,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TestConfiguration []TestConfigurationParameters `json:"testConfiguration" tf:"test_configuration,omitempty"`

	// +kubebuilder:validation:Required
	TestGroup []TestGroupParameters `json:"testGroup" tf:"test_group,omitempty"`
}

type RequestHeaderObservation struct {
}

type RequestHeaderParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type SuccessThresholdObservation struct {
}

type SuccessThresholdParameters struct {

	// +kubebuilder:validation:Optional
	ChecksFailedPercent *int64 `json:"checksFailedPercent,omitempty" tf:"checks_failed_percent,omitempty"`

	// +kubebuilder:validation:Optional
	RoundTripTimeMs *float64 `json:"roundTripTimeMs,omitempty" tf:"round_trip_time_ms,omitempty"`
}

type TCPConfigurationObservation struct {
}

type TCPConfigurationParameters struct {

	// +kubebuilder:validation:Required
	Port *int64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	TraceRouteEnabled *bool `json:"traceRouteEnabled,omitempty" tf:"trace_route_enabled,omitempty"`
}

type TestConfigurationObservation struct {
}

type TestConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	HTTPConfiguration []HTTPConfigurationParameters `json:"httpConfiguration,omitempty" tf:"http_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	IcmpConfiguration []IcmpConfigurationParameters `json:"icmpConfiguration,omitempty" tf:"icmp_configuration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PreferredIPVersion *string `json:"preferredIpVersion,omitempty" tf:"preferred_ip_version,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	SuccessThreshold []SuccessThresholdParameters `json:"successThreshold,omitempty" tf:"success_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	TCPConfiguration []TCPConfigurationParameters `json:"tcpConfiguration,omitempty" tf:"tcp_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	TestFrequencyInSeconds *int64 `json:"testFrequencyInSeconds,omitempty" tf:"test_frequency_in_seconds,omitempty"`
}

type TestGroupObservation struct {
}

type TestGroupParameters struct {

	// +kubebuilder:validation:Required
	DestinationEndpoints []*string `json:"destinationEndpoints" tf:"destination_endpoints,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	SourceEndpoints []*string `json:"sourceEndpoints" tf:"source_endpoints,omitempty"`

	// +kubebuilder:validation:Required
	TestConfigurationNames []*string `json:"testConfigurationNames" tf:"test_configuration_names,omitempty"`
}

// NetworkConnectionMonitorSpec defines the desired state of NetworkConnectionMonitor
type NetworkConnectionMonitorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkConnectionMonitorParameters `json:"forProvider"`
}

// NetworkConnectionMonitorStatus defines the observed state of NetworkConnectionMonitor.
type NetworkConnectionMonitorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkConnectionMonitorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkConnectionMonitor is the Schema for the NetworkConnectionMonitors API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetworkConnectionMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkConnectionMonitorSpec   `json:"spec"`
	Status            NetworkConnectionMonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkConnectionMonitorList contains a list of NetworkConnectionMonitors
type NetworkConnectionMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConnectionMonitor `json:"items"`
}

// Repository type metadata.
var (
	NetworkConnectionMonitorKind             = "NetworkConnectionMonitor"
	NetworkConnectionMonitorGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkConnectionMonitorKind}.String()
	NetworkConnectionMonitorKindAPIVersion   = NetworkConnectionMonitorKind + "." + GroupVersion.String()
	NetworkConnectionMonitorGroupVersionKind = GroupVersion.WithKind(NetworkConnectionMonitorKind)
)

func init() {
	SchemeBuilder.Register(&NetworkConnectionMonitor{}, &NetworkConnectionMonitorList{})
}

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

type CircuitObservation struct {
	ExpressRouteID *string `json:"expressRouteId,omitempty" tf:"express_route_id,omitempty"`

	ExpressRoutePrivatePeeringID *string `json:"expressRoutePrivatePeeringId,omitempty" tf:"express_route_private_peering_id,omitempty"`

	PrimarySubnetCidr *string `json:"primarySubnetCidr,omitempty" tf:"primary_subnet_cidr,omitempty"`

	SecondarySubnetCidr *string `json:"secondarySubnetCidr,omitempty" tf:"secondary_subnet_cidr,omitempty"`
}

type CircuitParameters struct {
}

type ManagementClusterObservation struct {
	Hosts []*string `json:"hosts,omitempty" tf:"hosts,omitempty"`

	ID *int64 `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagementClusterParameters struct {

	// +kubebuilder:validation:Required
	Size *int64 `json:"size" tf:"size,omitempty"`
}

type VmwarePrivateCloudObservation struct {
	Circuit []CircuitObservation `json:"circuit,omitempty" tf:"circuit,omitempty"`

	HcxCloudManagerEndpoint *string `json:"hcxCloudManagerEndpoint,omitempty" tf:"hcx_cloud_manager_endpoint,omitempty"`

	ManagementSubnetCidr *string `json:"managementSubnetCidr,omitempty" tf:"management_subnet_cidr,omitempty"`

	NsxtCertificateThumbprint *string `json:"nsxtCertificateThumbprint,omitempty" tf:"nsxt_certificate_thumbprint,omitempty"`

	NsxtManagerEndpoint *string `json:"nsxtManagerEndpoint,omitempty" tf:"nsxt_manager_endpoint,omitempty"`

	ProvisioningSubnetCidr *string `json:"provisioningSubnetCidr,omitempty" tf:"provisioning_subnet_cidr,omitempty"`

	VcenterCertificateThumbprint *string `json:"vcenterCertificateThumbprint,omitempty" tf:"vcenter_certificate_thumbprint,omitempty"`

	VcsaEndpoint *string `json:"vcsaEndpoint,omitempty" tf:"vcsa_endpoint,omitempty"`

	VmotionSubnetCidr *string `json:"vmotionSubnetCidr,omitempty" tf:"vmotion_subnet_cidr,omitempty"`
}

type VmwarePrivateCloudParameters struct {

	// +kubebuilder:validation:Optional
	InternetConnectionEnabled *bool `json:"internetConnectionEnabled,omitempty" tf:"internet_connection_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	ManagementCluster []ManagementClusterParameters `json:"managementCluster" tf:"management_cluster,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NetworkSubnetCidr *string `json:"networkSubnetCidr" tf:"network_subnet_cidr,omitempty"`

	// +kubebuilder:validation:Optional
	NsxtPasswordSecretRef v1.SecretKeySelector `json:"nsxtPasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VcenterPasswordSecretRef v1.SecretKeySelector `json:"vcenterPasswordSecretRef,omitempty" tf:"-"`
}

// VmwarePrivateCloudSpec defines the desired state of VmwarePrivateCloud
type VmwarePrivateCloudSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VmwarePrivateCloudParameters `json:"forProvider"`
}

// VmwarePrivateCloudStatus defines the observed state of VmwarePrivateCloud.
type VmwarePrivateCloudStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VmwarePrivateCloudObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VmwarePrivateCloud is the Schema for the VmwarePrivateClouds API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VmwarePrivateCloud struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VmwarePrivateCloudSpec   `json:"spec"`
	Status            VmwarePrivateCloudStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VmwarePrivateCloudList contains a list of VmwarePrivateClouds
type VmwarePrivateCloudList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VmwarePrivateCloud `json:"items"`
}

// Repository type metadata.
var (
	VmwarePrivateCloudKind             = "VmwarePrivateCloud"
	VmwarePrivateCloudGroupKind        = schema.GroupKind{Group: Group, Kind: VmwarePrivateCloudKind}.String()
	VmwarePrivateCloudKindAPIVersion   = VmwarePrivateCloudKind + "." + GroupVersion.String()
	VmwarePrivateCloudGroupVersionKind = GroupVersion.WithKind(VmwarePrivateCloudKind)
)

func init() {
	SchemeBuilder.Register(&VmwarePrivateCloud{}, &VmwarePrivateCloudList{})
}

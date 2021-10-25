// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationObservation) DeepCopyInto(out *ConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationObservation.
func (in *ConfigurationObservation) DeepCopy() *ConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationParameters) DeepCopyInto(out *ConfigurationParameters) {
	*out = *in
	if in.AssignmentType != nil {
		in, out := &in.AssignmentType, &out.AssignmentType
		*out = new(string)
		**out = **in
	}
	if in.ContentHash != nil {
		in, out := &in.ContentHash, &out.ContentHash
		*out = new(string)
		**out = **in
	}
	if in.ContentURI != nil {
		in, out := &in.ContentURI, &out.ContentURI
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = make([]ParameterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationParameters.
func (in *ConfigurationParameters) DeepCopy() *ConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObservation) DeepCopyInto(out *IdentityObservation) {
	*out = *in
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObservation.
func (in *IdentityObservation) DeepCopy() *IdentityObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityParameters) DeepCopyInto(out *IdentityParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityParameters.
func (in *IdentityParameters) DeepCopy() *IdentityParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterObservation) DeepCopyInto(out *ParameterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterObservation.
func (in *ParameterObservation) DeepCopy() *ParameterObservation {
	if in == nil {
		return nil
	}
	out := new(ParameterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterParameters) DeepCopyInto(out *ParameterParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterParameters.
func (in *ParameterParameters) DeepCopy() *ParameterParameters {
	if in == nil {
		return nil
	}
	out := new(ParameterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAssignment) DeepCopyInto(out *PolicyAssignment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAssignment.
func (in *PolicyAssignment) DeepCopy() *PolicyAssignment {
	if in == nil {
		return nil
	}
	out := new(PolicyAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyAssignment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAssignmentList) DeepCopyInto(out *PolicyAssignmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAssignmentList.
func (in *PolicyAssignmentList) DeepCopy() *PolicyAssignmentList {
	if in == nil {
		return nil
	}
	out := new(PolicyAssignmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyAssignmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAssignmentObservation) DeepCopyInto(out *PolicyAssignmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAssignmentObservation.
func (in *PolicyAssignmentObservation) DeepCopy() *PolicyAssignmentObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyAssignmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAssignmentParameters) DeepCopyInto(out *PolicyAssignmentParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EnforcementMode != nil {
		in, out := &in.EnforcementMode, &out.EnforcementMode
		*out = new(bool)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NotScopes != nil {
		in, out := &in.NotScopes, &out.NotScopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(string)
		**out = **in
	}
	if in.PolicyDefinitionID != nil {
		in, out := &in.PolicyDefinitionID, &out.PolicyDefinitionID
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAssignmentParameters.
func (in *PolicyAssignmentParameters) DeepCopy() *PolicyAssignmentParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyAssignmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAssignmentSpec) DeepCopyInto(out *PolicyAssignmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAssignmentSpec.
func (in *PolicyAssignmentSpec) DeepCopy() *PolicyAssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyAssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyAssignmentStatus) DeepCopyInto(out *PolicyAssignmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyAssignmentStatus.
func (in *PolicyAssignmentStatus) DeepCopy() *PolicyAssignmentStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyAssignmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinition) DeepCopyInto(out *PolicyDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinition.
func (in *PolicyDefinition) DeepCopy() *PolicyDefinition {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionGroupObservation) DeepCopyInto(out *PolicyDefinitionGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionGroupObservation.
func (in *PolicyDefinitionGroupObservation) DeepCopy() *PolicyDefinitionGroupObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionGroupParameters) DeepCopyInto(out *PolicyDefinitionGroupParameters) {
	*out = *in
	if in.AdditionalMetadataResourceID != nil {
		in, out := &in.AdditionalMetadataResourceID, &out.AdditionalMetadataResourceID
		*out = new(string)
		**out = **in
	}
	if in.Category != nil {
		in, out := &in.Category, &out.Category
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionGroupParameters.
func (in *PolicyDefinitionGroupParameters) DeepCopy() *PolicyDefinitionGroupParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionList) DeepCopyInto(out *PolicyDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionList.
func (in *PolicyDefinitionList) DeepCopy() *PolicyDefinitionList {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionObservation) DeepCopyInto(out *PolicyDefinitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionObservation.
func (in *PolicyDefinitionObservation) DeepCopy() *PolicyDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionParameters) DeepCopyInto(out *PolicyDefinitionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ManagementGroupID != nil {
		in, out := &in.ManagementGroupID, &out.ManagementGroupID
		*out = new(string)
		**out = **in
	}
	if in.ManagementGroupName != nil {
		in, out := &in.ManagementGroupName, &out.ManagementGroupName
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(string)
		**out = **in
	}
	if in.PolicyRule != nil {
		in, out := &in.PolicyRule, &out.PolicyRule
		*out = new(string)
		**out = **in
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionParameters.
func (in *PolicyDefinitionParameters) DeepCopy() *PolicyDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionReferenceObservation) DeepCopyInto(out *PolicyDefinitionReferenceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionReferenceObservation.
func (in *PolicyDefinitionReferenceObservation) DeepCopy() *PolicyDefinitionReferenceObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionReferenceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionReferenceParameters) DeepCopyInto(out *PolicyDefinitionReferenceParameters) {
	*out = *in
	if in.ParameterValues != nil {
		in, out := &in.ParameterValues, &out.ParameterValues
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.PolicyDefinitionID != nil {
		in, out := &in.PolicyDefinitionID, &out.PolicyDefinitionID
		*out = new(string)
		**out = **in
	}
	if in.PolicyGroupNames != nil {
		in, out := &in.PolicyGroupNames, &out.PolicyGroupNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ReferenceID != nil {
		in, out := &in.ReferenceID, &out.ReferenceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionReferenceParameters.
func (in *PolicyDefinitionReferenceParameters) DeepCopy() *PolicyDefinitionReferenceParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionReferenceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionSpec) DeepCopyInto(out *PolicyDefinitionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionSpec.
func (in *PolicyDefinitionSpec) DeepCopy() *PolicyDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionStatus) DeepCopyInto(out *PolicyDefinitionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionStatus.
func (in *PolicyDefinitionStatus) DeepCopy() *PolicyDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRemediation) DeepCopyInto(out *PolicyRemediation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRemediation.
func (in *PolicyRemediation) DeepCopy() *PolicyRemediation {
	if in == nil {
		return nil
	}
	out := new(PolicyRemediation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyRemediation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRemediationList) DeepCopyInto(out *PolicyRemediationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyRemediation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRemediationList.
func (in *PolicyRemediationList) DeepCopy() *PolicyRemediationList {
	if in == nil {
		return nil
	}
	out := new(PolicyRemediationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyRemediationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRemediationObservation) DeepCopyInto(out *PolicyRemediationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRemediationObservation.
func (in *PolicyRemediationObservation) DeepCopy() *PolicyRemediationObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyRemediationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRemediationParameters) DeepCopyInto(out *PolicyRemediationParameters) {
	*out = *in
	if in.LocationFilters != nil {
		in, out := &in.LocationFilters, &out.LocationFilters
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PolicyAssignmentID != nil {
		in, out := &in.PolicyAssignmentID, &out.PolicyAssignmentID
		*out = new(string)
		**out = **in
	}
	if in.PolicyDefinitionReferenceID != nil {
		in, out := &in.PolicyDefinitionReferenceID, &out.PolicyDefinitionReferenceID
		*out = new(string)
		**out = **in
	}
	if in.ResourceDiscoveryMode != nil {
		in, out := &in.ResourceDiscoveryMode, &out.ResourceDiscoveryMode
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRemediationParameters.
func (in *PolicyRemediationParameters) DeepCopy() *PolicyRemediationParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyRemediationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRemediationSpec) DeepCopyInto(out *PolicyRemediationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRemediationSpec.
func (in *PolicyRemediationSpec) DeepCopy() *PolicyRemediationSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyRemediationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRemediationStatus) DeepCopyInto(out *PolicyRemediationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRemediationStatus.
func (in *PolicyRemediationStatus) DeepCopy() *PolicyRemediationStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyRemediationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySetDefinition) DeepCopyInto(out *PolicySetDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySetDefinition.
func (in *PolicySetDefinition) DeepCopy() *PolicySetDefinition {
	if in == nil {
		return nil
	}
	out := new(PolicySetDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicySetDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySetDefinitionList) DeepCopyInto(out *PolicySetDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicySetDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySetDefinitionList.
func (in *PolicySetDefinitionList) DeepCopy() *PolicySetDefinitionList {
	if in == nil {
		return nil
	}
	out := new(PolicySetDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicySetDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySetDefinitionObservation) DeepCopyInto(out *PolicySetDefinitionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySetDefinitionObservation.
func (in *PolicySetDefinitionObservation) DeepCopy() *PolicySetDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(PolicySetDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySetDefinitionParameters) DeepCopyInto(out *PolicySetDefinitionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ManagementGroupID != nil {
		in, out := &in.ManagementGroupID, &out.ManagementGroupID
		*out = new(string)
		**out = **in
	}
	if in.ManagementGroupName != nil {
		in, out := &in.ManagementGroupName, &out.ManagementGroupName
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(string)
		**out = **in
	}
	if in.PolicyDefinitionGroup != nil {
		in, out := &in.PolicyDefinitionGroup, &out.PolicyDefinitionGroup
		*out = make([]PolicyDefinitionGroupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PolicyDefinitionReference != nil {
		in, out := &in.PolicyDefinitionReference, &out.PolicyDefinitionReference
		*out = make([]PolicyDefinitionReferenceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PolicyDefinitions != nil {
		in, out := &in.PolicyDefinitions, &out.PolicyDefinitions
		*out = new(string)
		**out = **in
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySetDefinitionParameters.
func (in *PolicySetDefinitionParameters) DeepCopy() *PolicySetDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(PolicySetDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySetDefinitionSpec) DeepCopyInto(out *PolicySetDefinitionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySetDefinitionSpec.
func (in *PolicySetDefinitionSpec) DeepCopy() *PolicySetDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(PolicySetDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySetDefinitionStatus) DeepCopyInto(out *PolicySetDefinitionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySetDefinitionStatus.
func (in *PolicySetDefinitionStatus) DeepCopy() *PolicySetDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(PolicySetDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVirtualMachineConfigurationAssignment) DeepCopyInto(out *PolicyVirtualMachineConfigurationAssignment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVirtualMachineConfigurationAssignment.
func (in *PolicyVirtualMachineConfigurationAssignment) DeepCopy() *PolicyVirtualMachineConfigurationAssignment {
	if in == nil {
		return nil
	}
	out := new(PolicyVirtualMachineConfigurationAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyVirtualMachineConfigurationAssignment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVirtualMachineConfigurationAssignmentList) DeepCopyInto(out *PolicyVirtualMachineConfigurationAssignmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PolicyVirtualMachineConfigurationAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVirtualMachineConfigurationAssignmentList.
func (in *PolicyVirtualMachineConfigurationAssignmentList) DeepCopy() *PolicyVirtualMachineConfigurationAssignmentList {
	if in == nil {
		return nil
	}
	out := new(PolicyVirtualMachineConfigurationAssignmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyVirtualMachineConfigurationAssignmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVirtualMachineConfigurationAssignmentObservation) DeepCopyInto(out *PolicyVirtualMachineConfigurationAssignmentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVirtualMachineConfigurationAssignmentObservation.
func (in *PolicyVirtualMachineConfigurationAssignmentObservation) DeepCopy() *PolicyVirtualMachineConfigurationAssignmentObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyVirtualMachineConfigurationAssignmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVirtualMachineConfigurationAssignmentParameters) DeepCopyInto(out *PolicyVirtualMachineConfigurationAssignmentParameters) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]ConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineID != nil {
		in, out := &in.VirtualMachineID, &out.VirtualMachineID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVirtualMachineConfigurationAssignmentParameters.
func (in *PolicyVirtualMachineConfigurationAssignmentParameters) DeepCopy() *PolicyVirtualMachineConfigurationAssignmentParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyVirtualMachineConfigurationAssignmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVirtualMachineConfigurationAssignmentSpec) DeepCopyInto(out *PolicyVirtualMachineConfigurationAssignmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVirtualMachineConfigurationAssignmentSpec.
func (in *PolicyVirtualMachineConfigurationAssignmentSpec) DeepCopy() *PolicyVirtualMachineConfigurationAssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(PolicyVirtualMachineConfigurationAssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyVirtualMachineConfigurationAssignmentStatus) DeepCopyInto(out *PolicyVirtualMachineConfigurationAssignmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyVirtualMachineConfigurationAssignmentStatus.
func (in *PolicyVirtualMachineConfigurationAssignmentStatus) DeepCopy() *PolicyVirtualMachineConfigurationAssignmentStatus {
	if in == nil {
		return nil
	}
	out := new(PolicyVirtualMachineConfigurationAssignmentStatus)
	in.DeepCopyInto(out)
	return out
}

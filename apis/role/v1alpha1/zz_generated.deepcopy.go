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
func (in *PermissionsObservation) DeepCopyInto(out *PermissionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionsObservation.
func (in *PermissionsObservation) DeepCopy() *PermissionsObservation {
	if in == nil {
		return nil
	}
	out := new(PermissionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PermissionsParameters) DeepCopyInto(out *PermissionsParameters) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DataActions != nil {
		in, out := &in.DataActions, &out.DataActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NotActions != nil {
		in, out := &in.NotActions, &out.NotActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.NotDataActions != nil {
		in, out := &in.NotDataActions, &out.NotDataActions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PermissionsParameters.
func (in *PermissionsParameters) DeepCopy() *PermissionsParameters {
	if in == nil {
		return nil
	}
	out := new(PermissionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignment) DeepCopyInto(out *RoleAssignment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignment.
func (in *RoleAssignment) DeepCopy() *RoleAssignment {
	if in == nil {
		return nil
	}
	out := new(RoleAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleAssignment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentList) DeepCopyInto(out *RoleAssignmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentList.
func (in *RoleAssignmentList) DeepCopy() *RoleAssignmentList {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleAssignmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentObservation) DeepCopyInto(out *RoleAssignmentObservation) {
	*out = *in
	if in.PrincipalType != nil {
		in, out := &in.PrincipalType, &out.PrincipalType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentObservation.
func (in *RoleAssignmentObservation) DeepCopy() *RoleAssignmentObservation {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentParameters) DeepCopyInto(out *RoleAssignmentParameters) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.ConditionVersion != nil {
		in, out := &in.ConditionVersion, &out.ConditionVersion
		*out = new(string)
		**out = **in
	}
	if in.DelegatedManagedIdentityResourceID != nil {
		in, out := &in.DelegatedManagedIdentityResourceID, &out.DelegatedManagedIdentityResourceID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.RoleDefinitionID != nil {
		in, out := &in.RoleDefinitionID, &out.RoleDefinitionID
		*out = new(string)
		**out = **in
	}
	if in.RoleDefinitionName != nil {
		in, out := &in.RoleDefinitionName, &out.RoleDefinitionName
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.SkipServicePrincipalAadCheck != nil {
		in, out := &in.SkipServicePrincipalAadCheck, &out.SkipServicePrincipalAadCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentParameters.
func (in *RoleAssignmentParameters) DeepCopy() *RoleAssignmentParameters {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentSpec) DeepCopyInto(out *RoleAssignmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentSpec.
func (in *RoleAssignmentSpec) DeepCopy() *RoleAssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAssignmentStatus) DeepCopyInto(out *RoleAssignmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAssignmentStatus.
func (in *RoleAssignmentStatus) DeepCopy() *RoleAssignmentStatus {
	if in == nil {
		return nil
	}
	out := new(RoleAssignmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleDefinition) DeepCopyInto(out *RoleDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleDefinition.
func (in *RoleDefinition) DeepCopy() *RoleDefinition {
	if in == nil {
		return nil
	}
	out := new(RoleDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleDefinitionList) DeepCopyInto(out *RoleDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RoleDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleDefinitionList.
func (in *RoleDefinitionList) DeepCopy() *RoleDefinitionList {
	if in == nil {
		return nil
	}
	out := new(RoleDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RoleDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleDefinitionObservation) DeepCopyInto(out *RoleDefinitionObservation) {
	*out = *in
	if in.RoleDefinitionResourceID != nil {
		in, out := &in.RoleDefinitionResourceID, &out.RoleDefinitionResourceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleDefinitionObservation.
func (in *RoleDefinitionObservation) DeepCopy() *RoleDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(RoleDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleDefinitionParameters) DeepCopyInto(out *RoleDefinitionParameters) {
	*out = *in
	if in.AssignableScopes != nil {
		in, out := &in.AssignableScopes, &out.AssignableScopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]PermissionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RoleDefinitionID != nil {
		in, out := &in.RoleDefinitionID, &out.RoleDefinitionID
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleDefinitionParameters.
func (in *RoleDefinitionParameters) DeepCopy() *RoleDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(RoleDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleDefinitionSpec) DeepCopyInto(out *RoleDefinitionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleDefinitionSpec.
func (in *RoleDefinitionSpec) DeepCopy() *RoleDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(RoleDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleDefinitionStatus) DeepCopyInto(out *RoleDefinitionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleDefinitionStatus.
func (in *RoleDefinitionStatus) DeepCopy() *RoleDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(RoleDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

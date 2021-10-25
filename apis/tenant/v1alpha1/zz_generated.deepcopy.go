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
func (in *TenantTemplateDeployment) DeepCopyInto(out *TenantTemplateDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantTemplateDeployment.
func (in *TenantTemplateDeployment) DeepCopy() *TenantTemplateDeployment {
	if in == nil {
		return nil
	}
	out := new(TenantTemplateDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantTemplateDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantTemplateDeploymentList) DeepCopyInto(out *TenantTemplateDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TenantTemplateDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantTemplateDeploymentList.
func (in *TenantTemplateDeploymentList) DeepCopy() *TenantTemplateDeploymentList {
	if in == nil {
		return nil
	}
	out := new(TenantTemplateDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TenantTemplateDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantTemplateDeploymentObservation) DeepCopyInto(out *TenantTemplateDeploymentObservation) {
	*out = *in
	if in.OutputContent != nil {
		in, out := &in.OutputContent, &out.OutputContent
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantTemplateDeploymentObservation.
func (in *TenantTemplateDeploymentObservation) DeepCopy() *TenantTemplateDeploymentObservation {
	if in == nil {
		return nil
	}
	out := new(TenantTemplateDeploymentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantTemplateDeploymentParameters) DeepCopyInto(out *TenantTemplateDeploymentParameters) {
	*out = *in
	if in.DebugLevel != nil {
		in, out := &in.DebugLevel, &out.DebugLevel
		*out = new(string)
		**out = **in
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
	if in.ParametersContent != nil {
		in, out := &in.ParametersContent, &out.ParametersContent
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
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
	if in.TemplateContent != nil {
		in, out := &in.TemplateContent, &out.TemplateContent
		*out = new(string)
		**out = **in
	}
	if in.TemplateSpecVersionID != nil {
		in, out := &in.TemplateSpecVersionID, &out.TemplateSpecVersionID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantTemplateDeploymentParameters.
func (in *TenantTemplateDeploymentParameters) DeepCopy() *TenantTemplateDeploymentParameters {
	if in == nil {
		return nil
	}
	out := new(TenantTemplateDeploymentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantTemplateDeploymentSpec) DeepCopyInto(out *TenantTemplateDeploymentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantTemplateDeploymentSpec.
func (in *TenantTemplateDeploymentSpec) DeepCopy() *TenantTemplateDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(TenantTemplateDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantTemplateDeploymentStatus) DeepCopyInto(out *TenantTemplateDeploymentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantTemplateDeploymentStatus.
func (in *TenantTemplateDeploymentStatus) DeepCopy() *TenantTemplateDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(TenantTemplateDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

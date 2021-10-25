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
func (in *CustomHeaderObservation) DeepCopyInto(out *CustomHeaderObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHeaderObservation.
func (in *CustomHeaderObservation) DeepCopy() *CustomHeaderObservation {
	if in == nil {
		return nil
	}
	out := new(CustomHeaderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHeaderParameters) DeepCopyInto(out *CustomHeaderParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHeaderParameters.
func (in *CustomHeaderParameters) DeepCopy() *CustomHeaderParameters {
	if in == nil {
		return nil
	}
	out := new(CustomHeaderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSConfigObservation) DeepCopyInto(out *DNSConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSConfigObservation.
func (in *DNSConfigObservation) DeepCopy() *DNSConfigObservation {
	if in == nil {
		return nil
	}
	out := new(DNSConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSConfigParameters) DeepCopyInto(out *DNSConfigParameters) {
	*out = *in
	if in.RelativeName != nil {
		in, out := &in.RelativeName, &out.RelativeName
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSConfigParameters.
func (in *DNSConfigParameters) DeepCopy() *DNSConfigParameters {
	if in == nil {
		return nil
	}
	out := new(DNSConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorConfigCustomHeaderObservation) DeepCopyInto(out *MonitorConfigCustomHeaderObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorConfigCustomHeaderObservation.
func (in *MonitorConfigCustomHeaderObservation) DeepCopy() *MonitorConfigCustomHeaderObservation {
	if in == nil {
		return nil
	}
	out := new(MonitorConfigCustomHeaderObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorConfigCustomHeaderParameters) DeepCopyInto(out *MonitorConfigCustomHeaderParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorConfigCustomHeaderParameters.
func (in *MonitorConfigCustomHeaderParameters) DeepCopy() *MonitorConfigCustomHeaderParameters {
	if in == nil {
		return nil
	}
	out := new(MonitorConfigCustomHeaderParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorConfigObservation) DeepCopyInto(out *MonitorConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorConfigObservation.
func (in *MonitorConfigObservation) DeepCopy() *MonitorConfigObservation {
	if in == nil {
		return nil
	}
	out := new(MonitorConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitorConfigParameters) DeepCopyInto(out *MonitorConfigParameters) {
	*out = *in
	if in.CustomHeader != nil {
		in, out := &in.CustomHeader, &out.CustomHeader
		*out = make([]MonitorConfigCustomHeaderParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExpectedStatusCodeRanges != nil {
		in, out := &in.ExpectedStatusCodeRanges, &out.ExpectedStatusCodeRanges
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IntervalInSeconds != nil {
		in, out := &in.IntervalInSeconds, &out.IntervalInSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.TimeoutInSeconds != nil {
		in, out := &in.TimeoutInSeconds, &out.TimeoutInSeconds
		*out = new(int64)
		**out = **in
	}
	if in.ToleratedNumberOfFailures != nil {
		in, out := &in.ToleratedNumberOfFailures, &out.ToleratedNumberOfFailures
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitorConfigParameters.
func (in *MonitorConfigParameters) DeepCopy() *MonitorConfigParameters {
	if in == nil {
		return nil
	}
	out := new(MonitorConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetObservation) DeepCopyInto(out *SubnetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetObservation.
func (in *SubnetObservation) DeepCopy() *SubnetObservation {
	if in == nil {
		return nil
	}
	out := new(SubnetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetParameters) DeepCopyInto(out *SubnetParameters) {
	*out = *in
	if in.First != nil {
		in, out := &in.First, &out.First
		*out = new(string)
		**out = **in
	}
	if in.Last != nil {
		in, out := &in.Last, &out.Last
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetParameters.
func (in *SubnetParameters) DeepCopy() *SubnetParameters {
	if in == nil {
		return nil
	}
	out := new(SubnetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagerEndpoint) DeepCopyInto(out *TrafficManagerEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagerEndpoint.
func (in *TrafficManagerEndpoint) DeepCopy() *TrafficManagerEndpoint {
	if in == nil {
		return nil
	}
	out := new(TrafficManagerEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrafficManagerEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagerEndpointList) DeepCopyInto(out *TrafficManagerEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrafficManagerEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagerEndpointList.
func (in *TrafficManagerEndpointList) DeepCopy() *TrafficManagerEndpointList {
	if in == nil {
		return nil
	}
	out := new(TrafficManagerEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrafficManagerEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagerEndpointObservation) DeepCopyInto(out *TrafficManagerEndpointObservation) {
	*out = *in
	if in.EndpointMonitorStatus != nil {
		in, out := &in.EndpointMonitorStatus, &out.EndpointMonitorStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagerEndpointObservation.
func (in *TrafficManagerEndpointObservation) DeepCopy() *TrafficManagerEndpointObservation {
	if in == nil {
		return nil
	}
	out := new(TrafficManagerEndpointObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagerEndpointParameters) DeepCopyInto(out *TrafficManagerEndpointParameters) {
	*out = *in
	if in.CustomHeader != nil {
		in, out := &in.CustomHeader, &out.CustomHeader
		*out = make([]CustomHeaderParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EndpointLocation != nil {
		in, out := &in.EndpointLocation, &out.EndpointLocation
		*out = new(string)
		**out = **in
	}
	if in.EndpointStatus != nil {
		in, out := &in.EndpointStatus, &out.EndpointStatus
		*out = new(string)
		**out = **in
	}
	if in.GeoMappings != nil {
		in, out := &in.GeoMappings, &out.GeoMappings
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MinChildEndpoints != nil {
		in, out := &in.MinChildEndpoints, &out.MinChildEndpoints
		*out = new(int64)
		**out = **in
	}
	if in.MinimumRequiredChildEndpointsIPv4 != nil {
		in, out := &in.MinimumRequiredChildEndpointsIPv4, &out.MinimumRequiredChildEndpointsIPv4
		*out = new(int64)
		**out = **in
	}
	if in.MinimumRequiredChildEndpointsIPv6 != nil {
		in, out := &in.MinimumRequiredChildEndpointsIPv6, &out.MinimumRequiredChildEndpointsIPv6
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.ProfileName != nil {
		in, out := &in.ProfileName, &out.ProfileName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = make([]SubnetParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.TargetResourceID != nil {
		in, out := &in.TargetResourceID, &out.TargetResourceID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagerEndpointParameters.
func (in *TrafficManagerEndpointParameters) DeepCopy() *TrafficManagerEndpointParameters {
	if in == nil {
		return nil
	}
	out := new(TrafficManagerEndpointParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagerEndpointSpec) DeepCopyInto(out *TrafficManagerEndpointSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagerEndpointSpec.
func (in *TrafficManagerEndpointSpec) DeepCopy() *TrafficManagerEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(TrafficManagerEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagerEndpointStatus) DeepCopyInto(out *TrafficManagerEndpointStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagerEndpointStatus.
func (in *TrafficManagerEndpointStatus) DeepCopy() *TrafficManagerEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(TrafficManagerEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagerProfile) DeepCopyInto(out *TrafficManagerProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagerProfile.
func (in *TrafficManagerProfile) DeepCopy() *TrafficManagerProfile {
	if in == nil {
		return nil
	}
	out := new(TrafficManagerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrafficManagerProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagerProfileList) DeepCopyInto(out *TrafficManagerProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrafficManagerProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagerProfileList.
func (in *TrafficManagerProfileList) DeepCopy() *TrafficManagerProfileList {
	if in == nil {
		return nil
	}
	out := new(TrafficManagerProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrafficManagerProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagerProfileObservation) DeepCopyInto(out *TrafficManagerProfileObservation) {
	*out = *in
	if in.Fqdn != nil {
		in, out := &in.Fqdn, &out.Fqdn
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagerProfileObservation.
func (in *TrafficManagerProfileObservation) DeepCopy() *TrafficManagerProfileObservation {
	if in == nil {
		return nil
	}
	out := new(TrafficManagerProfileObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagerProfileParameters) DeepCopyInto(out *TrafficManagerProfileParameters) {
	*out = *in
	if in.DNSConfig != nil {
		in, out := &in.DNSConfig, &out.DNSConfig
		*out = make([]DNSConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaxReturn != nil {
		in, out := &in.MaxReturn, &out.MaxReturn
		*out = new(int64)
		**out = **in
	}
	if in.MonitorConfig != nil {
		in, out := &in.MonitorConfig, &out.MonitorConfig
		*out = make([]MonitorConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ProfileStatus != nil {
		in, out := &in.ProfileStatus, &out.ProfileStatus
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
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
	if in.TrafficRoutingMethod != nil {
		in, out := &in.TrafficRoutingMethod, &out.TrafficRoutingMethod
		*out = new(string)
		**out = **in
	}
	if in.TrafficViewEnabled != nil {
		in, out := &in.TrafficViewEnabled, &out.TrafficViewEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagerProfileParameters.
func (in *TrafficManagerProfileParameters) DeepCopy() *TrafficManagerProfileParameters {
	if in == nil {
		return nil
	}
	out := new(TrafficManagerProfileParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagerProfileSpec) DeepCopyInto(out *TrafficManagerProfileSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagerProfileSpec.
func (in *TrafficManagerProfileSpec) DeepCopy() *TrafficManagerProfileSpec {
	if in == nil {
		return nil
	}
	out := new(TrafficManagerProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficManagerProfileStatus) DeepCopyInto(out *TrafficManagerProfileStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficManagerProfileStatus.
func (in *TrafficManagerProfileStatus) DeepCopy() *TrafficManagerProfileStatus {
	if in == nil {
		return nil
	}
	out := new(TrafficManagerProfileStatus)
	in.DeepCopyInto(out)
	return out
}

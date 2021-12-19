//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/core/v1"
	metav1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/meta/v1"
	apicorev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointApplyConfiguration) DeepCopyInto(out *EndpointApplyConfiguration) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = new(EndpointConditionsApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.TargetRef != nil {
		in, out := &in.TargetRef, &out.TargetRef
		*out = new(corev1.ObjectReferenceApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.DeprecatedTopology != nil {
		in, out := &in.DeprecatedTopology, &out.DeprecatedTopology
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeName != nil {
		in, out := &in.NodeName, &out.NodeName
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	if in.Hints != nil {
		in, out := &in.Hints, &out.Hints
		*out = new(EndpointHintsApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointApplyConfiguration.
func (in *EndpointApplyConfiguration) DeepCopy() *EndpointApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(EndpointApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointConditionsApplyConfiguration) DeepCopyInto(out *EndpointConditionsApplyConfiguration) {
	*out = *in
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = new(bool)
		**out = **in
	}
	if in.Serving != nil {
		in, out := &in.Serving, &out.Serving
		*out = new(bool)
		**out = **in
	}
	if in.Terminating != nil {
		in, out := &in.Terminating, &out.Terminating
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointConditionsApplyConfiguration.
func (in *EndpointConditionsApplyConfiguration) DeepCopy() *EndpointConditionsApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(EndpointConditionsApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointHintsApplyConfiguration) DeepCopyInto(out *EndpointHintsApplyConfiguration) {
	*out = *in
	if in.ForZones != nil {
		in, out := &in.ForZones, &out.ForZones
		*out = make([]ForZoneApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointHintsApplyConfiguration.
func (in *EndpointHintsApplyConfiguration) DeepCopy() *EndpointHintsApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(EndpointHintsApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointPortApplyConfiguration) DeepCopyInto(out *EndpointPortApplyConfiguration) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(apicorev1.Protocol)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.AppProtocol != nil {
		in, out := &in.AppProtocol, &out.AppProtocol
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointPortApplyConfiguration.
func (in *EndpointPortApplyConfiguration) DeepCopy() *EndpointPortApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(EndpointPortApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSliceApplyConfiguration) DeepCopyInto(out *EndpointSliceApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(metav1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.AddressType != nil {
		in, out := &in.AddressType, &out.AddressType
		*out = new(discoveryv1.AddressType)
		**out = **in
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]EndpointApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]EndpointPortApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSliceApplyConfiguration.
func (in *EndpointSliceApplyConfiguration) DeepCopy() *EndpointSliceApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(EndpointSliceApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForZoneApplyConfiguration) DeepCopyInto(out *ForZoneApplyConfiguration) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForZoneApplyConfiguration.
func (in *ForZoneApplyConfiguration) DeepCopy() *ForZoneApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ForZoneApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

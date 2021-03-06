//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/zoetrope/ssa-helper/applyconfigurations/meta/v1"
	apiserverinternalv1alpha1 "k8s.io/api/apiserverinternal/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerStorageVersionApplyConfiguration) DeepCopyInto(out *ServerStorageVersionApplyConfiguration) {
	*out = *in
	if in.APIServerID != nil {
		in, out := &in.APIServerID, &out.APIServerID
		*out = new(string)
		**out = **in
	}
	if in.EncodingVersion != nil {
		in, out := &in.EncodingVersion, &out.EncodingVersion
		*out = new(string)
		**out = **in
	}
	if in.DecodableVersions != nil {
		in, out := &in.DecodableVersions, &out.DecodableVersions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerStorageVersionApplyConfiguration.
func (in *ServerStorageVersionApplyConfiguration) DeepCopy() *ServerStorageVersionApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServerStorageVersionApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageVersionApplyConfiguration) DeepCopyInto(out *StorageVersionApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(v1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(apiserverinternalv1alpha1.StorageVersionSpec)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(StorageVersionStatusApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageVersionApplyConfiguration.
func (in *StorageVersionApplyConfiguration) DeepCopy() *StorageVersionApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(StorageVersionApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageVersionConditionApplyConfiguration) DeepCopyInto(out *StorageVersionConditionApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(apiserverinternalv1alpha1.StorageVersionConditionType)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(apiserverinternalv1alpha1.ConditionStatus)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageVersionConditionApplyConfiguration.
func (in *StorageVersionConditionApplyConfiguration) DeepCopy() *StorageVersionConditionApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(StorageVersionConditionApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageVersionStatusApplyConfiguration) DeepCopyInto(out *StorageVersionStatusApplyConfiguration) {
	*out = *in
	if in.StorageVersions != nil {
		in, out := &in.StorageVersions, &out.StorageVersions
		*out = make([]ServerStorageVersionApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CommonEncodingVersion != nil {
		in, out := &in.CommonEncodingVersion, &out.CommonEncodingVersion
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]StorageVersionConditionApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageVersionStatusApplyConfiguration.
func (in *StorageVersionStatusApplyConfiguration) DeepCopy() *StorageVersionStatusApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(StorageVersionStatusApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

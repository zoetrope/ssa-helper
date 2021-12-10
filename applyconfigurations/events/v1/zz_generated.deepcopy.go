//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	corev1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/core/v1"
	metav1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventApplyConfiguration) DeepCopyInto(out *EventApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(metav1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.EventTime != nil {
		in, out := &in.EventTime, &out.EventTime
		*out = (*in).DeepCopy()
	}
	if in.Series != nil {
		in, out := &in.Series, &out.Series
		*out = new(EventSeriesApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ReportingController != nil {
		in, out := &in.ReportingController, &out.ReportingController
		*out = new(string)
		**out = **in
	}
	if in.ReportingInstance != nil {
		in, out := &in.ReportingInstance, &out.ReportingInstance
		*out = new(string)
		**out = **in
	}
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Reason != nil {
		in, out := &in.Reason, &out.Reason
		*out = new(string)
		**out = **in
	}
	if in.Regarding != nil {
		in, out := &in.Regarding, &out.Regarding
		*out = new(corev1.ObjectReferenceApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Related != nil {
		in, out := &in.Related, &out.Related
		*out = new(corev1.ObjectReferenceApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Note != nil {
		in, out := &in.Note, &out.Note
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.DeprecatedSource != nil {
		in, out := &in.DeprecatedSource, &out.DeprecatedSource
		*out = new(corev1.EventSourceApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.DeprecatedFirstTimestamp != nil {
		in, out := &in.DeprecatedFirstTimestamp, &out.DeprecatedFirstTimestamp
		*out = (*in).DeepCopy()
	}
	if in.DeprecatedLastTimestamp != nil {
		in, out := &in.DeprecatedLastTimestamp, &out.DeprecatedLastTimestamp
		*out = (*in).DeepCopy()
	}
	if in.DeprecatedCount != nil {
		in, out := &in.DeprecatedCount, &out.DeprecatedCount
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventApplyConfiguration.
func (in *EventApplyConfiguration) DeepCopy() *EventApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(EventApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSeriesApplyConfiguration) DeepCopyInto(out *EventSeriesApplyConfiguration) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int32)
		**out = **in
	}
	if in.LastObservedTime != nil {
		in, out := &in.LastObservedTime, &out.LastObservedTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSeriesApplyConfiguration.
func (in *EventSeriesApplyConfiguration) DeepCopy() *EventSeriesApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(EventSeriesApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}
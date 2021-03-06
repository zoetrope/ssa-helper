//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	batchv1 "github.com/zoetrope/ssa-helper/applyconfigurations/batch/v1"
	corev1 "github.com/zoetrope/ssa-helper/applyconfigurations/core/v1"
	v1 "github.com/zoetrope/ssa-helper/applyconfigurations/meta/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobApplyConfiguration) DeepCopyInto(out *CronJobApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(v1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(CronJobSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(CronJobStatusApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobApplyConfiguration.
func (in *CronJobApplyConfiguration) DeepCopy() *CronJobApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(CronJobApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobSpecApplyConfiguration) DeepCopyInto(out *CronJobSpecApplyConfiguration) {
	*out = *in
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(string)
		**out = **in
	}
	if in.StartingDeadlineSeconds != nil {
		in, out := &in.StartingDeadlineSeconds, &out.StartingDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.ConcurrencyPolicy != nil {
		in, out := &in.ConcurrencyPolicy, &out.ConcurrencyPolicy
		*out = new(batchv1beta1.ConcurrencyPolicy)
		**out = **in
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
	if in.JobTemplate != nil {
		in, out := &in.JobTemplate, &out.JobTemplate
		*out = new(JobTemplateSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.SuccessfulJobsHistoryLimit != nil {
		in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.FailedJobsHistoryLimit != nil {
		in, out := &in.FailedJobsHistoryLimit, &out.FailedJobsHistoryLimit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobSpecApplyConfiguration.
func (in *CronJobSpecApplyConfiguration) DeepCopy() *CronJobSpecApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(CronJobSpecApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobStatusApplyConfiguration) DeepCopyInto(out *CronJobStatusApplyConfiguration) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = make([]corev1.ObjectReferenceApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		*out = (*in).DeepCopy()
	}
	if in.LastSuccessfulTime != nil {
		in, out := &in.LastSuccessfulTime, &out.LastSuccessfulTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobStatusApplyConfiguration.
func (in *CronJobStatusApplyConfiguration) DeepCopy() *CronJobStatusApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(CronJobStatusApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTemplateSpecApplyConfiguration) DeepCopyInto(out *JobTemplateSpecApplyConfiguration) {
	*out = *in
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(v1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(batchv1.JobSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTemplateSpecApplyConfiguration.
func (in *JobTemplateSpecApplyConfiguration) DeepCopy() *JobTemplateSpecApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(JobTemplateSpecApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

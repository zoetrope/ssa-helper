//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	applyconfigurationscorev1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/core/v1"
	metav1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRevisionApplyConfiguration) DeepCopyInto(out *ControllerRevisionApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(metav1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRevisionApplyConfiguration.
func (in *ControllerRevisionApplyConfiguration) DeepCopy() *ControllerRevisionApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerRevisionApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetApplyConfiguration) DeepCopyInto(out *DaemonSetApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(metav1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(DaemonSetSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(DaemonSetStatusApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetApplyConfiguration.
func (in *DaemonSetApplyConfiguration) DeepCopy() *DaemonSetApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(DaemonSetApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetConditionApplyConfiguration) DeepCopyInto(out *DaemonSetConditionApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(appsv1.DaemonSetConditionType)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(corev1.ConditionStatus)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetConditionApplyConfiguration.
func (in *DaemonSetConditionApplyConfiguration) DeepCopy() *DaemonSetConditionApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(DaemonSetConditionApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetSpecApplyConfiguration) DeepCopyInto(out *DaemonSetSpecApplyConfiguration) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelectorApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(applyconfigurationscorev1.PodTemplateSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.UpdateStrategy != nil {
		in, out := &in.UpdateStrategy, &out.UpdateStrategy
		*out = new(DaemonSetUpdateStrategyApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.MinReadySeconds != nil {
		in, out := &in.MinReadySeconds, &out.MinReadySeconds
		*out = new(int32)
		**out = **in
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetSpecApplyConfiguration.
func (in *DaemonSetSpecApplyConfiguration) DeepCopy() *DaemonSetSpecApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(DaemonSetSpecApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetStatusApplyConfiguration) DeepCopyInto(out *DaemonSetStatusApplyConfiguration) {
	*out = *in
	if in.CurrentNumberScheduled != nil {
		in, out := &in.CurrentNumberScheduled, &out.CurrentNumberScheduled
		*out = new(int32)
		**out = **in
	}
	if in.NumberMisscheduled != nil {
		in, out := &in.NumberMisscheduled, &out.NumberMisscheduled
		*out = new(int32)
		**out = **in
	}
	if in.DesiredNumberScheduled != nil {
		in, out := &in.DesiredNumberScheduled, &out.DesiredNumberScheduled
		*out = new(int32)
		**out = **in
	}
	if in.NumberReady != nil {
		in, out := &in.NumberReady, &out.NumberReady
		*out = new(int32)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.UpdatedNumberScheduled != nil {
		in, out := &in.UpdatedNumberScheduled, &out.UpdatedNumberScheduled
		*out = new(int32)
		**out = **in
	}
	if in.NumberAvailable != nil {
		in, out := &in.NumberAvailable, &out.NumberAvailable
		*out = new(int32)
		**out = **in
	}
	if in.NumberUnavailable != nil {
		in, out := &in.NumberUnavailable, &out.NumberUnavailable
		*out = new(int32)
		**out = **in
	}
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]DaemonSetConditionApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetStatusApplyConfiguration.
func (in *DaemonSetStatusApplyConfiguration) DeepCopy() *DaemonSetStatusApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(DaemonSetStatusApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetUpdateStrategyApplyConfiguration) DeepCopyInto(out *DaemonSetUpdateStrategyApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(appsv1.DaemonSetUpdateStrategyType)
		**out = **in
	}
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(RollingUpdateDaemonSetApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetUpdateStrategyApplyConfiguration.
func (in *DaemonSetUpdateStrategyApplyConfiguration) DeepCopy() *DaemonSetUpdateStrategyApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(DaemonSetUpdateStrategyApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentApplyConfiguration) DeepCopyInto(out *DeploymentApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(metav1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(DeploymentSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(DeploymentStatusApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentApplyConfiguration.
func (in *DeploymentApplyConfiguration) DeepCopy() *DeploymentApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeploymentApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentConditionApplyConfiguration) DeepCopyInto(out *DeploymentConditionApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(appsv1.DeploymentConditionType)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(corev1.ConditionStatus)
		**out = **in
	}
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = (*in).DeepCopy()
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentConditionApplyConfiguration.
func (in *DeploymentConditionApplyConfiguration) DeepCopy() *DeploymentConditionApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeploymentConditionApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSpecApplyConfiguration) DeepCopyInto(out *DeploymentSpecApplyConfiguration) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelectorApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(applyconfigurationscorev1.PodTemplateSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(DeploymentStrategyApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.MinReadySeconds != nil {
		in, out := &in.MinReadySeconds, &out.MinReadySeconds
		*out = new(int32)
		**out = **in
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.ProgressDeadlineSeconds != nil {
		in, out := &in.ProgressDeadlineSeconds, &out.ProgressDeadlineSeconds
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSpecApplyConfiguration.
func (in *DeploymentSpecApplyConfiguration) DeepCopy() *DeploymentSpecApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeploymentSpecApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatusApplyConfiguration) DeepCopyInto(out *DeploymentStatusApplyConfiguration) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.UpdatedReplicas != nil {
		in, out := &in.UpdatedReplicas, &out.UpdatedReplicas
		*out = new(int32)
		**out = **in
	}
	if in.ReadyReplicas != nil {
		in, out := &in.ReadyReplicas, &out.ReadyReplicas
		*out = new(int32)
		**out = **in
	}
	if in.AvailableReplicas != nil {
		in, out := &in.AvailableReplicas, &out.AvailableReplicas
		*out = new(int32)
		**out = **in
	}
	if in.UnavailableReplicas != nil {
		in, out := &in.UnavailableReplicas, &out.UnavailableReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]DeploymentConditionApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatusApplyConfiguration.
func (in *DeploymentStatusApplyConfiguration) DeepCopy() *DeploymentStatusApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatusApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStrategyApplyConfiguration) DeepCopyInto(out *DeploymentStrategyApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(appsv1.DeploymentStrategyType)
		**out = **in
	}
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(RollingUpdateDeploymentApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStrategyApplyConfiguration.
func (in *DeploymentStrategyApplyConfiguration) DeepCopy() *DeploymentStrategyApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(DeploymentStrategyApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSetApplyConfiguration) DeepCopyInto(out *ReplicaSetApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(metav1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(ReplicaSetSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ReplicaSetStatusApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSetApplyConfiguration.
func (in *ReplicaSetApplyConfiguration) DeepCopy() *ReplicaSetApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ReplicaSetApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSetConditionApplyConfiguration) DeepCopyInto(out *ReplicaSetConditionApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(appsv1.ReplicaSetConditionType)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(corev1.ConditionStatus)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSetConditionApplyConfiguration.
func (in *ReplicaSetConditionApplyConfiguration) DeepCopy() *ReplicaSetConditionApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ReplicaSetConditionApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSetSpecApplyConfiguration) DeepCopyInto(out *ReplicaSetSpecApplyConfiguration) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.MinReadySeconds != nil {
		in, out := &in.MinReadySeconds, &out.MinReadySeconds
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelectorApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(applyconfigurationscorev1.PodTemplateSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSetSpecApplyConfiguration.
func (in *ReplicaSetSpecApplyConfiguration) DeepCopy() *ReplicaSetSpecApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ReplicaSetSpecApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSetStatusApplyConfiguration) DeepCopyInto(out *ReplicaSetStatusApplyConfiguration) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.FullyLabeledReplicas != nil {
		in, out := &in.FullyLabeledReplicas, &out.FullyLabeledReplicas
		*out = new(int32)
		**out = **in
	}
	if in.ReadyReplicas != nil {
		in, out := &in.ReadyReplicas, &out.ReadyReplicas
		*out = new(int32)
		**out = **in
	}
	if in.AvailableReplicas != nil {
		in, out := &in.AvailableReplicas, &out.AvailableReplicas
		*out = new(int32)
		**out = **in
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ReplicaSetConditionApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSetStatusApplyConfiguration.
func (in *ReplicaSetStatusApplyConfiguration) DeepCopy() *ReplicaSetStatusApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ReplicaSetStatusApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateDaemonSetApplyConfiguration) DeepCopyInto(out *RollingUpdateDaemonSetApplyConfiguration) {
	*out = *in
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxSurge != nil {
		in, out := &in.MaxSurge, &out.MaxSurge
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateDaemonSetApplyConfiguration.
func (in *RollingUpdateDaemonSetApplyConfiguration) DeepCopy() *RollingUpdateDaemonSetApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateDaemonSetApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateDeploymentApplyConfiguration) DeepCopyInto(out *RollingUpdateDeploymentApplyConfiguration) {
	*out = *in
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxSurge != nil {
		in, out := &in.MaxSurge, &out.MaxSurge
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateDeploymentApplyConfiguration.
func (in *RollingUpdateDeploymentApplyConfiguration) DeepCopy() *RollingUpdateDeploymentApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateDeploymentApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RollingUpdateStatefulSetStrategyApplyConfiguration) DeepCopyInto(out *RollingUpdateStatefulSetStrategyApplyConfiguration) {
	*out = *in
	if in.Partition != nil {
		in, out := &in.Partition, &out.Partition
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RollingUpdateStatefulSetStrategyApplyConfiguration.
func (in *RollingUpdateStatefulSetStrategyApplyConfiguration) DeepCopy() *RollingUpdateStatefulSetStrategyApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(RollingUpdateStatefulSetStrategyApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetApplyConfiguration) DeepCopyInto(out *StatefulSetApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(metav1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(StatefulSetSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(StatefulSetStatusApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetApplyConfiguration.
func (in *StatefulSetApplyConfiguration) DeepCopy() *StatefulSetApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(StatefulSetApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetConditionApplyConfiguration) DeepCopyInto(out *StatefulSetConditionApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(appsv1.StatefulSetConditionType)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(corev1.ConditionStatus)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetConditionApplyConfiguration.
func (in *StatefulSetConditionApplyConfiguration) DeepCopy() *StatefulSetConditionApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(StatefulSetConditionApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetSpecApplyConfiguration) DeepCopyInto(out *StatefulSetSpecApplyConfiguration) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelectorApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(applyconfigurationscorev1.PodTemplateSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]applyconfigurationscorev1.PersistentVolumeClaimApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.PodManagementPolicy != nil {
		in, out := &in.PodManagementPolicy, &out.PodManagementPolicy
		*out = new(appsv1.PodManagementPolicyType)
		**out = **in
	}
	if in.UpdateStrategy != nil {
		in, out := &in.UpdateStrategy, &out.UpdateStrategy
		*out = new(StatefulSetUpdateStrategyApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.RevisionHistoryLimit != nil {
		in, out := &in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = **in
	}
	if in.MinReadySeconds != nil {
		in, out := &in.MinReadySeconds, &out.MinReadySeconds
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetSpecApplyConfiguration.
func (in *StatefulSetSpecApplyConfiguration) DeepCopy() *StatefulSetSpecApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(StatefulSetSpecApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetStatusApplyConfiguration) DeepCopyInto(out *StatefulSetStatusApplyConfiguration) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.ReadyReplicas != nil {
		in, out := &in.ReadyReplicas, &out.ReadyReplicas
		*out = new(int32)
		**out = **in
	}
	if in.CurrentReplicas != nil {
		in, out := &in.CurrentReplicas, &out.CurrentReplicas
		*out = new(int32)
		**out = **in
	}
	if in.UpdatedReplicas != nil {
		in, out := &in.UpdatedReplicas, &out.UpdatedReplicas
		*out = new(int32)
		**out = **in
	}
	if in.CurrentRevision != nil {
		in, out := &in.CurrentRevision, &out.CurrentRevision
		*out = new(string)
		**out = **in
	}
	if in.UpdateRevision != nil {
		in, out := &in.UpdateRevision, &out.UpdateRevision
		*out = new(string)
		**out = **in
	}
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]StatefulSetConditionApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AvailableReplicas != nil {
		in, out := &in.AvailableReplicas, &out.AvailableReplicas
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetStatusApplyConfiguration.
func (in *StatefulSetStatusApplyConfiguration) DeepCopy() *StatefulSetStatusApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(StatefulSetStatusApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetUpdateStrategyApplyConfiguration) DeepCopyInto(out *StatefulSetUpdateStrategyApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(appsv1.StatefulSetUpdateStrategyType)
		**out = **in
	}
	if in.RollingUpdate != nil {
		in, out := &in.RollingUpdate, &out.RollingUpdate
		*out = new(RollingUpdateStatefulSetStrategyApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetUpdateStrategyApplyConfiguration.
func (in *StatefulSetUpdateStrategyApplyConfiguration) DeepCopy() *StatefulSetUpdateStrategyApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(StatefulSetUpdateStrategyApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

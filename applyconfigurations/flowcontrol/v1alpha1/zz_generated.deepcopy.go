//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/meta/v1"
	flowcontrolv1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowDistinguisherMethodApplyConfiguration) DeepCopyInto(out *FlowDistinguisherMethodApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(flowcontrolv1alpha1.FlowDistinguisherMethodType)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowDistinguisherMethodApplyConfiguration.
func (in *FlowDistinguisherMethodApplyConfiguration) DeepCopy() *FlowDistinguisherMethodApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(FlowDistinguisherMethodApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowSchemaApplyConfiguration) DeepCopyInto(out *FlowSchemaApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(v1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(FlowSchemaSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(FlowSchemaStatusApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowSchemaApplyConfiguration.
func (in *FlowSchemaApplyConfiguration) DeepCopy() *FlowSchemaApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(FlowSchemaApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowSchemaConditionApplyConfiguration) DeepCopyInto(out *FlowSchemaConditionApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(flowcontrolv1alpha1.FlowSchemaConditionType)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(flowcontrolv1alpha1.ConditionStatus)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowSchemaConditionApplyConfiguration.
func (in *FlowSchemaConditionApplyConfiguration) DeepCopy() *FlowSchemaConditionApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(FlowSchemaConditionApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowSchemaSpecApplyConfiguration) DeepCopyInto(out *FlowSchemaSpecApplyConfiguration) {
	*out = *in
	if in.PriorityLevelConfiguration != nil {
		in, out := &in.PriorityLevelConfiguration, &out.PriorityLevelConfiguration
		*out = new(PriorityLevelConfigurationReferenceApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.MatchingPrecedence != nil {
		in, out := &in.MatchingPrecedence, &out.MatchingPrecedence
		*out = new(int32)
		**out = **in
	}
	if in.DistinguisherMethod != nil {
		in, out := &in.DistinguisherMethod, &out.DistinguisherMethod
		*out = new(FlowDistinguisherMethodApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PolicyRulesWithSubjectsApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowSchemaSpecApplyConfiguration.
func (in *FlowSchemaSpecApplyConfiguration) DeepCopy() *FlowSchemaSpecApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(FlowSchemaSpecApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowSchemaStatusApplyConfiguration) DeepCopyInto(out *FlowSchemaStatusApplyConfiguration) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]FlowSchemaConditionApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowSchemaStatusApplyConfiguration.
func (in *FlowSchemaStatusApplyConfiguration) DeepCopy() *FlowSchemaStatusApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(FlowSchemaStatusApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSubjectApplyConfiguration) DeepCopyInto(out *GroupSubjectApplyConfiguration) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSubjectApplyConfiguration.
func (in *GroupSubjectApplyConfiguration) DeepCopy() *GroupSubjectApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(GroupSubjectApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitResponseApplyConfiguration) DeepCopyInto(out *LimitResponseApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(flowcontrolv1alpha1.LimitResponseType)
		**out = **in
	}
	if in.Queuing != nil {
		in, out := &in.Queuing, &out.Queuing
		*out = new(QueuingConfigurationApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitResponseApplyConfiguration.
func (in *LimitResponseApplyConfiguration) DeepCopy() *LimitResponseApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(LimitResponseApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LimitedPriorityLevelConfigurationApplyConfiguration) DeepCopyInto(out *LimitedPriorityLevelConfigurationApplyConfiguration) {
	*out = *in
	if in.AssuredConcurrencyShares != nil {
		in, out := &in.AssuredConcurrencyShares, &out.AssuredConcurrencyShares
		*out = new(int32)
		**out = **in
	}
	if in.LimitResponse != nil {
		in, out := &in.LimitResponse, &out.LimitResponse
		*out = new(LimitResponseApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LimitedPriorityLevelConfigurationApplyConfiguration.
func (in *LimitedPriorityLevelConfigurationApplyConfiguration) DeepCopy() *LimitedPriorityLevelConfigurationApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(LimitedPriorityLevelConfigurationApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NonResourcePolicyRuleApplyConfiguration) DeepCopyInto(out *NonResourcePolicyRuleApplyConfiguration) {
	*out = *in
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NonResourceURLs != nil {
		in, out := &in.NonResourceURLs, &out.NonResourceURLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NonResourcePolicyRuleApplyConfiguration.
func (in *NonResourcePolicyRuleApplyConfiguration) DeepCopy() *NonResourcePolicyRuleApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(NonResourcePolicyRuleApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyRulesWithSubjectsApplyConfiguration) DeepCopyInto(out *PolicyRulesWithSubjectsApplyConfiguration) {
	*out = *in
	if in.Subjects != nil {
		in, out := &in.Subjects, &out.Subjects
		*out = make([]SubjectApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceRules != nil {
		in, out := &in.ResourceRules, &out.ResourceRules
		*out = make([]ResourcePolicyRuleApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NonResourceRules != nil {
		in, out := &in.NonResourceRules, &out.NonResourceRules
		*out = make([]NonResourcePolicyRuleApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyRulesWithSubjectsApplyConfiguration.
func (in *PolicyRulesWithSubjectsApplyConfiguration) DeepCopy() *PolicyRulesWithSubjectsApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(PolicyRulesWithSubjectsApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityLevelConfigurationApplyConfiguration) DeepCopyInto(out *PriorityLevelConfigurationApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(v1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(PriorityLevelConfigurationSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(PriorityLevelConfigurationStatusApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityLevelConfigurationApplyConfiguration.
func (in *PriorityLevelConfigurationApplyConfiguration) DeepCopy() *PriorityLevelConfigurationApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(PriorityLevelConfigurationApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityLevelConfigurationConditionApplyConfiguration) DeepCopyInto(out *PriorityLevelConfigurationConditionApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(flowcontrolv1alpha1.PriorityLevelConfigurationConditionType)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(flowcontrolv1alpha1.ConditionStatus)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityLevelConfigurationConditionApplyConfiguration.
func (in *PriorityLevelConfigurationConditionApplyConfiguration) DeepCopy() *PriorityLevelConfigurationConditionApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(PriorityLevelConfigurationConditionApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityLevelConfigurationReferenceApplyConfiguration) DeepCopyInto(out *PriorityLevelConfigurationReferenceApplyConfiguration) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityLevelConfigurationReferenceApplyConfiguration.
func (in *PriorityLevelConfigurationReferenceApplyConfiguration) DeepCopy() *PriorityLevelConfigurationReferenceApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(PriorityLevelConfigurationReferenceApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityLevelConfigurationSpecApplyConfiguration) DeepCopyInto(out *PriorityLevelConfigurationSpecApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(flowcontrolv1alpha1.PriorityLevelEnablement)
		**out = **in
	}
	if in.Limited != nil {
		in, out := &in.Limited, &out.Limited
		*out = new(LimitedPriorityLevelConfigurationApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityLevelConfigurationSpecApplyConfiguration.
func (in *PriorityLevelConfigurationSpecApplyConfiguration) DeepCopy() *PriorityLevelConfigurationSpecApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(PriorityLevelConfigurationSpecApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PriorityLevelConfigurationStatusApplyConfiguration) DeepCopyInto(out *PriorityLevelConfigurationStatusApplyConfiguration) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PriorityLevelConfigurationConditionApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PriorityLevelConfigurationStatusApplyConfiguration.
func (in *PriorityLevelConfigurationStatusApplyConfiguration) DeepCopy() *PriorityLevelConfigurationStatusApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(PriorityLevelConfigurationStatusApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueuingConfigurationApplyConfiguration) DeepCopyInto(out *QueuingConfigurationApplyConfiguration) {
	*out = *in
	if in.Queues != nil {
		in, out := &in.Queues, &out.Queues
		*out = new(int32)
		**out = **in
	}
	if in.HandSize != nil {
		in, out := &in.HandSize, &out.HandSize
		*out = new(int32)
		**out = **in
	}
	if in.QueueLengthLimit != nil {
		in, out := &in.QueueLengthLimit, &out.QueueLengthLimit
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueuingConfigurationApplyConfiguration.
func (in *QueuingConfigurationApplyConfiguration) DeepCopy() *QueuingConfigurationApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(QueuingConfigurationApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicyRuleApplyConfiguration) DeepCopyInto(out *ResourcePolicyRuleApplyConfiguration) {
	*out = *in
	if in.Verbs != nil {
		in, out := &in.Verbs, &out.Verbs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.APIGroups != nil {
		in, out := &in.APIGroups, &out.APIGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterScope != nil {
		in, out := &in.ClusterScope, &out.ClusterScope
		*out = new(bool)
		**out = **in
	}
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicyRuleApplyConfiguration.
func (in *ResourcePolicyRuleApplyConfiguration) DeepCopy() *ResourcePolicyRuleApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicyRuleApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountSubjectApplyConfiguration) DeepCopyInto(out *ServiceAccountSubjectApplyConfiguration) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountSubjectApplyConfiguration.
func (in *ServiceAccountSubjectApplyConfiguration) DeepCopy() *ServiceAccountSubjectApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountSubjectApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubjectApplyConfiguration) DeepCopyInto(out *SubjectApplyConfiguration) {
	*out = *in
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(flowcontrolv1alpha1.SubjectKind)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(UserSubjectApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = new(GroupSubjectApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(ServiceAccountSubjectApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubjectApplyConfiguration.
func (in *SubjectApplyConfiguration) DeepCopy() *SubjectApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(SubjectApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSubjectApplyConfiguration) DeepCopyInto(out *UserSubjectApplyConfiguration) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSubjectApplyConfiguration.
func (in *UserSubjectApplyConfiguration) DeepCopy() *UserSubjectApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(UserSubjectApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

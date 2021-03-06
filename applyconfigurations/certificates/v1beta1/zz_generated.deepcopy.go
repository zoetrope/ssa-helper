//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "github.com/zoetrope/ssa-helper/applyconfigurations/meta/v1"
	certificatesv1beta1 "k8s.io/api/certificates/v1beta1"
	corev1 "k8s.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSigningRequestApplyConfiguration) DeepCopyInto(out *CertificateSigningRequestApplyConfiguration) {
	*out = *in
	in.TypeMetaApplyConfiguration.DeepCopyInto(&out.TypeMetaApplyConfiguration)
	if in.ObjectMetaApplyConfiguration != nil {
		in, out := &in.ObjectMetaApplyConfiguration, &out.ObjectMetaApplyConfiguration
		*out = new(v1.ObjectMetaApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(CertificateSigningRequestSpecApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(CertificateSigningRequestStatusApplyConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSigningRequestApplyConfiguration.
func (in *CertificateSigningRequestApplyConfiguration) DeepCopy() *CertificateSigningRequestApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(CertificateSigningRequestApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSigningRequestConditionApplyConfiguration) DeepCopyInto(out *CertificateSigningRequestConditionApplyConfiguration) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(certificatesv1beta1.RequestConditionType)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(corev1.ConditionStatus)
		**out = **in
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
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = (*in).DeepCopy()
	}
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSigningRequestConditionApplyConfiguration.
func (in *CertificateSigningRequestConditionApplyConfiguration) DeepCopy() *CertificateSigningRequestConditionApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(CertificateSigningRequestConditionApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSigningRequestSpecApplyConfiguration) DeepCopyInto(out *CertificateSigningRequestSpecApplyConfiguration) {
	*out = *in
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.SignerName != nil {
		in, out := &in.SignerName, &out.SignerName
		*out = new(string)
		**out = **in
	}
	if in.ExpirationSeconds != nil {
		in, out := &in.ExpirationSeconds, &out.ExpirationSeconds
		*out = new(int32)
		**out = **in
	}
	if in.Usages != nil {
		in, out := &in.Usages, &out.Usages
		*out = make([]certificatesv1beta1.KeyUsage, len(*in))
		copy(*out, *in)
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	if in.UID != nil {
		in, out := &in.UID, &out.UID
		*out = new(string)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extra != nil {
		in, out := &in.Extra, &out.Extra
		*out = make(map[string]certificatesv1beta1.ExtraValue, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(certificatesv1beta1.ExtraValue, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSigningRequestSpecApplyConfiguration.
func (in *CertificateSigningRequestSpecApplyConfiguration) DeepCopy() *CertificateSigningRequestSpecApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(CertificateSigningRequestSpecApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateSigningRequestStatusApplyConfiguration) DeepCopyInto(out *CertificateSigningRequestStatusApplyConfiguration) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CertificateSigningRequestConditionApplyConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateSigningRequestStatusApplyConfiguration.
func (in *CertificateSigningRequestStatusApplyConfiguration) DeepCopy() *CertificateSigningRequestStatusApplyConfiguration {
	if in == nil {
		return nil
	}
	out := new(CertificateSigningRequestStatusApplyConfiguration)
	in.DeepCopyInto(out)
	return out
}

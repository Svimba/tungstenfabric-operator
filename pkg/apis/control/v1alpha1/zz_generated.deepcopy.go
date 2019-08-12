// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecCtx) DeepCopyInto(out *SecCtx) {
	*out = *in
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]v1.Capability, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecCtx.
func (in *SecCtx) DeepCopy() *SecCtx {
	if in == nil {
		return nil
	}
	out := new(SecCtx)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFControl) DeepCopyInto(out *TFControl) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFControl.
func (in *TFControl) DeepCopy() *TFControl {
	if in == nil {
		return nil
	}
	out := new(TFControl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TFControl) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFControlControlSpec) DeepCopyInto(out *TFControlControlSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]Port, len(*in))
		copy(*out, *in)
	}
	if in.EnvList != nil {
		in, out := &in.EnvList, &out.EnvList
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFControlControlSpec.
func (in *TFControlControlSpec) DeepCopy() *TFControlControlSpec {
	if in == nil {
		return nil
	}
	out := new(TFControlControlSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFControlDnsSpec) DeepCopyInto(out *TFControlDnsSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]Port, len(*in))
		copy(*out, *in)
	}
	if in.EnvList != nil {
		in, out := &in.EnvList, &out.EnvList
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFControlDnsSpec.
func (in *TFControlDnsSpec) DeepCopy() *TFControlDnsSpec {
	if in == nil {
		return nil
	}
	out := new(TFControlDnsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFControlList) DeepCopyInto(out *TFControlList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TFControl, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFControlList.
func (in *TFControlList) DeepCopy() *TFControlList {
	if in == nil {
		return nil
	}
	out := new(TFControlList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TFControlList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFControlNamedSpec) DeepCopyInto(out *TFControlNamedSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]Port, len(*in))
		copy(*out, *in)
	}
	in.SecurityContext.DeepCopyInto(&out.SecurityContext)
	if in.EnvList != nil {
		in, out := &in.EnvList, &out.EnvList
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFControlNamedSpec.
func (in *TFControlNamedSpec) DeepCopy() *TFControlNamedSpec {
	if in == nil {
		return nil
	}
	out := new(TFControlNamedSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFControlSpec) DeepCopyInto(out *TFControlSpec) {
	*out = *in
	in.ControlSpec.DeepCopyInto(&out.ControlSpec)
	in.NamedSpec.DeepCopyInto(&out.NamedSpec)
	in.DnsSpec.DeepCopyInto(&out.DnsSpec)
	if in.ConfigMapList != nil {
		in, out := &in.ConfigMapList, &out.ConfigMapList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EnvList != nil {
		in, out := &in.EnvList, &out.EnvList
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFControlSpec.
func (in *TFControlSpec) DeepCopy() *TFControlSpec {
	if in == nil {
		return nil
	}
	out := new(TFControlSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFControlStatus) DeepCopyInto(out *TFControlStatus) {
	*out = *in
	if in.ConfigMapList != nil {
		in, out := &in.ConfigMapList, &out.ConfigMapList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFControlStatus.
func (in *TFControlStatus) DeepCopy() *TFControlStatus {
	if in == nil {
		return nil
	}
	out := new(TFControlStatus)
	in.DeepCopyInto(out)
	return out
}

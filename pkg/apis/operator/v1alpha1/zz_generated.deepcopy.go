// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	configv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/config/v1alpha1"
	controlv1alpha1 "github.com/Svimba/tungstenfabric-operator/pkg/apis/control/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

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
func (in *TFOperator) DeepCopyInto(out *TFOperator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFOperator.
func (in *TFOperator) DeepCopy() *TFOperator {
	if in == nil {
		return nil
	}
	out := new(TFOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TFOperator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFOperatorList) DeepCopyInto(out *TFOperatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TFOperator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFOperatorList.
func (in *TFOperatorList) DeepCopy() *TFOperatorList {
	if in == nil {
		return nil
	}
	out := new(TFOperatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TFOperatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFOperatorSpec) DeepCopyInto(out *TFOperatorSpec) {
	*out = *in
	if in.TFConfig != nil {
		in, out := &in.TFConfig, &out.TFConfig
		*out = new(configv1alpha1.TFConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TFControl != nil {
		in, out := &in.TFControl, &out.TFControl
		*out = new(controlv1alpha1.TFControlSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFOperatorSpec.
func (in *TFOperatorSpec) DeepCopy() *TFOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(TFOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFOperatorStatus) DeepCopyInto(out *TFOperatorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFOperatorStatus.
func (in *TFOperatorStatus) DeepCopy() *TFOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(TFOperatorStatus)
	in.DeepCopyInto(out)
	return out
}

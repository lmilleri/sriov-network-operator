// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetwork) DeepCopyInto(out *SriovNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetwork.
func (in *SriovNetwork) DeepCopy() *SriovNetwork {
	if in == nil {
		return nil
	}
	out := new(SriovNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkList) DeepCopyInto(out *SriovNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SriovNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkList.
func (in *SriovNetworkList) DeepCopy() *SriovNetworkList {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNodePolicy) DeepCopyInto(out *SriovNetworkNodePolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNodePolicy.
func (in *SriovNetworkNodePolicy) DeepCopy() *SriovNetworkNodePolicy {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNodePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovNetworkNodePolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNodePolicyList) DeepCopyInto(out *SriovNetworkNodePolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SriovNetworkNodePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNodePolicyList.
func (in *SriovNetworkNodePolicyList) DeepCopy() *SriovNetworkNodePolicyList {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNodePolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovNetworkNodePolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNodePolicySpec) DeepCopyInto(out *SriovNetworkNodePolicySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNodePolicySpec.
func (in *SriovNetworkNodePolicySpec) DeepCopy() *SriovNetworkNodePolicySpec {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNodePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkNodePolicyStatus) DeepCopyInto(out *SriovNetworkNodePolicyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkNodePolicyStatus.
func (in *SriovNetworkNodePolicyStatus) DeepCopy() *SriovNetworkNodePolicyStatus {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkNodePolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkSpec) DeepCopyInto(out *SriovNetworkSpec) {
	*out = *in
	out.IPAM = in.IPAM
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkSpec.
func (in *SriovNetworkSpec) DeepCopy() *SriovNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNetworkStatus) DeepCopyInto(out *SriovNetworkStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNetworkStatus.
func (in *SriovNetworkStatus) DeepCopy() *SriovNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(SriovNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNodeState) DeepCopyInto(out *SriovNodeState) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNodeState.
func (in *SriovNodeState) DeepCopy() *SriovNodeState {
	if in == nil {
		return nil
	}
	out := new(SriovNodeState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovNodeState) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNodeStateList) DeepCopyInto(out *SriovNodeStateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SriovNodeState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNodeStateList.
func (in *SriovNodeStateList) DeepCopy() *SriovNodeStateList {
	if in == nil {
		return nil
	}
	out := new(SriovNodeStateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SriovNodeStateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNodeStateSpec) DeepCopyInto(out *SriovNodeStateSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNodeStateSpec.
func (in *SriovNodeStateSpec) DeepCopy() *SriovNodeStateSpec {
	if in == nil {
		return nil
	}
	out := new(SriovNodeStateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SriovNodeStateStatus) DeepCopyInto(out *SriovNodeStateStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SriovNodeStateStatus.
func (in *SriovNodeStateStatus) DeepCopy() *SriovNodeStateStatus {
	if in == nil {
		return nil
	}
	out := new(SriovNodeStateStatus)
	in.DeepCopyInto(out)
	return out
}

// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package tenancyv1alpha1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using Namespace within kubernetes types, where deepcopy-gen is used.
func (in *Namespace) DeepCopyInto(out *Namespace) {
	p := proto.Clone(in).(*Namespace)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Namespace. Required by controller-gen.
func (in *Namespace) DeepCopy() *Namespace {
	if in == nil {
		return nil
	}
	out := new(Namespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Namespace. Required by controller-gen.
func (in *Namespace) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

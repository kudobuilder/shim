// +build !ignore_autogenerated

/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KUDOOperator) DeepCopyInto(out *KUDOOperator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KUDOOperator.
func (in *KUDOOperator) DeepCopy() *KUDOOperator {
	if in == nil {
		return nil
	}
	out := new(KUDOOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShimInstance) DeepCopyInto(out *ShimInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShimInstance.
func (in *ShimInstance) DeepCopy() *ShimInstance {
	if in == nil {
		return nil
	}
	out := new(ShimInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShimInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShimInstanceList) DeepCopyInto(out *ShimInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ShimInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShimInstanceList.
func (in *ShimInstanceList) DeepCopy() *ShimInstanceList {
	if in == nil {
		return nil
	}
	out := new(ShimInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShimInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShimInstanceSpec) DeepCopyInto(out *ShimInstanceSpec) {
	*out = *in
	out.KUDOOperator = in.KUDOOperator
	in.CRDSpec.DeepCopyInto(&out.CRDSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShimInstanceSpec.
func (in *ShimInstanceSpec) DeepCopy() *ShimInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(ShimInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShimInstanceStatus) DeepCopyInto(out *ShimInstanceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShimInstanceStatus.
func (in *ShimInstanceStatus) DeepCopy() *ShimInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(ShimInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

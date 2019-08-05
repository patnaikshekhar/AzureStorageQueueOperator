// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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
func (in *AzureConnectionStringSecretRef) DeepCopyInto(out *AzureConnectionStringSecretRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureConnectionStringSecretRef.
func (in *AzureConnectionStringSecretRef) DeepCopy() *AzureConnectionStringSecretRef {
	if in == nil {
		return nil
	}
	out := new(AzureConnectionStringSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureQueue) DeepCopyInto(out *AzureQueue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Status = in.Status
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureQueue.
func (in *AzureQueue) DeepCopy() *AzureQueue {
	if in == nil {
		return nil
	}
	out := new(AzureQueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureQueue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureQueueList) DeepCopyInto(out *AzureQueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureQueue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureQueueList.
func (in *AzureQueueList) DeepCopy() *AzureQueueList {
	if in == nil {
		return nil
	}
	out := new(AzureQueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureQueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureQueueSpec) DeepCopyInto(out *AzureQueueSpec) {
	*out = *in
	out.ConnectionStringSecretRef = in.ConnectionStringSecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureQueueSpec.
func (in *AzureQueueSpec) DeepCopy() *AzureQueueSpec {
	if in == nil {
		return nil
	}
	out := new(AzureQueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureQueueStatus) DeepCopyInto(out *AzureQueueStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureQueueStatus.
func (in *AzureQueueStatus) DeepCopy() *AzureQueueStatus {
	if in == nil {
		return nil
	}
	out := new(AzureQueueStatus)
	in.DeepCopyInto(out)
	return out
}

// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainService) DeepCopyInto(out *ActiveDirectoryDomainService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainService.
func (in *ActiveDirectoryDomainService) DeepCopy() *ActiveDirectoryDomainService {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveDirectoryDomainService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceList) DeepCopyInto(out *ActiveDirectoryDomainServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActiveDirectoryDomainService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceList.
func (in *ActiveDirectoryDomainServiceList) DeepCopy() *ActiveDirectoryDomainServiceList {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveDirectoryDomainServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceObservation) DeepCopyInto(out *ActiveDirectoryDomainServiceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceObservation.
func (in *ActiveDirectoryDomainServiceObservation) DeepCopy() *ActiveDirectoryDomainServiceObservation {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceParameters) DeepCopyInto(out *ActiveDirectoryDomainServiceParameters) {
	*out = *in
	if in.FilteredSyncEnabled != nil {
		in, out := &in.FilteredSyncEnabled, &out.FilteredSyncEnabled
		*out = new(bool)
		**out = **in
	}
	if in.InitialReplicaSet != nil {
		in, out := &in.InitialReplicaSet, &out.InitialReplicaSet
		*out = make([]InitialReplicaSetParameters, len(*in))
		copy(*out, *in)
	}
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = make([]NotificationsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecureLdap != nil {
		in, out := &in.SecureLdap, &out.SecureLdap
		*out = make([]SecureLdapParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = make([]SecurityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceParameters.
func (in *ActiveDirectoryDomainServiceParameters) DeepCopy() *ActiveDirectoryDomainServiceParameters {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceReplicaSet) DeepCopyInto(out *ActiveDirectoryDomainServiceReplicaSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceReplicaSet.
func (in *ActiveDirectoryDomainServiceReplicaSet) DeepCopy() *ActiveDirectoryDomainServiceReplicaSet {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceReplicaSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveDirectoryDomainServiceReplicaSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceReplicaSetList) DeepCopyInto(out *ActiveDirectoryDomainServiceReplicaSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActiveDirectoryDomainServiceReplicaSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceReplicaSetList.
func (in *ActiveDirectoryDomainServiceReplicaSetList) DeepCopy() *ActiveDirectoryDomainServiceReplicaSetList {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceReplicaSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveDirectoryDomainServiceReplicaSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceReplicaSetObservation) DeepCopyInto(out *ActiveDirectoryDomainServiceReplicaSetObservation) {
	*out = *in
	if in.DomainControllerIPAddresses != nil {
		in, out := &in.DomainControllerIPAddresses, &out.DomainControllerIPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceReplicaSetObservation.
func (in *ActiveDirectoryDomainServiceReplicaSetObservation) DeepCopy() *ActiveDirectoryDomainServiceReplicaSetObservation {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceReplicaSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceReplicaSetParameters) DeepCopyInto(out *ActiveDirectoryDomainServiceReplicaSetParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceReplicaSetParameters.
func (in *ActiveDirectoryDomainServiceReplicaSetParameters) DeepCopy() *ActiveDirectoryDomainServiceReplicaSetParameters {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceReplicaSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceReplicaSetSpec) DeepCopyInto(out *ActiveDirectoryDomainServiceReplicaSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	out.ForProvider = in.ForProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceReplicaSetSpec.
func (in *ActiveDirectoryDomainServiceReplicaSetSpec) DeepCopy() *ActiveDirectoryDomainServiceReplicaSetSpec {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceReplicaSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceReplicaSetStatus) DeepCopyInto(out *ActiveDirectoryDomainServiceReplicaSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceReplicaSetStatus.
func (in *ActiveDirectoryDomainServiceReplicaSetStatus) DeepCopy() *ActiveDirectoryDomainServiceReplicaSetStatus {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceReplicaSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceSpec) DeepCopyInto(out *ActiveDirectoryDomainServiceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceSpec.
func (in *ActiveDirectoryDomainServiceSpec) DeepCopy() *ActiveDirectoryDomainServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryDomainServiceStatus) DeepCopyInto(out *ActiveDirectoryDomainServiceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryDomainServiceStatus.
func (in *ActiveDirectoryDomainServiceStatus) DeepCopy() *ActiveDirectoryDomainServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryDomainServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialReplicaSetObservation) DeepCopyInto(out *InitialReplicaSetObservation) {
	*out = *in
	if in.DomainControllerIPAddresses != nil {
		in, out := &in.DomainControllerIPAddresses, &out.DomainControllerIPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialReplicaSetObservation.
func (in *InitialReplicaSetObservation) DeepCopy() *InitialReplicaSetObservation {
	if in == nil {
		return nil
	}
	out := new(InitialReplicaSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InitialReplicaSetParameters) DeepCopyInto(out *InitialReplicaSetParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InitialReplicaSetParameters.
func (in *InitialReplicaSetParameters) DeepCopy() *InitialReplicaSetParameters {
	if in == nil {
		return nil
	}
	out := new(InitialReplicaSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationsObservation) DeepCopyInto(out *NotificationsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationsObservation.
func (in *NotificationsObservation) DeepCopy() *NotificationsObservation {
	if in == nil {
		return nil
	}
	out := new(NotificationsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotificationsParameters) DeepCopyInto(out *NotificationsParameters) {
	*out = *in
	if in.AdditionalRecipients != nil {
		in, out := &in.AdditionalRecipients, &out.AdditionalRecipients
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotifyDcAdmins != nil {
		in, out := &in.NotifyDcAdmins, &out.NotifyDcAdmins
		*out = new(bool)
		**out = **in
	}
	if in.NotifyGlobalAdmins != nil {
		in, out := &in.NotifyGlobalAdmins, &out.NotifyGlobalAdmins
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotificationsParameters.
func (in *NotificationsParameters) DeepCopy() *NotificationsParameters {
	if in == nil {
		return nil
	}
	out := new(NotificationsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureLdapObservation) DeepCopyInto(out *SecureLdapObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureLdapObservation.
func (in *SecureLdapObservation) DeepCopy() *SecureLdapObservation {
	if in == nil {
		return nil
	}
	out := new(SecureLdapObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecureLdapParameters) DeepCopyInto(out *SecureLdapParameters) {
	*out = *in
	if in.ExternalAccessEnabled != nil {
		in, out := &in.ExternalAccessEnabled, &out.ExternalAccessEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecureLdapParameters.
func (in *SecureLdapParameters) DeepCopy() *SecureLdapParameters {
	if in == nil {
		return nil
	}
	out := new(SecureLdapParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityObservation) DeepCopyInto(out *SecurityObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityObservation.
func (in *SecurityObservation) DeepCopy() *SecurityObservation {
	if in == nil {
		return nil
	}
	out := new(SecurityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityParameters) DeepCopyInto(out *SecurityParameters) {
	*out = *in
	if in.NtlmV1Enabled != nil {
		in, out := &in.NtlmV1Enabled, &out.NtlmV1Enabled
		*out = new(bool)
		**out = **in
	}
	if in.SyncKerberosPasswords != nil {
		in, out := &in.SyncKerberosPasswords, &out.SyncKerberosPasswords
		*out = new(bool)
		**out = **in
	}
	if in.SyncNtlmPasswords != nil {
		in, out := &in.SyncNtlmPasswords, &out.SyncNtlmPasswords
		*out = new(bool)
		**out = **in
	}
	if in.SyncOnPremPasswords != nil {
		in, out := &in.SyncOnPremPasswords, &out.SyncOnPremPasswords
		*out = new(bool)
		**out = **in
	}
	if in.TLSV1Enabled != nil {
		in, out := &in.TLSV1Enabled, &out.TLSV1Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityParameters.
func (in *SecurityParameters) DeepCopy() *SecurityParameters {
	if in == nil {
		return nil
	}
	out := new(SecurityParameters)
	in.DeepCopyInto(out)
	return out
}

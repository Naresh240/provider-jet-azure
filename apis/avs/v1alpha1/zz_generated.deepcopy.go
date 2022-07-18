//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Crossplane Authors.

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
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CircuitObservation) DeepCopyInto(out *CircuitObservation) {
	*out = *in
	if in.ExpressRouteID != nil {
		in, out := &in.ExpressRouteID, &out.ExpressRouteID
		*out = new(string)
		**out = **in
	}
	if in.ExpressRoutePrivatePeeringID != nil {
		in, out := &in.ExpressRoutePrivatePeeringID, &out.ExpressRoutePrivatePeeringID
		*out = new(string)
		**out = **in
	}
	if in.PrimarySubnetCidr != nil {
		in, out := &in.PrimarySubnetCidr, &out.PrimarySubnetCidr
		*out = new(string)
		**out = **in
	}
	if in.SecondarySubnetCidr != nil {
		in, out := &in.SecondarySubnetCidr, &out.SecondarySubnetCidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CircuitObservation.
func (in *CircuitObservation) DeepCopy() *CircuitObservation {
	if in == nil {
		return nil
	}
	out := new(CircuitObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CircuitParameters) DeepCopyInto(out *CircuitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CircuitParameters.
func (in *CircuitParameters) DeepCopy() *CircuitParameters {
	if in == nil {
		return nil
	}
	out := new(CircuitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementClusterObservation) DeepCopyInto(out *ManagementClusterObservation) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementClusterObservation.
func (in *ManagementClusterObservation) DeepCopy() *ManagementClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ManagementClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementClusterParameters) DeepCopyInto(out *ManagementClusterParameters) {
	*out = *in
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementClusterParameters.
func (in *ManagementClusterParameters) DeepCopy() *ManagementClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ManagementClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwareCluster) DeepCopyInto(out *VMwareCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwareCluster.
func (in *VMwareCluster) DeepCopy() *VMwareCluster {
	if in == nil {
		return nil
	}
	out := new(VMwareCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMwareCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwareClusterList) DeepCopyInto(out *VMwareClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VMwareCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwareClusterList.
func (in *VMwareClusterList) DeepCopy() *VMwareClusterList {
	if in == nil {
		return nil
	}
	out := new(VMwareClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMwareClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwareClusterObservation) DeepCopyInto(out *VMwareClusterObservation) {
	*out = *in
	if in.ClusterNumber != nil {
		in, out := &in.ClusterNumber, &out.ClusterNumber
		*out = new(float64)
		**out = **in
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwareClusterObservation.
func (in *VMwareClusterObservation) DeepCopy() *VMwareClusterObservation {
	if in == nil {
		return nil
	}
	out := new(VMwareClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwareClusterParameters) DeepCopyInto(out *VMwareClusterParameters) {
	*out = *in
	if in.ClusterNodeCount != nil {
		in, out := &in.ClusterNodeCount, &out.ClusterNodeCount
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.VMwareCloudID != nil {
		in, out := &in.VMwareCloudID, &out.VMwareCloudID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwareClusterParameters.
func (in *VMwareClusterParameters) DeepCopy() *VMwareClusterParameters {
	if in == nil {
		return nil
	}
	out := new(VMwareClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwareClusterSpec) DeepCopyInto(out *VMwareClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwareClusterSpec.
func (in *VMwareClusterSpec) DeepCopy() *VMwareClusterSpec {
	if in == nil {
		return nil
	}
	out := new(VMwareClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwareClusterStatus) DeepCopyInto(out *VMwareClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwareClusterStatus.
func (in *VMwareClusterStatus) DeepCopy() *VMwareClusterStatus {
	if in == nil {
		return nil
	}
	out := new(VMwareClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwareExpressRouteAuthorization) DeepCopyInto(out *VMwareExpressRouteAuthorization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwareExpressRouteAuthorization.
func (in *VMwareExpressRouteAuthorization) DeepCopy() *VMwareExpressRouteAuthorization {
	if in == nil {
		return nil
	}
	out := new(VMwareExpressRouteAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMwareExpressRouteAuthorization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwareExpressRouteAuthorizationList) DeepCopyInto(out *VMwareExpressRouteAuthorizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VMwareExpressRouteAuthorization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwareExpressRouteAuthorizationList.
func (in *VMwareExpressRouteAuthorizationList) DeepCopy() *VMwareExpressRouteAuthorizationList {
	if in == nil {
		return nil
	}
	out := new(VMwareExpressRouteAuthorizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMwareExpressRouteAuthorizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwareExpressRouteAuthorizationObservation) DeepCopyInto(out *VMwareExpressRouteAuthorizationObservation) {
	*out = *in
	if in.ExpressRouteAuthorizationID != nil {
		in, out := &in.ExpressRouteAuthorizationID, &out.ExpressRouteAuthorizationID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwareExpressRouteAuthorizationObservation.
func (in *VMwareExpressRouteAuthorizationObservation) DeepCopy() *VMwareExpressRouteAuthorizationObservation {
	if in == nil {
		return nil
	}
	out := new(VMwareExpressRouteAuthorizationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwareExpressRouteAuthorizationParameters) DeepCopyInto(out *VMwareExpressRouteAuthorizationParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PrivateCloudID != nil {
		in, out := &in.PrivateCloudID, &out.PrivateCloudID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwareExpressRouteAuthorizationParameters.
func (in *VMwareExpressRouteAuthorizationParameters) DeepCopy() *VMwareExpressRouteAuthorizationParameters {
	if in == nil {
		return nil
	}
	out := new(VMwareExpressRouteAuthorizationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwareExpressRouteAuthorizationSpec) DeepCopyInto(out *VMwareExpressRouteAuthorizationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwareExpressRouteAuthorizationSpec.
func (in *VMwareExpressRouteAuthorizationSpec) DeepCopy() *VMwareExpressRouteAuthorizationSpec {
	if in == nil {
		return nil
	}
	out := new(VMwareExpressRouteAuthorizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwareExpressRouteAuthorizationStatus) DeepCopyInto(out *VMwareExpressRouteAuthorizationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwareExpressRouteAuthorizationStatus.
func (in *VMwareExpressRouteAuthorizationStatus) DeepCopy() *VMwareExpressRouteAuthorizationStatus {
	if in == nil {
		return nil
	}
	out := new(VMwareExpressRouteAuthorizationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwarePrivateCloud) DeepCopyInto(out *VMwarePrivateCloud) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwarePrivateCloud.
func (in *VMwarePrivateCloud) DeepCopy() *VMwarePrivateCloud {
	if in == nil {
		return nil
	}
	out := new(VMwarePrivateCloud)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMwarePrivateCloud) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwarePrivateCloudList) DeepCopyInto(out *VMwarePrivateCloudList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VMwarePrivateCloud, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwarePrivateCloudList.
func (in *VMwarePrivateCloudList) DeepCopy() *VMwarePrivateCloudList {
	if in == nil {
		return nil
	}
	out := new(VMwarePrivateCloudList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VMwarePrivateCloudList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwarePrivateCloudObservation) DeepCopyInto(out *VMwarePrivateCloudObservation) {
	*out = *in
	if in.Circuit != nil {
		in, out := &in.Circuit, &out.Circuit
		*out = make([]CircuitObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HcxCloudManagerEndpoint != nil {
		in, out := &in.HcxCloudManagerEndpoint, &out.HcxCloudManagerEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ManagementCluster != nil {
		in, out := &in.ManagementCluster, &out.ManagementCluster
		*out = make([]ManagementClusterObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ManagementSubnetCidr != nil {
		in, out := &in.ManagementSubnetCidr, &out.ManagementSubnetCidr
		*out = new(string)
		**out = **in
	}
	if in.NsxtCertificateThumbprint != nil {
		in, out := &in.NsxtCertificateThumbprint, &out.NsxtCertificateThumbprint
		*out = new(string)
		**out = **in
	}
	if in.NsxtManagerEndpoint != nil {
		in, out := &in.NsxtManagerEndpoint, &out.NsxtManagerEndpoint
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningSubnetCidr != nil {
		in, out := &in.ProvisioningSubnetCidr, &out.ProvisioningSubnetCidr
		*out = new(string)
		**out = **in
	}
	if in.VcenterCertificateThumbprint != nil {
		in, out := &in.VcenterCertificateThumbprint, &out.VcenterCertificateThumbprint
		*out = new(string)
		**out = **in
	}
	if in.VcsaEndpoint != nil {
		in, out := &in.VcsaEndpoint, &out.VcsaEndpoint
		*out = new(string)
		**out = **in
	}
	if in.VmotionSubnetCidr != nil {
		in, out := &in.VmotionSubnetCidr, &out.VmotionSubnetCidr
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwarePrivateCloudObservation.
func (in *VMwarePrivateCloudObservation) DeepCopy() *VMwarePrivateCloudObservation {
	if in == nil {
		return nil
	}
	out := new(VMwarePrivateCloudObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwarePrivateCloudParameters) DeepCopyInto(out *VMwarePrivateCloudParameters) {
	*out = *in
	if in.InternetConnectionEnabled != nil {
		in, out := &in.InternetConnectionEnabled, &out.InternetConnectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ManagementCluster != nil {
		in, out := &in.ManagementCluster, &out.ManagementCluster
		*out = make([]ManagementClusterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkSubnetCidr != nil {
		in, out := &in.NetworkSubnetCidr, &out.NetworkSubnetCidr
		*out = new(string)
		**out = **in
	}
	if in.NsxtPasswordSecretRef != nil {
		in, out := &in.NsxtPasswordSecretRef, &out.NsxtPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VcenterPasswordSecretRef != nil {
		in, out := &in.VcenterPasswordSecretRef, &out.VcenterPasswordSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwarePrivateCloudParameters.
func (in *VMwarePrivateCloudParameters) DeepCopy() *VMwarePrivateCloudParameters {
	if in == nil {
		return nil
	}
	out := new(VMwarePrivateCloudParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwarePrivateCloudSpec) DeepCopyInto(out *VMwarePrivateCloudSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwarePrivateCloudSpec.
func (in *VMwarePrivateCloudSpec) DeepCopy() *VMwarePrivateCloudSpec {
	if in == nil {
		return nil
	}
	out := new(VMwarePrivateCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMwarePrivateCloudStatus) DeepCopyInto(out *VMwarePrivateCloudStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMwarePrivateCloudStatus.
func (in *VMwarePrivateCloudStatus) DeepCopy() *VMwarePrivateCloudStatus {
	if in == nil {
		return nil
	}
	out := new(VMwarePrivateCloudStatus)
	in.DeepCopyInto(out)
	return out
}
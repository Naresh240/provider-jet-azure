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
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
	*out = *in
	if in.ClusterNumber != nil {
		in, out := &in.ClusterNumber, &out.ClusterNumber
		*out = new(int64)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
	*out = *in
	if in.ClusterNodeCount != nil {
		in, out := &in.ClusterNodeCount, &out.ClusterNodeCount
		*out = new(int64)
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
	if in.VmwareCloudID != nil {
		in, out := &in.VmwareCloudID, &out.VmwareCloudID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpressRouteAuthorization) DeepCopyInto(out *ExpressRouteAuthorization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpressRouteAuthorization.
func (in *ExpressRouteAuthorization) DeepCopy() *ExpressRouteAuthorization {
	if in == nil {
		return nil
	}
	out := new(ExpressRouteAuthorization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExpressRouteAuthorization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpressRouteAuthorizationList) DeepCopyInto(out *ExpressRouteAuthorizationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExpressRouteAuthorization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpressRouteAuthorizationList.
func (in *ExpressRouteAuthorizationList) DeepCopy() *ExpressRouteAuthorizationList {
	if in == nil {
		return nil
	}
	out := new(ExpressRouteAuthorizationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExpressRouteAuthorizationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpressRouteAuthorizationObservation) DeepCopyInto(out *ExpressRouteAuthorizationObservation) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpressRouteAuthorizationObservation.
func (in *ExpressRouteAuthorizationObservation) DeepCopy() *ExpressRouteAuthorizationObservation {
	if in == nil {
		return nil
	}
	out := new(ExpressRouteAuthorizationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpressRouteAuthorizationParameters) DeepCopyInto(out *ExpressRouteAuthorizationParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpressRouteAuthorizationParameters.
func (in *ExpressRouteAuthorizationParameters) DeepCopy() *ExpressRouteAuthorizationParameters {
	if in == nil {
		return nil
	}
	out := new(ExpressRouteAuthorizationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpressRouteAuthorizationSpec) DeepCopyInto(out *ExpressRouteAuthorizationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpressRouteAuthorizationSpec.
func (in *ExpressRouteAuthorizationSpec) DeepCopy() *ExpressRouteAuthorizationSpec {
	if in == nil {
		return nil
	}
	out := new(ExpressRouteAuthorizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExpressRouteAuthorizationStatus) DeepCopyInto(out *ExpressRouteAuthorizationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExpressRouteAuthorizationStatus.
func (in *ExpressRouteAuthorizationStatus) DeepCopy() *ExpressRouteAuthorizationStatus {
	if in == nil {
		return nil
	}
	out := new(ExpressRouteAuthorizationStatus)
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
		*out = new(int64)
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
		*out = new(int64)
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
func (in *PrivateCloud) DeepCopyInto(out *PrivateCloud) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCloud.
func (in *PrivateCloud) DeepCopy() *PrivateCloud {
	if in == nil {
		return nil
	}
	out := new(PrivateCloud)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateCloud) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCloudList) DeepCopyInto(out *PrivateCloudList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrivateCloud, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCloudList.
func (in *PrivateCloudList) DeepCopy() *PrivateCloudList {
	if in == nil {
		return nil
	}
	out := new(PrivateCloudList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateCloudList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCloudObservation) DeepCopyInto(out *PrivateCloudObservation) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCloudObservation.
func (in *PrivateCloudObservation) DeepCopy() *PrivateCloudObservation {
	if in == nil {
		return nil
	}
	out := new(PrivateCloudObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCloudParameters) DeepCopyInto(out *PrivateCloudParameters) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCloudParameters.
func (in *PrivateCloudParameters) DeepCopy() *PrivateCloudParameters {
	if in == nil {
		return nil
	}
	out := new(PrivateCloudParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCloudSpec) DeepCopyInto(out *PrivateCloudSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCloudSpec.
func (in *PrivateCloudSpec) DeepCopy() *PrivateCloudSpec {
	if in == nil {
		return nil
	}
	out := new(PrivateCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateCloudStatus) DeepCopyInto(out *PrivateCloudStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateCloudStatus.
func (in *PrivateCloudStatus) DeepCopy() *PrivateCloudStatus {
	if in == nil {
		return nil
	}
	out := new(PrivateCloudStatus)
	in.DeepCopyInto(out)
	return out
}
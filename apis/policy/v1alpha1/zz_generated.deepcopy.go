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
func (in *Assignment) DeepCopyInto(out *Assignment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Assignment.
func (in *Assignment) DeepCopy() *Assignment {
	if in == nil {
		return nil
	}
	out := new(Assignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Assignment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssignmentList) DeepCopyInto(out *AssignmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Assignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssignmentList.
func (in *AssignmentList) DeepCopy() *AssignmentList {
	if in == nil {
		return nil
	}
	out := new(AssignmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AssignmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssignmentObservation) DeepCopyInto(out *AssignmentObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssignmentObservation.
func (in *AssignmentObservation) DeepCopy() *AssignmentObservation {
	if in == nil {
		return nil
	}
	out := new(AssignmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssignmentParameters) DeepCopyInto(out *AssignmentParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EnforcementMode != nil {
		in, out := &in.EnforcementMode, &out.EnforcementMode
		*out = new(bool)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = make([]IdentityParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NotScopes != nil {
		in, out := &in.NotScopes, &out.NotScopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(string)
		**out = **in
	}
	if in.PolicyDefinitionID != nil {
		in, out := &in.PolicyDefinitionID, &out.PolicyDefinitionID
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssignmentParameters.
func (in *AssignmentParameters) DeepCopy() *AssignmentParameters {
	if in == nil {
		return nil
	}
	out := new(AssignmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssignmentSpec) DeepCopyInto(out *AssignmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssignmentSpec.
func (in *AssignmentSpec) DeepCopy() *AssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(AssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AssignmentStatus) DeepCopyInto(out *AssignmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AssignmentStatus.
func (in *AssignmentStatus) DeepCopy() *AssignmentStatus {
	if in == nil {
		return nil
	}
	out := new(AssignmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationObservation) DeepCopyInto(out *ConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationObservation.
func (in *ConfigurationObservation) DeepCopy() *ConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationParameters) DeepCopyInto(out *ConfigurationParameters) {
	*out = *in
	if in.AssignmentType != nil {
		in, out := &in.AssignmentType, &out.AssignmentType
		*out = new(string)
		**out = **in
	}
	if in.ContentHash != nil {
		in, out := &in.ContentHash, &out.ContentHash
		*out = new(string)
		**out = **in
	}
	if in.ContentURI != nil {
		in, out := &in.ContentURI, &out.ContentURI
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = make([]ParameterParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationParameters.
func (in *ConfigurationParameters) DeepCopy() *ConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Definition) DeepCopyInto(out *Definition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Definition.
func (in *Definition) DeepCopy() *Definition {
	if in == nil {
		return nil
	}
	out := new(Definition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Definition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionList) DeepCopyInto(out *DefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Definition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionList.
func (in *DefinitionList) DeepCopy() *DefinitionList {
	if in == nil {
		return nil
	}
	out := new(DefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionObservation) DeepCopyInto(out *DefinitionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionObservation.
func (in *DefinitionObservation) DeepCopy() *DefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(DefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionParameters) DeepCopyInto(out *DefinitionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ManagementGroupID != nil {
		in, out := &in.ManagementGroupID, &out.ManagementGroupID
		*out = new(string)
		**out = **in
	}
	if in.ManagementGroupName != nil {
		in, out := &in.ManagementGroupName, &out.ManagementGroupName
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(string)
		**out = **in
	}
	if in.PolicyRule != nil {
		in, out := &in.PolicyRule, &out.PolicyRule
		*out = new(string)
		**out = **in
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionParameters.
func (in *DefinitionParameters) DeepCopy() *DefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(DefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionSpec) DeepCopyInto(out *DefinitionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionSpec.
func (in *DefinitionSpec) DeepCopy() *DefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(DefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefinitionStatus) DeepCopyInto(out *DefinitionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefinitionStatus.
func (in *DefinitionStatus) DeepCopy() *DefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(DefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityObservation) DeepCopyInto(out *IdentityObservation) {
	*out = *in
	if in.PrincipalID != nil {
		in, out := &in.PrincipalID, &out.PrincipalID
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityObservation.
func (in *IdentityObservation) DeepCopy() *IdentityObservation {
	if in == nil {
		return nil
	}
	out := new(IdentityObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityParameters) DeepCopyInto(out *IdentityParameters) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityParameters.
func (in *IdentityParameters) DeepCopy() *IdentityParameters {
	if in == nil {
		return nil
	}
	out := new(IdentityParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterObservation) DeepCopyInto(out *ParameterObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterObservation.
func (in *ParameterObservation) DeepCopy() *ParameterObservation {
	if in == nil {
		return nil
	}
	out := new(ParameterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterParameters) DeepCopyInto(out *ParameterParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterParameters.
func (in *ParameterParameters) DeepCopy() *ParameterParameters {
	if in == nil {
		return nil
	}
	out := new(ParameterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionGroupObservation) DeepCopyInto(out *PolicyDefinitionGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionGroupObservation.
func (in *PolicyDefinitionGroupObservation) DeepCopy() *PolicyDefinitionGroupObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionGroupParameters) DeepCopyInto(out *PolicyDefinitionGroupParameters) {
	*out = *in
	if in.AdditionalMetadataResourceID != nil {
		in, out := &in.AdditionalMetadataResourceID, &out.AdditionalMetadataResourceID
		*out = new(string)
		**out = **in
	}
	if in.Category != nil {
		in, out := &in.Category, &out.Category
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionGroupParameters.
func (in *PolicyDefinitionGroupParameters) DeepCopy() *PolicyDefinitionGroupParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionReferenceObservation) DeepCopyInto(out *PolicyDefinitionReferenceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionReferenceObservation.
func (in *PolicyDefinitionReferenceObservation) DeepCopy() *PolicyDefinitionReferenceObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionReferenceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyDefinitionReferenceParameters) DeepCopyInto(out *PolicyDefinitionReferenceParameters) {
	*out = *in
	if in.ParameterValues != nil {
		in, out := &in.ParameterValues, &out.ParameterValues
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
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
	if in.PolicyDefinitionID != nil {
		in, out := &in.PolicyDefinitionID, &out.PolicyDefinitionID
		*out = new(string)
		**out = **in
	}
	if in.PolicyGroupNames != nil {
		in, out := &in.PolicyGroupNames, &out.PolicyGroupNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ReferenceID != nil {
		in, out := &in.ReferenceID, &out.ReferenceID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyDefinitionReferenceParameters.
func (in *PolicyDefinitionReferenceParameters) DeepCopy() *PolicyDefinitionReferenceParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyDefinitionReferenceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Remediation) DeepCopyInto(out *Remediation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Remediation.
func (in *Remediation) DeepCopy() *Remediation {
	if in == nil {
		return nil
	}
	out := new(Remediation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Remediation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemediationList) DeepCopyInto(out *RemediationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Remediation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemediationList.
func (in *RemediationList) DeepCopy() *RemediationList {
	if in == nil {
		return nil
	}
	out := new(RemediationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemediationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemediationObservation) DeepCopyInto(out *RemediationObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemediationObservation.
func (in *RemediationObservation) DeepCopy() *RemediationObservation {
	if in == nil {
		return nil
	}
	out := new(RemediationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemediationParameters) DeepCopyInto(out *RemediationParameters) {
	*out = *in
	if in.LocationFilters != nil {
		in, out := &in.LocationFilters, &out.LocationFilters
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PolicyAssignmentID != nil {
		in, out := &in.PolicyAssignmentID, &out.PolicyAssignmentID
		*out = new(string)
		**out = **in
	}
	if in.PolicyDefinitionReferenceID != nil {
		in, out := &in.PolicyDefinitionReferenceID, &out.PolicyDefinitionReferenceID
		*out = new(string)
		**out = **in
	}
	if in.ResourceDiscoveryMode != nil {
		in, out := &in.ResourceDiscoveryMode, &out.ResourceDiscoveryMode
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemediationParameters.
func (in *RemediationParameters) DeepCopy() *RemediationParameters {
	if in == nil {
		return nil
	}
	out := new(RemediationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemediationSpec) DeepCopyInto(out *RemediationSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemediationSpec.
func (in *RemediationSpec) DeepCopy() *RemediationSpec {
	if in == nil {
		return nil
	}
	out := new(RemediationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemediationStatus) DeepCopyInto(out *RemediationStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemediationStatus.
func (in *RemediationStatus) DeepCopy() *RemediationStatus {
	if in == nil {
		return nil
	}
	out := new(RemediationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SetDefinition) DeepCopyInto(out *SetDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SetDefinition.
func (in *SetDefinition) DeepCopy() *SetDefinition {
	if in == nil {
		return nil
	}
	out := new(SetDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SetDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SetDefinitionList) DeepCopyInto(out *SetDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SetDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SetDefinitionList.
func (in *SetDefinitionList) DeepCopy() *SetDefinitionList {
	if in == nil {
		return nil
	}
	out := new(SetDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SetDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SetDefinitionObservation) DeepCopyInto(out *SetDefinitionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SetDefinitionObservation.
func (in *SetDefinitionObservation) DeepCopy() *SetDefinitionObservation {
	if in == nil {
		return nil
	}
	out := new(SetDefinitionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SetDefinitionParameters) DeepCopyInto(out *SetDefinitionParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ManagementGroupID != nil {
		in, out := &in.ManagementGroupID, &out.ManagementGroupID
		*out = new(string)
		**out = **in
	}
	if in.ManagementGroupName != nil {
		in, out := &in.ManagementGroupName, &out.ManagementGroupName
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(string)
		**out = **in
	}
	if in.PolicyDefinitionGroup != nil {
		in, out := &in.PolicyDefinitionGroup, &out.PolicyDefinitionGroup
		*out = make([]PolicyDefinitionGroupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PolicyDefinitionReference != nil {
		in, out := &in.PolicyDefinitionReference, &out.PolicyDefinitionReference
		*out = make([]PolicyDefinitionReferenceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PolicyDefinitions != nil {
		in, out := &in.PolicyDefinitions, &out.PolicyDefinitions
		*out = new(string)
		**out = **in
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SetDefinitionParameters.
func (in *SetDefinitionParameters) DeepCopy() *SetDefinitionParameters {
	if in == nil {
		return nil
	}
	out := new(SetDefinitionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SetDefinitionSpec) DeepCopyInto(out *SetDefinitionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SetDefinitionSpec.
func (in *SetDefinitionSpec) DeepCopy() *SetDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(SetDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SetDefinitionStatus) DeepCopyInto(out *SetDefinitionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SetDefinitionStatus.
func (in *SetDefinitionStatus) DeepCopy() *SetDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(SetDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineConfigurationAssignment) DeepCopyInto(out *VirtualMachineConfigurationAssignment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineConfigurationAssignment.
func (in *VirtualMachineConfigurationAssignment) DeepCopy() *VirtualMachineConfigurationAssignment {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineConfigurationAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineConfigurationAssignment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineConfigurationAssignmentList) DeepCopyInto(out *VirtualMachineConfigurationAssignmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineConfigurationAssignment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineConfigurationAssignmentList.
func (in *VirtualMachineConfigurationAssignmentList) DeepCopy() *VirtualMachineConfigurationAssignmentList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineConfigurationAssignmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineConfigurationAssignmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineConfigurationAssignmentObservation) DeepCopyInto(out *VirtualMachineConfigurationAssignmentObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineConfigurationAssignmentObservation.
func (in *VirtualMachineConfigurationAssignmentObservation) DeepCopy() *VirtualMachineConfigurationAssignmentObservation {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineConfigurationAssignmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineConfigurationAssignmentParameters) DeepCopyInto(out *VirtualMachineConfigurationAssignmentParameters) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]ConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.VirtualMachineID != nil {
		in, out := &in.VirtualMachineID, &out.VirtualMachineID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineConfigurationAssignmentParameters.
func (in *VirtualMachineConfigurationAssignmentParameters) DeepCopy() *VirtualMachineConfigurationAssignmentParameters {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineConfigurationAssignmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineConfigurationAssignmentSpec) DeepCopyInto(out *VirtualMachineConfigurationAssignmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineConfigurationAssignmentSpec.
func (in *VirtualMachineConfigurationAssignmentSpec) DeepCopy() *VirtualMachineConfigurationAssignmentSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineConfigurationAssignmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineConfigurationAssignmentStatus) DeepCopyInto(out *VirtualMachineConfigurationAssignmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineConfigurationAssignmentStatus.
func (in *VirtualMachineConfigurationAssignmentStatus) DeepCopy() *VirtualMachineConfigurationAssignmentStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineConfigurationAssignmentStatus)
	in.DeepCopyInto(out)
	return out
}
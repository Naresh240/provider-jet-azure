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
func (in *ApplicationFirewallPolicy) DeepCopyInto(out *ApplicationFirewallPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationFirewallPolicy.
func (in *ApplicationFirewallPolicy) DeepCopy() *ApplicationFirewallPolicy {
	if in == nil {
		return nil
	}
	out := new(ApplicationFirewallPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationFirewallPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationFirewallPolicyList) DeepCopyInto(out *ApplicationFirewallPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ApplicationFirewallPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationFirewallPolicyList.
func (in *ApplicationFirewallPolicyList) DeepCopy() *ApplicationFirewallPolicyList {
	if in == nil {
		return nil
	}
	out := new(ApplicationFirewallPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationFirewallPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationFirewallPolicyObservation) DeepCopyInto(out *ApplicationFirewallPolicyObservation) {
	*out = *in
	if in.HTTPListenerIds != nil {
		in, out := &in.HTTPListenerIds, &out.HTTPListenerIds
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
	if in.PathBasedRuleIds != nil {
		in, out := &in.PathBasedRuleIds, &out.PathBasedRuleIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationFirewallPolicyObservation.
func (in *ApplicationFirewallPolicyObservation) DeepCopy() *ApplicationFirewallPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(ApplicationFirewallPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationFirewallPolicyParameters) DeepCopyInto(out *ApplicationFirewallPolicyParameters) {
	*out = *in
	if in.CustomRules != nil {
		in, out := &in.CustomRules, &out.CustomRules
		*out = make([]CustomRulesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ManagedRules != nil {
		in, out := &in.ManagedRules, &out.ManagedRules
		*out = make([]ManagedRulesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PolicySettings != nil {
		in, out := &in.PolicySettings, &out.PolicySettings
		*out = make([]PolicySettingsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationFirewallPolicyParameters.
func (in *ApplicationFirewallPolicyParameters) DeepCopy() *ApplicationFirewallPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationFirewallPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationFirewallPolicySpec) DeepCopyInto(out *ApplicationFirewallPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationFirewallPolicySpec.
func (in *ApplicationFirewallPolicySpec) DeepCopy() *ApplicationFirewallPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationFirewallPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationFirewallPolicyStatus) DeepCopyInto(out *ApplicationFirewallPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationFirewallPolicyStatus.
func (in *ApplicationFirewallPolicyStatus) DeepCopy() *ApplicationFirewallPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationFirewallPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRulesObservation) DeepCopyInto(out *CustomRulesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRulesObservation.
func (in *CustomRulesObservation) DeepCopy() *CustomRulesObservation {
	if in == nil {
		return nil
	}
	out := new(CustomRulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRulesParameters) DeepCopyInto(out *CustomRulesParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.MatchConditions != nil {
		in, out := &in.MatchConditions, &out.MatchConditions
		*out = make([]MatchConditionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.RuleType != nil {
		in, out := &in.RuleType, &out.RuleType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRulesParameters.
func (in *CustomRulesParameters) DeepCopy() *CustomRulesParameters {
	if in == nil {
		return nil
	}
	out := new(CustomRulesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExclusionObservation) DeepCopyInto(out *ExclusionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExclusionObservation.
func (in *ExclusionObservation) DeepCopy() *ExclusionObservation {
	if in == nil {
		return nil
	}
	out := new(ExclusionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExclusionParameters) DeepCopyInto(out *ExclusionParameters) {
	*out = *in
	if in.MatchVariable != nil {
		in, out := &in.MatchVariable, &out.MatchVariable
		*out = new(string)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(string)
		**out = **in
	}
	if in.SelectorMatchOperator != nil {
		in, out := &in.SelectorMatchOperator, &out.SelectorMatchOperator
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExclusionParameters.
func (in *ExclusionParameters) DeepCopy() *ExclusionParameters {
	if in == nil {
		return nil
	}
	out := new(ExclusionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedRuleSetObservation) DeepCopyInto(out *ManagedRuleSetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedRuleSetObservation.
func (in *ManagedRuleSetObservation) DeepCopy() *ManagedRuleSetObservation {
	if in == nil {
		return nil
	}
	out := new(ManagedRuleSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedRuleSetParameters) DeepCopyInto(out *ManagedRuleSetParameters) {
	*out = *in
	if in.RuleGroupOverride != nil {
		in, out := &in.RuleGroupOverride, &out.RuleGroupOverride
		*out = make([]RuleGroupOverrideParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedRuleSetParameters.
func (in *ManagedRuleSetParameters) DeepCopy() *ManagedRuleSetParameters {
	if in == nil {
		return nil
	}
	out := new(ManagedRuleSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedRulesObservation) DeepCopyInto(out *ManagedRulesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedRulesObservation.
func (in *ManagedRulesObservation) DeepCopy() *ManagedRulesObservation {
	if in == nil {
		return nil
	}
	out := new(ManagedRulesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedRulesParameters) DeepCopyInto(out *ManagedRulesParameters) {
	*out = *in
	if in.Exclusion != nil {
		in, out := &in.Exclusion, &out.Exclusion
		*out = make([]ExclusionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ManagedRuleSet != nil {
		in, out := &in.ManagedRuleSet, &out.ManagedRuleSet
		*out = make([]ManagedRuleSetParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedRulesParameters.
func (in *ManagedRulesParameters) DeepCopy() *ManagedRulesParameters {
	if in == nil {
		return nil
	}
	out := new(ManagedRulesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchConditionsObservation) DeepCopyInto(out *MatchConditionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchConditionsObservation.
func (in *MatchConditionsObservation) DeepCopy() *MatchConditionsObservation {
	if in == nil {
		return nil
	}
	out := new(MatchConditionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchConditionsParameters) DeepCopyInto(out *MatchConditionsParameters) {
	*out = *in
	if in.MatchValues != nil {
		in, out := &in.MatchValues, &out.MatchValues
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.MatchVariables != nil {
		in, out := &in.MatchVariables, &out.MatchVariables
		*out = make([]MatchVariablesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NegationCondition != nil {
		in, out := &in.NegationCondition, &out.NegationCondition
		*out = new(bool)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Transforms != nil {
		in, out := &in.Transforms, &out.Transforms
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchConditionsParameters.
func (in *MatchConditionsParameters) DeepCopy() *MatchConditionsParameters {
	if in == nil {
		return nil
	}
	out := new(MatchConditionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchVariablesObservation) DeepCopyInto(out *MatchVariablesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchVariablesObservation.
func (in *MatchVariablesObservation) DeepCopy() *MatchVariablesObservation {
	if in == nil {
		return nil
	}
	out := new(MatchVariablesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchVariablesParameters) DeepCopyInto(out *MatchVariablesParameters) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(string)
		**out = **in
	}
	if in.VariableName != nil {
		in, out := &in.VariableName, &out.VariableName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchVariablesParameters.
func (in *MatchVariablesParameters) DeepCopy() *MatchVariablesParameters {
	if in == nil {
		return nil
	}
	out := new(MatchVariablesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySettingsObservation) DeepCopyInto(out *PolicySettingsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySettingsObservation.
func (in *PolicySettingsObservation) DeepCopy() *PolicySettingsObservation {
	if in == nil {
		return nil
	}
	out := new(PolicySettingsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySettingsParameters) DeepCopyInto(out *PolicySettingsParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.FileUploadLimitInMb != nil {
		in, out := &in.FileUploadLimitInMb, &out.FileUploadLimitInMb
		*out = new(int64)
		**out = **in
	}
	if in.MaxRequestBodySizeInKb != nil {
		in, out := &in.MaxRequestBodySizeInKb, &out.MaxRequestBodySizeInKb
		*out = new(int64)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.RequestBodyCheck != nil {
		in, out := &in.RequestBodyCheck, &out.RequestBodyCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySettingsParameters.
func (in *PolicySettingsParameters) DeepCopy() *PolicySettingsParameters {
	if in == nil {
		return nil
	}
	out := new(PolicySettingsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleGroupOverrideObservation) DeepCopyInto(out *RuleGroupOverrideObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleGroupOverrideObservation.
func (in *RuleGroupOverrideObservation) DeepCopy() *RuleGroupOverrideObservation {
	if in == nil {
		return nil
	}
	out := new(RuleGroupOverrideObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleGroupOverrideParameters) DeepCopyInto(out *RuleGroupOverrideParameters) {
	*out = *in
	if in.DisabledRules != nil {
		in, out := &in.DisabledRules, &out.DisabledRules
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RuleGroupName != nil {
		in, out := &in.RuleGroupName, &out.RuleGroupName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleGroupOverrideParameters.
func (in *RuleGroupOverrideParameters) DeepCopy() *RuleGroupOverrideParameters {
	if in == nil {
		return nil
	}
	out := new(RuleGroupOverrideParameters)
	in.DeepCopyInto(out)
	return out
}
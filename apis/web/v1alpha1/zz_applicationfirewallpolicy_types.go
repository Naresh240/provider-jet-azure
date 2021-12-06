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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ApplicationFirewallPolicyObservation struct {
	HTTPListenerIds []*string `json:"httpListenerIds,omitempty" tf:"http_listener_ids,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PathBasedRuleIds []*string `json:"pathBasedRuleIds,omitempty" tf:"path_based_rule_ids,omitempty"`
}

type ApplicationFirewallPolicyParameters struct {

	// +kubebuilder:validation:Optional
	CustomRules []CustomRulesParameters `json:"customRules,omitempty" tf:"custom_rules,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	ManagedRules []ManagedRulesParameters `json:"managedRules" tf:"managed_rules,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PolicySettings []PolicySettingsParameters `json:"policySettings,omitempty" tf:"policy_settings,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CustomRulesObservation struct {
}

type CustomRulesParameters struct {

	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	MatchConditions []MatchConditionsParameters `json:"matchConditions" tf:"match_conditions,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *int64 `json:"priority" tf:"priority,omitempty"`

	// +kubebuilder:validation:Required
	RuleType *string `json:"ruleType" tf:"rule_type,omitempty"`
}

type ExclusionObservation struct {
}

type ExclusionParameters struct {

	// +kubebuilder:validation:Required
	MatchVariable *string `json:"matchVariable" tf:"match_variable,omitempty"`

	// +kubebuilder:validation:Required
	Selector *string `json:"selector" tf:"selector,omitempty"`

	// +kubebuilder:validation:Required
	SelectorMatchOperator *string `json:"selectorMatchOperator" tf:"selector_match_operator,omitempty"`
}

type ManagedRuleSetObservation struct {
}

type ManagedRuleSetParameters struct {

	// +kubebuilder:validation:Optional
	RuleGroupOverride []RuleGroupOverrideParameters `json:"ruleGroupOverride,omitempty" tf:"rule_group_override,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type ManagedRulesObservation struct {
}

type ManagedRulesParameters struct {

	// +kubebuilder:validation:Optional
	Exclusion []ExclusionParameters `json:"exclusion,omitempty" tf:"exclusion,omitempty"`

	// +kubebuilder:validation:Required
	ManagedRuleSet []ManagedRuleSetParameters `json:"managedRuleSet" tf:"managed_rule_set,omitempty"`
}

type MatchConditionsObservation struct {
}

type MatchConditionsParameters struct {

	// +kubebuilder:validation:Required
	MatchValues []*string `json:"matchValues" tf:"match_values,omitempty"`

	// +kubebuilder:validation:Required
	MatchVariables []MatchVariablesParameters `json:"matchVariables" tf:"match_variables,omitempty"`

	// +kubebuilder:validation:Optional
	NegationCondition *bool `json:"negationCondition,omitempty" tf:"negation_condition,omitempty"`

	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// +kubebuilder:validation:Optional
	Transforms []*string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type MatchVariablesObservation struct {
}

type MatchVariablesParameters struct {

	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// +kubebuilder:validation:Required
	VariableName *string `json:"variableName" tf:"variable_name,omitempty"`
}

type PolicySettingsObservation struct {
}

type PolicySettingsParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	FileUploadLimitInMb *int64 `json:"fileUploadLimitInMb,omitempty" tf:"file_upload_limit_in_mb,omitempty"`

	// +kubebuilder:validation:Optional
	MaxRequestBodySizeInKb *int64 `json:"maxRequestBodySizeInKb,omitempty" tf:"max_request_body_size_in_kb,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	RequestBodyCheck *bool `json:"requestBodyCheck,omitempty" tf:"request_body_check,omitempty"`
}

type RuleGroupOverrideObservation struct {
}

type RuleGroupOverrideParameters struct {

	// +kubebuilder:validation:Required
	DisabledRules []*string `json:"disabledRules" tf:"disabled_rules,omitempty"`

	// +kubebuilder:validation:Required
	RuleGroupName *string `json:"ruleGroupName" tf:"rule_group_name,omitempty"`
}

// ApplicationFirewallPolicySpec defines the desired state of ApplicationFirewallPolicy
type ApplicationFirewallPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationFirewallPolicyParameters `json:"forProvider"`
}

// ApplicationFirewallPolicyStatus defines the observed state of ApplicationFirewallPolicy.
type ApplicationFirewallPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationFirewallPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationFirewallPolicy is the Schema for the ApplicationFirewallPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ApplicationFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationFirewallPolicySpec   `json:"spec"`
	Status            ApplicationFirewallPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationFirewallPolicyList contains a list of ApplicationFirewallPolicys
type ApplicationFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationFirewallPolicy `json:"items"`
}

// Repository type metadata.
var (
	ApplicationFirewallPolicy_Kind             = "ApplicationFirewallPolicy"
	ApplicationFirewallPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationFirewallPolicy_Kind}.String()
	ApplicationFirewallPolicy_KindAPIVersion   = ApplicationFirewallPolicy_Kind + "." + CRDGroupVersion.String()
	ApplicationFirewallPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationFirewallPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationFirewallPolicy{}, &ApplicationFirewallPolicyList{})
}
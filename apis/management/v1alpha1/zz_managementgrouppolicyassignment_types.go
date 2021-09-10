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
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type IdentityObservation struct {
	PrincipalID string `json:"principalId" tf:"principal_id"`

	TenantID string `json:"tenantId" tf:"tenant_id"`
}

type IdentityParameters struct {
	Type *string `json:"type,omitempty" tf:"type"`
}

type ManagementGroupPolicyAssignmentObservation struct {
}

type ManagementGroupPolicyAssignmentParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`

	Enforce *bool `json:"enforce,omitempty" tf:"enforce"`

	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	Location *string `json:"location,omitempty" tf:"location"`

	ManagementGroupID string `json:"managementGroupId" tf:"management_group_id"`

	Metadata *string `json:"metadata,omitempty" tf:"metadata"`

	Name string `json:"name" tf:"name"`

	NotScopes []string `json:"notScopes,omitempty" tf:"not_scopes"`

	Parameters *string `json:"parameters,omitempty" tf:"parameters"`

	PolicyDefinitionID string `json:"policyDefinitionId" tf:"policy_definition_id"`
}

// ManagementGroupPolicyAssignmentSpec defines the desired state of ManagementGroupPolicyAssignment
type ManagementGroupPolicyAssignmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ManagementGroupPolicyAssignmentParameters `json:"forProvider"`
}

// ManagementGroupPolicyAssignmentStatus defines the observed state of ManagementGroupPolicyAssignment.
type ManagementGroupPolicyAssignmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ManagementGroupPolicyAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementGroupPolicyAssignment is the Schema for the ManagementGroupPolicyAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagementGroupPolicyAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementGroupPolicyAssignmentSpec   `json:"spec"`
	Status            ManagementGroupPolicyAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementGroupPolicyAssignmentList contains a list of ManagementGroupPolicyAssignments
type ManagementGroupPolicyAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementGroupPolicyAssignment `json:"items"`
}

// Repository type metadata.
var (
	ManagementGroupPolicyAssignmentKind             = "ManagementGroupPolicyAssignment"
	ManagementGroupPolicyAssignmentGroupKind        = schema.GroupKind{Group: Group, Kind: ManagementGroupPolicyAssignmentKind}.String()
	ManagementGroupPolicyAssignmentKindAPIVersion   = ManagementGroupPolicyAssignmentKind + "." + GroupVersion.String()
	ManagementGroupPolicyAssignmentGroupVersionKind = GroupVersion.WithKind(ManagementGroupPolicyAssignmentKind)
)

func init() {
	SchemeBuilder.Register(&ManagementGroupPolicyAssignment{}, &ManagementGroupPolicyAssignmentList{})
}

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

type KustoDatabasePrincipalAssignmentObservation struct {
	PrincipalName string `json:"principalName" tf:"principal_name"`

	TenantName string `json:"tenantName" tf:"tenant_name"`
}

type KustoDatabasePrincipalAssignmentParameters struct {
	ClusterName string `json:"clusterName" tf:"cluster_name"`

	DatabaseName string `json:"databaseName" tf:"database_name"`

	Name string `json:"name" tf:"name"`

	PrincipalID string `json:"principalId" tf:"principal_id"`

	PrincipalType string `json:"principalType" tf:"principal_type"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Role string `json:"role" tf:"role"`

	TenantID string `json:"tenantId" tf:"tenant_id"`
}

// KustoDatabasePrincipalAssignmentSpec defines the desired state of KustoDatabasePrincipalAssignment
type KustoDatabasePrincipalAssignmentSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KustoDatabasePrincipalAssignmentParameters `json:"forProvider"`
}

// KustoDatabasePrincipalAssignmentStatus defines the observed state of KustoDatabasePrincipalAssignment.
type KustoDatabasePrincipalAssignmentStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KustoDatabasePrincipalAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KustoDatabasePrincipalAssignment is the Schema for the KustoDatabasePrincipalAssignments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type KustoDatabasePrincipalAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KustoDatabasePrincipalAssignmentSpec   `json:"spec"`
	Status            KustoDatabasePrincipalAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KustoDatabasePrincipalAssignmentList contains a list of KustoDatabasePrincipalAssignments
type KustoDatabasePrincipalAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KustoDatabasePrincipalAssignment `json:"items"`
}

// Repository type metadata.
var (
	KustoDatabasePrincipalAssignmentKind             = "KustoDatabasePrincipalAssignment"
	KustoDatabasePrincipalAssignmentGroupKind        = schema.GroupKind{Group: Group, Kind: KustoDatabasePrincipalAssignmentKind}.String()
	KustoDatabasePrincipalAssignmentKindAPIVersion   = KustoDatabasePrincipalAssignmentKind + "." + GroupVersion.String()
	KustoDatabasePrincipalAssignmentGroupVersionKind = GroupVersion.WithKind(KustoDatabasePrincipalAssignmentKind)
)

func init() {
	SchemeBuilder.Register(&KustoDatabasePrincipalAssignment{}, &KustoDatabasePrincipalAssignmentList{})
}

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

type VirtualHubRouteTableObservation struct {
}

type VirtualHubRouteTableParameters struct {

	// +kubebuilder:validation:Optional
	Labels []string `json:"labels,omitempty" tf:"labels"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	Route []VirtualHubRouteTableRouteParameters `json:"route,omitempty" tf:"route"`

	// +kubebuilder:validation:Required
	VirtualHubID string `json:"virtualHubId" tf:"virtual_hub_id"`
}

type VirtualHubRouteTableRouteObservation struct {
}

type VirtualHubRouteTableRouteParameters struct {

	// +kubebuilder:validation:Required
	Destinations []string `json:"destinations" tf:"destinations"`

	// +kubebuilder:validation:Required
	DestinationsType string `json:"destinationsType" tf:"destinations_type"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	NextHop string `json:"nextHop" tf:"next_hop"`

	// +kubebuilder:validation:Optional
	NextHopType *string `json:"nextHopType,omitempty" tf:"next_hop_type"`
}

// VirtualHubRouteTableSpec defines the desired state of VirtualHubRouteTable
type VirtualHubRouteTableSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VirtualHubRouteTableParameters `json:"forProvider"`
}

// VirtualHubRouteTableStatus defines the observed state of VirtualHubRouteTable.
type VirtualHubRouteTableStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VirtualHubRouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubRouteTable is the Schema for the VirtualHubRouteTables API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualHubRouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualHubRouteTableSpec   `json:"spec"`
	Status            VirtualHubRouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubRouteTableList contains a list of VirtualHubRouteTables
type VirtualHubRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualHubRouteTable `json:"items"`
}

// Repository type metadata.
var (
	VirtualHubRouteTableKind             = "VirtualHubRouteTable"
	VirtualHubRouteTableGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualHubRouteTableKind}.String()
	VirtualHubRouteTableKindAPIVersion   = VirtualHubRouteTableKind + "." + GroupVersion.String()
	VirtualHubRouteTableGroupVersionKind = GroupVersion.WithKind(VirtualHubRouteTableKind)
)

func init() {
	SchemeBuilder.Register(&VirtualHubRouteTable{}, &VirtualHubRouteTableList{})
}
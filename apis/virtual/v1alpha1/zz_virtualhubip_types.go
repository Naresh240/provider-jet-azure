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

type VirtualHubIpObservation struct {
}

type VirtualHubIpParameters struct {
	Name string `json:"name" tf:"name"`

	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address"`

	PrivateIPAllocationMethod *string `json:"privateIpAllocationMethod,omitempty" tf:"private_ip_allocation_method"`

	PublicIPAddressID *string `json:"publicIpAddressId,omitempty" tf:"public_ip_address_id"`

	SubnetID string `json:"subnetId" tf:"subnet_id"`

	VirtualHubID string `json:"virtualHubId" tf:"virtual_hub_id"`
}

// VirtualHubIpSpec defines the desired state of VirtualHubIp
type VirtualHubIpSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VirtualHubIpParameters `json:"forProvider"`
}

// VirtualHubIpStatus defines the observed state of VirtualHubIp.
type VirtualHubIpStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VirtualHubIpObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubIp is the Schema for the VirtualHubIps API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualHubIp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualHubIpSpec   `json:"spec"`
	Status            VirtualHubIpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubIpList contains a list of VirtualHubIps
type VirtualHubIpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualHubIp `json:"items"`
}

// Repository type metadata.
var (
	VirtualHubIpKind             = "VirtualHubIp"
	VirtualHubIpGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualHubIpKind}.String()
	VirtualHubIpKindAPIVersion   = VirtualHubIpKind + "." + GroupVersion.String()
	VirtualHubIpGroupVersionKind = GroupVersion.WithKind(VirtualHubIpKind)
)

func init() {
	SchemeBuilder.Register(&VirtualHubIp{}, &VirtualHubIpList{})
}

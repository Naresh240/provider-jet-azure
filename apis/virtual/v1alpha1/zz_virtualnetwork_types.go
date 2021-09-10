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

type DdosProtectionPlanObservation struct {
}

type DdosProtectionPlanParameters struct {
	Enable bool `json:"enable" tf:"enable"`

	ID string `json:"id" tf:"id"`
}

type SubnetObservation struct {
	ID string `json:"id" tf:"id"`
}

type SubnetParameters struct {
	AddressPrefix string `json:"addressPrefix" tf:"address_prefix"`

	Name string `json:"name" tf:"name"`

	SecurityGroup *string `json:"securityGroup,omitempty" tf:"security_group"`
}

type VirtualNetworkObservation struct {
	GUID string `json:"guid" tf:"guid"`
}

type VirtualNetworkParameters struct {
	AddressSpace []string `json:"addressSpace" tf:"address_space"`

	BgpCommunity *string `json:"bgpCommunity,omitempty" tf:"bgp_community"`

	DNSServers []string `json:"dnsServers,omitempty" tf:"dns_servers"`

	DdosProtectionPlan []DdosProtectionPlanParameters `json:"ddosProtectionPlan,omitempty" tf:"ddos_protection_plan"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Subnet []SubnetParameters `json:"subnet,omitempty" tf:"subnet"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	VMProtectionEnabled *bool `json:"vmProtectionEnabled,omitempty" tf:"vm_protection_enabled"`
}

// VirtualNetworkSpec defines the desired state of VirtualNetwork
type VirtualNetworkSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VirtualNetworkParameters `json:"forProvider"`
}

// VirtualNetworkStatus defines the observed state of VirtualNetwork.
type VirtualNetworkStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VirtualNetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetwork is the Schema for the VirtualNetworks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkSpec   `json:"spec"`
	Status            VirtualNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkList contains a list of VirtualNetworks
type VirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetwork `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetworkKind             = "VirtualNetwork"
	VirtualNetworkGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualNetworkKind}.String()
	VirtualNetworkKindAPIVersion   = VirtualNetworkKind + "." + GroupVersion.String()
	VirtualNetworkGroupVersionKind = GroupVersion.WithKind(VirtualNetworkKind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetwork{}, &VirtualNetworkList{})
}

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

type VirtualNetworkDnsServersObservation struct {
}

type VirtualNetworkDnsServersParameters struct {
	DNSServers []string `json:"dnsServers,omitempty" tf:"dns_servers"`

	VirtualNetworkID string `json:"virtualNetworkId" tf:"virtual_network_id"`
}

// VirtualNetworkDnsServersSpec defines the desired state of VirtualNetworkDnsServers
type VirtualNetworkDnsServersSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VirtualNetworkDnsServersParameters `json:"forProvider"`
}

// VirtualNetworkDnsServersStatus defines the observed state of VirtualNetworkDnsServers.
type VirtualNetworkDnsServersStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VirtualNetworkDnsServersObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkDnsServers is the Schema for the VirtualNetworkDnsServerss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetworkDnsServers struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkDnsServersSpec   `json:"spec"`
	Status            VirtualNetworkDnsServersStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkDnsServersList contains a list of VirtualNetworkDnsServerss
type VirtualNetworkDnsServersList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworkDnsServers `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetworkDnsServersKind             = "VirtualNetworkDnsServers"
	VirtualNetworkDnsServersGroupKind        = schema.GroupKind{Group: Group, Kind: VirtualNetworkDnsServersKind}.String()
	VirtualNetworkDnsServersKindAPIVersion   = VirtualNetworkDnsServersKind + "." + GroupVersion.String()
	VirtualNetworkDnsServersGroupVersionKind = GroupVersion.WithKind(VirtualNetworkDnsServersKind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetworkDnsServers{}, &VirtualNetworkDnsServersList{})
}

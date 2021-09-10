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

type BgpSettingsObservation struct {
	BgpPeeringAddress string `json:"bgpPeeringAddress" tf:"bgp_peering_address"`
}

type BgpSettingsParameters struct {
	Asn int64 `json:"asn" tf:"asn"`

	Instance0BgpPeeringAddress []Instance0BgpPeeringAddressParameters `json:"instance0BgpPeeringAddress,omitempty" tf:"instance_0_bgp_peering_address"`

	Instance1BgpPeeringAddress []Instance1BgpPeeringAddressParameters `json:"instance1BgpPeeringAddress,omitempty" tf:"instance_1_bgp_peering_address"`

	PeerWeight int64 `json:"peerWeight" tf:"peer_weight"`
}

type Instance0BgpPeeringAddressObservation struct {
	DefaultIps []string `json:"defaultIps" tf:"default_ips"`

	IPConfigurationID string `json:"ipConfigurationId" tf:"ip_configuration_id"`

	TunnelIps []string `json:"tunnelIps" tf:"tunnel_ips"`
}

type Instance0BgpPeeringAddressParameters struct {
	CustomIps []string `json:"customIps" tf:"custom_ips"`
}

type Instance1BgpPeeringAddressObservation struct {
	DefaultIps []string `json:"defaultIps" tf:"default_ips"`

	IPConfigurationID string `json:"ipConfigurationId" tf:"ip_configuration_id"`

	TunnelIps []string `json:"tunnelIps" tf:"tunnel_ips"`
}

type Instance1BgpPeeringAddressParameters struct {
	CustomIps []string `json:"customIps" tf:"custom_ips"`
}

type VpnGatewayObservation struct {
}

type VpnGatewayParameters struct {
	BgpSettings []BgpSettingsParameters `json:"bgpSettings,omitempty" tf:"bgp_settings"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	ScaleUnit *int64 `json:"scaleUnit,omitempty" tf:"scale_unit"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	VirtualHubID string `json:"virtualHubId" tf:"virtual_hub_id"`
}

// VpnGatewaySpec defines the desired state of VpnGateway
type VpnGatewaySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VpnGatewayParameters `json:"forProvider"`
}

// VpnGatewayStatus defines the observed state of VpnGateway.
type VpnGatewayStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          VpnGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VpnGateway is the Schema for the VpnGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VpnGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpnGatewaySpec   `json:"spec"`
	Status            VpnGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpnGatewayList contains a list of VpnGateways
type VpnGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpnGateway `json:"items"`
}

// Repository type metadata.
var (
	VpnGatewayKind             = "VpnGateway"
	VpnGatewayGroupKind        = schema.GroupKind{Group: Group, Kind: VpnGatewayKind}.String()
	VpnGatewayKindAPIVersion   = VpnGatewayKind + "." + GroupVersion.String()
	VpnGatewayGroupVersionKind = GroupVersion.WithKind(VpnGatewayKind)
)

func init() {
	SchemeBuilder.Register(&VpnGateway{}, &VpnGatewayList{})
}

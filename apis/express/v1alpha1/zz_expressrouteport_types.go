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

type ExpressRoutePortObservation struct {
	Ethertype string `json:"ethertype" tf:"ethertype"`

	GUID string `json:"guid" tf:"guid"`

	Mtu string `json:"mtu" tf:"mtu"`
}

type ExpressRoutePortParameters struct {
	BandwidthInGbps int64 `json:"bandwidthInGbps" tf:"bandwidth_in_gbps"`

	Encapsulation string `json:"encapsulation" tf:"encapsulation"`

	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	Link1 []Link1Parameters `json:"link1,omitempty" tf:"link1"`

	Link2 []Link2Parameters `json:"link2,omitempty" tf:"link2"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	PeeringLocation string `json:"peeringLocation" tf:"peering_location"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type IdentityObservation struct {
}

type IdentityParameters struct {
	IdentityIds []string `json:"identityIds,omitempty" tf:"identity_ids"`

	Type string `json:"type" tf:"type"`
}

type Link1Observation struct {
	ConnectorType string `json:"connectorType" tf:"connector_type"`

	ID string `json:"id" tf:"id"`

	InterfaceName string `json:"interfaceName" tf:"interface_name"`

	PatchPanelID string `json:"patchPanelId" tf:"patch_panel_id"`

	RackID string `json:"rackId" tf:"rack_id"`

	RouterName string `json:"routerName" tf:"router_name"`
}

type Link1Parameters struct {
	AdminEnabled *bool `json:"adminEnabled,omitempty" tf:"admin_enabled"`

	MacsecCakKeyvaultSecretID *string `json:"macsecCakKeyvaultSecretId,omitempty" tf:"macsec_cak_keyvault_secret_id"`

	MacsecCipher *string `json:"macsecCipher,omitempty" tf:"macsec_cipher"`

	MacsecCknKeyvaultSecretID *string `json:"macsecCknKeyvaultSecretId,omitempty" tf:"macsec_ckn_keyvault_secret_id"`
}

type Link2Observation struct {
	ConnectorType string `json:"connectorType" tf:"connector_type"`

	ID string `json:"id" tf:"id"`

	InterfaceName string `json:"interfaceName" tf:"interface_name"`

	PatchPanelID string `json:"patchPanelId" tf:"patch_panel_id"`

	RackID string `json:"rackId" tf:"rack_id"`

	RouterName string `json:"routerName" tf:"router_name"`
}

type Link2Parameters struct {
	AdminEnabled *bool `json:"adminEnabled,omitempty" tf:"admin_enabled"`

	MacsecCakKeyvaultSecretID *string `json:"macsecCakKeyvaultSecretId,omitempty" tf:"macsec_cak_keyvault_secret_id"`

	MacsecCipher *string `json:"macsecCipher,omitempty" tf:"macsec_cipher"`

	MacsecCknKeyvaultSecretID *string `json:"macsecCknKeyvaultSecretId,omitempty" tf:"macsec_ckn_keyvault_secret_id"`
}

// ExpressRoutePortSpec defines the desired state of ExpressRoutePort
type ExpressRoutePortSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ExpressRoutePortParameters `json:"forProvider"`
}

// ExpressRoutePortStatus defines the observed state of ExpressRoutePort.
type ExpressRoutePortStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ExpressRoutePortObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRoutePort is the Schema for the ExpressRoutePorts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRoutePort struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRoutePortSpec   `json:"spec"`
	Status            ExpressRoutePortStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRoutePortList contains a list of ExpressRoutePorts
type ExpressRoutePortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRoutePort `json:"items"`
}

// Repository type metadata.
var (
	ExpressRoutePortKind             = "ExpressRoutePort"
	ExpressRoutePortGroupKind        = schema.GroupKind{Group: Group, Kind: ExpressRoutePortKind}.String()
	ExpressRoutePortKindAPIVersion   = ExpressRoutePortKind + "." + GroupVersion.String()
	ExpressRoutePortGroupVersionKind = GroupVersion.WithKind(ExpressRoutePortKind)
)

func init() {
	SchemeBuilder.Register(&ExpressRoutePort{}, &ExpressRoutePortList{})
}

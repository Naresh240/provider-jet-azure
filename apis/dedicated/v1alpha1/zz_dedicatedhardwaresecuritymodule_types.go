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

type DedicatedHardwareSecurityModuleObservation struct {
}

type DedicatedHardwareSecurityModuleParameters struct {
	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	NetworkProfile []NetworkProfileParameters `json:"networkProfile" tf:"network_profile"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	SkuName string `json:"skuName" tf:"sku_name"`

	StampID *string `json:"stampId,omitempty" tf:"stamp_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type NetworkProfileObservation struct {
}

type NetworkProfileParameters struct {
	NetworkInterfacePrivateIPAddresses []string `json:"networkInterfacePrivateIpAddresses" tf:"network_interface_private_ip_addresses"`

	SubnetID string `json:"subnetId" tf:"subnet_id"`
}

// DedicatedHardwareSecurityModuleSpec defines the desired state of DedicatedHardwareSecurityModule
type DedicatedHardwareSecurityModuleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DedicatedHardwareSecurityModuleParameters `json:"forProvider"`
}

// DedicatedHardwareSecurityModuleStatus defines the observed state of DedicatedHardwareSecurityModule.
type DedicatedHardwareSecurityModuleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DedicatedHardwareSecurityModuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedHardwareSecurityModule is the Schema for the DedicatedHardwareSecurityModules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DedicatedHardwareSecurityModule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedHardwareSecurityModuleSpec   `json:"spec"`
	Status            DedicatedHardwareSecurityModuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedHardwareSecurityModuleList contains a list of DedicatedHardwareSecurityModules
type DedicatedHardwareSecurityModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedHardwareSecurityModule `json:"items"`
}

// Repository type metadata.
var (
	DedicatedHardwareSecurityModuleKind             = "DedicatedHardwareSecurityModule"
	DedicatedHardwareSecurityModuleGroupKind        = schema.GroupKind{Group: Group, Kind: DedicatedHardwareSecurityModuleKind}.String()
	DedicatedHardwareSecurityModuleKindAPIVersion   = DedicatedHardwareSecurityModuleKind + "." + GroupVersion.String()
	DedicatedHardwareSecurityModuleGroupVersionKind = GroupVersion.WithKind(DedicatedHardwareSecurityModuleKind)
)

func init() {
	SchemeBuilder.Register(&DedicatedHardwareSecurityModule{}, &DedicatedHardwareSecurityModuleList{})
}

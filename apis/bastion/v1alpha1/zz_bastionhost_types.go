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

type BastionHostObservation struct {
	DNSName string `json:"dnsName" tf:"dns_name"`
}

type BastionHostParameters struct {
	IPConfiguration []IPConfigurationParameters `json:"ipConfiguration,omitempty" tf:"ip_configuration"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

type IPConfigurationObservation struct {
}

type IPConfigurationParameters struct {
	Name string `json:"name" tf:"name"`

	PublicIPAddressID string `json:"publicIpAddressId" tf:"public_ip_address_id"`

	SubnetID string `json:"subnetId" tf:"subnet_id"`
}

// BastionHostSpec defines the desired state of BastionHost
type BastionHostSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BastionHostParameters `json:"forProvider"`
}

// BastionHostStatus defines the observed state of BastionHost.
type BastionHostStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BastionHostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BastionHost is the Schema for the BastionHosts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BastionHost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BastionHostSpec   `json:"spec"`
	Status            BastionHostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BastionHostList contains a list of BastionHosts
type BastionHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BastionHost `json:"items"`
}

// Repository type metadata.
var (
	BastionHostKind             = "BastionHost"
	BastionHostGroupKind        = schema.GroupKind{Group: Group, Kind: BastionHostKind}.String()
	BastionHostKindAPIVersion   = BastionHostKind + "." + GroupVersion.String()
	BastionHostGroupVersionKind = GroupVersion.WithKind(BastionHostKind)
)

func init() {
	SchemeBuilder.Register(&BastionHost{}, &BastionHostList{})
}

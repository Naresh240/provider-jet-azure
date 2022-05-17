/*
Copyright 2022 The Crossplane Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActiveDirectoryDomainServiceReplicaSetObservation struct {
	DomainControllerIPAddresses []*string `json:"domainControllerIpAddresses,omitempty" tf:"domain_controller_ip_addresses,omitempty"`

	ExternalAccessIPAddress *string `json:"externalAccessIpAddress,omitempty" tf:"external_access_ip_address,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ServiceStatus *string `json:"serviceStatus,omitempty" tf:"service_status,omitempty"`
}

type ActiveDirectoryDomainServiceReplicaSetParameters struct {

	// +kubebuilder:validation:Required
	DomainServiceID *string `json:"domainServiceId" tf:"domain_service_id,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// ActiveDirectoryDomainServiceReplicaSetSpec defines the desired state of ActiveDirectoryDomainServiceReplicaSet
type ActiveDirectoryDomainServiceReplicaSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActiveDirectoryDomainServiceReplicaSetParameters `json:"forProvider"`
}

// ActiveDirectoryDomainServiceReplicaSetStatus defines the observed state of ActiveDirectoryDomainServiceReplicaSet.
type ActiveDirectoryDomainServiceReplicaSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActiveDirectoryDomainServiceReplicaSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ActiveDirectoryDomainServiceReplicaSet is the Schema for the ActiveDirectoryDomainServiceReplicaSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ActiveDirectoryDomainServiceReplicaSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActiveDirectoryDomainServiceReplicaSetSpec   `json:"spec"`
	Status            ActiveDirectoryDomainServiceReplicaSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActiveDirectoryDomainServiceReplicaSetList contains a list of ActiveDirectoryDomainServiceReplicaSets
type ActiveDirectoryDomainServiceReplicaSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActiveDirectoryDomainServiceReplicaSet `json:"items"`
}

// Repository type metadata.
var (
	ActiveDirectoryDomainServiceReplicaSet_Kind             = "ActiveDirectoryDomainServiceReplicaSet"
	ActiveDirectoryDomainServiceReplicaSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActiveDirectoryDomainServiceReplicaSet_Kind}.String()
	ActiveDirectoryDomainServiceReplicaSet_KindAPIVersion   = ActiveDirectoryDomainServiceReplicaSet_Kind + "." + CRDGroupVersion.String()
	ActiveDirectoryDomainServiceReplicaSet_GroupVersionKind = CRDGroupVersion.WithKind(ActiveDirectoryDomainServiceReplicaSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ActiveDirectoryDomainServiceReplicaSet{}, &ActiveDirectoryDomainServiceReplicaSetList{})
}

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

type DnsARecordObservation struct {
	Fqdn string `json:"fqdn" tf:"fqdn"`
}

type DnsARecordParameters struct {
	Name string `json:"name" tf:"name"`

	Records []string `json:"records,omitempty" tf:"records"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	TTL int64 `json:"ttl" tf:"ttl"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id"`

	ZoneName string `json:"zoneName" tf:"zone_name"`
}

// DnsARecordSpec defines the desired state of DnsARecord
type DnsARecordSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DnsARecordParameters `json:"forProvider"`
}

// DnsARecordStatus defines the observed state of DnsARecord.
type DnsARecordStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DnsARecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DnsARecord is the Schema for the DnsARecords API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DnsARecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsARecordSpec   `json:"spec"`
	Status            DnsARecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DnsARecordList contains a list of DnsARecords
type DnsARecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsARecord `json:"items"`
}

// Repository type metadata.
var (
	DnsARecordKind             = "DnsARecord"
	DnsARecordGroupKind        = schema.GroupKind{Group: Group, Kind: DnsARecordKind}.String()
	DnsARecordKindAPIVersion   = DnsARecordKind + "." + GroupVersion.String()
	DnsARecordGroupVersionKind = GroupVersion.WithKind(DnsARecordKind)
)

func init() {
	SchemeBuilder.Register(&DnsARecord{}, &DnsARecordList{})
}

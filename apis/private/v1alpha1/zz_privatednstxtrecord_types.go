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

type PrivateDnsTxtRecordObservation struct {
	Fqdn string `json:"fqdn" tf:"fqdn"`
}

type PrivateDnsTxtRecordParameters struct {
	Name string `json:"name" tf:"name"`

	Record []PrivateDnsTxtRecordRecordParameters `json:"record" tf:"record"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	TTL int64 `json:"ttl" tf:"ttl"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	ZoneName string `json:"zoneName" tf:"zone_name"`
}

type PrivateDnsTxtRecordRecordObservation struct {
}

type PrivateDnsTxtRecordRecordParameters struct {
	Value string `json:"value" tf:"value"`
}

// PrivateDnsTxtRecordSpec defines the desired state of PrivateDnsTxtRecord
type PrivateDnsTxtRecordSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       PrivateDnsTxtRecordParameters `json:"forProvider"`
}

// PrivateDnsTxtRecordStatus defines the observed state of PrivateDnsTxtRecord.
type PrivateDnsTxtRecordStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          PrivateDnsTxtRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDnsTxtRecord is the Schema for the PrivateDnsTxtRecords API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDnsTxtRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDnsTxtRecordSpec   `json:"spec"`
	Status            PrivateDnsTxtRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDnsTxtRecordList contains a list of PrivateDnsTxtRecords
type PrivateDnsTxtRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDnsTxtRecord `json:"items"`
}

// Repository type metadata.
var (
	PrivateDnsTxtRecordKind             = "PrivateDnsTxtRecord"
	PrivateDnsTxtRecordGroupKind        = schema.GroupKind{Group: Group, Kind: PrivateDnsTxtRecordKind}.String()
	PrivateDnsTxtRecordKindAPIVersion   = PrivateDnsTxtRecordKind + "." + GroupVersion.String()
	PrivateDnsTxtRecordGroupVersionKind = GroupVersion.WithKind(PrivateDnsTxtRecordKind)
)

func init() {
	SchemeBuilder.Register(&PrivateDnsTxtRecord{}, &PrivateDnsTxtRecordList{})
}

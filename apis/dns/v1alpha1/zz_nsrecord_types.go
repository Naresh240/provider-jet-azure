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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NsRecordObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NsRecordParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Records []*string `json:"records" tf:"records,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	TTL *int64 `json:"ttl" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	ZoneName *string `json:"zoneName" tf:"zone_name,omitempty"`
}

// NsRecordSpec defines the desired state of NsRecord
type NsRecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NsRecordParameters `json:"forProvider"`
}

// NsRecordStatus defines the observed state of NsRecord.
type NsRecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NsRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NsRecord is the Schema for the NsRecords API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type NsRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NsRecordSpec   `json:"spec"`
	Status            NsRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NsRecordList contains a list of NsRecords
type NsRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsRecord `json:"items"`
}

// Repository type metadata.
var (
	NsRecord_Kind             = "NsRecord"
	NsRecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NsRecord_Kind}.String()
	NsRecord_KindAPIVersion   = NsRecord_Kind + "." + CRDGroupVersion.String()
	NsRecord_GroupVersionKind = CRDGroupVersion.WithKind(NsRecord_Kind)
)

func init() {
	SchemeBuilder.Register(&NsRecord{}, &NsRecordList{})
}
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

type DataFactoryTriggerBlobEventObservation struct {
}

type DataFactoryTriggerBlobEventParameters struct {
	Activated *bool `json:"activated,omitempty" tf:"activated"`

	AdditionalProperties map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties"`

	Annotations []string `json:"annotations,omitempty" tf:"annotations"`

	BlobPathBeginsWith *string `json:"blobPathBeginsWith,omitempty" tf:"blob_path_begins_with"`

	BlobPathEndsWith *string `json:"blobPathEndsWith,omitempty" tf:"blob_path_ends_with"`

	DataFactoryID string `json:"dataFactoryId" tf:"data_factory_id"`

	Description *string `json:"description,omitempty" tf:"description"`

	Events []string `json:"events" tf:"events"`

	IgnoreEmptyBlobs *bool `json:"ignoreEmptyBlobs,omitempty" tf:"ignore_empty_blobs"`

	Name string `json:"name" tf:"name"`

	Pipeline []PipelineParameters `json:"pipeline" tf:"pipeline"`

	StorageAccountID string `json:"storageAccountId" tf:"storage_account_id"`
}

type PipelineObservation struct {
}

type PipelineParameters struct {
	Name string `json:"name" tf:"name"`

	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`
}

// DataFactoryTriggerBlobEventSpec defines the desired state of DataFactoryTriggerBlobEvent
type DataFactoryTriggerBlobEventSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DataFactoryTriggerBlobEventParameters `json:"forProvider"`
}

// DataFactoryTriggerBlobEventStatus defines the observed state of DataFactoryTriggerBlobEvent.
type DataFactoryTriggerBlobEventStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DataFactoryTriggerBlobEventObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryTriggerBlobEvent is the Schema for the DataFactoryTriggerBlobEvents API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFactoryTriggerBlobEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryTriggerBlobEventSpec   `json:"spec"`
	Status            DataFactoryTriggerBlobEventStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryTriggerBlobEventList contains a list of DataFactoryTriggerBlobEvents
type DataFactoryTriggerBlobEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFactoryTriggerBlobEvent `json:"items"`
}

// Repository type metadata.
var (
	DataFactoryTriggerBlobEventKind             = "DataFactoryTriggerBlobEvent"
	DataFactoryTriggerBlobEventGroupKind        = schema.GroupKind{Group: Group, Kind: DataFactoryTriggerBlobEventKind}.String()
	DataFactoryTriggerBlobEventKindAPIVersion   = DataFactoryTriggerBlobEventKind + "." + GroupVersion.String()
	DataFactoryTriggerBlobEventGroupVersionKind = GroupVersion.WithKind(DataFactoryTriggerBlobEventKind)
)

func init() {
	SchemeBuilder.Register(&DataFactoryTriggerBlobEvent{}, &DataFactoryTriggerBlobEventList{})
}

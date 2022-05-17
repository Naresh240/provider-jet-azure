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

type TriggerBlobEventObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TriggerBlobEventParameters struct {

	// +kubebuilder:validation:Optional
	Activated *bool `json:"activated,omitempty" tf:"activated,omitempty"`

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Optional
	BlobPathBeginsWith *string `json:"blobPathBeginsWith,omitempty" tf:"blob_path_begins_with,omitempty"`

	// +kubebuilder:validation:Optional
	BlobPathEndsWith *string `json:"blobPathEndsWith,omitempty" tf:"blob_path_ends_with,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryID *string `json:"dataFactoryId" tf:"data_factory_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Events []*string `json:"events" tf:"events,omitempty"`

	// +kubebuilder:validation:Optional
	IgnoreEmptyBlobs *bool `json:"ignoreEmptyBlobs,omitempty" tf:"ignore_empty_blobs,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Pipeline []TriggerBlobEventPipelineParameters `json:"pipeline" tf:"pipeline,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

type TriggerBlobEventPipelineObservation struct {
}

type TriggerBlobEventPipelineParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

// TriggerBlobEventSpec defines the desired state of TriggerBlobEvent
type TriggerBlobEventSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TriggerBlobEventParameters `json:"forProvider"`
}

// TriggerBlobEventStatus defines the observed state of TriggerBlobEvent.
type TriggerBlobEventStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TriggerBlobEventObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerBlobEvent is the Schema for the TriggerBlobEvents API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type TriggerBlobEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TriggerBlobEventSpec   `json:"spec"`
	Status            TriggerBlobEventStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TriggerBlobEventList contains a list of TriggerBlobEvents
type TriggerBlobEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TriggerBlobEvent `json:"items"`
}

// Repository type metadata.
var (
	TriggerBlobEvent_Kind             = "TriggerBlobEvent"
	TriggerBlobEvent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TriggerBlobEvent_Kind}.String()
	TriggerBlobEvent_KindAPIVersion   = TriggerBlobEvent_Kind + "." + CRDGroupVersion.String()
	TriggerBlobEvent_GroupVersionKind = CRDGroupVersion.WithKind(TriggerBlobEvent_Kind)
)

func init() {
	SchemeBuilder.Register(&TriggerBlobEvent{}, &TriggerBlobEventList{})
}

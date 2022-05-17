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

type KeyPropertyObservation struct {
}

type KeyPropertyParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ReferenceDataSetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ReferenceDataSetParameters struct {

	// +kubebuilder:validation:Optional
	DataStringComparisonBehavior *string `json:"dataStringComparisonBehavior,omitempty" tf:"data_string_comparison_behavior,omitempty"`

	// +kubebuilder:validation:Required
	KeyProperty []KeyPropertyParameters `json:"keyProperty" tf:"key_property,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TimeSeriesInsightsEnvironmentID *string `json:"timeSeriesInsightsEnvironmentId" tf:"time_series_insights_environment_id,omitempty"`
}

// ReferenceDataSetSpec defines the desired state of ReferenceDataSet
type ReferenceDataSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReferenceDataSetParameters `json:"forProvider"`
}

// ReferenceDataSetStatus defines the observed state of ReferenceDataSet.
type ReferenceDataSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReferenceDataSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReferenceDataSet is the Schema for the ReferenceDataSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ReferenceDataSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReferenceDataSetSpec   `json:"spec"`
	Status            ReferenceDataSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReferenceDataSetList contains a list of ReferenceDataSets
type ReferenceDataSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReferenceDataSet `json:"items"`
}

// Repository type metadata.
var (
	ReferenceDataSet_Kind             = "ReferenceDataSet"
	ReferenceDataSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReferenceDataSet_Kind}.String()
	ReferenceDataSet_KindAPIVersion   = ReferenceDataSet_Kind + "." + CRDGroupVersion.String()
	ReferenceDataSet_GroupVersionKind = CRDGroupVersion.WithKind(ReferenceDataSet_Kind)
)

func init() {
	SchemeBuilder.Register(&ReferenceDataSet{}, &ReferenceDataSetList{})
}

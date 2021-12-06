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

type AnalyzerEdgeModuleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AnalyzerEdgeModuleParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	VideoAnalyzerName *string `json:"videoAnalyzerName" tf:"video_analyzer_name,omitempty"`
}

// AnalyzerEdgeModuleSpec defines the desired state of AnalyzerEdgeModule
type AnalyzerEdgeModuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AnalyzerEdgeModuleParameters `json:"forProvider"`
}

// AnalyzerEdgeModuleStatus defines the observed state of AnalyzerEdgeModule.
type AnalyzerEdgeModuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AnalyzerEdgeModuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyzerEdgeModule is the Schema for the AnalyzerEdgeModules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AnalyzerEdgeModule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalyzerEdgeModuleSpec   `json:"spec"`
	Status            AnalyzerEdgeModuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyzerEdgeModuleList contains a list of AnalyzerEdgeModules
type AnalyzerEdgeModuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnalyzerEdgeModule `json:"items"`
}

// Repository type metadata.
var (
	AnalyzerEdgeModule_Kind             = "AnalyzerEdgeModule"
	AnalyzerEdgeModule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AnalyzerEdgeModule_Kind}.String()
	AnalyzerEdgeModule_KindAPIVersion   = AnalyzerEdgeModule_Kind + "." + CRDGroupVersion.String()
	AnalyzerEdgeModule_GroupVersionKind = CRDGroupVersion.WithKind(AnalyzerEdgeModule_Kind)
)

func init() {
	SchemeBuilder.Register(&AnalyzerEdgeModule{}, &AnalyzerEdgeModuleList{})
}
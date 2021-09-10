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

type AutoPauseObservation struct {
}

type AutoPauseParameters struct {
	DelayInMinutes int64 `json:"delayInMinutes" tf:"delay_in_minutes"`
}

type AutoScaleObservation struct {
}

type AutoScaleParameters struct {
	MaxNodeCount int64 `json:"maxNodeCount" tf:"max_node_count"`

	MinNodeCount int64 `json:"minNodeCount" tf:"min_node_count"`
}

type LibraryRequirementObservation struct {
}

type LibraryRequirementParameters struct {
	Content string `json:"content" tf:"content"`

	Filename string `json:"filename" tf:"filename"`
}

type SynapseSparkPoolObservation struct {
}

type SynapseSparkPoolParameters struct {
	AutoPause []AutoPauseParameters `json:"autoPause,omitempty" tf:"auto_pause"`

	AutoScale []AutoScaleParameters `json:"autoScale,omitempty" tf:"auto_scale"`

	LibraryRequirement []LibraryRequirementParameters `json:"libraryRequirement,omitempty" tf:"library_requirement"`

	Name string `json:"name" tf:"name"`

	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`

	NodeSize string `json:"nodeSize" tf:"node_size"`

	NodeSizeFamily string `json:"nodeSizeFamily" tf:"node_size_family"`

	SparkEventsFolder *string `json:"sparkEventsFolder,omitempty" tf:"spark_events_folder"`

	SparkLogFolder *string `json:"sparkLogFolder,omitempty" tf:"spark_log_folder"`

	SparkVersion *string `json:"sparkVersion,omitempty" tf:"spark_version"`

	SynapseWorkspaceID string `json:"synapseWorkspaceId" tf:"synapse_workspace_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

// SynapseSparkPoolSpec defines the desired state of SynapseSparkPool
type SynapseSparkPoolSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SynapseSparkPoolParameters `json:"forProvider"`
}

// SynapseSparkPoolStatus defines the observed state of SynapseSparkPool.
type SynapseSparkPoolStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SynapseSparkPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SynapseSparkPool is the Schema for the SynapseSparkPools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SynapseSparkPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SynapseSparkPoolSpec   `json:"spec"`
	Status            SynapseSparkPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SynapseSparkPoolList contains a list of SynapseSparkPools
type SynapseSparkPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SynapseSparkPool `json:"items"`
}

// Repository type metadata.
var (
	SynapseSparkPoolKind             = "SynapseSparkPool"
	SynapseSparkPoolGroupKind        = schema.GroupKind{Group: Group, Kind: SynapseSparkPoolKind}.String()
	SynapseSparkPoolKindAPIVersion   = SynapseSparkPoolKind + "." + GroupVersion.String()
	SynapseSparkPoolGroupVersionKind = GroupVersion.WithKind(SynapseSparkPoolKind)
)

func init() {
	SchemeBuilder.Register(&SynapseSparkPool{}, &SynapseSparkPoolList{})
}

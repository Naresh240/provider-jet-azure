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

type GroupTemplateDeploymentObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	OutputContent *string `json:"outputContent,omitempty" tf:"output_content,omitempty"`
}

type GroupTemplateDeploymentParameters struct {

	// +kubebuilder:validation:Optional
	DebugLevel *string `json:"debugLevel,omitempty" tf:"debug_level,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	ManagementGroupID *string `json:"managementGroupId" tf:"management_group_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ParametersContent *string `json:"parametersContent,omitempty" tf:"parameters_content,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateContent *string `json:"templateContent,omitempty" tf:"template_content,omitempty"`

	// +kubebuilder:validation:Optional
	TemplateSpecVersionID *string `json:"templateSpecVersionId,omitempty" tf:"template_spec_version_id,omitempty"`
}

// GroupTemplateDeploymentSpec defines the desired state of GroupTemplateDeployment
type GroupTemplateDeploymentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupTemplateDeploymentParameters `json:"forProvider"`
}

// GroupTemplateDeploymentStatus defines the observed state of GroupTemplateDeployment.
type GroupTemplateDeploymentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupTemplateDeploymentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupTemplateDeployment is the Schema for the GroupTemplateDeployments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type GroupTemplateDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupTemplateDeploymentSpec   `json:"spec"`
	Status            GroupTemplateDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupTemplateDeploymentList contains a list of GroupTemplateDeployments
type GroupTemplateDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupTemplateDeployment `json:"items"`
}

// Repository type metadata.
var (
	GroupTemplateDeployment_Kind             = "GroupTemplateDeployment"
	GroupTemplateDeployment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupTemplateDeployment_Kind}.String()
	GroupTemplateDeployment_KindAPIVersion   = GroupTemplateDeployment_Kind + "." + CRDGroupVersion.String()
	GroupTemplateDeployment_GroupVersionKind = CRDGroupVersion.WithKind(GroupTemplateDeployment_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupTemplateDeployment{}, &GroupTemplateDeploymentList{})
}
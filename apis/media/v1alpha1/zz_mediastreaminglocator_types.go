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

type ContentKeyObservation struct {
}

type ContentKeyParameters struct {
	ContentKeyID *string `json:"contentKeyId,omitempty" tf:"content_key_id"`

	LabelReferenceInStreamingPolicy *string `json:"labelReferenceInStreamingPolicy,omitempty" tf:"label_reference_in_streaming_policy"`

	PolicyName *string `json:"policyName,omitempty" tf:"policy_name"`

	Type *string `json:"type,omitempty" tf:"type"`

	Value *string `json:"value,omitempty" tf:"value"`
}

type MediaStreamingLocatorObservation struct {
}

type MediaStreamingLocatorParameters struct {
	AlternativeMediaID *string `json:"alternativeMediaId,omitempty" tf:"alternative_media_id"`

	AssetName string `json:"assetName" tf:"asset_name"`

	ContentKey []ContentKeyParameters `json:"contentKey,omitempty" tf:"content_key"`

	DefaultContentKeyPolicyName *string `json:"defaultContentKeyPolicyName,omitempty" tf:"default_content_key_policy_name"`

	EndTime *string `json:"endTime,omitempty" tf:"end_time"`

	MediaServicesAccountName string `json:"mediaServicesAccountName" tf:"media_services_account_name"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	StartTime *string `json:"startTime,omitempty" tf:"start_time"`

	StreamingLocatorID *string `json:"streamingLocatorId,omitempty" tf:"streaming_locator_id"`

	StreamingPolicyName string `json:"streamingPolicyName" tf:"streaming_policy_name"`
}

// MediaStreamingLocatorSpec defines the desired state of MediaStreamingLocator
type MediaStreamingLocatorSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MediaStreamingLocatorParameters `json:"forProvider"`
}

// MediaStreamingLocatorStatus defines the observed state of MediaStreamingLocator.
type MediaStreamingLocatorStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MediaStreamingLocatorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MediaStreamingLocator is the Schema for the MediaStreamingLocators API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MediaStreamingLocator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MediaStreamingLocatorSpec   `json:"spec"`
	Status            MediaStreamingLocatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MediaStreamingLocatorList contains a list of MediaStreamingLocators
type MediaStreamingLocatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MediaStreamingLocator `json:"items"`
}

// Repository type metadata.
var (
	MediaStreamingLocatorKind             = "MediaStreamingLocator"
	MediaStreamingLocatorGroupKind        = schema.GroupKind{Group: Group, Kind: MediaStreamingLocatorKind}.String()
	MediaStreamingLocatorKindAPIVersion   = MediaStreamingLocatorKind + "." + GroupVersion.String()
	MediaStreamingLocatorGroupVersionKind = GroupVersion.WithKind(MediaStreamingLocatorKind)
)

func init() {
	SchemeBuilder.Register(&MediaStreamingLocator{}, &MediaStreamingLocatorList{})
}

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

type EventgridSystemTopicObservation struct {
	MetricArmResourceID string `json:"metricArmResourceId" tf:"metric_arm_resource_id"`
}

type EventgridSystemTopicParameters struct {
	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	SourceArmResourceID string `json:"sourceArmResourceId" tf:"source_arm_resource_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TopicType string `json:"topicType" tf:"topic_type"`
}

// EventgridSystemTopicSpec defines the desired state of EventgridSystemTopic
type EventgridSystemTopicSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EventgridSystemTopicParameters `json:"forProvider"`
}

// EventgridSystemTopicStatus defines the observed state of EventgridSystemTopic.
type EventgridSystemTopicStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EventgridSystemTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventgridSystemTopic is the Schema for the EventgridSystemTopics API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EventgridSystemTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventgridSystemTopicSpec   `json:"spec"`
	Status            EventgridSystemTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventgridSystemTopicList contains a list of EventgridSystemTopics
type EventgridSystemTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventgridSystemTopic `json:"items"`
}

// Repository type metadata.
var (
	EventgridSystemTopicKind             = "EventgridSystemTopic"
	EventgridSystemTopicGroupKind        = schema.GroupKind{Group: Group, Kind: EventgridSystemTopicKind}.String()
	EventgridSystemTopicKindAPIVersion   = EventgridSystemTopicKind + "." + GroupVersion.String()
	EventgridSystemTopicGroupVersionKind = GroupVersion.WithKind(EventgridSystemTopicKind)
)

func init() {
	SchemeBuilder.Register(&EventgridSystemTopic{}, &EventgridSystemTopicList{})
}

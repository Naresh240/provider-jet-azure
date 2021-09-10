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

type LogicAppActionHttpObservation struct {
}

type LogicAppActionHttpParameters struct {
	Body *string `json:"body,omitempty" tf:"body"`

	Headers map[string]string `json:"headers,omitempty" tf:"headers"`

	LogicAppID string `json:"logicAppId" tf:"logic_app_id"`

	Method string `json:"method" tf:"method"`

	Name string `json:"name" tf:"name"`

	RunAfter []RunAfterParameters `json:"runAfter,omitempty" tf:"run_after"`

	URI string `json:"uri" tf:"uri"`
}

type RunAfterObservation struct {
}

type RunAfterParameters struct {
	ActionName string `json:"actionName" tf:"action_name"`

	ActionResult string `json:"actionResult" tf:"action_result"`
}

// LogicAppActionHttpSpec defines the desired state of LogicAppActionHttp
type LogicAppActionHttpSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LogicAppActionHttpParameters `json:"forProvider"`
}

// LogicAppActionHttpStatus defines the observed state of LogicAppActionHttp.
type LogicAppActionHttpStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LogicAppActionHttpObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogicAppActionHttp is the Schema for the LogicAppActionHttps API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LogicAppActionHttp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogicAppActionHttpSpec   `json:"spec"`
	Status            LogicAppActionHttpStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogicAppActionHttpList contains a list of LogicAppActionHttps
type LogicAppActionHttpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogicAppActionHttp `json:"items"`
}

// Repository type metadata.
var (
	LogicAppActionHttpKind             = "LogicAppActionHttp"
	LogicAppActionHttpGroupKind        = schema.GroupKind{Group: Group, Kind: LogicAppActionHttpKind}.String()
	LogicAppActionHttpKindAPIVersion   = LogicAppActionHttpKind + "." + GroupVersion.String()
	LogicAppActionHttpGroupVersionKind = GroupVersion.WithKind(LogicAppActionHttpKind)
)

func init() {
	SchemeBuilder.Register(&LogicAppActionHttp{}, &LogicAppActionHttpList{})
}

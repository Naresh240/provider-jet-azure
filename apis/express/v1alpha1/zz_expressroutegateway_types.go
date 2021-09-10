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

type ExpressRouteGatewayObservation struct {
}

type ExpressRouteGatewayParameters struct {
	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	ScaleUnits int64 `json:"scaleUnits" tf:"scale_units"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	VirtualHubID string `json:"virtualHubId" tf:"virtual_hub_id"`
}

// ExpressRouteGatewaySpec defines the desired state of ExpressRouteGateway
type ExpressRouteGatewaySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ExpressRouteGatewayParameters `json:"forProvider"`
}

// ExpressRouteGatewayStatus defines the observed state of ExpressRouteGateway.
type ExpressRouteGatewayStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ExpressRouteGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteGateway is the Schema for the ExpressRouteGateways API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRouteGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteGatewaySpec   `json:"spec"`
	Status            ExpressRouteGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteGatewayList contains a list of ExpressRouteGateways
type ExpressRouteGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRouteGateway `json:"items"`
}

// Repository type metadata.
var (
	ExpressRouteGatewayKind             = "ExpressRouteGateway"
	ExpressRouteGatewayGroupKind        = schema.GroupKind{Group: Group, Kind: ExpressRouteGatewayKind}.String()
	ExpressRouteGatewayKindAPIVersion   = ExpressRouteGatewayKind + "." + GroupVersion.String()
	ExpressRouteGatewayGroupVersionKind = GroupVersion.WithKind(ExpressRouteGatewayKind)
)

func init() {
	SchemeBuilder.Register(&ExpressRouteGateway{}, &ExpressRouteGatewayList{})
}

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

type ApiManagementGatewayApiObservation struct {
}

type ApiManagementGatewayApiParameters struct {
	APIID string `json:"apiId" tf:"api_id"`

	GatewayID string `json:"gatewayId" tf:"gateway_id"`
}

// ApiManagementGatewayApiSpec defines the desired state of ApiManagementGatewayApi
type ApiManagementGatewayApiSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiManagementGatewayApiParameters `json:"forProvider"`
}

// ApiManagementGatewayApiStatus defines the observed state of ApiManagementGatewayApi.
type ApiManagementGatewayApiStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiManagementGatewayApiObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementGatewayApi is the Schema for the ApiManagementGatewayApis API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApiManagementGatewayApi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementGatewayApiSpec   `json:"spec"`
	Status            ApiManagementGatewayApiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiManagementGatewayApiList contains a list of ApiManagementGatewayApis
type ApiManagementGatewayApiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiManagementGatewayApi `json:"items"`
}

// Repository type metadata.
var (
	ApiManagementGatewayApiKind             = "ApiManagementGatewayApi"
	ApiManagementGatewayApiGroupKind        = schema.GroupKind{Group: Group, Kind: ApiManagementGatewayApiKind}.String()
	ApiManagementGatewayApiKindAPIVersion   = ApiManagementGatewayApiKind + "." + GroupVersion.String()
	ApiManagementGatewayApiGroupVersionKind = GroupVersion.WithKind(ApiManagementGatewayApiKind)
)

func init() {
	SchemeBuilder.Register(&ApiManagementGatewayApi{}, &ApiManagementGatewayApiList{})
}

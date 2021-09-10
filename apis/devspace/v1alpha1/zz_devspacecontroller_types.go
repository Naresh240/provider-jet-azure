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

type DevspaceControllerObservation struct {
	DataPlaneFqdn string `json:"dataPlaneFqdn" tf:"data_plane_fqdn"`

	HostSuffix string `json:"hostSuffix" tf:"host_suffix"`
}

type DevspaceControllerParameters struct {
	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	SkuName string `json:"skuName" tf:"sku_name"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TargetContainerHostCredentialsBase64 string `json:"targetContainerHostCredentialsBase64" tf:"target_container_host_credentials_base64"`

	TargetContainerHostResourceID string `json:"targetContainerHostResourceId" tf:"target_container_host_resource_id"`
}

// DevspaceControllerSpec defines the desired state of DevspaceController
type DevspaceControllerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DevspaceControllerParameters `json:"forProvider"`
}

// DevspaceControllerStatus defines the observed state of DevspaceController.
type DevspaceControllerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DevspaceControllerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DevspaceController is the Schema for the DevspaceControllers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DevspaceController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevspaceControllerSpec   `json:"spec"`
	Status            DevspaceControllerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DevspaceControllerList contains a list of DevspaceControllers
type DevspaceControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevspaceController `json:"items"`
}

// Repository type metadata.
var (
	DevspaceControllerKind             = "DevspaceController"
	DevspaceControllerGroupKind        = schema.GroupKind{Group: Group, Kind: DevspaceControllerKind}.String()
	DevspaceControllerKindAPIVersion   = DevspaceControllerKind + "." + GroupVersion.String()
	DevspaceControllerGroupVersionKind = GroupVersion.WithKind(DevspaceControllerKind)
)

func init() {
	SchemeBuilder.Register(&DevspaceController{}, &DevspaceControllerList{})
}

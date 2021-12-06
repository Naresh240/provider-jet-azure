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

type SyncCloudEndpointObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SyncCloudEndpointParameters struct {

	// +kubebuilder:validation:Required
	FileShareName *string `json:"fileShareName" tf:"file_share_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountTenantID *string `json:"storageAccountTenantId,omitempty" tf:"storage_account_tenant_id,omitempty"`

	// +kubebuilder:validation:Required
	StorageSyncGroupID *string `json:"storageSyncGroupId" tf:"storage_sync_group_id,omitempty"`
}

// SyncCloudEndpointSpec defines the desired state of SyncCloudEndpoint
type SyncCloudEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SyncCloudEndpointParameters `json:"forProvider"`
}

// SyncCloudEndpointStatus defines the observed state of SyncCloudEndpoint.
type SyncCloudEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SyncCloudEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SyncCloudEndpoint is the Schema for the SyncCloudEndpoints API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type SyncCloudEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SyncCloudEndpointSpec   `json:"spec"`
	Status            SyncCloudEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SyncCloudEndpointList contains a list of SyncCloudEndpoints
type SyncCloudEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SyncCloudEndpoint `json:"items"`
}

// Repository type metadata.
var (
	SyncCloudEndpoint_Kind             = "SyncCloudEndpoint"
	SyncCloudEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SyncCloudEndpoint_Kind}.String()
	SyncCloudEndpoint_KindAPIVersion   = SyncCloudEndpoint_Kind + "." + CRDGroupVersion.String()
	SyncCloudEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(SyncCloudEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&SyncCloudEndpoint{}, &SyncCloudEndpointList{})
}
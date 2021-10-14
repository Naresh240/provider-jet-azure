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

type StorageContainerObservation struct {
	HasImmutabilityPolicy *bool `json:"hasImmutabilityPolicy,omitempty" tf:"has_immutability_policy,omitempty"`

	HasLegalHold *bool `json:"hasLegalHold,omitempty" tf:"has_legal_hold,omitempty"`

	ResourceManagerID *string `json:"resourceManagerId,omitempty" tf:"resource_manager_id,omitempty"`
}

type StorageContainerParameters struct {

	// +kubebuilder:validation:Optional
	ContainerAccessType *string `json:"containerAccessType,omitempty" tf:"container_access_type,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=StorageAccount
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-tf-azure/apis/rconfig.ExtractResourceName()
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`
}

// StorageContainerSpec defines the desired state of StorageContainer
type StorageContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StorageContainerParameters `json:"forProvider"`
}

// StorageContainerStatus defines the observed state of StorageContainer.
type StorageContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StorageContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StorageContainer is the Schema for the StorageContainers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StorageContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageContainerSpec   `json:"spec"`
	Status            StorageContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageContainerList contains a list of StorageContainers
type StorageContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageContainer `json:"items"`
}

// Repository type metadata.
var (
	StorageContainerKind             = "StorageContainer"
	StorageContainerGroupKind        = schema.GroupKind{Group: Group, Kind: StorageContainerKind}.String()
	StorageContainerKindAPIVersion   = StorageContainerKind + "." + GroupVersion.String()
	StorageContainerGroupVersionKind = GroupVersion.WithKind(StorageContainerKind)
)

func init() {
	SchemeBuilder.Register(&StorageContainer{}, &StorageContainerList{})
}

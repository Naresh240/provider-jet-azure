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

type StorageShareFileObservation struct {
}

type StorageShareFileParameters struct {
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition"`

	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding"`

	ContentMd5 *string `json:"contentMd5,omitempty" tf:"content_md5"`

	ContentType *string `json:"contentType,omitempty" tf:"content_type"`

	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata"`

	Name string `json:"name" tf:"name"`

	Path *string `json:"path,omitempty" tf:"path"`

	Source *string `json:"source,omitempty" tf:"source"`

	StorageShareID string `json:"storageShareId" tf:"storage_share_id"`
}

// StorageShareFileSpec defines the desired state of StorageShareFile
type StorageShareFileSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       StorageShareFileParameters `json:"forProvider"`
}

// StorageShareFileStatus defines the observed state of StorageShareFile.
type StorageShareFileStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          StorageShareFileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StorageShareFile is the Schema for the StorageShareFiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type StorageShareFile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageShareFileSpec   `json:"spec"`
	Status            StorageShareFileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageShareFileList contains a list of StorageShareFiles
type StorageShareFileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageShareFile `json:"items"`
}

// Repository type metadata.
var (
	StorageShareFileKind             = "StorageShareFile"
	StorageShareFileGroupKind        = schema.GroupKind{Group: Group, Kind: StorageShareFileKind}.String()
	StorageShareFileKindAPIVersion   = StorageShareFileKind + "." + GroupVersion.String()
	StorageShareFileGroupVersionKind = GroupVersion.WithKind(StorageShareFileKind)
)

func init() {
	SchemeBuilder.Register(&StorageShareFile{}, &StorageShareFileList{})
}

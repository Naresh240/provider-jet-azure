/*
Copyright 2022 The Crossplane Authors.

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

type DataSetDataLakeGen1Observation struct {
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataSetDataLakeGen1Parameters struct {

	// +kubebuilder:validation:Required
	DataLakeStoreID *string `json:"dataLakeStoreId" tf:"data_lake_store_id,omitempty"`

	// +kubebuilder:validation:Required
	DataShareID *string `json:"dataShareId" tf:"data_share_id,omitempty"`

	// +kubebuilder:validation:Optional
	FileName *string `json:"fileName,omitempty" tf:"file_name,omitempty"`

	// +kubebuilder:validation:Required
	FolderPath *string `json:"folderPath" tf:"folder_path,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// DataSetDataLakeGen1Spec defines the desired state of DataSetDataLakeGen1
type DataSetDataLakeGen1Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetDataLakeGen1Parameters `json:"forProvider"`
}

// DataSetDataLakeGen1Status defines the observed state of DataSetDataLakeGen1.
type DataSetDataLakeGen1Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetDataLakeGen1Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetDataLakeGen1 is the Schema for the DataSetDataLakeGen1s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DataSetDataLakeGen1 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetDataLakeGen1Spec   `json:"spec"`
	Status            DataSetDataLakeGen1Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetDataLakeGen1List contains a list of DataSetDataLakeGen1s
type DataSetDataLakeGen1List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetDataLakeGen1 `json:"items"`
}

// Repository type metadata.
var (
	DataSetDataLakeGen1_Kind             = "DataSetDataLakeGen1"
	DataSetDataLakeGen1_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetDataLakeGen1_Kind}.String()
	DataSetDataLakeGen1_KindAPIVersion   = DataSetDataLakeGen1_Kind + "." + CRDGroupVersion.String()
	DataSetDataLakeGen1_GroupVersionKind = CRDGroupVersion.WithKind(DataSetDataLakeGen1_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetDataLakeGen1{}, &DataSetDataLakeGen1List{})
}
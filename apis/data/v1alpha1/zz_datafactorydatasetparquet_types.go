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

type DataFactoryDatasetParquetAzureBlobStorageLocationObservation struct {
}

type DataFactoryDatasetParquetAzureBlobStorageLocationParameters struct {
	Container string `json:"container" tf:"container"`

	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled"`

	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled"`

	Filename *string `json:"filename,omitempty" tf:"filename"`

	Path string `json:"path" tf:"path"`
}

type DataFactoryDatasetParquetHTTPServerLocationObservation struct {
}

type DataFactoryDatasetParquetHTTPServerLocationParameters struct {
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled"`

	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled"`

	Filename string `json:"filename" tf:"filename"`

	Path string `json:"path" tf:"path"`

	RelativeURL string `json:"relativeUrl" tf:"relative_url"`
}

type DataFactoryDatasetParquetObservation struct {
}

type DataFactoryDatasetParquetParameters struct {
	AdditionalProperties map[string]string `json:"additionalProperties,omitempty" tf:"additional_properties"`

	Annotations []string `json:"annotations,omitempty" tf:"annotations"`

	AzureBlobStorageLocation []DataFactoryDatasetParquetAzureBlobStorageLocationParameters `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location"`

	CompressionCodec *string `json:"compressionCodec,omitempty" tf:"compression_codec"`

	CompressionLevel *string `json:"compressionLevel,omitempty" tf:"compression_level"`

	DataFactoryName string `json:"dataFactoryName" tf:"data_factory_name"`

	Description *string `json:"description,omitempty" tf:"description"`

	Folder *string `json:"folder,omitempty" tf:"folder"`

	HTTPServerLocation []DataFactoryDatasetParquetHTTPServerLocationParameters `json:"httpServerLocation,omitempty" tf:"http_server_location"`

	LinkedServiceName string `json:"linkedServiceName" tf:"linked_service_name"`

	Name string `json:"name" tf:"name"`

	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	SchemaColumn []DataFactoryDatasetParquetSchemaColumnParameters `json:"schemaColumn,omitempty" tf:"schema_column"`
}

type DataFactoryDatasetParquetSchemaColumnObservation struct {
}

type DataFactoryDatasetParquetSchemaColumnParameters struct {
	Description *string `json:"description,omitempty" tf:"description"`

	Name string `json:"name" tf:"name"`

	Type *string `json:"type,omitempty" tf:"type"`
}

// DataFactoryDatasetParquetSpec defines the desired state of DataFactoryDatasetParquet
type DataFactoryDatasetParquetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DataFactoryDatasetParquetParameters `json:"forProvider"`
}

// DataFactoryDatasetParquetStatus defines the observed state of DataFactoryDatasetParquet.
type DataFactoryDatasetParquetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DataFactoryDatasetParquetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryDatasetParquet is the Schema for the DataFactoryDatasetParquets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataFactoryDatasetParquet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactoryDatasetParquetSpec   `json:"spec"`
	Status            DataFactoryDatasetParquetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataFactoryDatasetParquetList contains a list of DataFactoryDatasetParquets
type DataFactoryDatasetParquetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataFactoryDatasetParquet `json:"items"`
}

// Repository type metadata.
var (
	DataFactoryDatasetParquetKind             = "DataFactoryDatasetParquet"
	DataFactoryDatasetParquetGroupKind        = schema.GroupKind{Group: Group, Kind: DataFactoryDatasetParquetKind}.String()
	DataFactoryDatasetParquetKindAPIVersion   = DataFactoryDatasetParquetKind + "." + GroupVersion.String()
	DataFactoryDatasetParquetGroupVersionKind = GroupVersion.WithKind(DataFactoryDatasetParquetKind)
)

func init() {
	SchemeBuilder.Register(&DataFactoryDatasetParquet{}, &DataFactoryDatasetParquetList{})
}

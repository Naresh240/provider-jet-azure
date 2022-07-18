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

type DataSetJSONAzureBlobStorageLocationObservation struct {
}

type DataSetJSONAzureBlobStorageLocationParameters struct {

	// +kubebuilder:validation:Required
	Container *string `json:"container" tf:"container,omitempty"`

	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Filename *string `json:"filename" tf:"filename,omitempty"`

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`
}

type DataSetJSONHTTPServerLocationObservation struct {
}

type DataSetJSONHTTPServerLocationParameters struct {

	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// +kubebuilder:validation:Required
	Filename *string `json:"filename" tf:"filename,omitempty"`

	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// +kubebuilder:validation:Required
	RelativeURL *string `json:"relativeUrl" tf:"relative_url,omitempty"`
}

type DataSetJSONObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataSetJSONParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Optional
	AzureBlobStorageLocation []DataSetJSONAzureBlobStorageLocationParameters `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPServerLocation []DataSetJSONHTTPServerLocationParameters `json:"httpServerLocation,omitempty" tf:"http_server_location,omitempty"`

	// +kubebuilder:validation:Required
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SchemaColumn []DataSetJSONSchemaColumnParameters `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetJSONSchemaColumnObservation struct {
}

type DataSetJSONSchemaColumnParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DataSetJSONSpec defines the desired state of DataSetJSON
type DataSetJSONSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetJSONParameters `json:"forProvider"`
}

// DataSetJSONStatus defines the observed state of DataSetJSON.
type DataSetJSONStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetJSONObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetJSON is the Schema for the DataSetJSONs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DataSetJSON struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetJSONSpec   `json:"spec"`
	Status            DataSetJSONStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetJSONList contains a list of DataSetJSONs
type DataSetJSONList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetJSON `json:"items"`
}

// Repository type metadata.
var (
	DataSetJSON_Kind             = "DataSetJSON"
	DataSetJSON_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetJSON_Kind}.String()
	DataSetJSON_KindAPIVersion   = DataSetJSON_Kind + "." + CRDGroupVersion.String()
	DataSetJSON_GroupVersionKind = CRDGroupVersion.WithKind(DataSetJSON_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetJSON{}, &DataSetJSONList{})
}
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

type DataSetSnowflakeObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataSetSnowflakeParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryName *string `json:"dataFactoryName" tf:"data_factory_name,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

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
	SchemaColumn []DataSetSnowflakeSchemaColumnParameters `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`

	// +kubebuilder:validation:Optional
	SchemaName *string `json:"schemaName,omitempty" tf:"schema_name,omitempty"`

	// +kubebuilder:validation:Optional
	StructureColumn []StructureColumnParameters `json:"structureColumn,omitempty" tf:"structure_column,omitempty"`

	// +kubebuilder:validation:Optional
	TableName *string `json:"tableName,omitempty" tf:"table_name,omitempty"`
}

type DataSetSnowflakeSchemaColumnObservation struct {
}

type DataSetSnowflakeSchemaColumnParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Precision *float64 `json:"precision,omitempty" tf:"precision,omitempty"`

	// +kubebuilder:validation:Optional
	Scale *float64 `json:"scale,omitempty" tf:"scale,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type StructureColumnObservation struct {
}

type StructureColumnParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DataSetSnowflakeSpec defines the desired state of DataSetSnowflake
type DataSetSnowflakeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetSnowflakeParameters `json:"forProvider"`
}

// DataSetSnowflakeStatus defines the observed state of DataSetSnowflake.
type DataSetSnowflakeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetSnowflakeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetSnowflake is the Schema for the DataSetSnowflakes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type DataSetSnowflake struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetSnowflakeSpec   `json:"spec"`
	Status            DataSetSnowflakeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetSnowflakeList contains a list of DataSetSnowflakes
type DataSetSnowflakeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetSnowflake `json:"items"`
}

// Repository type metadata.
var (
	DataSetSnowflake_Kind             = "DataSetSnowflake"
	DataSetSnowflake_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetSnowflake_Kind}.String()
	DataSetSnowflake_KindAPIVersion   = DataSetSnowflake_Kind + "." + CRDGroupVersion.String()
	DataSetSnowflake_GroupVersionKind = CRDGroupVersion.WithKind(DataSetSnowflake_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetSnowflake{}, &DataSetSnowflakeList{})
}

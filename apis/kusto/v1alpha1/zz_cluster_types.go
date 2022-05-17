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

type ClusterObservation struct {
	DataIngestionURI *string `json:"dataIngestionUri,omitempty" tf:"data_ingestion_uri,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ClusterParameters struct {

	// +kubebuilder:validation:Optional
	DoubleEncryptionEnabled *bool `json:"doubleEncryptionEnabled,omitempty" tf:"double_encryption_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EnableDiskEncryption *bool `json:"enableDiskEncryption,omitempty" tf:"enable_disk_encryption,omitempty"`

	// +kubebuilder:validation:Optional
	EnablePurge *bool `json:"enablePurge,omitempty" tf:"enable_purge,omitempty"`

	// +kubebuilder:validation:Optional
	EnableStreamingIngest *bool `json:"enableStreamingIngest,omitempty" tf:"enable_streaming_ingest,omitempty"`

	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Optional
	LanguageExtensions []*string `json:"languageExtensions,omitempty" tf:"language_extensions,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OptimizedAutoScale []OptimizedAutoScaleParameters `json:"optimizedAutoScale,omitempty" tf:"optimized_auto_scale,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Sku []SkuParameters `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TrustedExternalTenants []*string `json:"trustedExternalTenants,omitempty" tf:"trusted_external_tenants,omitempty"`

	// +kubebuilder:validation:Optional
	VirtualNetworkConfiguration []VirtualNetworkConfigurationParameters `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type IdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type OptimizedAutoScaleObservation struct {
}

type OptimizedAutoScaleParameters struct {

	// +kubebuilder:validation:Required
	MaximumInstances *float64 `json:"maximumInstances" tf:"maximum_instances,omitempty"`

	// +kubebuilder:validation:Required
	MinimumInstances *float64 `json:"minimumInstances" tf:"minimum_instances,omitempty"`
}

type SkuObservation struct {
}

type SkuParameters struct {

	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type VirtualNetworkConfigurationObservation struct {
}

type VirtualNetworkConfigurationParameters struct {

	// +kubebuilder:validation:Required
	DataManagementPublicIPID *string `json:"dataManagementPublicIpId" tf:"data_management_public_ip_id,omitempty"`

	// +kubebuilder:validation:Required
	EnginePublicIPID *string `json:"enginePublicIpId" tf:"engine_public_ip_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/network/v1alpha2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}

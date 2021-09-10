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

type IdentityObservation struct {
	PrincipalID string `json:"principalId" tf:"principal_id"`

	TenantID string `json:"tenantId" tf:"tenant_id"`
}

type IdentityParameters struct {
	IdentityIds []string `json:"identityIds,omitempty" tf:"identity_ids"`

	Type string `json:"type" tf:"type"`
}

type KustoClusterObservation struct {
	DataIngestionURI string `json:"dataIngestionUri" tf:"data_ingestion_uri"`

	URI string `json:"uri" tf:"uri"`
}

type KustoClusterParameters struct {
	DoubleEncryptionEnabled *bool `json:"doubleEncryptionEnabled,omitempty" tf:"double_encryption_enabled"`

	EnableDiskEncryption *bool `json:"enableDiskEncryption,omitempty" tf:"enable_disk_encryption"`

	EnablePurge *bool `json:"enablePurge,omitempty" tf:"enable_purge"`

	EnableStreamingIngest *bool `json:"enableStreamingIngest,omitempty" tf:"enable_streaming_ingest"`

	Engine *string `json:"engine,omitempty" tf:"engine"`

	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	LanguageExtensions []string `json:"languageExtensions,omitempty" tf:"language_extensions"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	OptimizedAutoScale []OptimizedAutoScaleParameters `json:"optimizedAutoScale,omitempty" tf:"optimized_auto_scale"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Sku []SkuParameters `json:"sku" tf:"sku"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	TrustedExternalTenants []string `json:"trustedExternalTenants,omitempty" tf:"trusted_external_tenants"`

	VirtualNetworkConfiguration []VirtualNetworkConfigurationParameters `json:"virtualNetworkConfiguration,omitempty" tf:"virtual_network_configuration"`

	Zones []string `json:"zones,omitempty" tf:"zones"`
}

type OptimizedAutoScaleObservation struct {
}

type OptimizedAutoScaleParameters struct {
	MaximumInstances int64 `json:"maximumInstances" tf:"maximum_instances"`

	MinimumInstances int64 `json:"minimumInstances" tf:"minimum_instances"`
}

type SkuObservation struct {
}

type SkuParameters struct {
	Capacity *int64 `json:"capacity,omitempty" tf:"capacity"`

	Name string `json:"name" tf:"name"`
}

type VirtualNetworkConfigurationObservation struct {
}

type VirtualNetworkConfigurationParameters struct {
	DataManagementPublicIPID string `json:"dataManagementPublicIpId" tf:"data_management_public_ip_id"`

	EnginePublicIPID string `json:"enginePublicIpId" tf:"engine_public_ip_id"`

	SubnetID string `json:"subnetId" tf:"subnet_id"`
}

// KustoClusterSpec defines the desired state of KustoCluster
type KustoClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KustoClusterParameters `json:"forProvider"`
}

// KustoClusterStatus defines the observed state of KustoCluster.
type KustoClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KustoClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KustoCluster is the Schema for the KustoClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type KustoCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KustoClusterSpec   `json:"spec"`
	Status            KustoClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KustoClusterList contains a list of KustoClusters
type KustoClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KustoCluster `json:"items"`
}

// Repository type metadata.
var (
	KustoClusterKind             = "KustoCluster"
	KustoClusterGroupKind        = schema.GroupKind{Group: Group, Kind: KustoClusterKind}.String()
	KustoClusterKindAPIVersion   = KustoClusterKind + "." + GroupVersion.String()
	KustoClusterGroupVersionKind = GroupVersion.WithKind(KustoClusterKind)
)

func init() {
	SchemeBuilder.Register(&KustoCluster{}, &KustoClusterList{})
}

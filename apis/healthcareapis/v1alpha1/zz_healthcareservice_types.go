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

type AuthenticationConfigurationObservation struct {
}

type AuthenticationConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// +kubebuilder:validation:Optional
	Authority *string `json:"authority,omitempty" tf:"authority,omitempty"`

	// +kubebuilder:validation:Optional
	SmartProxyEnabled *bool `json:"smartProxyEnabled,omitempty" tf:"smart_proxy_enabled,omitempty"`
}

type CorsConfigurationObservation struct {
}

type CorsConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AllowCredentials *bool `json:"allowCredentials,omitempty" tf:"allow_credentials,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// +kubebuilder:validation:Optional
	MaxAgeInSeconds *float64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds,omitempty"`
}

type HealthcareServiceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HealthcareServiceParameters struct {

	// +kubebuilder:validation:Optional
	AccessPolicyObjectIds []*string `json:"accessPolicyObjectIds,omitempty" tf:"access_policy_object_ids,omitempty"`

	// +kubebuilder:validation:Optional
	AuthenticationConfiguration []AuthenticationConfigurationParameters `json:"authenticationConfiguration,omitempty" tf:"authentication_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	CorsConfiguration []CorsConfigurationParameters `json:"corsConfiguration,omitempty" tf:"cors_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	CosmosDBKeyVaultKeyVersionlessID *string `json:"cosmosdbKeyVaultKeyVersionlessId,omitempty" tf:"cosmosdb_key_vault_key_versionless_id,omitempty"`

	// +kubebuilder:validation:Optional
	CosmosDBThroughput *float64 `json:"cosmosdbThroughput,omitempty" tf:"cosmosdb_throughput,omitempty"`

	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// HealthcareServiceSpec defines the desired state of HealthcareService
type HealthcareServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HealthcareServiceParameters `json:"forProvider"`
}

// HealthcareServiceStatus defines the observed state of HealthcareService.
type HealthcareServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HealthcareServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HealthcareService is the Schema for the HealthcareServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type HealthcareService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HealthcareServiceSpec   `json:"spec"`
	Status            HealthcareServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthcareServiceList contains a list of HealthcareServices
type HealthcareServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthcareService `json:"items"`
}

// Repository type metadata.
var (
	HealthcareService_Kind             = "HealthcareService"
	HealthcareService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HealthcareService_Kind}.String()
	HealthcareService_KindAPIVersion   = HealthcareService_Kind + "." + CRDGroupVersion.String()
	HealthcareService_GroupVersionKind = CRDGroupVersion.WithKind(HealthcareService_Kind)
)

func init() {
	SchemeBuilder.Register(&HealthcareService{}, &HealthcareServiceList{})
}
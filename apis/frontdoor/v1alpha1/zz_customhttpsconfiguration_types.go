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

type CustomHTTPSConfigurationObservation struct {
	MinimumTLSVersion *string `json:"minimumTlsVersion,omitempty" tf:"minimum_tls_version,omitempty"`

	ProvisioningState *string `json:"provisioningState,omitempty" tf:"provisioning_state,omitempty"`

	ProvisioningSubstate *string `json:"provisioningSubstate,omitempty" tf:"provisioning_substate,omitempty"`
}

type CustomHTTPSConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AzureKeyVaultCertificateSecretName *string `json:"azureKeyVaultCertificateSecretName,omitempty" tf:"azure_key_vault_certificate_secret_name,omitempty"`

	// +kubebuilder:validation:Optional
	AzureKeyVaultCertificateSecretVersion *string `json:"azureKeyVaultCertificateSecretVersion,omitempty" tf:"azure_key_vault_certificate_secret_version,omitempty"`

	// +kubebuilder:validation:Optional
	AzureKeyVaultCertificateVaultID *string `json:"azureKeyVaultCertificateVaultId,omitempty" tf:"azure_key_vault_certificate_vault_id,omitempty"`

	// +kubebuilder:validation:Optional
	CertificateSource *string `json:"certificateSource,omitempty" tf:"certificate_source,omitempty"`
}

type CustomHttpsConfigurationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CustomHttpsConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	CustomHTTPSConfiguration []CustomHTTPSConfigurationParameters `json:"customHttpsConfiguration,omitempty" tf:"custom_https_configuration,omitempty"`

	// +kubebuilder:validation:Required
	CustomHTTPSProvisioningEnabled *bool `json:"customHttpsProvisioningEnabled" tf:"custom_https_provisioning_enabled,omitempty"`

	// +kubebuilder:validation:Required
	FrontendEndpointID *string `json:"frontendEndpointId" tf:"frontend_endpoint_id,omitempty"`
}

// CustomHttpsConfigurationSpec defines the desired state of CustomHttpsConfiguration
type CustomHttpsConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomHttpsConfigurationParameters `json:"forProvider"`
}

// CustomHttpsConfigurationStatus defines the observed state of CustomHttpsConfiguration.
type CustomHttpsConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomHttpsConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CustomHttpsConfiguration is the Schema for the CustomHttpsConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type CustomHttpsConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomHttpsConfigurationSpec   `json:"spec"`
	Status            CustomHttpsConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomHttpsConfigurationList contains a list of CustomHttpsConfigurations
type CustomHttpsConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomHttpsConfiguration `json:"items"`
}

// Repository type metadata.
var (
	CustomHttpsConfiguration_Kind             = "CustomHttpsConfiguration"
	CustomHttpsConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomHttpsConfiguration_Kind}.String()
	CustomHttpsConfiguration_KindAPIVersion   = CustomHttpsConfiguration_Kind + "." + CRDGroupVersion.String()
	CustomHttpsConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(CustomHttpsConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomHttpsConfiguration{}, &CustomHttpsConfigurationList{})
}
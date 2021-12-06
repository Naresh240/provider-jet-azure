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

type ManagementCertificateObservation struct {
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`
}

type ManagementCertificateParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Optional
	DataSecretRef *v1.SecretKeySelector `json:"dataSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KeyVaultIdentityClientID *string `json:"keyVaultIdentityClientId,omitempty" tf:"key_vault_identity_client_id,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultSecretID *string `json:"keyVaultSecretId,omitempty" tf:"key_vault_secret_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// ManagementCertificateSpec defines the desired state of ManagementCertificate
type ManagementCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementCertificateParameters `json:"forProvider"`
}

// ManagementCertificateStatus defines the observed state of ManagementCertificate.
type ManagementCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementCertificate is the Schema for the ManagementCertificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ManagementCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementCertificateSpec   `json:"spec"`
	Status            ManagementCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementCertificateList contains a list of ManagementCertificates
type ManagementCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementCertificate `json:"items"`
}

// Repository type metadata.
var (
	ManagementCertificate_Kind             = "ManagementCertificate"
	ManagementCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagementCertificate_Kind}.String()
	ManagementCertificate_KindAPIVersion   = ManagementCertificate_Kind + "." + CRDGroupVersion.String()
	ManagementCertificate_GroupVersionKind = CRDGroupVersion.WithKind(ManagementCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagementCertificate{}, &ManagementCertificateList{})
}
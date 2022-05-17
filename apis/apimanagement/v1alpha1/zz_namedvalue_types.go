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

type NamedValueObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NamedValueParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Secret *bool `json:"secret,omitempty" tf:"secret,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	ValueFromKeyVault []ValueFromKeyVaultParameters `json:"valueFromKeyVault,omitempty" tf:"value_from_key_vault,omitempty"`

	// +kubebuilder:validation:Optional
	ValueSecretRef *v1.SecretKeySelector `json:"valueSecretRef,omitempty" tf:"-"`
}

type ValueFromKeyVaultObservation struct {
}

type ValueFromKeyVaultParameters struct {

	// +kubebuilder:validation:Optional
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// +kubebuilder:validation:Required
	SecretID *string `json:"secretId" tf:"secret_id,omitempty"`
}

// NamedValueSpec defines the desired state of NamedValue
type NamedValueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NamedValueParameters `json:"forProvider"`
}

// NamedValueStatus defines the observed state of NamedValue.
type NamedValueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NamedValueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NamedValue is the Schema for the NamedValues API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type NamedValue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamedValueSpec   `json:"spec"`
	Status            NamedValueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NamedValueList contains a list of NamedValues
type NamedValueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamedValue `json:"items"`
}

// Repository type metadata.
var (
	NamedValue_Kind             = "NamedValue"
	NamedValue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NamedValue_Kind}.String()
	NamedValue_KindAPIVersion   = NamedValue_Kind + "." + CRDGroupVersion.String()
	NamedValue_GroupVersionKind = CRDGroupVersion.WithKind(NamedValue_Kind)
)

func init() {
	SchemeBuilder.Register(&NamedValue{}, &NamedValueList{})
}

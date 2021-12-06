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

type ManagementAuthorizationServerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ManagementAuthorizationServerParameters struct {

	// +kubebuilder:validation:Required
	APIManagementName *string `json:"apiManagementName" tf:"api_management_name,omitempty"`

	// +kubebuilder:validation:Required
	AuthorizationEndpoint *string `json:"authorizationEndpoint" tf:"authorization_endpoint,omitempty"`

	// +kubebuilder:validation:Required
	AuthorizationMethods []*string `json:"authorizationMethods" tf:"authorization_methods,omitempty"`

	// +kubebuilder:validation:Optional
	BearerTokenSendingMethods []*string `json:"bearerTokenSendingMethods,omitempty" tf:"bearer_token_sending_methods,omitempty"`

	// +kubebuilder:validation:Optional
	ClientAuthenticationMethod []*string `json:"clientAuthenticationMethod,omitempty" tf:"client_authentication_method,omitempty"`

	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// +kubebuilder:validation:Required
	ClientRegistrationEndpoint *string `json:"clientRegistrationEndpoint" tf:"client_registration_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	ClientSecretSecretRef *v1.SecretKeySelector `json:"clientSecretSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DefaultScope *string `json:"defaultScope,omitempty" tf:"default_scope,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// +kubebuilder:validation:Required
	GrantTypes []*string `json:"grantTypes" tf:"grant_types,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceOwnerPasswordSecretRef *v1.SecretKeySelector `json:"resourceOwnerPasswordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceOwnerUsername *string `json:"resourceOwnerUsername,omitempty" tf:"resource_owner_username,omitempty"`

	// +kubebuilder:validation:Optional
	SupportState *bool `json:"supportState,omitempty" tf:"support_state,omitempty"`

	// +kubebuilder:validation:Optional
	TokenBodyParameter []TokenBodyParameterParameters `json:"tokenBodyParameter,omitempty" tf:"token_body_parameter,omitempty"`

	// +kubebuilder:validation:Optional
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`
}

type TokenBodyParameterObservation struct {
}

type TokenBodyParameterParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

// ManagementAuthorizationServerSpec defines the desired state of ManagementAuthorizationServer
type ManagementAuthorizationServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagementAuthorizationServerParameters `json:"forProvider"`
}

// ManagementAuthorizationServerStatus defines the observed state of ManagementAuthorizationServer.
type ManagementAuthorizationServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagementAuthorizationServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementAuthorizationServer is the Schema for the ManagementAuthorizationServers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ManagementAuthorizationServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementAuthorizationServerSpec   `json:"spec"`
	Status            ManagementAuthorizationServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagementAuthorizationServerList contains a list of ManagementAuthorizationServers
type ManagementAuthorizationServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagementAuthorizationServer `json:"items"`
}

// Repository type metadata.
var (
	ManagementAuthorizationServer_Kind             = "ManagementAuthorizationServer"
	ManagementAuthorizationServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagementAuthorizationServer_Kind}.String()
	ManagementAuthorizationServer_KindAPIVersion   = ManagementAuthorizationServer_Kind + "." + CRDGroupVersion.String()
	ManagementAuthorizationServer_GroupVersionKind = CRDGroupVersion.WithKind(ManagementAuthorizationServer_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagementAuthorizationServer{}, &ManagementAuthorizationServerList{})
}
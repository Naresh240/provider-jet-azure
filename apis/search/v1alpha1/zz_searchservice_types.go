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
	Type string `json:"type" tf:"type"`
}

type QueryKeysObservation struct {
	Key string `json:"key" tf:"key"`

	Name string `json:"name" tf:"name"`
}

type QueryKeysParameters struct {
}

type SearchServiceObservation struct {
	PrimaryKey string `json:"primaryKey" tf:"primary_key"`

	QueryKeys []QueryKeysObservation `json:"queryKeys" tf:"query_keys"`

	SecondaryKey string `json:"secondaryKey" tf:"secondary_key"`
}

type SearchServiceParameters struct {
	AllowedIps []string `json:"allowedIps,omitempty" tf:"allowed_ips"`

	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity"`

	Location string `json:"location" tf:"location"`

	Name string `json:"name" tf:"name"`

	PartitionCount *int64 `json:"partitionCount,omitempty" tf:"partition_count"`

	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled"`

	ReplicaCount *int64 `json:"replicaCount,omitempty" tf:"replica_count"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	Sku string `json:"sku" tf:"sku"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`
}

// SearchServiceSpec defines the desired state of SearchService
type SearchServiceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SearchServiceParameters `json:"forProvider"`
}

// SearchServiceStatus defines the observed state of SearchService.
type SearchServiceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SearchServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SearchService is the Schema for the SearchServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SearchService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SearchServiceSpec   `json:"spec"`
	Status            SearchServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SearchServiceList contains a list of SearchServices
type SearchServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SearchService `json:"items"`
}

// Repository type metadata.
var (
	SearchServiceKind             = "SearchService"
	SearchServiceGroupKind        = schema.GroupKind{Group: Group, Kind: SearchServiceKind}.String()
	SearchServiceKindAPIVersion   = SearchServiceKind + "." + GroupVersion.String()
	SearchServiceGroupVersionKind = GroupVersion.WithKind(SearchServiceKind)
)

func init() {
	SchemeBuilder.Register(&SearchService{}, &SearchServiceList{})
}

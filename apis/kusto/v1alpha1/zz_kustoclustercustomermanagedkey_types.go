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

type KustoClusterCustomerManagedKeyObservation struct {
}

type KustoClusterCustomerManagedKeyParameters struct {
	ClusterID string `json:"clusterId" tf:"cluster_id"`

	KeyName string `json:"keyName" tf:"key_name"`

	KeyVaultID string `json:"keyVaultId" tf:"key_vault_id"`

	KeyVersion string `json:"keyVersion" tf:"key_version"`

	UserIdentity *string `json:"userIdentity,omitempty" tf:"user_identity"`
}

// KustoClusterCustomerManagedKeySpec defines the desired state of KustoClusterCustomerManagedKey
type KustoClusterCustomerManagedKeySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       KustoClusterCustomerManagedKeyParameters `json:"forProvider"`
}

// KustoClusterCustomerManagedKeyStatus defines the observed state of KustoClusterCustomerManagedKey.
type KustoClusterCustomerManagedKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          KustoClusterCustomerManagedKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KustoClusterCustomerManagedKey is the Schema for the KustoClusterCustomerManagedKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type KustoClusterCustomerManagedKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KustoClusterCustomerManagedKeySpec   `json:"spec"`
	Status            KustoClusterCustomerManagedKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KustoClusterCustomerManagedKeyList contains a list of KustoClusterCustomerManagedKeys
type KustoClusterCustomerManagedKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KustoClusterCustomerManagedKey `json:"items"`
}

// Repository type metadata.
var (
	KustoClusterCustomerManagedKeyKind             = "KustoClusterCustomerManagedKey"
	KustoClusterCustomerManagedKeyGroupKind        = schema.GroupKind{Group: Group, Kind: KustoClusterCustomerManagedKeyKind}.String()
	KustoClusterCustomerManagedKeyKindAPIVersion   = KustoClusterCustomerManagedKeyKind + "." + GroupVersion.String()
	KustoClusterCustomerManagedKeyGroupVersionKind = GroupVersion.WithKind(KustoClusterCustomerManagedKeyKind)
)

func init() {
	SchemeBuilder.Register(&KustoClusterCustomerManagedKey{}, &KustoClusterCustomerManagedKeyList{})
}

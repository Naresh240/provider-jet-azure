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

type BackupContainerStorageAccountObservation struct {
}

type BackupContainerStorageAccountParameters struct {
	RecoveryVaultName string `json:"recoveryVaultName" tf:"recovery_vault_name"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`

	StorageAccountID string `json:"storageAccountId" tf:"storage_account_id"`
}

// BackupContainerStorageAccountSpec defines the desired state of BackupContainerStorageAccount
type BackupContainerStorageAccountSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BackupContainerStorageAccountParameters `json:"forProvider"`
}

// BackupContainerStorageAccountStatus defines the observed state of BackupContainerStorageAccount.
type BackupContainerStorageAccountStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BackupContainerStorageAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupContainerStorageAccount is the Schema for the BackupContainerStorageAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupContainerStorageAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupContainerStorageAccountSpec   `json:"spec"`
	Status            BackupContainerStorageAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupContainerStorageAccountList contains a list of BackupContainerStorageAccounts
type BackupContainerStorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupContainerStorageAccount `json:"items"`
}

// Repository type metadata.
var (
	BackupContainerStorageAccountKind             = "BackupContainerStorageAccount"
	BackupContainerStorageAccountGroupKind        = schema.GroupKind{Group: Group, Kind: BackupContainerStorageAccountKind}.String()
	BackupContainerStorageAccountKindAPIVersion   = BackupContainerStorageAccountKind + "." + GroupVersion.String()
	BackupContainerStorageAccountGroupVersionKind = GroupVersion.WithKind(BackupContainerStorageAccountKind)
)

func init() {
	SchemeBuilder.Register(&BackupContainerStorageAccount{}, &BackupContainerStorageAccountList{})
}

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

type ModuleObservation struct {
	Version string `json:"version" tf:"version"`
}

type ModuleParameters struct {
	Args *string `json:"args,omitempty" tf:"args"`

	Name string `json:"name" tf:"name"`
}

type RedisEnterpriseDatabaseObservation struct {
	PrimaryAccessKey string `json:"primaryAccessKey" tf:"primary_access_key"`

	SecondaryAccessKey string `json:"secondaryAccessKey" tf:"secondary_access_key"`
}

type RedisEnterpriseDatabaseParameters struct {
	ClientProtocol *string `json:"clientProtocol,omitempty" tf:"client_protocol"`

	ClusterID string `json:"clusterId" tf:"cluster_id"`

	ClusteringPolicy *string `json:"clusteringPolicy,omitempty" tf:"clustering_policy"`

	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy"`

	Module []ModuleParameters `json:"module,omitempty" tf:"module"`

	Name *string `json:"name,omitempty" tf:"name"`

	Port *int64 `json:"port,omitempty" tf:"port"`

	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

// RedisEnterpriseDatabaseSpec defines the desired state of RedisEnterpriseDatabase
type RedisEnterpriseDatabaseSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RedisEnterpriseDatabaseParameters `json:"forProvider"`
}

// RedisEnterpriseDatabaseStatus defines the observed state of RedisEnterpriseDatabase.
type RedisEnterpriseDatabaseStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RedisEnterpriseDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisEnterpriseDatabase is the Schema for the RedisEnterpriseDatabases API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type RedisEnterpriseDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisEnterpriseDatabaseSpec   `json:"spec"`
	Status            RedisEnterpriseDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisEnterpriseDatabaseList contains a list of RedisEnterpriseDatabases
type RedisEnterpriseDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisEnterpriseDatabase `json:"items"`
}

// Repository type metadata.
var (
	RedisEnterpriseDatabaseKind             = "RedisEnterpriseDatabase"
	RedisEnterpriseDatabaseGroupKind        = schema.GroupKind{Group: Group, Kind: RedisEnterpriseDatabaseKind}.String()
	RedisEnterpriseDatabaseKindAPIVersion   = RedisEnterpriseDatabaseKind + "." + GroupVersion.String()
	RedisEnterpriseDatabaseGroupVersionKind = GroupVersion.WithKind(RedisEnterpriseDatabaseKind)
)

func init() {
	SchemeBuilder.Register(&RedisEnterpriseDatabase{}, &RedisEnterpriseDatabaseList{})
}

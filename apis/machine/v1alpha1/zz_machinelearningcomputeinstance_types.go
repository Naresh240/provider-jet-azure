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

type AssignToUserObservation struct {
}

type AssignToUserParameters struct {
	ObjectID *string `json:"objectId,omitempty" tf:"object_id"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id"`
}

type MachineLearningComputeInstanceIdentityObservation struct {
	PrincipalID string `json:"principalId" tf:"principal_id"`

	TenantID string `json:"tenantId" tf:"tenant_id"`
}

type MachineLearningComputeInstanceIdentityParameters struct {
	IdentityIds []string `json:"identityIds,omitempty" tf:"identity_ids"`

	Type string `json:"type" tf:"type"`
}

type MachineLearningComputeInstanceObservation struct {
}

type MachineLearningComputeInstanceParameters struct {
	AssignToUser []AssignToUserParameters `json:"assignToUser,omitempty" tf:"assign_to_user"`

	AuthorizationType *string `json:"authorizationType,omitempty" tf:"authorization_type"`

	Description *string `json:"description,omitempty" tf:"description"`

	Identity []MachineLearningComputeInstanceIdentityParameters `json:"identity,omitempty" tf:"identity"`

	Location string `json:"location" tf:"location"`

	MachineLearningWorkspaceID string `json:"machineLearningWorkspaceId" tf:"machine_learning_workspace_id"`

	Name string `json:"name" tf:"name"`

	SSH []SSHParameters `json:"ssh,omitempty" tf:"ssh"`

	SubnetResourceID *string `json:"subnetResourceId,omitempty" tf:"subnet_resource_id"`

	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	VirtualMachineSize string `json:"virtualMachineSize" tf:"virtual_machine_size"`
}

type SSHObservation struct {
	Port int64 `json:"port" tf:"port"`

	Username string `json:"username" tf:"username"`
}

type SSHParameters struct {
	PublicKey string `json:"publicKey" tf:"public_key"`
}

// MachineLearningComputeInstanceSpec defines the desired state of MachineLearningComputeInstance
type MachineLearningComputeInstanceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MachineLearningComputeInstanceParameters `json:"forProvider"`
}

// MachineLearningComputeInstanceStatus defines the observed state of MachineLearningComputeInstance.
type MachineLearningComputeInstanceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MachineLearningComputeInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MachineLearningComputeInstance is the Schema for the MachineLearningComputeInstances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MachineLearningComputeInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MachineLearningComputeInstanceSpec   `json:"spec"`
	Status            MachineLearningComputeInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MachineLearningComputeInstanceList contains a list of MachineLearningComputeInstances
type MachineLearningComputeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MachineLearningComputeInstance `json:"items"`
}

// Repository type metadata.
var (
	MachineLearningComputeInstanceKind             = "MachineLearningComputeInstance"
	MachineLearningComputeInstanceGroupKind        = schema.GroupKind{Group: Group, Kind: MachineLearningComputeInstanceKind}.String()
	MachineLearningComputeInstanceKindAPIVersion   = MachineLearningComputeInstanceKind + "." + GroupVersion.String()
	MachineLearningComputeInstanceGroupVersionKind = GroupVersion.WithKind(MachineLearningComputeInstanceKind)
)

func init() {
	SchemeBuilder.Register(&MachineLearningComputeInstance{}, &MachineLearningComputeInstanceList{})
}

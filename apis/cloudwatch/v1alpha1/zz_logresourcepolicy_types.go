/*
Copyright 2021 The Crossplane Authors.

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

type LogResourcePolicyObservation struct {
}

type LogResourcePolicyParameters struct {

	// +kubebuilder:validation:Required
	PolicyDocument *string `json:"policyDocument" tf:"policy_document,omitempty"`

	// +kubebuilder:validation:Required
	PolicyName *string `json:"policyName" tf:"policy_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// LogResourcePolicySpec defines the desired state of LogResourcePolicy
type LogResourcePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LogResourcePolicyParameters `json:"forProvider"`
}

// LogResourcePolicyStatus defines the observed state of LogResourcePolicy.
type LogResourcePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LogResourcePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LogResourcePolicy is the Schema for the LogResourcePolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LogResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogResourcePolicySpec   `json:"spec"`
	Status            LogResourcePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LogResourcePolicyList contains a list of LogResourcePolicys
type LogResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LogResourcePolicy `json:"items"`
}

// Repository type metadata.
var (
	LogResourcePolicy_Kind             = "LogResourcePolicy"
	LogResourcePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LogResourcePolicy_Kind}.String()
	LogResourcePolicy_KindAPIVersion   = LogResourcePolicy_Kind + "." + CRDGroupVersion.String()
	LogResourcePolicy_GroupVersionKind = CRDGroupVersion.WithKind(LogResourcePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&LogResourcePolicy{}, &LogResourcePolicyList{})
}
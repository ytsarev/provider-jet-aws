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

type TopicPolicyObservation struct {
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`
}

type TopicPolicyParameters struct {

	// +kubebuilder:validation:Required
	Arn *string `json:"arn" tf:"arn,omitempty"`

	// +kubebuilder:validation:Required
	Policy *string `json:"policy" tf:"policy,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// TopicPolicySpec defines the desired state of TopicPolicy
type TopicPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TopicPolicyParameters `json:"forProvider"`
}

// TopicPolicyStatus defines the observed state of TopicPolicy.
type TopicPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TopicPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TopicPolicy is the Schema for the TopicPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type TopicPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TopicPolicySpec   `json:"spec"`
	Status            TopicPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TopicPolicyList contains a list of TopicPolicys
type TopicPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TopicPolicy `json:"items"`
}

// Repository type metadata.
var (
	TopicPolicy_Kind             = "TopicPolicy"
	TopicPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TopicPolicy_Kind}.String()
	TopicPolicy_KindAPIVersion   = TopicPolicy_Kind + "." + CRDGroupVersion.String()
	TopicPolicy_GroupVersionKind = CRDGroupVersion.WithKind(TopicPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&TopicPolicy{}, &TopicPolicyList{})
}
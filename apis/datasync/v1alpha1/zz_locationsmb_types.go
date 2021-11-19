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

type LocationSmbMountOptionsObservation struct {
}

type LocationSmbMountOptionsParameters struct {

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type LocationSmbObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type LocationSmbParameters struct {

	// +kubebuilder:validation:Required
	AgentArns []*string `json:"agentArns" tf:"agent_arns,omitempty"`

	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// +kubebuilder:validation:Optional
	MountOptions []LocationSmbMountOptionsParameters `json:"mountOptions,omitempty" tf:"mount_options,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ServerHostname *string `json:"serverHostname" tf:"server_hostname,omitempty"`

	// +kubebuilder:validation:Required
	Subdirectory *string `json:"subdirectory" tf:"subdirectory,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	User *string `json:"user" tf:"user,omitempty"`
}

// LocationSmbSpec defines the desired state of LocationSmb
type LocationSmbSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocationSmbParameters `json:"forProvider"`
}

// LocationSmbStatus defines the observed state of LocationSmb.
type LocationSmbStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocationSmbObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LocationSmb is the Schema for the LocationSmbs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type LocationSmb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LocationSmbSpec   `json:"spec"`
	Status            LocationSmbStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocationSmbList contains a list of LocationSmbs
type LocationSmbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocationSmb `json:"items"`
}

// Repository type metadata.
var (
	LocationSmb_Kind             = "LocationSmb"
	LocationSmb_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LocationSmb_Kind}.String()
	LocationSmb_KindAPIVersion   = LocationSmb_Kind + "." + CRDGroupVersion.String()
	LocationSmb_GroupVersionKind = CRDGroupVersion.WithKind(LocationSmb_Kind)
)

func init() {
	SchemeBuilder.Register(&LocationSmb{}, &LocationSmbList{})
}
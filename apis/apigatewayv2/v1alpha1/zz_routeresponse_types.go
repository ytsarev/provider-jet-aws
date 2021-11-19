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

type RouteResponseObservation struct {
}

type RouteResponseParameters struct {

	// +kubebuilder:validation:Required
	APIID *string `json:"apiId" tf:"api_id,omitempty"`

	// +kubebuilder:validation:Optional
	ModelSelectionExpression *string `json:"modelSelectionExpression,omitempty" tf:"model_selection_expression,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ResponseModels map[string]*string `json:"responseModels,omitempty" tf:"response_models,omitempty"`

	// +kubebuilder:validation:Required
	RouteID *string `json:"routeId" tf:"route_id,omitempty"`

	// +kubebuilder:validation:Required
	RouteResponseKey *string `json:"routeResponseKey" tf:"route_response_key,omitempty"`
}

// RouteResponseSpec defines the desired state of RouteResponse
type RouteResponseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteResponseParameters `json:"forProvider"`
}

// RouteResponseStatus defines the observed state of RouteResponse.
type RouteResponseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteResponseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteResponse is the Schema for the RouteResponses API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type RouteResponse struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteResponseSpec   `json:"spec"`
	Status            RouteResponseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteResponseList contains a list of RouteResponses
type RouteResponseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteResponse `json:"items"`
}

// Repository type metadata.
var (
	RouteResponse_Kind             = "RouteResponse"
	RouteResponse_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteResponse_Kind}.String()
	RouteResponse_KindAPIVersion   = RouteResponse_Kind + "." + CRDGroupVersion.String()
	RouteResponse_GroupVersionKind = CRDGroupVersion.WithKind(RouteResponse_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteResponse{}, &RouteResponseList{})
}
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

type DNSConfigObservation struct {
}

type DNSConfigParameters struct {

	// +kubebuilder:validation:Required
	DNSRecords []DNSRecordsParameters `json:"dnsRecords" tf:"dns_records,omitempty"`

	// +kubebuilder:validation:Required
	NamespaceID *string `json:"namespaceId" tf:"namespace_id,omitempty"`

	// +kubebuilder:validation:Optional
	RoutingPolicy *string `json:"routingPolicy,omitempty" tf:"routing_policy,omitempty"`
}

type DNSRecordsObservation struct {
}

type DNSRecordsParameters struct {

	// +kubebuilder:validation:Required
	TTL *int64 `json:"ttl" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type DiscoveryServiceObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type DiscoveryServiceParameters struct {

	// +kubebuilder:validation:Optional
	DNSConfig []DNSConfigParameters `json:"dnsConfig,omitempty" tf:"dns_config,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckConfig []HealthCheckConfigParameters `json:"healthCheckConfig,omitempty" tf:"health_check_config,omitempty"`

	// +kubebuilder:validation:Optional
	HealthCheckCustomConfig []HealthCheckCustomConfigParameters `json:"healthCheckCustomConfig,omitempty" tf:"health_check_custom_config,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type HealthCheckConfigObservation struct {
}

type HealthCheckConfigParameters struct {

	// +kubebuilder:validation:Optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`

	// +kubebuilder:validation:Optional
	ResourcePath *string `json:"resourcePath,omitempty" tf:"resource_path,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HealthCheckCustomConfigObservation struct {
}

type HealthCheckCustomConfigParameters struct {

	// +kubebuilder:validation:Optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`
}

// DiscoveryServiceSpec defines the desired state of DiscoveryService
type DiscoveryServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiscoveryServiceParameters `json:"forProvider"`
}

// DiscoveryServiceStatus defines the observed state of DiscoveryService.
type DiscoveryServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiscoveryServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DiscoveryService is the Schema for the DiscoveryServices API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type DiscoveryService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiscoveryServiceSpec   `json:"spec"`
	Status            DiscoveryServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiscoveryServiceList contains a list of DiscoveryServices
type DiscoveryServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryService `json:"items"`
}

// Repository type metadata.
var (
	DiscoveryService_Kind             = "DiscoveryService"
	DiscoveryService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiscoveryService_Kind}.String()
	DiscoveryService_KindAPIVersion   = DiscoveryService_Kind + "." + CRDGroupVersion.String()
	DiscoveryService_GroupVersionKind = CRDGroupVersion.WithKind(DiscoveryService_Kind)
)

func init() {
	SchemeBuilder.Register(&DiscoveryService{}, &DiscoveryServiceList{})
}
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

type ReportDefinitionObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type ReportDefinitionParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalArtifacts []*string `json:"additionalArtifacts,omitempty" tf:"additional_artifacts,omitempty"`

	// +kubebuilder:validation:Required
	AdditionalSchemaElements []*string `json:"additionalSchemaElements" tf:"additional_schema_elements,omitempty"`

	// +kubebuilder:validation:Required
	Compression *string `json:"compression" tf:"compression,omitempty"`

	// +kubebuilder:validation:Required
	Format *string `json:"format" tf:"format,omitempty"`

	// +kubebuilder:validation:Optional
	RefreshClosedReports *bool `json:"refreshClosedReports,omitempty" tf:"refresh_closed_reports,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ReportName *string `json:"reportName" tf:"report_name,omitempty"`

	// +kubebuilder:validation:Optional
	ReportVersioning *string `json:"reportVersioning,omitempty" tf:"report_versioning,omitempty"`

	// +kubebuilder:validation:Required
	S3Bucket *string `json:"s3Bucket" tf:"s3_bucket,omitempty"`

	// +kubebuilder:validation:Optional
	S3Prefix *string `json:"s3Prefix,omitempty" tf:"s3_prefix,omitempty"`

	// +kubebuilder:validation:Required
	S3Region *string `json:"s3Region" tf:"s3_region,omitempty"`

	// +kubebuilder:validation:Required
	TimeUnit *string `json:"timeUnit" tf:"time_unit,omitempty"`
}

// ReportDefinitionSpec defines the desired state of ReportDefinition
type ReportDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReportDefinitionParameters `json:"forProvider"`
}

// ReportDefinitionStatus defines the observed state of ReportDefinition.
type ReportDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReportDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ReportDefinition is the Schema for the ReportDefinitions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ReportDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReportDefinitionSpec   `json:"spec"`
	Status            ReportDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReportDefinitionList contains a list of ReportDefinitions
type ReportDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReportDefinition `json:"items"`
}

// Repository type metadata.
var (
	ReportDefinition_Kind             = "ReportDefinition"
	ReportDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ReportDefinition_Kind}.String()
	ReportDefinition_KindAPIVersion   = ReportDefinition_Kind + "." + CRDGroupVersion.String()
	ReportDefinition_GroupVersionKind = CRDGroupVersion.WithKind(ReportDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&ReportDefinition{}, &ReportDefinitionList{})
}
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

type ObjectCopyGrantObservation struct {
}

type ObjectCopyGrantParameters struct {

	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Permissions []*string `json:"permissions" tf:"permissions,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ObjectCopyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	RequestCharged *bool `json:"requestCharged,omitempty" tf:"request_charged,omitempty"`

	SourceVersionID *string `json:"sourceVersionId,omitempty" tf:"source_version_id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type ObjectCopyParameters struct {

	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// +kubebuilder:validation:Required
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	BucketKeyEnabled *bool `json:"bucketKeyEnabled,omitempty" tf:"bucket_key_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// +kubebuilder:validation:Optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// +kubebuilder:validation:Optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// +kubebuilder:validation:Optional
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// +kubebuilder:validation:Optional
	CopyIfMatch *string `json:"copyIfMatch,omitempty" tf:"copy_if_match,omitempty"`

	// +kubebuilder:validation:Optional
	CopyIfModifiedSince *string `json:"copyIfModifiedSince,omitempty" tf:"copy_if_modified_since,omitempty"`

	// +kubebuilder:validation:Optional
	CopyIfNoneMatch *string `json:"copyIfNoneMatch,omitempty" tf:"copy_if_none_match,omitempty"`

	// +kubebuilder:validation:Optional
	CopyIfUnmodifiedSince *string `json:"copyIfUnmodifiedSince,omitempty" tf:"copy_if_unmodified_since,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerAlgorithm *string `json:"customerAlgorithm,omitempty" tf:"customer_algorithm,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerKeyMd5 *string `json:"customerKeyMd5,omitempty" tf:"customer_key_md5,omitempty"`

	// +kubebuilder:validation:Optional
	CustomerKeySecretRef *v1.SecretKeySelector `json:"customerKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// +kubebuilder:validation:Optional
	ExpectedSourceBucketOwner *string `json:"expectedSourceBucketOwner,omitempty" tf:"expected_source_bucket_owner,omitempty"`

	// +kubebuilder:validation:Optional
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// +kubebuilder:validation:Optional
	Grant []ObjectCopyGrantParameters `json:"grant,omitempty" tf:"grant,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	KmsEncryptionContextSecretRef *v1.SecretKeySelector `json:"kmsEncryptionContextSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KmsKeyIDSecretRef *v1.SecretKeySelector `json:"kmsKeyIdSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Optional
	MetadataDirective *string `json:"metadataDirective,omitempty" tf:"metadata_directive,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectLockLegalHoldStatus *string `json:"objectLockLegalHoldStatus,omitempty" tf:"object_lock_legal_hold_status,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectLockMode *string `json:"objectLockMode,omitempty" tf:"object_lock_mode,omitempty"`

	// +kubebuilder:validation:Optional
	ObjectLockRetainUntilDate *string `json:"objectLockRetainUntilDate,omitempty" tf:"object_lock_retain_until_date,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer,omitempty"`

	// +kubebuilder:validation:Optional
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`

	// +kubebuilder:validation:Optional
	SourceCustomerAlgorithm *string `json:"sourceCustomerAlgorithm,omitempty" tf:"source_customer_algorithm,omitempty"`

	// +kubebuilder:validation:Optional
	SourceCustomerKeyMd5 *string `json:"sourceCustomerKeyMd5,omitempty" tf:"source_customer_key_md5,omitempty"`

	// +kubebuilder:validation:Optional
	SourceCustomerKeySecretRef *v1.SecretKeySelector `json:"sourceCustomerKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// +kubebuilder:validation:Optional
	TaggingDirective *string `json:"taggingDirective,omitempty" tf:"tagging_directive,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

// ObjectCopySpec defines the desired state of ObjectCopy
type ObjectCopySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectCopyParameters `json:"forProvider"`
}

// ObjectCopyStatus defines the observed state of ObjectCopy.
type ObjectCopyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectCopyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectCopy is the Schema for the ObjectCopys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type ObjectCopy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ObjectCopySpec   `json:"spec"`
	Status            ObjectCopyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectCopyList contains a list of ObjectCopys
type ObjectCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectCopy `json:"items"`
}

// Repository type metadata.
var (
	ObjectCopy_Kind             = "ObjectCopy"
	ObjectCopy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ObjectCopy_Kind}.String()
	ObjectCopy_KindAPIVersion   = ObjectCopy_Kind + "." + CRDGroupVersion.String()
	ObjectCopy_GroupVersionKind = CRDGroupVersion.WithKind(ObjectCopy_Kind)
)

func init() {
	SchemeBuilder.Register(&ObjectCopy{}, &ObjectCopyList{})
}
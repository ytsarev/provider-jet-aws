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

type AutoscalingGroupsObservation struct {
	Name *string `json:"name,omitempty" tf:"name"`
}

type AutoscalingGroupsParameters struct {
}

type LaunchTemplateObservation struct {
}

type LaunchTemplateParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version"`
}

type NodeGroupObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn"`

	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources"`

	Status *string `json:"status,omitempty" tf:"status"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type NodeGroupParameters struct {

	// +kubebuilder:validation:Optional
	AmiType *string `json:"amiType,omitempty" tf:"ami_type"`

	// +kubebuilder:validation:Optional
	CapacityType *string `json:"capacityType,omitempty" tf:"capacity_type"`

	// +crossplane:generate:reference:type=Cluster
	// +kubebuilder:validation:Optional
	ClusterName *string `json:"clusterName,omitempty" tf:"cluster_name"`

	// +kubebuilder:validation:Optional
	ClusterNameRef *v1.Reference `json:"clusterNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ClusterNameSelector *v1.Selector `json:"clusterNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DiskSize *int64 `json:"diskSize,omitempty" tf:"disk_size"`

	// +kubebuilder:validation:Optional
	ForceUpdateVersion *bool `json:"forceUpdateVersion,omitempty" tf:"force_update_version"`

	// +kubebuilder:validation:Optional
	InstanceTypes []*string `json:"instanceTypes,omitempty" tf:"instance_types"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels"`

	// +kubebuilder:validation:Optional
	LaunchTemplate []LaunchTemplateParameters `json:"launchTemplate,omitempty" tf:"launch_template"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1.Role
	// +kubebuilder:validation:Optional
	NodeRoleArn *string `json:"nodeRoleArn,omitempty" tf:"node_role_arn"`

	// +kubebuilder:validation:Optional
	NodeRoleArnRef *v1.Reference `json:"nodeRoleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NodeRoleArnSelector *v1.Selector `json:"nodeRoleArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReleaseVersion *string `json:"releaseVersion,omitempty" tf:"release_version"`

	// +kubebuilder:validation:Optional
	RemoteAccess []RemoteAccessParameters `json:"remoteAccess,omitempty" tf:"remote_access"`

	// +kubebuilder:validation:Required
	ScalingConfig []ScalingConfigParameters `json:"scalingConfig" tf:"scaling_config"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids"`

	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	Taint []TaintParameters `json:"taint,omitempty" tf:"taint"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type RemoteAccessObservation struct {
}

type RemoteAccessParameters struct {

	// +kubebuilder:validation:Optional
	Ec2SSHKey *string `json:"ec2SshKey,omitempty" tf:"ec2_ssh_key"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-tf-aws/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	SourceSecurityGroupIds []*string `json:"sourceSecurityGroupIds,omitempty" tf:"source_security_group_ids"`

	// +kubebuilder:validation:Optional
	SourceSecurityGroupIdsRefs []v1.Reference `json:"sourceSecurityGroupIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SourceSecurityGroupIdsSelector *v1.Selector `json:"sourceSecurityGroupIdsSelector,omitempty" tf:"-"`
}

type ResourcesObservation struct {
	AutoscalingGroups []AutoscalingGroupsObservation `json:"autoscalingGroups,omitempty" tf:"autoscaling_groups"`

	RemoteAccessSecurityGroupID *string `json:"remoteAccessSecurityGroupId,omitempty" tf:"remote_access_security_group_id"`
}

type ResourcesParameters struct {
}

type ScalingConfigObservation struct {
}

type ScalingConfigParameters struct {

	// +kubebuilder:validation:Required
	DesiredSize *int64 `json:"desiredSize" tf:"desired_size"`

	// +kubebuilder:validation:Required
	MaxSize *int64 `json:"maxSize" tf:"max_size"`

	// +kubebuilder:validation:Required
	MinSize *int64 `json:"minSize" tf:"min_size"`
}

type TaintObservation struct {
}

type TaintParameters struct {

	// +kubebuilder:validation:Required
	Effect *string `json:"effect" tf:"effect"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value"`
}

// NodeGroupSpec defines the desired state of NodeGroup
type NodeGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeGroupParameters `json:"forProvider"`
}

// NodeGroupStatus defines the observed state of NodeGroup.
type NodeGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodeGroup is the Schema for the NodeGroups API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type NodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeGroupSpec   `json:"spec"`
	Status            NodeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeGroupList contains a list of NodeGroups
type NodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeGroup `json:"items"`
}

// Repository type metadata.
var (
	NodeGroupKind             = "NodeGroup"
	NodeGroupGroupKind        = schema.GroupKind{Group: Group, Kind: NodeGroupKind}.String()
	NodeGroupKindAPIVersion   = NodeGroupKind + "." + GroupVersion.String()
	NodeGroupGroupVersionKind = GroupVersion.WithKind(NodeGroupKind)
)

func init() {
	SchemeBuilder.Register(&NodeGroup{}, &NodeGroupList{})
}

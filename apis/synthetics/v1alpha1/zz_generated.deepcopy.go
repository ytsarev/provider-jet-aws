// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Canary) DeepCopyInto(out *Canary) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Canary.
func (in *Canary) DeepCopy() *Canary {
	if in == nil {
		return nil
	}
	out := new(Canary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Canary) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryList) DeepCopyInto(out *CanaryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Canary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryList.
func (in *CanaryList) DeepCopy() *CanaryList {
	if in == nil {
		return nil
	}
	out := new(CanaryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CanaryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryObservation) DeepCopyInto(out *CanaryObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.EngineArn != nil {
		in, out := &in.EngineArn, &out.EngineArn
		*out = new(string)
		**out = **in
	}
	if in.SourceLocationArn != nil {
		in, out := &in.SourceLocationArn, &out.SourceLocationArn
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Timeline != nil {
		in, out := &in.Timeline, &out.Timeline
		*out = make([]TimelineObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryObservation.
func (in *CanaryObservation) DeepCopy() *CanaryObservation {
	if in == nil {
		return nil
	}
	out := new(CanaryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryParameters) DeepCopyInto(out *CanaryParameters) {
	*out = *in
	if in.ArtifactS3Location != nil {
		in, out := &in.ArtifactS3Location, &out.ArtifactS3Location
		*out = new(string)
		**out = **in
	}
	if in.ExecutionRoleArn != nil {
		in, out := &in.ExecutionRoleArn, &out.ExecutionRoleArn
		*out = new(string)
		**out = **in
	}
	if in.FailureRetentionPeriod != nil {
		in, out := &in.FailureRetentionPeriod, &out.FailureRetentionPeriod
		*out = new(int64)
		**out = **in
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.RunConfig != nil {
		in, out := &in.RunConfig, &out.RunConfig
		*out = make([]RunConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RuntimeVersion != nil {
		in, out := &in.RuntimeVersion, &out.RuntimeVersion
		*out = new(string)
		**out = **in
	}
	if in.S3Bucket != nil {
		in, out := &in.S3Bucket, &out.S3Bucket
		*out = new(string)
		**out = **in
	}
	if in.S3Key != nil {
		in, out := &in.S3Key, &out.S3Key
		*out = new(string)
		**out = **in
	}
	if in.S3Version != nil {
		in, out := &in.S3Version, &out.S3Version
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = make([]ScheduleParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartCanary != nil {
		in, out := &in.StartCanary, &out.StartCanary
		*out = new(bool)
		**out = **in
	}
	if in.SuccessRetentionPeriod != nil {
		in, out := &in.SuccessRetentionPeriod, &out.SuccessRetentionPeriod
		*out = new(int64)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.VpcConfig != nil {
		in, out := &in.VpcConfig, &out.VpcConfig
		*out = make([]VpcConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ZipFile != nil {
		in, out := &in.ZipFile, &out.ZipFile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryParameters.
func (in *CanaryParameters) DeepCopy() *CanaryParameters {
	if in == nil {
		return nil
	}
	out := new(CanaryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanarySpec) DeepCopyInto(out *CanarySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanarySpec.
func (in *CanarySpec) DeepCopy() *CanarySpec {
	if in == nil {
		return nil
	}
	out := new(CanarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStatus) DeepCopyInto(out *CanaryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStatus.
func (in *CanaryStatus) DeepCopy() *CanaryStatus {
	if in == nil {
		return nil
	}
	out := new(CanaryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunConfigObservation) DeepCopyInto(out *RunConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunConfigObservation.
func (in *RunConfigObservation) DeepCopy() *RunConfigObservation {
	if in == nil {
		return nil
	}
	out := new(RunConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunConfigParameters) DeepCopyInto(out *RunConfigParameters) {
	*out = *in
	if in.ActiveTracing != nil {
		in, out := &in.ActiveTracing, &out.ActiveTracing
		*out = new(bool)
		**out = **in
	}
	if in.MemoryInMb != nil {
		in, out := &in.MemoryInMb, &out.MemoryInMb
		*out = new(int64)
		**out = **in
	}
	if in.TimeoutInSeconds != nil {
		in, out := &in.TimeoutInSeconds, &out.TimeoutInSeconds
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunConfigParameters.
func (in *RunConfigParameters) DeepCopy() *RunConfigParameters {
	if in == nil {
		return nil
	}
	out := new(RunConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleObservation) DeepCopyInto(out *ScheduleObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleObservation.
func (in *ScheduleObservation) DeepCopy() *ScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(ScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleParameters) DeepCopyInto(out *ScheduleParameters) {
	*out = *in
	if in.DurationInSeconds != nil {
		in, out := &in.DurationInSeconds, &out.DurationInSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleParameters.
func (in *ScheduleParameters) DeepCopy() *ScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(ScheduleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimelineObservation) DeepCopyInto(out *TimelineObservation) {
	*out = *in
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.LastStarted != nil {
		in, out := &in.LastStarted, &out.LastStarted
		*out = new(string)
		**out = **in
	}
	if in.LastStopped != nil {
		in, out := &in.LastStopped, &out.LastStopped
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimelineObservation.
func (in *TimelineObservation) DeepCopy() *TimelineObservation {
	if in == nil {
		return nil
	}
	out := new(TimelineObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimelineParameters) DeepCopyInto(out *TimelineParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimelineParameters.
func (in *TimelineParameters) DeepCopy() *TimelineParameters {
	if in == nil {
		return nil
	}
	out := new(TimelineParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcConfigObservation) DeepCopyInto(out *VpcConfigObservation) {
	*out = *in
	if in.VpcID != nil {
		in, out := &in.VpcID, &out.VpcID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcConfigObservation.
func (in *VpcConfigObservation) DeepCopy() *VpcConfigObservation {
	if in == nil {
		return nil
	}
	out := new(VpcConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpcConfigParameters) DeepCopyInto(out *VpcConfigParameters) {
	*out = *in
	if in.SecurityGroupIds != nil {
		in, out := &in.SecurityGroupIds, &out.SecurityGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SubnetIds != nil {
		in, out := &in.SubnetIds, &out.SubnetIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpcConfigParameters.
func (in *VpcConfigParameters) DeepCopy() *VpcConfigParameters {
	if in == nil {
		return nil
	}
	out := new(VpcConfigParameters)
	in.DeepCopyInto(out)
	return out
}
//go:build !ignore_autogenerated

/*
Copyright 2023.

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
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Arg) DeepCopyInto(out *Arg) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Arg.
func (in *Arg) DeepCopy() *Arg {
	if in == nil {
		return nil
	}
	out := new(Arg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Card) DeepCopyInto(out *Card) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Card.
func (in *Card) DeepCopy() *Card {
	if in == nil {
		return nil
	}
	out := new(Card)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomArtifact) DeepCopyInto(out *CustomArtifact) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomArtifact.
func (in *CustomArtifact) DeepCopy() *CustomArtifact {
	if in == nil {
		return nil
	}
	out := new(CustomArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomArtifacts) DeepCopyInto(out *CustomArtifacts) {
	*out = *in
	if in.Templates != nil {
		in, out := &in.Templates, &out.Templates
		*out = make([]CustomArtifact, len(*in))
		copy(*out, *in)
	}
	if in.SystemPrompts != nil {
		in, out := &in.SystemPrompts, &out.SystemPrompts
		*out = make([]CustomArtifact, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomArtifacts.
func (in *CustomArtifacts) DeepCopy() *CustomArtifacts {
	if in == nil {
		return nil
	}
	out := new(CustomArtifacts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomTaskSource) DeepCopyInto(out *CustomTaskSource) {
	*out = *in
	in.GitSource.DeepCopyInto(&out.GitSource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomTaskSource.
func (in *CustomTaskSource) DeepCopy() *CustomTaskSource {
	if in == nil {
		return nil
	}
	out := new(CustomTaskSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomTasks) DeepCopyInto(out *CustomTasks) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomTasks.
func (in *CustomTasks) DeepCopy() *CustomTasks {
	if in == nil {
		return nil
	}
	out := new(CustomTasks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSource) DeepCopyInto(out *GitSource) {
	*out = *in
	if in.Branch != nil {
		in, out := &in.Branch, &out.Branch
		*out = new(string)
		**out = **in
	}
	if in.Commit != nil {
		in, out := &in.Commit, &out.Commit
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSource.
func (in *GitSource) DeepCopy() *GitSource {
	if in == nil {
		return nil
	}
	out := new(GitSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LMEvalContainer) DeepCopyInto(out *LMEvalContainer) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LMEvalContainer.
func (in *LMEvalContainer) DeepCopy() *LMEvalContainer {
	if in == nil {
		return nil
	}
	out := new(LMEvalContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LMEvalJob) DeepCopyInto(out *LMEvalJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LMEvalJob.
func (in *LMEvalJob) DeepCopy() *LMEvalJob {
	if in == nil {
		return nil
	}
	out := new(LMEvalJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LMEvalJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LMEvalJobList) DeepCopyInto(out *LMEvalJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LMEvalJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LMEvalJobList.
func (in *LMEvalJobList) DeepCopy() *LMEvalJobList {
	if in == nil {
		return nil
	}
	out := new(LMEvalJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LMEvalJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LMEvalJobSpec) DeepCopyInto(out *LMEvalJobSpec) {
	*out = *in
	if in.ModelArgs != nil {
		in, out := &in.ModelArgs, &out.ModelArgs
		*out = make([]Arg, len(*in))
		copy(*out, *in)
	}
	in.TaskList.DeepCopyInto(&out.TaskList)
	if in.NumFewShot != nil {
		in, out := &in.NumFewShot, &out.NumFewShot
		*out = new(int)
		**out = **in
	}
	if in.GenArgs != nil {
		in, out := &in.GenArgs, &out.GenArgs
		*out = make([]Arg, len(*in))
		copy(*out, *in)
	}
	if in.LogSamples != nil {
		in, out := &in.LogSamples, &out.LogSamples
		*out = new(bool)
		**out = **in
	}
	if in.BatchSize != nil {
		in, out := &in.BatchSize, &out.BatchSize
		*out = new(string)
		**out = **in
	}
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(LMEvalPodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = new(Outputs)
		(*in).DeepCopyInto(*out)
	}
	if in.Offline != nil {
		in, out := &in.Offline, &out.Offline
		*out = new(OfflineSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AllowOnline != nil {
		in, out := &in.AllowOnline, &out.AllowOnline
		*out = new(bool)
		**out = **in
	}
	if in.AllowCodeExecution != nil {
		in, out := &in.AllowCodeExecution, &out.AllowCodeExecution
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LMEvalJobSpec.
func (in *LMEvalJobSpec) DeepCopy() *LMEvalJobSpec {
	if in == nil {
		return nil
	}
	out := new(LMEvalJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LMEvalJobStatus) DeepCopyInto(out *LMEvalJobStatus) {
	*out = *in
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		*out = (*in).DeepCopy()
	}
	if in.CompleteTime != nil {
		in, out := &in.CompleteTime, &out.CompleteTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LMEvalJobStatus.
func (in *LMEvalJobStatus) DeepCopy() *LMEvalJobStatus {
	if in == nil {
		return nil
	}
	out := new(LMEvalJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LMEvalPodSpec) DeepCopyInto(out *LMEvalPodSpec) {
	*out = *in
	if in.Container != nil {
		in, out := &in.Container, &out.Container
		*out = new(LMEvalContainer)
		(*in).DeepCopyInto(*out)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SideCars != nil {
		in, out := &in.SideCars, &out.SideCars
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LMEvalPodSpec.
func (in *LMEvalPodSpec) DeepCopy() *LMEvalPodSpec {
	if in == nil {
		return nil
	}
	out := new(LMEvalPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OfflineS3Spec) DeepCopyInto(out *OfflineS3Spec) {
	*out = *in
	in.AccessKeyIdRef.DeepCopyInto(&out.AccessKeyIdRef)
	in.SecretAccessKeyRef.DeepCopyInto(&out.SecretAccessKeyRef)
	in.Bucket.DeepCopyInto(&out.Bucket)
	in.Region.DeepCopyInto(&out.Region)
	in.Endpoint.DeepCopyInto(&out.Endpoint)
	if in.VerifySSL != nil {
		in, out := &in.VerifySSL, &out.VerifySSL
		*out = new(bool)
		**out = **in
	}
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OfflineS3Spec.
func (in *OfflineS3Spec) DeepCopy() *OfflineS3Spec {
	if in == nil {
		return nil
	}
	out := new(OfflineS3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OfflineSpec) DeepCopyInto(out *OfflineSpec) {
	*out = *in
	in.StorageSpec.DeepCopyInto(&out.StorageSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OfflineSpec.
func (in *OfflineSpec) DeepCopy() *OfflineSpec {
	if in == nil {
		return nil
	}
	out := new(OfflineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OfflineStorageSpec) DeepCopyInto(out *OfflineStorageSpec) {
	*out = *in
	if in.PersistentVolumeClaimName != nil {
		in, out := &in.PersistentVolumeClaimName, &out.PersistentVolumeClaimName
		*out = new(string)
		**out = **in
	}
	if in.S3Spec != nil {
		in, out := &in.S3Spec, &out.S3Spec
		*out = new(OfflineS3Spec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OfflineStorageSpec.
func (in *OfflineStorageSpec) DeepCopy() *OfflineStorageSpec {
	if in == nil {
		return nil
	}
	out := new(OfflineStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Outputs) DeepCopyInto(out *Outputs) {
	*out = *in
	if in.PersistentVolumeClaimName != nil {
		in, out := &in.PersistentVolumeClaimName, &out.PersistentVolumeClaimName
		*out = new(string)
		**out = **in
	}
	if in.PersistentVolumeClaimManaged != nil {
		in, out := &in.PersistentVolumeClaimManaged, &out.PersistentVolumeClaimManaged
		*out = new(PersistentVolumeClaimManaged)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Outputs.
func (in *Outputs) DeepCopy() *Outputs {
	if in == nil {
		return nil
	}
	out := new(Outputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaimManaged) DeepCopyInto(out *PersistentVolumeClaimManaged) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaimManaged.
func (in *PersistentVolumeClaimManaged) DeepCopy() *PersistentVolumeClaimManaged {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaimManaged)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemPrompt) DeepCopyInto(out *SystemPrompt) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemPrompt.
func (in *SystemPrompt) DeepCopy() *SystemPrompt {
	if in == nil {
		return nil
	}
	out := new(SystemPrompt)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskList) DeepCopyInto(out *TaskList) {
	*out = *in
	if in.TaskNames != nil {
		in, out := &in.TaskNames, &out.TaskNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TaskRecipes != nil {
		in, out := &in.TaskRecipes, &out.TaskRecipes
		*out = make([]TaskRecipe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomArtifacts != nil {
		in, out := &in.CustomArtifacts, &out.CustomArtifacts
		*out = new(CustomArtifacts)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomTasks != nil {
		in, out := &in.CustomTasks, &out.CustomTasks
		*out = new(CustomTasks)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskList.
func (in *TaskList) DeepCopy() *TaskList {
	if in == nil {
		return nil
	}
	out := new(TaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskRecipe) DeepCopyInto(out *TaskRecipe) {
	*out = *in
	out.Card = in.Card
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(Template)
		**out = **in
	}
	if in.SystemPrompt != nil {
		in, out := &in.SystemPrompt, &out.SystemPrompt
		*out = new(SystemPrompt)
		**out = **in
	}
	if in.Task != nil {
		in, out := &in.Task, &out.Task
		*out = new(string)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.LoaderLimit != nil {
		in, out := &in.LoaderLimit, &out.LoaderLimit
		*out = new(int)
		**out = **in
	}
	if in.NumDemos != nil {
		in, out := &in.NumDemos, &out.NumDemos
		*out = new(int)
		**out = **in
	}
	if in.DemosPoolSize != nil {
		in, out := &in.DemosPoolSize, &out.DemosPoolSize
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskRecipe.
func (in *TaskRecipe) DeepCopy() *TaskRecipe {
	if in == nil {
		return nil
	}
	out := new(TaskRecipe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Template) DeepCopyInto(out *Template) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Template.
func (in *Template) DeepCopy() *Template {
	if in == nil {
		return nil
	}
	out := new(Template)
	in.DeepCopyInto(out)
	return out
}

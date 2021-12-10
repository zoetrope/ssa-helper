/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	corev1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/core/v1"
	v1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/meta/v1"
)

// DeploymentSpecApplyConfiguration represents an declarative configuration of the DeploymentSpec type for use
// with apply.
type DeploymentSpecApplyConfiguration struct {
	Replicas                *int32                                    `json:"replicas,omitempty"`
	Selector                *v1.LabelSelectorApplyConfiguration       `json:"selector,omitempty"`
	Template                *corev1.PodTemplateSpecApplyConfiguration `json:"template,omitempty"`
	Strategy                *DeploymentStrategyApplyConfiguration     `json:"strategy,omitempty"`
	MinReadySeconds         *int32                                    `json:"minReadySeconds,omitempty"`
	RevisionHistoryLimit    *int32                                    `json:"revisionHistoryLimit,omitempty"`
	Paused                  *bool                                     `json:"paused,omitempty"`
	RollbackTo              *RollbackConfigApplyConfiguration         `json:"rollbackTo,omitempty"`
	ProgressDeadlineSeconds *int32                                    `json:"progressDeadlineSeconds,omitempty"`
}

// DeploymentSpecApplyConfiguration constructs an declarative configuration of the DeploymentSpec type for use with
// apply.
func DeploymentSpec() *DeploymentSpecApplyConfiguration {
	return &DeploymentSpecApplyConfiguration{}
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *DeploymentSpecApplyConfiguration) WithReplicas(value int32) *DeploymentSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *DeploymentSpecApplyConfiguration) WithSelector(value *v1.LabelSelectorApplyConfiguration) *DeploymentSpecApplyConfiguration {
	b.Selector = value
	return b
}

// WithTemplate sets the Template field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Template field is set to the value of the last call.
func (b *DeploymentSpecApplyConfiguration) WithTemplate(value *corev1.PodTemplateSpecApplyConfiguration) *DeploymentSpecApplyConfiguration {
	b.Template = value
	return b
}

// WithStrategy sets the Strategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Strategy field is set to the value of the last call.
func (b *DeploymentSpecApplyConfiguration) WithStrategy(value *DeploymentStrategyApplyConfiguration) *DeploymentSpecApplyConfiguration {
	b.Strategy = value
	return b
}

// WithMinReadySeconds sets the MinReadySeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinReadySeconds field is set to the value of the last call.
func (b *DeploymentSpecApplyConfiguration) WithMinReadySeconds(value int32) *DeploymentSpecApplyConfiguration {
	b.MinReadySeconds = &value
	return b
}

// WithRevisionHistoryLimit sets the RevisionHistoryLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RevisionHistoryLimit field is set to the value of the last call.
func (b *DeploymentSpecApplyConfiguration) WithRevisionHistoryLimit(value int32) *DeploymentSpecApplyConfiguration {
	b.RevisionHistoryLimit = &value
	return b
}

// WithPaused sets the Paused field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Paused field is set to the value of the last call.
func (b *DeploymentSpecApplyConfiguration) WithPaused(value bool) *DeploymentSpecApplyConfiguration {
	b.Paused = &value
	return b
}

// WithRollbackTo sets the RollbackTo field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RollbackTo field is set to the value of the last call.
func (b *DeploymentSpecApplyConfiguration) WithRollbackTo(value *RollbackConfigApplyConfiguration) *DeploymentSpecApplyConfiguration {
	b.RollbackTo = value
	return b
}

// WithProgressDeadlineSeconds sets the ProgressDeadlineSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProgressDeadlineSeconds field is set to the value of the last call.
func (b *DeploymentSpecApplyConfiguration) WithProgressDeadlineSeconds(value int32) *DeploymentSpecApplyConfiguration {
	b.ProgressDeadlineSeconds = &value
	return b
}

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

package v2beta2

import (
	v1 "github.com/zoetrope/ssa-helper/applyconfigurations/meta/v1"
)

// MetricIdentifierApplyConfiguration represents an declarative configuration of the MetricIdentifier type for use
// with apply.
type MetricIdentifierApplyConfiguration struct {
	Name     *string                             `json:"name,omitempty"`
	Selector *v1.LabelSelectorApplyConfiguration `json:"selector,omitempty"`
}

// MetricIdentifierApplyConfiguration constructs an declarative configuration of the MetricIdentifier type for use with
// apply.
func MetricIdentifier() *MetricIdentifierApplyConfiguration {
	return &MetricIdentifierApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *MetricIdentifierApplyConfiguration) WithName(value string) *MetricIdentifierApplyConfiguration {
	b.Name = &value
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *MetricIdentifierApplyConfiguration) WithSelector(value *v1.LabelSelectorApplyConfiguration) *MetricIdentifierApplyConfiguration {
	b.Selector = value
	return b
}

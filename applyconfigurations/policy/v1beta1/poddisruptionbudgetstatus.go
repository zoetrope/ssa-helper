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
	metav1 "github.com/zoetrope/ac-deepcopy/applyconfigurations/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodDisruptionBudgetStatusApplyConfiguration represents an declarative configuration of the PodDisruptionBudgetStatus type for use
// with apply.
type PodDisruptionBudgetStatusApplyConfiguration struct {
	ObservedGeneration *int64                               `json:"observedGeneration,omitempty"`
	DisruptedPods      map[string]v1.Time                   `json:"disruptedPods,omitempty"`
	DisruptionsAllowed *int32                               `json:"disruptionsAllowed,omitempty"`
	CurrentHealthy     *int32                               `json:"currentHealthy,omitempty"`
	DesiredHealthy     *int32                               `json:"desiredHealthy,omitempty"`
	ExpectedPods       *int32                               `json:"expectedPods,omitempty"`
	Conditions         []metav1.ConditionApplyConfiguration `json:"conditions,omitempty"`
}

// PodDisruptionBudgetStatusApplyConfiguration constructs an declarative configuration of the PodDisruptionBudgetStatus type for use with
// apply.
func PodDisruptionBudgetStatus() *PodDisruptionBudgetStatusApplyConfiguration {
	return &PodDisruptionBudgetStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *PodDisruptionBudgetStatusApplyConfiguration) WithObservedGeneration(value int64) *PodDisruptionBudgetStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithDisruptedPods puts the entries into the DisruptedPods field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the DisruptedPods field,
// overwriting an existing map entries in DisruptedPods field with the same key.
func (b *PodDisruptionBudgetStatusApplyConfiguration) WithDisruptedPods(entries map[string]v1.Time) *PodDisruptionBudgetStatusApplyConfiguration {
	if b.DisruptedPods == nil && len(entries) > 0 {
		b.DisruptedPods = make(map[string]v1.Time, len(entries))
	}
	for k, v := range entries {
		b.DisruptedPods[k] = v
	}
	return b
}

// WithDisruptionsAllowed sets the DisruptionsAllowed field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DisruptionsAllowed field is set to the value of the last call.
func (b *PodDisruptionBudgetStatusApplyConfiguration) WithDisruptionsAllowed(value int32) *PodDisruptionBudgetStatusApplyConfiguration {
	b.DisruptionsAllowed = &value
	return b
}

// WithCurrentHealthy sets the CurrentHealthy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CurrentHealthy field is set to the value of the last call.
func (b *PodDisruptionBudgetStatusApplyConfiguration) WithCurrentHealthy(value int32) *PodDisruptionBudgetStatusApplyConfiguration {
	b.CurrentHealthy = &value
	return b
}

// WithDesiredHealthy sets the DesiredHealthy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DesiredHealthy field is set to the value of the last call.
func (b *PodDisruptionBudgetStatusApplyConfiguration) WithDesiredHealthy(value int32) *PodDisruptionBudgetStatusApplyConfiguration {
	b.DesiredHealthy = &value
	return b
}

// WithExpectedPods sets the ExpectedPods field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExpectedPods field is set to the value of the last call.
func (b *PodDisruptionBudgetStatusApplyConfiguration) WithExpectedPods(value int32) *PodDisruptionBudgetStatusApplyConfiguration {
	b.ExpectedPods = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *PodDisruptionBudgetStatusApplyConfiguration) WithConditions(values ...*metav1.ConditionApplyConfiguration) *PodDisruptionBudgetStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

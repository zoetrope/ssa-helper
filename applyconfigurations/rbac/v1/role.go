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

package v1

import (
	"errors"

	internal "github.com/zoetrope/ssa-helper/applyconfigurations/internal"
	v1 "github.com/zoetrope/ssa-helper/applyconfigurations/meta/v1"
	apirbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// RoleApplyConfiguration represents an declarative configuration of the Role type for use
// with apply.
type RoleApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Rules                            []PolicyRuleApplyConfiguration `json:"rules,omitempty"`
}

// Role constructs an declarative configuration of the Role type for use with
// apply.
func Role(name, namespace string) *RoleApplyConfiguration {
	b := &RoleApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("Role")
	b.WithAPIVersion("rbac.authorization.k8s.io/v1")
	return b
}

// ExtractRole extracts the applied configuration owned by fieldManager from
// role. If no managedFields are found in role for fieldManager, a
// RoleApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// role must be a unmodified Role API object that was retrieved from the Kubernetes API.
// ExtractRole provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractRole(role *apirbacv1.Role, fieldManager string) (*RoleApplyConfiguration, error) {
	return extractRole(role, fieldManager, "")
}

// ExtractRoleStatus is the same as ExtractRole except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractRoleStatus(role *apirbacv1.Role, fieldManager string) (*RoleApplyConfiguration, error) {
	return extractRole(role, fieldManager, "status")
}

func extractRole(role *apirbacv1.Role, fieldManager string, subresource string) (*RoleApplyConfiguration, error) {
	b := &RoleApplyConfiguration{}
	err := managedfields.ExtractInto(role, internal.Parser().Type("io.k8s.api.rbac.v1.Role"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(role.Name)
	b.WithNamespace(role.Namespace)

	b.WithKind("Role")
	b.WithAPIVersion("rbac.authorization.k8s.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithKind(value string) *RoleApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithAPIVersion(value string) *RoleApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithName(value string) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithGenerateName(value string) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithNamespace(value string) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithSelfLink sets the SelfLink field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SelfLink field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithSelfLink(value string) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.SelfLink = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithUID(value types.UID) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithResourceVersion(value string) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithGeneration(value int64) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithCreationTimestamp(value metav1.Time) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *RoleApplyConfiguration) WithLabels(entries map[string]string) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *RoleApplyConfiguration) WithAnnotations(entries map[string]string) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *RoleApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *RoleApplyConfiguration) WithFinalizers(values ...string) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

// WithClusterName sets the ClusterName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClusterName field is set to the value of the last call.
func (b *RoleApplyConfiguration) WithClusterName(value string) *RoleApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ClusterName = &value
	return b
}

func (b *RoleApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithRules adds the given value to the Rules field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Rules field.
func (b *RoleApplyConfiguration) WithRules(values ...*PolicyRuleApplyConfiguration) *RoleApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRules")
		}
		b.Rules = append(b.Rules, *values[i])
	}
	return b
}
func (b *RoleApplyConfiguration) Original() client.Object {
	return &apirbacv1.Role{}
}

func (b *RoleApplyConfiguration) Extract(obj client.Object, fieldManager string, subresource string) (*RoleApplyConfiguration, error) {
	return extractRole(obj.(*apirbacv1.Role), fieldManager, subresource)
}
func (b *RoleApplyConfiguration) ObjectKey() (client.ObjectKey, error) {
	if b.Namespace == nil {
		return client.ObjectKey{}, errors.New("The RoleApplyConfiguration namespace should not be empty.")
	}
	if b.Name == nil {
		return client.ObjectKey{}, errors.New("The RoleApplyConfiguration name should not be empty.")
	}
	return client.ObjectKey{
		Name:      *b.Name,
		Namespace: *b.Namespace,
	}, nil
}

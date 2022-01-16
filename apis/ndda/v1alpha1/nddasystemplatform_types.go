/*
Copyright 2021 NDD.

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

package v1alpha1

import (
	"reflect"

	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SystemPlatform struct
type SystemPlatform struct {
	Config *SystemPlatformConfig `json:"config,omitempty"`
}

// SystemPlatformConfig struct
type SystemPlatformConfig struct {
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4294967295
	Index *uint32 `json:"index,omitempty"`
	// +kubebuilder:validation:Enum=`SRL`;`SROS`
	Kind E_SystemPlatformKind `json:"kind,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// A SystemPlatformSpec defines the desired state of a SystemPlatform.
type SystemPlatformSpec struct {
	SystemPlatform *SystemPlatform `json:"platform,omitempty"`
}

// A SystemPlatformStatus represents the observed state of a SystemPlatform.
type SystemPlatformStatus struct {
	nddv1.ConditionedStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// NddaSystemPlatform is the Schema for the SystemPlatform API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
type NddaSystemPlatform struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SystemPlatformSpec   `json:"spec,omitempty"`
	Status SystemPlatformStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NddaSystemPlatformList contains a list of SystemPlatforms
type NddaSystemPlatformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NddaSystemPlatform `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NddaSystemPlatform{}, &NddaSystemPlatformList{})
}

// SystemPlatform type metadata.
var (
	SystemPlatformKindKind         = reflect.TypeOf(NddaSystemPlatform{}).Name()
	SystemPlatformGroupKind        = schema.GroupKind{Group: Group, Kind: SystemPlatformKindKind}.String()
	SystemPlatformKindAPIVersion   = SystemPlatformKindKind + "." + GroupVersion.String()
	SystemPlatformGroupVersionKind = GroupVersion.WithKind(SystemPlatformKindKind)
)

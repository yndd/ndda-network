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

// Interface struct
type Interface struct {
	Config *InterfaceConfig `json:"config,omitempty"`
	Name   *string          `json:"name"`
}

// InterfaceConfig struct
type InterfaceConfig struct {
	// +kubebuilder:validation:Enum=`INTERFACE`;`IRB`;`LOOPBACK`;`MPLS`;`VXLAN`
	Kind         E_InterfaceKind `json:"kind,omitempty"`
	Lacp         *bool           `json:"lacp,omitempty"`
	LacpFallback *bool           `json:"lacp-fallback,omitempty"`
	Lag          *bool           `json:"lag,omitempty"`
	LagMember    *bool           `json:"lag-member,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	LagName *string `json:"lag-name,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name,omitempty"`
}

// A InterfaceSpec defines the desired state of a Interface.
type InterfaceSpec struct {
	Interface *Interface `json:"interface,omitempty"`
}

// A InterfaceStatus represents the observed state of a Interface.
type InterfaceStatus struct {
	nddv1.ConditionedStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// NddaInterface is the Schema for the Interface API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
type NddaInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InterfaceSpec   `json:"spec,omitempty"`
	Status InterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NddaInterfaceList contains a list of Interfaces
type NddaInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NddaInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NddaInterface{}, &NddaInterfaceList{})
}

// Interface type metadata.
var (
	InterfaceKindKind         = reflect.TypeOf(NddaInterface{}).Name()
	InterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: InterfaceKindKind}.String()
	InterfaceKindAPIVersion   = InterfaceKindKind + "." + GroupVersion.String()
	InterfaceGroupVersionKind = GroupVersion.WithKind(InterfaceKindKind)
)

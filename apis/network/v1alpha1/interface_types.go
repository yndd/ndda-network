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

const (
	// InterfaceFinalizer is the name of the finalizer added to
	// Interface structure to block delete operations until the physical node can be
	// deprovisioned.
	InterfaceFinalizer string = "interface.network.ndda.yndd.io"

	LabelInterfaceKindKey string = "ndda-interface-kind"
)

// Interface struct
type NetworkInterface struct {
	Name         *string `json:"name,omitempty"`
	Kind         *string `json:"kind,omitempty"`
	Lag          *bool   `json:"lag,omitempty"`
	LagMember    *bool   `json:"lag-member,omitempty"`
	LagName      *string `json:"lag-name,omitempty"`
	Lacp         *bool   `json:"lacp,omitempty"`
	LacpFallback *bool   `json:"lacp-fallback,omitempty"`
}

// A InterfaceSpec defines the desired state of a Interface.
type InterfaceSpec struct {
	//nddv1.ResourceSpec `json:",inline"`
	TopologyName  *string           `json:"topology-name,omitempty"`
	NodeName      *string           `json:"node-name,omitempty"`
	EndpointGroup *string           `json:"endpoint-group,omitempty"`
	Interface     *NetworkInterface `json:"interface,omitempty"`
}

// A InterfaceStatus represents the observed state of a InterfaceSpec.
type InterfaceStatus struct {
	nddv1.ConditionedStatus `json:",inline"`
	ControllerRef           nddv1.Reference `json:"controllerRef,omitempty"`
	//Interface                     *NddoInterfaceInterface     `json:"Interface,omitempty"`
}

// +kubebuilder:object:root=true

// Interface is the Schema for the Interface API
// +kubebuilder:subresource:status
type Interface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InterfaceSpec   `json:"spec,omitempty"`
	Status InterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InterfaceList contains a list of Interfaces
type InterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Interface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Interface{}, &InterfaceList{})
}

// Interface type metadata.
var (
	InterfaceKindKind         = reflect.TypeOf(Interface{}).Name()
	InterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: InterfaceKindKind}.String()
	InterfaceKindAPIVersion   = InterfaceKindKind + "." + GroupVersion.String()
	InterfaceGroupVersionKind = GroupVersion.WithKind(InterfaceKindKind)
)

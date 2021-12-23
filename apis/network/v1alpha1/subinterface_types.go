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
	// SubInterfaceFinalizer is the name of the finalizer added to
	// SubInterface structure to block delete operations until the physical node can be
	// deprovisioned.
	SubInterfaceFinalizer string = "subinterface.network.ndda.yndd.io"

	LabelSubInterfaceKindKey = "ndda-sub-interface-kind"
)

// NetworkSubInterface struct
type NetworkSubInterface struct {
	Index    *string   `json:"index,omitempty"`
	Kind     *string   `json:"kind,omitempty"`
	Tagging  *string   `json:"tagging,omitempty"`
	OuterTag *uint32   `json:"outer-tag,omitempty"`
	InnerTag *uint32   `json:"inner-tag,omitempty"`
	Ipv4     []*string `json:"ipv4,omitempty"`
	Ipv6     []*string `json:"ipv6,omitempty"`
}

// A SubInterfaceSpec defines the desired state of a SubInterface.
type SubInterfaceSpec struct {
	//nddv1.ResourceSpec `json:",inline"`
	TopologyName  *string `json:"topology-name,omitempty"`
	NodeName      *string `json:"node-name,omitempty"`
	InterfaceName *string `json:"interface-name,omitempty"`
	//EndpointGroup *string              `json:"endpoint-group,omitempty"`
	SubInterface *NetworkSubInterface `json:"sub-interface,omitempty"`
}

// A SubInterfaceStatus represents the observed state of a SubInterfaceSpec.
type SubInterfaceStatus struct {
	nddv1.ConditionedStatus `json:",inline"`
	ControllerRef           nddv1.Reference `json:"controllerRef,omitempty"`
	//Interface                     *NddoInterfaceInterface     `json:"Interface,omitempty"`
}

// +kubebuilder:object:root=true

// SubInterface is the Schema for the SubInterface API
// +kubebuilder:subresource:status
type SubInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SubInterfaceSpec   `json:"spec,omitempty"`
	Status SubInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubInterfaceList contains a list of SubInterfaceList
type SubInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SubInterface{}, &SubInterfaceList{})
}

// SubInterface type metadata.
var (
	SubInterfaceKindKind         = reflect.TypeOf(SubInterface{}).Name()
	SubInterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: SubInterfaceKindKind}.String()
	SubInterfaceKindAPIVersion   = SubInterfaceKindKind + "." + GroupVersion.String()
	SubInterfaceGroupVersionKind = GroupVersion.WithKind(SubInterfaceKindKind)
)

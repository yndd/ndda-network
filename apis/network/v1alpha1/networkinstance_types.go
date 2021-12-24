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
	// NetworkInstanceFinalizer is the name of the finalizer added to
	// NetworkInstance structure to block delete operations until the physical node can be
	// deprovisioned.
	NetworkInstanceFinalizer string = "networkinstance.network.ndda.yndd.io"

	LabelNetworkInstanceKindKey = "ndda-network-instance-kind"
)

// NetworkNetworkInstance struct
type NetworkNetworkInstance struct {
	Name      *string                            `json:"name,omitempty"`
	Kind      *string                            `json:"kind,omitempty"`
	Interface []*NetworkNetworkInstanceInterface `json:"interface,omitempty"`
}

type NetworkNetworkInstanceInterface struct {
	Name *string `json:"name,omitempty"`
	Kind *string `json:"kind,omitempty"`
}

// A NetworkInstanceSpec defines the desired state of a NetworkInstance.
type NetworkInstanceSpec struct {
	//nddv1.ResourceSpec `json:",inline"`
	TopologyName    *string                 `json:"topology-name,omitempty"`
	NodeName        *string                 `json:"node-name,omitempty"`
	NetworkInstance *NetworkNetworkInstance `json:"network-instance,omitempty"`
}

// A NetworkInstanceStatus represents the observed state of a NetworkInstanceSpec.
type NetworkInstanceStatus struct {
	nddv1.ConditionedStatus `json:",inline"`
	ControllerRef           nddv1.Reference `json:"controllerRef,omitempty"`
	//Interface                     *NddoInterfaceInterface     `json:"Interface,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInstance is the Schema for the NetworkInstance API
// +kubebuilder:subresource:status
type NetworkInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkInstanceSpec   `json:"spec,omitempty"`
	Status NetworkInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInstanceList contains a list of NetworkInstanceList
type NetworkInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkInstance{}, &NetworkInstanceList{})
}

// NetworkInstance type metadata.
var (
	NetworkInstanceKindKind         = reflect.TypeOf(NetworkInstance{}).Name()
	NetworkInstanceGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkInstanceKindKind}.String()
	NetworkInstanceKindAPIVersion   = NetworkInstanceKindKind + "." + GroupVersion.String()
	NetworkInstanceGroupVersionKind = GroupVersion.WithKind(NetworkInstanceKindKind)
)

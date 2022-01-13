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

// NetworkInstance struct
type NetworkInstance struct {
	Config *NetworkInstanceConfig `json:"config,omitempty"`
	Name   *string                `json:"name"`
}

// NetworkInstanceConfig struct
type NetworkInstanceConfig struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	Index     *uint32                           `json:"index,omitempty"`
	Interface []*NetworkInstanceConfigInterface `json:"interface,omitempty"`
	// +kubebuilder:validation:Enum=`BRIDGED`;`ROUTED`
	Kind E_NetworkInstanceKind `json:"kind,omitempty"`
	// kubebuilder:validation:MinLength=1
	// kubebuilder:validation:MaxLength=255
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="[A-Za-z0-9 !@#$^&()|+=`~.,'/_:;?-]*"
	Name *string `json:"name,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	RouterId *string `json:"router-id,omitempty"`
}

// NetworkInstanceConfigInterface struct
type NetworkInstanceConfigInterface struct {
	Name *string `json:"name"`
}

// A NetworkInstanceSpec defines the desired state of a NetworkInstance.
type NetworkInstanceSpec struct {
	NetworkInstance *NetworkInstance `json:"network-instance,omitempty"`
}

// A NetworkInstanceStatus represents the observed state of a NetworkInstance.
type NetworkInstanceStatus struct {
	nddv1.ConditionedStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// NetworkNetworkInstance is the Schema for the NetworkInstance API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
type NetworkNetworkInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkInstanceSpec   `json:"spec,omitempty"`
	Status NetworkInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkNetworkInstanceList contains a list of NetworkInstances
type NetworkNetworkInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkNetworkInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkNetworkInstance{}, &NetworkNetworkInstanceList{})
}

// NetworkInstance type metadata.
var (
	NetworkInstanceKindKind         = reflect.TypeOf(NetworkNetworkInstance{}).Name()
	NetworkInstanceGroupKind        = schema.GroupKind{Group: Group, Kind: NetworkInstanceKindKind}.String()
	NetworkInstanceKindAPIVersion   = NetworkInstanceKindKind + "." + GroupVersion.String()
	NetworkInstanceGroupVersionKind = GroupVersion.WithKind(NetworkInstanceKindKind)
)

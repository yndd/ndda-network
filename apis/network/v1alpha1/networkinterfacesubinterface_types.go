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

// InterfaceSubinterface struct
type InterfaceSubinterface struct {
	Config *InterfaceSubinterfaceConfig `json:"config,omitempty"`
	Index  *string                      `json:"index"`
	Ipv4   []*InterfaceSubinterfaceIpv4 `json:"ipv4,omitempty"`
	Ipv6   []*InterfaceSubinterfaceIpv6 `json:"ipv6,omitempty"`
}

// InterfaceSubinterfaceConfig struct
type InterfaceSubinterfaceConfig struct {
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=9999
	Index *uint32 `json:"index,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4094
	InnerVlanId *uint16 `json:"inner-vlan-id,omitempty"`
	// +kubebuilder:validation:Enum=`BRIDGED`;`ROUTED`
	Kind E_InterfaceSubinterfaceKind `json:"kind,omitempty"`
	// kubebuilder:validation:Minimum=1
	// kubebuilder:validation:Maximum=4094
	OuterVlanId *uint16 `json:"outer-vlan-id,omitempty"`
}

// InterfaceSubinterfaceIpv4 struct
type InterfaceSubinterfaceIpv4 struct {
	Config   *InterfaceSubinterfaceIpv4Config `json:"config,omitempty"`
	IpPrefix *string                          `json:"ip-prefix"`
}

// InterfaceSubinterfaceIpv4Config struct
type InterfaceSubinterfaceIpv4Config struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])`
	IpAddress *string `json:"ip-address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	IpCidr *string `json:"ip-cidr,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))`
	IpPrefix *string `json:"ip-prefix,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	PrefixLength *uint32 `json:"prefix-length,omitempty"`
}

// InterfaceSubinterfaceIpv6 struct
type InterfaceSubinterfaceIpv6 struct {
	Config   *InterfaceSubinterfaceIpv6Config `json:"config,omitempty"`
	IpPrefix *string                          `json:"ip-prefix"`
}

// InterfaceSubinterfaceIpv6Config struct
type InterfaceSubinterfaceIpv6Config struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))`
	IpAddress *string `json:"ip-address,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpCidr *string `json:"ip-cidr,omitempty"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=`((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	IpPrefix *string `json:"ip-prefix,omitempty"`
	// kubebuilder:validation:Minimum=0
	// kubebuilder:validation:Maximum=4294967295
	PrefixLength *uint32 `json:"prefix-length,omitempty"`
}

// A InterfaceSubinterfaceSpec defines the desired state of a InterfaceSubinterface.
type InterfaceSubinterfaceSpec struct {
	InterfaceName         *string                `json:"interface-name"`
	InterfaceSubinterface *InterfaceSubinterface `json:"subinterface,omitempty"`
}

// A InterfaceSubinterfaceStatus represents the observed state of a InterfaceSubinterface.
type InterfaceSubinterfaceStatus struct {
	nddv1.ConditionedStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceSubinterface is the Schema for the InterfaceSubinterface API
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.conditions[?(@.kind=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNC",type="string",JSONPath=".status.conditions[?(@.kind=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
type NetworkInterfaceSubinterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InterfaceSubinterfaceSpec   `json:"spec,omitempty"`
	Status InterfaceSubinterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceSubinterfaceList contains a list of InterfaceSubinterfaces
type NetworkInterfaceSubinterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterfaceSubinterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkInterfaceSubinterface{}, &NetworkInterfaceSubinterfaceList{})
}

// InterfaceSubinterface type metadata.
var (
	InterfaceSubinterfaceKindKind         = reflect.TypeOf(NetworkInterfaceSubinterface{}).Name()
	InterfaceSubinterfaceGroupKind        = schema.GroupKind{Group: Group, Kind: InterfaceSubinterfaceKindKind}.String()
	InterfaceSubinterfaceKindAPIVersion   = InterfaceSubinterfaceKindKind + "." + GroupVersion.String()
	InterfaceSubinterfaceGroupVersionKind = GroupVersion.WithKind(InterfaceSubinterfaceKindKind)
)

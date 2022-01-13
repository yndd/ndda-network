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
	"github.com/yndd/nddo-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ IFNetworkInterfaceSubinterfaceList = &NetworkInterfaceSubinterfaceList{}

// +k8s:deepcopy-gen=false
type IFNetworkInterfaceSubinterfaceList interface {
	client.ObjectList

	GetInterfaceSubinterfaces() []IFNetworkInterfaceSubinterface
}

func (x *NetworkInterfaceSubinterfaceList) GetInterfaceSubinterfaces() []IFNetworkInterfaceSubinterface {
	xs := make([]IFNetworkInterfaceSubinterface, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFNetworkInterfaceSubinterface = &NetworkInterfaceSubinterface{}

// +k8s:deepcopy-gen=false
type IFNetworkInterfaceSubinterface interface {
	resource.Object
	resource.Conditioned

	GetCondition(ct nddv1.ConditionKind) nddv1.Condition
	SetConditions(c ...nddv1.Condition)
	// getters based on labels
	GetOwner() string
	GetDeploymentPolicy() string
	GetDeviceName() string
	GetEndpointGroup() string
	GetOrganization() string
	GetDeployment() string
	GetAvailabilityZone() string
	// getters based on type
	GetSubinterfaceConfig() InterfaceSubinterfaceConfig
	GetSubinterfaceIndex() string
	GetSubinterfaceIpv4() []*InterfaceSubinterfaceIpv4
}

// GetCondition
func (x *NetworkInterfaceSubinterface) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *NetworkInterfaceSubinterface) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *NetworkInterfaceSubinterface) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkInterfaceSubinterface) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkInterfaceSubinterface) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkInterfaceSubinterface) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkInterfaceSubinterface) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkInterfaceSubinterface) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkInterfaceSubinterface) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *NetworkInterfaceSubinterface) GetSubinterfaceConfig() InterfaceSubinterfaceConfig {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Config).IsZero() {
		return InterfaceSubinterfaceConfig{}
	}
	return *x.Spec.InterfaceSubinterface.Config
}
func (x *NetworkInterfaceSubinterface) GetSubinterfaceIndex() string {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Index).IsZero() {
		return ""
	}
	return *x.Spec.InterfaceSubinterface.Index
}
func (x *NetworkInterfaceSubinterface) GetSubinterfaceIpv4() []*InterfaceSubinterfaceIpv4 {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Ipv4).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Ipv4
}
func (x *NetworkInterfaceSubinterface) GetSubinterfaceIpv6() []*InterfaceSubinterfaceIpv6 {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Ipv6).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Ipv6
}

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

var _ IFNddaInterfaceSubinterfaceList = &NddaInterfaceSubinterfaceList{}

// +k8s:deepcopy-gen=false
type IFNddaInterfaceSubinterfaceList interface {
	client.ObjectList

	GetInterfaceSubinterfaces() []IFNddaInterfaceSubinterface
}

func (x *NddaInterfaceSubinterfaceList) GetInterfaceSubinterfaces() []IFNddaInterfaceSubinterface {
	xs := make([]IFNddaInterfaceSubinterface, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFNddaInterfaceSubinterface = &NddaInterfaceSubinterface{}

// +k8s:deepcopy-gen=false
type IFNddaInterfaceSubinterface interface {
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
func (x *NddaInterfaceSubinterface) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *NddaInterfaceSubinterface) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *NddaInterfaceSubinterface) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaInterfaceSubinterface) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaInterfaceSubinterface) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaInterfaceSubinterface) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaInterfaceSubinterface) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaInterfaceSubinterface) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaInterfaceSubinterface) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *NddaInterfaceSubinterface) GetSubinterfaceConfig() InterfaceSubinterfaceConfig {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Config).IsZero() {
		return InterfaceSubinterfaceConfig{}
	}
	return *x.Spec.InterfaceSubinterface.Config
}
func (x *NddaInterfaceSubinterface) GetSubinterfaceIndex() string {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Index).IsZero() {
		return ""
	}
	return *x.Spec.InterfaceSubinterface.Index
}
func (x *NddaInterfaceSubinterface) GetSubinterfaceIpv4() []*InterfaceSubinterfaceIpv4 {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Ipv4).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Ipv4
}
func (x *NddaInterfaceSubinterface) GetSubinterfaceIpv6() []*InterfaceSubinterfaceIpv6 {
	if reflect.ValueOf(x.Spec.InterfaceSubinterface.Ipv6).IsZero() {
		return nil
	}
	return x.Spec.InterfaceSubinterface.Ipv6
}

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

var _ IFNetworkInterfaceList = &NetworkInterfaceList{}

// +k8s:deepcopy-gen=false
type IFNetworkInterfaceList interface {
	client.ObjectList

	GetInterfaces() []IFNetworkInterface
}

func (x *NetworkInterfaceList) GetInterfaces() []IFNetworkInterface {
	xs := make([]IFNetworkInterface, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFNetworkInterface = &NetworkInterface{}

// +k8s:deepcopy-gen=false
type IFNetworkInterface interface {
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
	GetInterfaceConfig() InterfaceConfig
	GetInterfaceName() string
	GetInterfaceConfigKind() E_InterfaceKind
	GetInterfaceConfigLacp() bool
	GetInterfaceConfigLacpFallback() bool
	GetInterfaceConfigLag() bool
	GetInterfaceConfigLagMember() bool
	GetInterfaceConfigLagName() string
	GetInterfaceConfigName() string
}

// GetCondition
func (x *NetworkInterface) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *NetworkInterface) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *NetworkInterface) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkInterface) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkInterface) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkInterface) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkInterface) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkInterface) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkInterface) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *NetworkInterface) GetInterfaceConfig() InterfaceConfig {
	if reflect.ValueOf(x.Spec.Interface.Config).IsZero() {
		return InterfaceConfig{}
	}
	return *x.Spec.Interface.Config
}
func (x *NetworkInterface) GetInterfaceName() string {
	if reflect.ValueOf(x.Spec.Interface.Name).IsZero() {
		return ""
	}
	return *x.Spec.Interface.Name
}
func (x *NetworkInterface) GetInterfaceConfigKind() E_InterfaceKind {
	if reflect.ValueOf(x.Spec.Interface.Config.Kind).IsZero() {
		return ""
	}
	return x.Spec.Interface.Config.Kind
}
func (x *NetworkInterface) GetInterfaceConfigLacp() bool {
	if reflect.ValueOf(x.Spec.Interface.Config.Lacp).IsZero() {
		return false
	}
	return *x.Spec.Interface.Config.Lacp
}
func (x *NetworkInterface) GetInterfaceConfigLacpFallback() bool {
	if reflect.ValueOf(x.Spec.Interface.Config.LacpFallback).IsZero() {
		return false
	}
	return *x.Spec.Interface.Config.LacpFallback
}
func (x *NetworkInterface) GetInterfaceConfigLag() bool {
	if reflect.ValueOf(x.Spec.Interface.Config.Lag).IsZero() {
		return false
	}
	return *x.Spec.Interface.Config.Lag
}
func (x *NetworkInterface) GetInterfaceConfigLagMember() bool {
	if reflect.ValueOf(x.Spec.Interface.Config.LagMember).IsZero() {
		return false
	}
	return *x.Spec.Interface.Config.LagMember
}
func (x *NetworkInterface) GetInterfaceConfigLagName() string {
	if reflect.ValueOf(x.Spec.Interface.Config.LagName).IsZero() {
		return ""
	}
	return *x.Spec.Interface.Config.LagName
}
func (x *NetworkInterface) GetInterfaceConfigName() string {
	if reflect.ValueOf(x.Spec.Interface.Config.Name).IsZero() {
		return ""
	}
	return *x.Spec.Interface.Config.Name
}

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

var _ IFNddaInterfaceList = &NddaInterfaceList{}

// +k8s:deepcopy-gen=false
type IFNddaInterfaceList interface {
	client.ObjectList

	GetInterfaces() []IFNddaInterface
}

func (x *NddaInterfaceList) GetInterfaces() []IFNddaInterface {
	xs := make([]IFNddaInterface, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFNddaInterface = &NddaInterface{}

// +k8s:deepcopy-gen=false
type IFNddaInterface interface {
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
func (x *NddaInterface) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *NddaInterface) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *NddaInterface) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaInterface) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaInterface) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaInterface) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaInterface) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaInterface) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaInterface) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *NddaInterface) GetInterfaceConfig() InterfaceConfig {
	if reflect.ValueOf(x.Spec.Interface.Config).IsZero() {
		return InterfaceConfig{}
	}
	return *x.Spec.Interface.Config
}
func (x *NddaInterface) GetInterfaceName() string {
	if reflect.ValueOf(x.Spec.Interface.Name).IsZero() {
		return ""
	}
	return *x.Spec.Interface.Name
}
func (x *NddaInterface) GetInterfaceConfigKind() E_InterfaceKind {
	if reflect.ValueOf(x.Spec.Interface.Config.Kind).IsZero() {
		return ""
	}
	return x.Spec.Interface.Config.Kind
}
func (x *NddaInterface) GetInterfaceConfigLacp() bool {
	if reflect.ValueOf(x.Spec.Interface.Config.Lacp).IsZero() {
		return false
	}
	return *x.Spec.Interface.Config.Lacp
}
func (x *NddaInterface) GetInterfaceConfigLacpFallback() bool {
	if reflect.ValueOf(x.Spec.Interface.Config.LacpFallback).IsZero() {
		return false
	}
	return *x.Spec.Interface.Config.LacpFallback
}
func (x *NddaInterface) GetInterfaceConfigLag() bool {
	if reflect.ValueOf(x.Spec.Interface.Config.Lag).IsZero() {
		return false
	}
	return *x.Spec.Interface.Config.Lag
}
func (x *NddaInterface) GetInterfaceConfigLagMember() bool {
	if reflect.ValueOf(x.Spec.Interface.Config.LagMember).IsZero() {
		return false
	}
	return *x.Spec.Interface.Config.LagMember
}
func (x *NddaInterface) GetInterfaceConfigLagName() string {
	if reflect.ValueOf(x.Spec.Interface.Config.LagName).IsZero() {
		return ""
	}
	return *x.Spec.Interface.Config.LagName
}
func (x *NddaInterface) GetInterfaceConfigName() string {
	if reflect.ValueOf(x.Spec.Interface.Config.Name).IsZero() {
		return ""
	}
	return *x.Spec.Interface.Config.Name
}

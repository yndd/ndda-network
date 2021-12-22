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
	"github.com/yndd/ndd-runtime/pkg/resource"
)

var _ If = &Interface{}

// +k8s:deepcopy-gen=false
type If interface {
	resource.Object
	resource.Conditioned

	GetControllerReference() nddv1.Reference
	SetControllerReference(c nddv1.Reference)

	GetTopologyName() string
	GetNodeName() string
	GetEndpointGroup() string

	GetInterfaceName() string
	GetLag() bool
	GetLagMember() bool
	GetLagName() string
	GetLacp() bool
	GetLacpFallback() bool
}

// GetCondition of this Network Node.
func (x *Interface) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions of the Network Node.
func (x *Interface) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

// GetControllerReference of the Network Node.
func (x *Interface) GetControllerReference() nddv1.Reference {
	return x.Status.ControllerRef
}

// SetControllerReference of the Network Node.
func (x *Interface) SetControllerReference(c nddv1.Reference) {
	x.Status.ControllerRef = c
}

func (x *Interface) GetTopologyName() string {
	if reflect.ValueOf(x.Spec.TopologyName).IsZero() {
		return ""
	}
	return *x.Spec.TopologyName
}

func (x *Interface) GetNodeName() string {
	if reflect.ValueOf(x.Spec.NodeName).IsZero() {
		return ""
	}
	return *x.Spec.NodeName
}

func (x *Interface) GetEndpointGroup() string {
	if reflect.ValueOf(x.Spec.EndpointGroup).IsZero() {
		return ""
	}
	return *x.Spec.EndpointGroup
}

func (x *Interface) GetInterfaceName() string {
	if reflect.ValueOf(x.Spec.Interface.Name).IsZero() {
		return ""
	}
	return *x.Spec.Interface.Name
}

func (x *Interface) GetLag() bool {
	if reflect.ValueOf(x.Spec.Interface.Lag).IsZero() {
		return false
	}
	return *x.Spec.Interface.Lag
}

func (x *Interface) GetLagMember() bool {
	if reflect.ValueOf(x.Spec.Interface.LagMember).IsZero() {
		return false
	}
	return *x.Spec.Interface.LagMember
}

func (x *Interface) GetLagName() string {
	if reflect.ValueOf(x.Spec.Interface.LagName).IsZero() {
		return ""
	}
	return *x.Spec.Interface.LagName
}

func (x *Interface) GetLacp() bool {
	if reflect.ValueOf(x.Spec.Interface.Lacp).IsZero() {
		return false
	}
	return *x.Spec.Interface.Lacp
}

func (x *Interface) GetLacpFallback() bool {
	if reflect.ValueOf(x.Spec.Interface.LacpFallback).IsZero() {
		return false
	}
	return *x.Spec.Interface.LacpFallback
}

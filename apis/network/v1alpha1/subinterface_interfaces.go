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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ SiList = &SubInterfaceList{}

// +k8s:deepcopy-gen=false
type SiList interface {
	client.ObjectList

	GetSubInterfaces() []Si
}

func (x *SubInterfaceList) GetSubInterfaces() []Si {
	xs := make([]Si, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ Si = &SubInterface{}

// +k8s:deepcopy-gen=false
type Si interface {
	resource.Object
	resource.Conditioned

	GetControllerReference() nddv1.Reference
	SetControllerReference(c nddv1.Reference)

	GetTopologyName() string
	GetNodeName() string
	//GetEndpointGroup() string
	GetInterfaceName() string

	GetSubInterfaceIndex() string
	GetKind() string
	GetTagging() string
	GetOuterTag() uint32
	GetInnerTag() uint32
	GetIpv4() []*string
	GetIpv6() []*string
}

// GetCondition of this Network Node.
func (x *SubInterface) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions of the Network Node.
func (x *SubInterface) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

// GetControllerReference of the Network Node.
func (x *SubInterface) GetControllerReference() nddv1.Reference {
	return x.Status.ControllerRef
}

// SetControllerReference of the Network Node.
func (x *SubInterface) SetControllerReference(c nddv1.Reference) {
	x.Status.ControllerRef = c
}

func (x *SubInterface) GetTopologyName() string {
	if reflect.ValueOf(x.Spec.TopologyName).IsZero() {
		return ""
	}
	return *x.Spec.TopologyName
}

func (x *SubInterface) GetNodeName() string {
	if reflect.ValueOf(x.Spec.NodeName).IsZero() {
		return ""
	}
	return *x.Spec.NodeName
}

func (x *SubInterface) GetInterfaceName() string {
	if reflect.ValueOf(x.Spec.InterfaceName).IsZero() {
		return ""
	}
	return *x.Spec.InterfaceName
}

func (x *SubInterface) GetSubInterfaceIndex() string {
	if reflect.ValueOf(x.Spec.SubInterface.Index).IsZero() {
		return ""
	}
	return *x.Spec.SubInterface.Index
}

func (x *SubInterface) GetKind() string {
	if reflect.ValueOf(x.Spec.SubInterface.Kind).IsZero() {
		return ""
	}
	return *x.Spec.SubInterface.Kind
}

func (x *SubInterface) GetTagging() string {
	if reflect.ValueOf(x.Spec.SubInterface.Tagging).IsZero() {
		return ""
	}
	return *x.Spec.SubInterface.Tagging
}

func (x *SubInterface) GetOuterTag() uint32 {
	if reflect.ValueOf(x.Spec.SubInterface.OuterTag).IsZero() {
		return 9999
	}
	return *x.Spec.SubInterface.OuterTag
}

func (x *SubInterface) GetInnerTag() uint32 {
	if reflect.ValueOf(x.Spec.SubInterface.OuterTag).IsZero() {
		return 9999
	}
	return *x.Spec.SubInterface.OuterTag
}

func (x *SubInterface) GetIpv4() []*string {
	if reflect.ValueOf(x.Spec.SubInterface.Ipv4).IsZero() {
		return make([]*string, 0)
	}
	return x.Spec.SubInterface.Ipv4
}

func (x *SubInterface) GetIpv6() []*string {
	if reflect.ValueOf(x.Spec.SubInterface.Ipv6).IsZero() {
		return make([]*string, 0)
	}
	return x.Spec.SubInterface.Ipv6
}

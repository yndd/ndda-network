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

var _ IFNddaNetworkInstanceList = &NddaNetworkInstanceList{}

// +k8s:deepcopy-gen=false
type IFNddaNetworkInstanceList interface {
	client.ObjectList

	GetNetworkInstances() []IFNddaNetworkInstance
}

func (x *NddaNetworkInstanceList) GetNetworkInstances() []IFNddaNetworkInstance {
	xs := make([]IFNddaNetworkInstance, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFNddaNetworkInstance = &NddaNetworkInstance{}

// +k8s:deepcopy-gen=false
type IFNddaNetworkInstance interface {
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
	GetNetworkInstanceConfig() NetworkInstanceConfig
	GetNetworkInstanceName() string
	GetNetworkInstanceConfigIndex() uint32
	GetNetworkInstanceConfigInterface() []*NetworkInstanceConfigInterface
}

// GetCondition
func (x *NddaNetworkInstance) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *NddaNetworkInstance) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *NddaNetworkInstance) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaNetworkInstance) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaNetworkInstance) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaNetworkInstance) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaNetworkInstance) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaNetworkInstance) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaNetworkInstance) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *NddaNetworkInstance) GetNetworkInstanceConfig() NetworkInstanceConfig {
	if reflect.ValueOf(x.Spec.NetworkInstance.Config).IsZero() {
		return NetworkInstanceConfig{}
	}
	return *x.Spec.NetworkInstance.Config
}
func (x *NddaNetworkInstance) GetNetworkInstanceName() string {
	if reflect.ValueOf(x.Spec.NetworkInstance.Name).IsZero() {
		return ""
	}
	return *x.Spec.NetworkInstance.Name
}
func (x *NddaNetworkInstance) GetNetworkInstanceConfigIndex() uint32 {
	if reflect.ValueOf(x.Spec.NetworkInstance.Config.Index).IsZero() {
		return 0
	}
	return *x.Spec.NetworkInstance.Config.Index
}
func (x *NddaNetworkInstance) GetNetworkInstanceConfigInterface() []*NetworkInstanceConfigInterface {
	if reflect.ValueOf(x.Spec.NetworkInstance.Config.Interface).IsZero() {
		return nil
	}
	return x.Spec.NetworkInstance.Config.Interface
}

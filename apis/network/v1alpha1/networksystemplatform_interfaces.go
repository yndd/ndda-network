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

var _ IFNetworkSystemPlatformList = &NetworkSystemPlatformList{}

// +k8s:deepcopy-gen=false
type IFNetworkSystemPlatformList interface {
	client.ObjectList

	GetSystemPlatforms() []IFNetworkSystemPlatform
}

func (x *NetworkSystemPlatformList) GetSystemPlatforms() []IFNetworkSystemPlatform {
	xs := make([]IFNetworkSystemPlatform, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFNetworkSystemPlatform = &NetworkSystemPlatform{}

// +k8s:deepcopy-gen=false
type IFNetworkSystemPlatform interface {
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
	GetPlatformConfig() SystemPlatformConfig
	GetPlatformConfigIndex() uint32
	GetPlatformConfigKind() E_SystemPlatformKind
	GetPlatformConfigName() string
	GetPlatformConfigVersion() string
}

// GetCondition
func (x *NetworkSystemPlatform) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *NetworkSystemPlatform) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *NetworkSystemPlatform) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkSystemPlatform) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkSystemPlatform) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkSystemPlatform) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkSystemPlatform) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkSystemPlatform) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NetworkSystemPlatform) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *NetworkSystemPlatform) GetPlatformConfig() SystemPlatformConfig {
	if reflect.ValueOf(x.Spec.SystemPlatform.Config).IsZero() {
		return SystemPlatformConfig{}
	}
	return *x.Spec.SystemPlatform.Config
}
func (x *NetworkSystemPlatform) GetPlatformConfigIndex() uint32 {
	if reflect.ValueOf(x.Spec.SystemPlatform.Config.Index).IsZero() {
		return 0
	}
	return *x.Spec.SystemPlatform.Config.Index
}
func (x *NetworkSystemPlatform) GetPlatformConfigKind() E_SystemPlatformKind {
	if reflect.ValueOf(x.Spec.SystemPlatform.Config.Kind).IsZero() {
		return ""
	}
	return x.Spec.SystemPlatform.Config.Kind
}
func (x *NetworkSystemPlatform) GetPlatformConfigName() string {
	if reflect.ValueOf(x.Spec.SystemPlatform.Config.Name).IsZero() {
		return ""
	}
	return *x.Spec.SystemPlatform.Config.Name
}
func (x *NetworkSystemPlatform) GetPlatformConfigVersion() string {
	if reflect.ValueOf(x.Spec.SystemPlatform.Config.Version).IsZero() {
		return ""
	}
	return *x.Spec.SystemPlatform.Config.Version
}

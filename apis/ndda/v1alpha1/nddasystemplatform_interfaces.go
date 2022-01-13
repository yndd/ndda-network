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

var _ IFNddaSystemPlatformList = &NddaSystemPlatformList{}

// +k8s:deepcopy-gen=false
type IFNddaSystemPlatformList interface {
	client.ObjectList

	GetSystemPlatforms() []IFNddaSystemPlatform
}

func (x *NddaSystemPlatformList) GetSystemPlatforms() []IFNddaSystemPlatform {
	xs := make([]IFNddaSystemPlatform, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ IFNddaSystemPlatform = &NddaSystemPlatform{}

// +k8s:deepcopy-gen=false
type IFNddaSystemPlatform interface {
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
func (x *NddaSystemPlatform) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions
func (x *NddaSystemPlatform) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

func (x *NddaSystemPlatform) GetOwner() string {
	if s, ok := x.GetLabels()[LabelNddaOwner]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaSystemPlatform) GetDeploymentPolicy() string {
	if s, ok := x.GetLabels()[LabelNddaDeploymentPolicy]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaSystemPlatform) GetDeviceName() string {
	if s, ok := x.GetLabels()[LabelNddaDevice]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaSystemPlatform) GetEndpointGroup() string {
	if s, ok := x.GetLabels()[LabelNddaEndpointGroup]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaSystemPlatform) GetOrganization() string {
	if s, ok := x.GetLabels()[LabelNddaOrganization]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaSystemPlatform) GetDeployment() string {
	if s, ok := x.GetLabels()[LabelNddaDeployment]; !ok {
		return ""
	} else {
		return s
	}
}

func (x *NddaSystemPlatform) GetAvailabilityZone() string {
	if s, ok := x.GetLabels()[LabelNddaAvailabilityZone]; !ok {
		return ""
	} else {
		return s
	}
}
func (x *NddaSystemPlatform) GetPlatformConfig() SystemPlatformConfig {
	if reflect.ValueOf(x.Spec.SystemPlatform.Config).IsZero() {
		return SystemPlatformConfig{}
	}
	return *x.Spec.SystemPlatform.Config
}
func (x *NddaSystemPlatform) GetPlatformConfigIndex() uint32 {
	if reflect.ValueOf(x.Spec.SystemPlatform.Config.Index).IsZero() {
		return 0
	}
	return *x.Spec.SystemPlatform.Config.Index
}
func (x *NddaSystemPlatform) GetPlatformConfigKind() E_SystemPlatformKind {
	if reflect.ValueOf(x.Spec.SystemPlatform.Config.Kind).IsZero() {
		return ""
	}
	return x.Spec.SystemPlatform.Config.Kind
}
func (x *NddaSystemPlatform) GetPlatformConfigName() string {
	if reflect.ValueOf(x.Spec.SystemPlatform.Config.Name).IsZero() {
		return ""
	}
	return *x.Spec.SystemPlatform.Config.Name
}
func (x *NddaSystemPlatform) GetPlatformConfigVersion() string {
	if reflect.ValueOf(x.Spec.SystemPlatform.Config.Version).IsZero() {
		return ""
	}
	return *x.Spec.SystemPlatform.Config.Version
}

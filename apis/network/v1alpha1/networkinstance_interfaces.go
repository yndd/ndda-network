package v1alpha1

import (
	"reflect"

	nddv1 "github.com/yndd/ndd-runtime/apis/common/v1"
	"github.com/yndd/ndd-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ NiList = &NetworkInstanceList{}

// +k8s:deepcopy-gen=false
type NiList interface {
	client.ObjectList

	GetNetworkInstance() []Ni
}

func (x *NetworkInstanceList) GetNetworkInstance() []Ni {
	xs := make([]Ni, len(x.Items))
	for i, r := range x.Items {
		r := r // Pin range variable so we can take its address.
		xs[i] = &r
	}
	return xs
}

var _ Ni = &NetworkInstance{}

// +k8s:deepcopy-gen=false
type Ni interface {
	resource.Object
	resource.Conditioned

	GetControllerReference() nddv1.Reference
	SetControllerReference(c nddv1.Reference)

	GetTopologyName() string
	GetNodeName() string

	GetName() string
	GetKind() string
	//GetInterfaces() []*NetworkNetworkInstanceInterface
	GetInterfaces() map[string]string
}

// GetCondition of this Network Node.
func (x *NetworkInstance) GetCondition(ct nddv1.ConditionKind) nddv1.Condition {
	return x.Status.GetCondition(ct)
}

// SetConditions of the Network Node.
func (x *NetworkInstance) SetConditions(c ...nddv1.Condition) {
	x.Status.SetConditions(c...)
}

// GetControllerReference of the Network Node.
func (x *NetworkInstance) GetControllerReference() nddv1.Reference {
	return x.Status.ControllerRef
}

// SetControllerReference of the Network Node.
func (x *NetworkInstance) SetControllerReference(c nddv1.Reference) {
	x.Status.ControllerRef = c
}

func (x *NetworkInstance) GetTopologyName() string {
	if reflect.ValueOf(x.Spec.TopologyName).IsZero() {
		return ""
	}
	return *x.Spec.TopologyName
}

func (x *NetworkInstance) GetNodeName() string {
	if reflect.ValueOf(x.Spec.NodeName).IsZero() {
		return ""
	}
	return *x.Spec.NodeName
}

func (x *NetworkInstance) GetName() string {
	if reflect.ValueOf(x.Spec.NetworkInstance.Name).IsZero() {
		return ""
	}
	return *x.Spec.NetworkInstance.Name
}

func (x *NetworkInstance) GetKind() string {
	if reflect.ValueOf(x.Spec.NetworkInstance.Kind).IsZero() {
		return ""
	}
	return *x.Spec.NetworkInstance.Kind
}

/*
func (x *NetworkInstance) GetInterfaces() []*NetworkNetworkInstanceInterface {
	i := make([]*NetworkNetworkInstanceInterface, 0)
	if reflect.ValueOf(x.Spec.NetworkInstance.Interface).IsZero() {
		return i
	}
	return x.Spec.NetworkInstance.Interface
}
*/

func (x *NetworkInstance) GetInterfaces() map[string]string {
	i := make(map[string]string)
	if reflect.ValueOf(x.Spec.NetworkInstance.Interface).IsZero() {
		return i
	}
	for _, itfce := range x.Spec.NetworkInstance.Interface {
		i[*itfce.Name] = *itfce.Kind
	}
	return i
}

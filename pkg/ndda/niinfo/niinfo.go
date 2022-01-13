package niinfo

import (
	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
)

/*
type NiInfo interface {
	GetItfceName() string
	GetItfceIndex() uint32
	GetItfceKind() networkv1alpha1.E_InterfaceKind
	GetInnerVlanId() uint16
	GetOuterVlanId() uint16
	GetIpv4Prefixes() []*string
	GetIpv6Prefixes() []*string
	SetItfceName(string)
	SetItfceIndex(uint32)
	SetItfceKind(networkv1alpha1.E_InterfaceKind)
	SetInnerVlanId(uint16)
	SetOuterVlanId(uint16)
	SetIpv4Prefixes([]*string)
	SetIpv6Prefixes([]*string)
}
*/

type NiInfoOption func(*NiInfo)

/*
func WithNiName(s string) ItfceInfoOption {
	return func(r *itfceInfo) {
		r.itfceName = &s
	}
}

func WithItfceIndex(s uint32) ItfceInfoOption {
	return func(r *itfceInfo) {
		r.itfceIndex = &s
	}
}

func WithItfceKind(s networkv1alpha1.E_InterfaceKind) ItfceInfoOption {
	return func(r *itfceInfo) {
		r.itfceKind = s
	}
}
*/

func NewItfceInfo(opts ...NiInfoOption) *NiInfo {
	i := &NiInfo{}

	for _, f := range opts {
		f(i)
	}

	return i
}

type NiInfo struct {
	Name  *string
	Index *uint32
	Kind  networkv1alpha1.E_NetworkInstanceKind
}

func (x *NiInfo) GetNiName() string {
	return *x.Name
}

func (x *NiInfo) GetNiIndex() uint32 {
	return *x.Index
}

func (x *NiInfo) GetNiKind() networkv1alpha1.E_NetworkInstanceKind {
	return x.Kind
}

func (x *NiInfo) SetNiName(s string) {
	x.Name = &s
}

func (x *NiInfo) SetNiIndex(s uint32) {
	x.Index = &s
}

func (x *NiInfo) SetNiKind(s networkv1alpha1.E_NetworkInstanceKind) {
	x.Kind = s
}

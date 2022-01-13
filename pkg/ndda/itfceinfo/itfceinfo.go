package itfceinfo

import (
	"reflect"

	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
)

type ItfceInfo interface {
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

type ItfceInfoOption func(*itfceInfo)

func WithItfceName(s string) ItfceInfoOption {
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

func WithInnerVlanId(s uint16) ItfceInfoOption {
	return func(r *itfceInfo) {
		r.innerVlanId = &s
	}
}

func WithOuterVlanId(s uint16) ItfceInfoOption {
	return func(r *itfceInfo) {
		r.outerVlanId = &s
	}
}

func WithIpv4Prefixes(s []*string) ItfceInfoOption {
	return func(r *itfceInfo) {
		r.ipv4Prefixes = s
	}
}

func WithIpv6Prefixes(s []*string) ItfceInfoOption {
	return func(r *itfceInfo) {
		r.ipv6Prefixes = s
	}
}

func NewItfceInfo(opts ...ItfceInfoOption) ItfceInfo {
	i := &itfceInfo{
		ipv4Prefixes: make([]*string, 0),
		ipv6Prefixes: make([]*string, 0),
	}

	for _, f := range opts {
		f(i)
	}

	return i
}

type itfceInfo struct {
	itfceName    *string
	itfceIndex   *uint32
	itfceKind    networkv1alpha1.E_InterfaceKind
	innerVlanId  *uint16
	outerVlanId  *uint16
	ipv4Prefixes []*string
	ipv6Prefixes []*string
}

func (x *itfceInfo) GetItfceName() string {
	return *x.itfceName
}

func (x *itfceInfo) GetItfceIndex() uint32 {
	return *x.itfceIndex
}

func (x *itfceInfo) GetItfceKind() networkv1alpha1.E_InterfaceKind {
	return x.itfceKind
}

func (x *itfceInfo) GetInnerVlanId() uint16 {
	if reflect.ValueOf(x.innerVlanId).IsZero() {
		return 9999
	}
	return *x.innerVlanId
}

func (x *itfceInfo) GetOuterVlanId() uint16 {
	if reflect.ValueOf(x.outerVlanId).IsZero() {
		return 9999
	}
	return *x.outerVlanId
}

func (x *itfceInfo) GetIpv4Prefixes() []*string {
	return x.ipv4Prefixes
}

func (x *itfceInfo) GetIpv6Prefixes() []*string {
	return x.ipv6Prefixes
}

func (x *itfceInfo) SetItfceName(s string) {
	x.itfceName = &s
}

func (x *itfceInfo) SetItfceIndex(s uint32) {
	x.itfceIndex = &s
}

func (x *itfceInfo) SetItfceKind(s networkv1alpha1.E_InterfaceKind) {
	x.itfceKind = s
}

func (x *itfceInfo) SetInnerVlanId(s uint16) {
	x.innerVlanId = &s
}

func (x *itfceInfo) SetOuterVlanId(s uint16) {
	x.outerVlanId = &s
}

func (x *itfceInfo) SetIpv4Prefixes(s []*string) {
	x.ipv4Prefixes = s
}

func (x *itfceInfo) SetIpv6Prefixes(s []*string) {
	x.ipv6Prefixes = s
}

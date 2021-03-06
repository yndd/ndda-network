package niinfo

import (
	"strings"
	//networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
	"github.com/yndd/ndda-network/pkg/ygotndda"
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

func GetBdName(bdName string) string {
	return strings.Join([]string{bdName, strings.ToLower(ygotndda.NddaCommon_NiKind_BRIDGED.String())}, "-")
	//return strings.Join([]string{bdName, strings.ToLower(string(networkv1alpha1.E_NetworkInstanceKind_BRIDGED))}, "-")
}

func GetRtName(rtName string) string {
	return strings.Join([]string{rtName, strings.ToLower(ygotndda.NddaCommon_NiKind_ROUTED.String())}, "-")
	//return strings.Join([]string{rtName, strings.ToLower(string(networkv1alpha1.E_NetworkInstanceKind_ROUTED))}, "-")
}

func NewNiInfo(opts ...NiInfoOption) *NiInfo {
	i := &NiInfo{}

	for _, f := range opts {
		f(i)
	}

	return i
}

type NiInfo struct {
	Name     *string
	Index    *uint32
	Registry *string
}

func (x *NiInfo) GetNiName() string {
	return *x.Name
}

func (x *NiInfo) GetNiIndex() uint32 {
	return *x.Index
}

func (x *NiInfo) GetNiRegistry() string {
	return *x.Registry
}

func (x *NiInfo) SetNiName(s string) {
	x.Name = &s
}

func (x *NiInfo) SetNiIndex(s uint32) {
	x.Index = &s
}

func (x *NiInfo) SetNiRegistry(s string) {
	x.Registry = &s
}

func (x *NiInfo) GetNiKind() ygotndda.E_NddaCommon_NiKind {
	if strings.HasSuffix(*x.Name, strings.ToLower(ygotndda.NddaCommon_NiKind_BRIDGED.String())) {
		return ygotndda.NddaCommon_NiKind_BRIDGED
	}
	if strings.HasSuffix(*x.Name, strings.ToLower(ygotndda.NddaCommon_NiKind_ROUTED.String())) {
		return ygotndda.NddaCommon_NiKind_ROUTED
	}
	return ygotndda.NddaCommon_NiKind_UNSET
}

/*
func (x *NiInfo) GetNiKind() networkv1alpha1.E_NetworkInstanceKind {
	if strings.HasSuffix(*x.Name, strings.ToLower(string(networkv1alpha1.E_NetworkInstanceKind_BRIDGED))) {
		return networkv1alpha1.E_NetworkInstanceKind_BRIDGED
	}
	if strings.HasSuffix(*x.Name, strings.ToLower(string(networkv1alpha1.E_NetworkInstanceKind_ROUTED))) {
		return networkv1alpha1.E_NetworkInstanceKind_ROUTED
	}
	return ""
}
*/

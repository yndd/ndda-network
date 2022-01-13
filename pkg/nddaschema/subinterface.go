package ndda

import (
	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
)

type SubInterface interface {
	Update(d *networkv1alpha1.InterfaceSubinterface)
	// add method to add ipv4 and ipv6
}

func NewSubInterface(p Interface, index string) SubInterface {
	return &subinterface{
		// parent
		parent: p,
		// children
		// data with key
		Subinterface: &networkv1alpha1.InterfaceSubinterface{
			Index: &index,
		},
	}
}

type subinterface struct {
	// parent
	parent Interface
	// children
	// Data
	Subinterface *networkv1alpha1.InterfaceSubinterface
}

func (x *subinterface) Update(d *networkv1alpha1.InterfaceSubinterface) {
	x.Subinterface = d
}

func (x *subinterface) AddIPv4(ai *networkv1alpha1.InterfaceSubinterfaceIpv4) {
	x.Subinterface.Ipv4 = append(x.Subinterface.Ipv4, ai)
}

func (x *subinterface) AddIPv6(ai *networkv1alpha1.InterfaceSubinterfaceIpv6) {
	x.Subinterface.Ipv6 = append(x.Subinterface.Ipv6, ai)
}

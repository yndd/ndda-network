package ndda

import (
	nddav1alpha1 "github.com/yndd/ndda-network/apis/ndda/v1alpha1"
)

type SubInterface interface {
	Update(d *nddav1alpha1.InterfaceSubinterface)
}

func NewSubInterface(p Interface, index string) SubInterface{
	return &subinterface{
		// parent
		parent:       p,
		// children
		// data with key
		Subinterface: &nddav1alpha1.InterfaceSubinterface{
			Index: &index,
		},
	}
}

type subinterface struct {
	// parent
	parent       Interface
	// children
	// Data
	Subinterface *nddav1alpha1.InterfaceSubinterface
}

func (x *subinterface) Update(d *nddav1alpha1.InterfaceSubinterface) {
	x.Subinterface = d
}
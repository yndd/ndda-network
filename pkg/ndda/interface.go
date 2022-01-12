package ndda

import (
	nddav1alpha1 "github.com/yndd/ndda-network/apis/ndda/v1alpha1"
)

type Interface interface {
	// children
	NewSubInterface(index string) SubInterface
	Update(x *nddav1alpha1.Interface)
}

func NewInterface(p Device, name string) Interface {
	return &itfce{
		// parent
		parent:       p,
		// children
		Subinterface: make(map[string]SubInterface),
		// Data key
		//Interface: &nddav1alpha1.Interface{
		//	Name: &name,
		//},
	}
}

type itfce struct {
	// parent
	parent Device
	// children
	Subinterface map[string]SubInterface
	// Data
	Interface *nddav1alpha1.Interface
}

func (x *itfce) NewSubInterface(index string) SubInterface {
	if _, ok := x.Subinterface[index]; !ok {
		x.Subinterface[index] = NewSubInterface(x, index)
	}
	return x.Subinterface[index]
}

func (x *itfce) Update(d *nddav1alpha1.Interface) {
	x.Interface = d
}
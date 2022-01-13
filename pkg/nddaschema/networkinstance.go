package ndda

import (
	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
)

type NetworkInstance interface {
	Update(d *networkv1alpha1.NetworkInstance)
	// add methods to add list entries e.g. AddInterface in network instance
}

func NewNetworkInstance(p Device, name string) NetworkInstance {
	return &networkInstance{
		// parent
		parent: p,
		// children
		// data key
		//networkInstance: &networkv1alpha1.NetworkInstance{
		//	Name: &name,
		//},
	}
}

type networkInstance struct {
	// parent
	parent Device
	// children
	// data
	networkInstance *networkv1alpha1.NetworkInstance
}

func (x *networkInstance) Update(d *networkv1alpha1.NetworkInstance) {
	x.networkInstance = d
}

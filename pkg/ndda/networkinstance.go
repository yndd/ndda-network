package ndda

import (
	nddav1alpha1 "github.com/yndd/ndda-network/apis/ndda/v1alpha1"
)

type NetworkInstance interface {
}

func NewNetworkInstance(p Device, name string) NetworkInstance {
	return &networkInstance{
		// parent
		parent: p,
		// children
		// data key
		//networkInstance: &nddav1alpha1.NetworkInstance{
		//	Name: &name,
		//},
	}
}

type networkInstance struct {
	// parent
	parent Device
	// children
	// data
	networkInstance *nddav1alpha1.NetworkInstance
}

func (x *networkInstance) Update(d *nddav1alpha1.NetworkInstance) {
	x.networkInstance = d
}

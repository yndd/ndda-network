package ndda

import (
	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
)

type SystemPlatform interface {
	Update(d *networkv1alpha1.SystemPlatform)
}

func NewSystemPlatform(p Device) SystemPlatform {
	return &systemPlatform{
		parent: p,
	}
}

type systemPlatform struct {
	// parent
	parent Device
	// children
	// Data
	SystemPlatform *networkv1alpha1.SystemPlatform
}

func (x *systemPlatform) Update(d *networkv1alpha1.SystemPlatform) {
	x.SystemPlatform = d
}

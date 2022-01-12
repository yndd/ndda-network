package ndda

import (
	nddav1alpha1 "github.com/yndd/ndda-network/apis/ndda/v1alpha1"
)

type SystemPlatform interface {
	Update(d *nddav1alpha1.SystemPlatform)
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
	SystemPlatform *nddav1alpha1.SystemPlatform
}

func (x *systemPlatform) Update(d *nddav1alpha1.SystemPlatform) {
	x.SystemPlatform = d
}

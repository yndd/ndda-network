/*
Copyright 2021 NDD.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nddaschema

import (
	nddav1alpha1 "github.com/yndd/ndda-network/apis/ndda/v1alpha1"
)

type SystemPlatform interface {
	// methods children
	// methods data
	Update(x *nddav1alpha1.SystemPlatform)
}

func NewSystemPlatform(p Device, key string) SystemPlatform {
	return &systemplatform{
		// parent
		parent: p,
		// children
		// data key
		//SystemPlatform: &nddav1alpha1.SystemPlatform{
		//	Name: &name,
		//},
	}
}

type systemplatform struct {
	// parent
	parent Device
	// children
	// Data
	SystemPlatform *nddav1alpha1.SystemPlatform
}

// children
// Data
func (x *systemplatform) Update(d *nddav1alpha1.SystemPlatform) {
	x.SystemPlatform = d
}

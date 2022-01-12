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

package ndda

import (
	nddav1alpha1 "github.com/yndd/ndda-network/apis/ndda/v1alpha1"
)

type NetworkInstance interface {
	// methods children
	// methods data
	Update(x *nddav1alpha1.NetworkInstance)
	AddNetworkInstanceInterface(ai *nddav1alpha1.NetworkInstanceConfigInterface)
}

func NewNetworkInstance(p Device, key string) NetworkInstance {
	return &networkinstance{
		// parent
		parent: p,
		// children
		// data key
		//NetworkInstance: &nddav1alpha1.NetworkInstance{
		//	Name: &name,
		//},
	}
}

type networkinstance struct {
	// parent
	parent Device
	// children
	// Data
	NetworkInstance *nddav1alpha1.NetworkInstance
}

// children
// Data
func (x *networkinstance) Update(d *nddav1alpha1.NetworkInstance) {
	x.NetworkInstance = d
}

// NetworkInstance interface network-instance-config NetworkInstance [network-instance config]
func (x *networkinstance) AddNetworkInstanceInterface(ai *nddav1alpha1.NetworkInstanceConfigInterface) {
	x.NetworkInstance.Config.Interface = append(x.NetworkInstance.Config.Interface, ai)
}

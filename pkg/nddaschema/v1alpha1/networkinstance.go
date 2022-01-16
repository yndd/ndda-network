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
	"github.com/yndd/nddo-runtime/pkg/resource"

	nddav1alpha1 "github.com/yndd/ndda-network/apis/ndda/v1alpha1"
)

type NetworkInstance interface {
	// methods children
	// methods data
	Update(x *nddav1alpha1.NetworkInstance)
	Get() *nddav1alpha1.NetworkInstance
	GetKey() string
	AddNetworkInstanceInterface(ai *nddav1alpha1.NetworkInstanceConfigInterface)
	// methods schema
}

func NewNetworkInstance(c resource.ClientApplicator, p Device, key string) NetworkInstance {
	return &networkinstance{
		// k8s client
		client: c,
		// key
		key: key,
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
	// k8s client
	client resource.ClientApplicator
	// key
	key string
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

func (x *networkinstance) Get() *nddav1alpha1.NetworkInstance {
	return x.NetworkInstance
}

func (x *networkinstance) GetKey() string {
	return x.key
}

// NetworkInstance interface network-instance-config NetworkInstance [network-instance config]
func (x *networkinstance) AddNetworkInstanceInterface(ai *nddav1alpha1.NetworkInstanceConfigInterface) {
	x.NetworkInstance.Config.Interface = append(x.NetworkInstance.Config.Interface, ai)
}

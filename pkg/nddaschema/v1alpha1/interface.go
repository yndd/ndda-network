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

type Interface interface {
	// methods children
	NewInterfaceSubinterface(c resource.ClientApplicator, key string) InterfaceSubinterface
	GetInterfaceSubinterfaces() map[string]InterfaceSubinterface
	// methods data
	Update(x *nddav1alpha1.Interface)
	Get() *nddav1alpha1.Interface
	GetKey() string
	// methods schema
}

func NewInterface(c resource.ClientApplicator, p Device, key string) Interface {
	return &itfce{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		InterfaceSubinterface: make(map[string]InterfaceSubinterface),
		// data key
		//Interface: &nddav1alpha1.Interface{
		//	Name: &name,
		//},
	}
}

type itfce struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	InterfaceSubinterface map[string]InterfaceSubinterface
	// Data
	Interface *nddav1alpha1.Interface
}

// children
func (x *itfce) NewInterfaceSubinterface(c resource.ClientApplicator, key string) InterfaceSubinterface {
	if _, ok := x.InterfaceSubinterface[key]; !ok {
		x.InterfaceSubinterface[key] = NewInterfaceSubinterface(c, x, key)
	}
	return x.InterfaceSubinterface[key]
}
func (x *itfce) GetInterfaceSubinterfaces() map[string]InterfaceSubinterface {
	return x.InterfaceSubinterface
}

// Data
func (x *itfce) Update(d *nddav1alpha1.Interface) {
	x.Interface = d
}

func (x *itfce) Get() *nddav1alpha1.Interface {
	return x.Interface
}

func (x *itfce) GetKey() string {
	return x.key
}

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

type Device interface {
	// methods children
	NewInterface(key string) Interface
	NewNetworkInstance(key string) NetworkInstance
	NewSystemPlatform(key string) SystemPlatform
	// methods data
}

func NewDevice(p Schema, key string) Device {
	return &device{
		// parent
		parent: p,
		// children
		Interface:       make(map[string]Interface),
		NetworkInstance: make(map[string]NetworkInstance),
		SystemPlatform:  make(map[string]SystemPlatform),
		// data key
		//Device: &nddav1alpha1.Device{
		//	Name: &name,
		//},
	}
}

type device struct {
	// parent
	parent Schema
	// children
	Interface       map[string]Interface
	NetworkInstance map[string]NetworkInstance
	SystemPlatform  map[string]SystemPlatform
	// Data
}

// children
func (x *device) NewInterface(key string) Interface {
	if _, ok := x.Interface[key]; !ok {
		x.Interface[key] = NewInterface(x, key)
	}
	return x.Interface[key]
}
func (x *device) NewNetworkInstance(key string) NetworkInstance {
	if _, ok := x.NetworkInstance[key]; !ok {
		x.NetworkInstance[key] = NewNetworkInstance(x, key)
	}
	return x.NetworkInstance[key]
}
func (x *device) NewSystemPlatform(key string) SystemPlatform {
	if _, ok := x.SystemPlatform[key]; !ok {
		x.SystemPlatform[key] = NewSystemPlatform(x, key)
	}
	return x.SystemPlatform[key]
}

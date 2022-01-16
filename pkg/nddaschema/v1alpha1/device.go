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
)

type Device interface {
	// methods children
	NewInterface(c resource.ClientApplicator, key string) Interface
	NewNetworkInstance(c resource.ClientApplicator, key string) NetworkInstance
	NewSystemPlatform(c resource.ClientApplicator, key string) SystemPlatform
	GetInterfaces() map[string]Interface
	GetNetworkInstances() map[string]NetworkInstance
	GetSystemPlatforms() map[string]SystemPlatform
	// methods data
	// methods schema
}

func NewDevice(c resource.ClientApplicator, p Schema, key string) Device {
	return &device{
		// k8s client
		client: c,
		// key
		key: key,
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
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Schema
	// children
	Interface       map[string]Interface
	NetworkInstance map[string]NetworkInstance
	SystemPlatform  map[string]SystemPlatform
	// Data
}

// children
func (x *device) NewInterface(c resource.ClientApplicator, key string) Interface {
	if _, ok := x.Interface[key]; !ok {
		x.Interface[key] = NewInterface(c, x, key)
	}
	return x.Interface[key]
}
func (x *device) NewNetworkInstance(c resource.ClientApplicator, key string) NetworkInstance {
	if _, ok := x.NetworkInstance[key]; !ok {
		x.NetworkInstance[key] = NewNetworkInstance(c, x, key)
	}
	return x.NetworkInstance[key]
}
func (x *device) NewSystemPlatform(c resource.ClientApplicator, key string) SystemPlatform {
	if _, ok := x.SystemPlatform[key]; !ok {
		x.SystemPlatform[key] = NewSystemPlatform(c, x, key)
	}
	return x.SystemPlatform[key]
}
func (x *device) GetInterfaces() map[string]Interface {
	return x.Interface
}
func (x *device) GetNetworkInstances() map[string]NetworkInstance {
	return x.NetworkInstance
}
func (x *device) GetSystemPlatforms() map[string]SystemPlatform {
	return x.SystemPlatform
}

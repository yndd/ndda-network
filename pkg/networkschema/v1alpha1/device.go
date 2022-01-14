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

package networkschema

import (
	"fmt"
	"strings"
)

type Device interface {
	// methods children
	NewInterface(key string) Interface
	NewNetworkInstance(key string) NetworkInstance
	NewSystemPlatform(key string) SystemPlatform
	GetInterfaces() map[string]Interface
	GetNetworkInstances() map[string]NetworkInstance
	GetSystemPlatforms() map[string]SystemPlatform
	// methods data
	Print(nodeName string, n int)
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
		//Device: &networkv1alpha1.Device{
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
func (x *device) GetInterfaces() map[string]Interface {
	return x.Interface
}
func (x *device) GetNetworkInstances() map[string]NetworkInstance {
	return x.NetworkInstance
}
func (x *device) GetSystemPlatforms() map[string]SystemPlatform {
	return x.SystemPlatform
}

func (x *device) Print(nodeName string, n int) {
	fmt.Printf("%s Node Name: %s\n", nodeName, strings.Repeat(" ", n))
	n++
	for itfceName, i := range x.GetInterfaces() {
		i.Print(itfceName, n)
	}
	for niName, ni := range x.GetNetworkInstances() {
		ni.Print(niName, n)
	}
}

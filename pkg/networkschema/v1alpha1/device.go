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
	"context"
	"fmt"
	"strings"

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
	Print(nodeName string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewDevice(c resource.ClientApplicator, p Schema, key string) Device {
	return &device{
		client: c,
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
	client resource.ClientApplicator
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
		x.Interface[key] = NewInterface(x.client, x, key)
	}
	return x.Interface[key]
}
func (x *device) NewNetworkInstance(c resource.ClientApplicator, key string) NetworkInstance {
	if _, ok := x.NetworkInstance[key]; !ok {
		x.NetworkInstance[key] = NewNetworkInstance(x.client, x, key)
	}
	return x.NetworkInstance[key]
}
func (x *device) NewSystemPlatform(c resource.ClientApplicator, key string) SystemPlatform {
	if _, ok := x.SystemPlatform[key]; !ok {
		x.SystemPlatform[key] = NewSystemPlatform(x.client, x, key)
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

func (x *device) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string) error {
	for _, i := range x.GetInterfaces() {
		if err := i.DeploySchema(ctx, mg, deviceName); err != nil {
			return err
		}
	}
	for _, ni := range x.GetNetworkInstances() {
		if err := ni.DeploySchema(ctx, mg, deviceName); err != nil {
			return err
		}
	}

	return nil
}

func (x *device) InitializeDummySchema() {
	i := x.NewInterface(x.client, "dummy")
	i.InitializeDummySchema()
	ni := x.NewNetworkInstance(x.client, "dummy")
	ni.InitializeDummySchema()
	p := x.NewSystemPlatform(x.client, "dummy")
	p.InitializeDummySchema()

}

func (x *device) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	for _, i := range x.GetInterfaces() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkInstances() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemPlatforms() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *device) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	for _, i := range x.GetInterfaces() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkInstances() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemPlatforms() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *device) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	for _, i := range x.GetInterfaces() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetNetworkInstances() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	for _, i := range x.GetSystemPlatforms() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	return nil
}

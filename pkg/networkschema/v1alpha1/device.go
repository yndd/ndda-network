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
	"encoding/json"
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
	GetKey() []string
	Get() interface{}
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
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
		//Device: &networkv1alpha1.Device{
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

// key type/method

type DeviceKey struct {
	Name string
}

func WithDeviceKey(key *DeviceKey) string {
	d, err := json.Marshal(key)
	if err != nil {
		return ""
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	return getKey(x1)
}

// methods children
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

// methods data
func (x *device) Get() interface{} {
	return nil
}

func (x *device) GetKey() []string {
	return strings.Split(x.key, ".")
}

// methods schema

func (x *device) Print(key string, n int) {
	if x.Get() != nil {
		return
	} else {
		fmt.Printf("%s Device: %s\n", strings.Repeat(" ", n), key)
	}

	n++
	for key, i := range x.GetInterfaces() {
		i.Print(key, n)
	}
	for key, i := range x.GetNetworkInstances() {
		i.Print(key, n)
	}
	for key, i := range x.GetSystemPlatforms() {
		i.Print(key, n)
	}
}

func (x *device) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		return nil
	}
	for _, r := range x.GetInterfaces() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetNetworkInstances() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}
	for _, r := range x.GetSystemPlatforms() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}

	return nil
}

func (x *device) InitializeDummySchema() {
	c0 := x.NewInterface(x.client, "dummy")
	c0.InitializeDummySchema()
	c1 := x.NewNetworkInstance(x.client, "dummy")
	c1.InitializeDummySchema()
	c2 := x.NewSystemPlatform(x.client, "dummy")
	c2.InitializeDummySchema()
}

func (x *device) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list

	// children
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
	// local CR validation

	// children
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
	// local CR deletion

	// children
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

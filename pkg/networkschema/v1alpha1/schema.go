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

type Schema interface {
	NewDevice(name string) Device
	GetDevices() map[string]Device
}

func NewSchema() Schema {
	return &schema{
		// parent nil
		// children
		devices: make(map[string]Device),
		// data key
	}
}

type schema struct {
	// parent is nil
	// children
	devices map[string]Device
	// data is nil
}

func (x *schema) NewDevice(name string) Device {
	if _, ok := x.devices[name]; !ok {
		x.devices[name] = NewDevice(x, name)
	}
	return x.devices[name]
}

func (x *schema) GetDevices() map[string]Device {
	return x.devices
}

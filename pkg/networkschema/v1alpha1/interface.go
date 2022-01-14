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

	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
)

type Interface interface {
	// methods children
	NewInterfaceSubinterface(key string) InterfaceSubinterface
	GetInterfaceSubinterfaces() map[string]InterfaceSubinterface
	// methods data
	Update(x *networkv1alpha1.Interface)

	Print(itfceName string, n int)
}

func NewInterface(p Device, key string) Interface {
	return &itfce{
		// parent
		parent: p,
		// children
		InterfaceSubinterface: make(map[string]InterfaceSubinterface),
		// data key
		//Interface: &networkv1alpha1.Interface{
		//	Name: &name,
		//},
	}
}

type itfce struct {
	// parent
	parent Device
	// children
	InterfaceSubinterface map[string]InterfaceSubinterface
	// Data
	Interface *networkv1alpha1.Interface
}

// children
func (x *itfce) NewInterfaceSubinterface(key string) InterfaceSubinterface {
	if _, ok := x.InterfaceSubinterface[key]; !ok {
		x.InterfaceSubinterface[key] = NewInterfaceSubinterface(x, key)
	}
	return x.InterfaceSubinterface[key]
}
func (x *itfce) GetInterfaceSubinterfaces() map[string]InterfaceSubinterface {
	return x.InterfaceSubinterface
}

// Data
func (x *itfce) Update(d *networkv1alpha1.Interface) {
	x.Interface = d
}

func (x *itfce) Print(itfceName string, n int) {
	fmt.Printf("%s Interface: %s Kind: %s LAG: %t, LAG Member: %t\n", strings.Repeat(" ", n), itfceName,  x.Interface.Config.Kind, *x.Interface.Config.Lag, *x.Interface.Config.LagMember)
	n++
	for subItfceName, i := range x.InterfaceSubinterface {
		i.Print(subItfceName, n)
	}
}

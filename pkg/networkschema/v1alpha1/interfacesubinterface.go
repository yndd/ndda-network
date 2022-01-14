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

type InterfaceSubinterface interface {
	// methods children
	// methods data
	Update(x *networkv1alpha1.InterfaceSubinterface)
	AddInterfaceSubinterfaceIpv4(ai *networkv1alpha1.InterfaceSubinterfaceIpv4)
	AddInterfaceSubinterfaceIpv6(ai *networkv1alpha1.InterfaceSubinterfaceIpv6)
	Print(itfceName string, n int)
}

func NewInterfaceSubinterface(p Interface, key string) InterfaceSubinterface {
	return &interfacesubinterface{
		// parent
		parent: p,
		// children
		// data key
		//InterfaceSubinterface: &networkv1alpha1.InterfaceSubinterface{
		//	Name: &name,
		//},
	}
}

type interfacesubinterface struct {
	// parent
	parent Interface
	// children
	// Data
	InterfaceSubinterface *networkv1alpha1.InterfaceSubinterface
}

// children
// Data
func (x *interfacesubinterface) Update(d *networkv1alpha1.InterfaceSubinterface) {
	x.InterfaceSubinterface = d
}

// InterfaceSubinterface ipv4 subinterface Subinterface [subinterface]
func (x *interfacesubinterface) AddInterfaceSubinterfaceIpv4(ai *networkv1alpha1.InterfaceSubinterfaceIpv4) {
	x.InterfaceSubinterface.Ipv4 = append(x.InterfaceSubinterface.Ipv4, ai)
}

// InterfaceSubinterface ipv6 subinterface Subinterface [subinterface]
func (x *interfacesubinterface) AddInterfaceSubinterfaceIpv6(ai *networkv1alpha1.InterfaceSubinterfaceIpv6) {
	x.InterfaceSubinterface.Ipv6 = append(x.InterfaceSubinterface.Ipv6, ai)
}

func (x *interfacesubinterface) Print(siName string, n int) {
	fmt.Printf("%s SubInterface: %s Kind: %s OuterTag: %d\n", strings.Repeat(" ", n), siName, x.InterfaceSubinterface.Config.Kind, *x.InterfaceSubinterface.DeepCopy().Config.OuterVlanId)
	n++
	fmt.Printf("%s Local Addressing Info\n", strings.Repeat(" ", n))
	for _, prefix := range x.InterfaceSubinterface.Ipv4 {
		fmt.Printf("%s IpPrefix: %s\n", strings.Repeat(" ", n), *prefix.IpPrefix)
	}
	for _, prefix := range x.InterfaceSubinterface.Ipv6 {
		fmt.Printf("%s IpPrefix: %s\n", strings.Repeat(" ", n), *prefix.IpPrefix)
	}
}

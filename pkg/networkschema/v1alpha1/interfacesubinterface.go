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

	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/meta"
	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	errCreateInterfaceSubInterface = "cannot create Interface SubInterface"
	errDeleteInterfaceSubInterface = "cannot delete Interface SubInterface"
	errGetInterfaceSubInterface    = "cannot get Interface SubInterface"
)

type InterfaceSubinterface interface {
	// methods children
	// methods data
	Update(x *networkv1alpha1.InterfaceSubinterface)
	AddInterfaceSubinterfaceIpv4(ai *networkv1alpha1.InterfaceSubinterfaceIpv4)
	AddInterfaceSubinterfaceIpv6(ai *networkv1alpha1.InterfaceSubinterfaceIpv6)
	Print(itfceName string, n int)
	ImplementSchema(ctx context.Context, mg resource.Managed, deviceName string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]interface{}) error
}

func NewInterfaceSubinterface(c resource.ClientApplicator, p Interface, key string) InterfaceSubinterface {
	newInterfaceSubinterfaceList := func() networkv1alpha1.IFNetworkInterfaceSubinterfaceList {
		return &networkv1alpha1.NetworkInterfaceSubinterfaceList{}
	}
	return &interfacesubinterface{
		client: c,
		// parent
		parent: p,
		// children
		// data key
		//InterfaceSubinterface: &networkv1alpha1.InterfaceSubinterface{
		//	Name: &name,
		//},
		newInterfaceSubInterfaceList: newInterfaceSubinterfaceList,
	}
}

type interfacesubinterface struct {
	client resource.ClientApplicator
	// parent
	parent Interface
	// children
	// Data
	InterfaceSubinterface *networkv1alpha1.InterfaceSubinterface

	newInterfaceSubInterfaceList func() networkv1alpha1.IFNetworkInterfaceSubinterfaceList
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

func (x *interfacesubinterface) ImplementSchema(ctx context.Context, mg resource.Managed, deviceName string) error {
	o := x.buildNddaNetworkInterfaceSubInterface(mg, deviceName)
	if err := x.client.Apply(ctx, o); err != nil {
		return errors.Wrap(err, errCreateInterfaceSubInterface)
	}
	return nil
}

func (x *interfacesubinterface) buildNddaNetworkInterfaceSubInterface(mg resource.Managed, deviceName string) *networkv1alpha1.NetworkInterfaceSubinterface {
	index := strings.ReplaceAll(*x.InterfaceSubinterface.Index, "/", "-")
	itfceName := strings.ReplaceAll(*x.parent.Get().Name, "/", "-")

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{deviceName, itfceName, index})

	return &networkv1alpha1.NetworkInterfaceSubinterface{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourceName,
			Namespace: mg.GetNamespace(),
			Labels: map[string]string{
				networkv1alpha1.LabelNddaDeploymentPolicy: string(mg.GetDeploymentPolicy()),
				networkv1alpha1.LabelNddaOwner:            odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind)),
				networkv1alpha1.LabelNddaDevice:           deviceName,
				networkv1alpha1.LabelNddaItfce:            itfceName,
			},
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: networkv1alpha1.InterfaceSubinterfaceSpec{
			InterfaceName:         x.parent.Get().Name,
			InterfaceSubinterface: x.InterfaceSubinterface,
		},
	}
}

func (x *interfacesubinterface) InitializeDummySchema() {
}

func (x *interfacesubinterface) ListResources(ctx context.Context, mg resource.Managed, resources map[string]interface{}) error {
	opts := []client.ListOption{
		client.MatchingLabels{networkv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newInterfaceSubInterfaceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetInterfaceSubinterfaces() {
		name := i.GetName()
		kind := strings.ToLower(i.GetObjectKind().GroupVersionKind().Kind)
		resources[strings.Join([]string{name, kind}, "/")] = "dummy"
	}
	return nil
}

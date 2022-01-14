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
)

const (
	errCreateInterface = "cannot create Interface"
	errDeleteInterface = "cannot delete Interface"
	errGetInterface    = "cannot get Interface"
)

type Interface interface {
	// methods children
	NewInterfaceSubinterface(c resource.ClientApplicator, key string) InterfaceSubinterface
	GetInterfaceSubinterfaces() map[string]InterfaceSubinterface
	// methods data
	Update(x *networkv1alpha1.Interface)
	Get() *networkv1alpha1.Interface

	Print(itfceName string, n int)
	ImplementSchema(ctx context.Context, mg resource.Managed, deviceName, deplPolicy string) error
}

func NewInterface(c resource.ClientApplicator, p Device, key string) Interface {
	return &itfce{
		client: c,
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
	client resource.ClientApplicator
	// parent
	parent Device
	// children
	InterfaceSubinterface map[string]InterfaceSubinterface
	// Data
	Interface *networkv1alpha1.Interface
}

// children
func (x *itfce) NewInterfaceSubinterface(c resource.ClientApplicator, key string) InterfaceSubinterface {
	if _, ok := x.InterfaceSubinterface[key]; !ok {
		x.InterfaceSubinterface[key] = NewInterfaceSubinterface(x.client, x, key)
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

func (x *itfce) Get() *networkv1alpha1.Interface {
	return x.Interface
}

func (x *itfce) Print(itfceName string, n int) {
	fmt.Printf("%s Interface: %s Kind: %s LAG: %t, LAG Member: %t\n", strings.Repeat(" ", n), itfceName, x.Interface.Config.Kind, *x.Interface.Config.Lag, *x.Interface.Config.LagMember)
	n++
	for subItfceName, i := range x.InterfaceSubinterface {
		i.Print(subItfceName, n)
	}
}

func (x *itfce) ImplementSchema(ctx context.Context, mg resource.Managed, deviceName, deplPolicy string) error {
	o := x.buildNddaNetworkInterface(mg, deviceName, deplPolicy)
	if err := x.client.Apply(ctx, o); err != nil {
		return errors.Wrap(err, errCreateInterface)
	}

	for _, r := range x.GetInterfaceSubinterfaces() {
		if err := r.ImplementSchema(ctx, mg, deviceName, deplPolicy); err != nil {
			return err
		}
	}

	return nil
}

func (x *itfce) buildNddaNetworkInterface(mg resource.Managed, deviceName, deplPolicy string) *networkv1alpha1.NetworkInterface {
	itfceName := strings.ReplaceAll(*x.Interface.Name, "/", "-")

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{deviceName, itfceName})

	return &networkv1alpha1.NetworkInterface{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourceName,
			Namespace: mg.GetNamespace(),
			Labels: map[string]string{
				networkv1alpha1.LabelNddaDeploymentPolicy: deplPolicy,
				networkv1alpha1.LabelNddaOwner:            odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind)),
				networkv1alpha1.LabelNddaDevice:           deviceName,
				networkv1alpha1.LabelNddaItfce:            itfceName,
			},
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: networkv1alpha1.InterfaceSpec{
			Interface: x.Interface,
		},
	}
}

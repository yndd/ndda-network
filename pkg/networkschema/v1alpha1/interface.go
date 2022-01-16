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

	"github.com/pkg/errors"
	"github.com/yndd/ndd-runtime/pkg/meta"
	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
	GetKey() string

	Print(itfceName string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewInterface(c resource.ClientApplicator, p Device, key string) Interface {
	newInterfaceList := func() networkv1alpha1.IFNetworkInterfaceList { return &networkv1alpha1.NetworkInterfaceList{} }
	return &itfce{
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		InterfaceSubinterface: make(map[string]InterfaceSubinterface),
		// data key
		//Interface: &networkv1alpha1.Interface{
		//	Name: &name,
		//},
		newInterfaceList: newInterfaceList,
	}
}

type itfce struct {
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	InterfaceSubinterface map[string]InterfaceSubinterface
	// Data
	Interface *networkv1alpha1.Interface

	newInterfaceList func() networkv1alpha1.IFNetworkInterfaceList
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

func (x *itfce) GetKey() string {
	return x.key
}

func (x *itfce) Print(itfceName string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.Interface)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s Interface: %s Data: %v\n", strings.Repeat(" ", n), itfceName, x1)
	} else {
		fmt.Printf("%s Interface: %s\n", strings.Repeat(" ", n), itfceName)
	}

	n++
	for subItfceName, i := range x.InterfaceSubinterface {
		i.Print(subItfceName, n)
	}
}

func (x *itfce) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildNddaNetworkInterface(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateInterface)
		}
	}
	for _, r := range x.GetInterfaceSubinterfaces() {
		if err := r.DeploySchema(ctx, mg, deviceName, labels); err != nil {
			return err
		}
	}

	return nil
}

func (x *itfce) buildNddaNetworkInterface(mg resource.Managed, deviceName string, labels map[string]string) *networkv1alpha1.NetworkInterface {
	itfceName := strings.ReplaceAll(*x.Interface.Name, "/", "-")

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{itfceName, deviceName})

	labels[networkv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[networkv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[networkv1alpha1.LabelNddaDevice] = deviceName
	labels[networkv1alpha1.LabelNddaItfce] = itfceName
	return &networkv1alpha1.NetworkInterface{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       mg.GetNamespace(),
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: networkv1alpha1.InterfaceSpec{
			Interface: x.Interface,
		},
	}
}

func (x *itfce) InitializeDummySchema() {
	si := x.NewInterfaceSubinterface(x.client, "dummy")
	si.InitializeDummySchema()

}

func (x *itfce) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	opts := []client.ListOption{
		client.MatchingLabels{networkv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newInterfaceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetInterfaces() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	for _, i := range x.GetInterfaceSubinterfaces() {
		if err := i.ListResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	return nil
}

func (x *itfce) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	if x.Get() != nil {
		itfceName := strings.ReplaceAll(*x.Interface.Name, "/", "-")

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{itfceName, deviceName})

		if r, ok := resources[networkv1alpha1.InterfaceKindKind]; ok {
			delete(r, resourceName)
		}
	}
	for _, i := range x.GetInterfaceSubinterfaces() {
		if err := i.ValidateResources(ctx, mg, deviceName, resources); err != nil {
			return err
		}
	}

	return nil
}

func (x *itfce) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	if res, ok := resources[networkv1alpha1.InterfaceKindKind]; ok {
		for resName := range res {
			o := &networkv1alpha1.NetworkInterface{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resName,
					Namespace: mg.GetNamespace(),
				},
			}
			if err := x.client.Delete(ctx, o); err != nil {
				return err
			}
		}
	}

	for _, i := range x.GetInterfaceSubinterfaces() {
		if err := i.DeleteResources(ctx, mg, resources); err != nil {
			return err
		}
	}
	return nil
}

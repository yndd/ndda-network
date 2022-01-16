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
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
)

const (
	errCreateNetworkInstance = "cannot create NetworkInstance"
	errDeleteNetworkInstance = "cannot delete NetworkInstance"
	errGetNetworkInstance    = "cannot get NetworkInstance"
)

type NetworkInstance interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *networkv1alpha1.NetworkInstance
	Update(x *networkv1alpha1.NetworkInstance)
	AddNetworkInstanceInterface(ai *networkv1alpha1.NetworkInstanceConfigInterface)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewNetworkInstance(c resource.ClientApplicator, p Device, key string) NetworkInstance {
	newNetworkInstanceList := func() networkv1alpha1.IFNetworkNetworkInstanceList {
		return &networkv1alpha1.NetworkNetworkInstanceList{}
	}
	return &networkinstance{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//NetworkInstance: &networkv1alpha1.NetworkInstance{
		//	Name: &name,
		//},
		newNetworkInstanceList: newNetworkInstanceList,
	}
}

type networkinstance struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	// Data
	NetworkInstance        *networkv1alpha1.NetworkInstance
	newNetworkInstanceList func() networkv1alpha1.IFNetworkNetworkInstanceList
}

// key type/method

type NetworkInstanceKey struct {
	Name string
}

func WithNetworkInstanceKey(key *NetworkInstanceKey) string {
	d, err := json.Marshal(key)
	if err != nil {
		return ""
	}
	var x1 interface{}
	json.Unmarshal(d, &x1)

	return getKey(x1)
}

// methods children
// Data methods
func (x *networkinstance) Update(d *networkv1alpha1.NetworkInstance) {
	x.NetworkInstance = d
}

// methods data
func (x *networkinstance) Get() *networkv1alpha1.NetworkInstance {
	return x.NetworkInstance
}

func (x *networkinstance) GetKey() []string {
	return strings.Split(x.key, ".")
}

// NetworkInstance interface network-instance-config NetworkInstance [network-instance config]
func (x *networkinstance) AddNetworkInstanceInterface(ai *networkv1alpha1.NetworkInstanceConfigInterface) {
	x.NetworkInstance.Config.Interface = append(x.NetworkInstance.Config.Interface, ai)
}

// methods schema

func (x *networkinstance) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.NetworkInstance)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s NetworkInstance: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s NetworkInstance: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *networkinstance) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateNetworkInstance)
		}
	}

	return nil
}
func (x *networkinstance) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *networkv1alpha1.NetworkNetworkInstance {
	key0 := strings.ReplaceAll(*x.NetworkInstance.Name, "/", "-")

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			strings.ToLower(key0),
			strings.ToLower(deviceName)})

	labels[networkv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[networkv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[networkv1alpha1.LabelNddaDevice] = deviceName
	//labels[networkv1alpha1.LabelNddaItfce] = itfceName
	return &networkv1alpha1.NetworkNetworkInstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       mg.GetNamespace(),
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: networkv1alpha1.NetworkInstanceSpec{
			NetworkInstance: x.NetworkInstance,
		},
	}
}

func (x *networkinstance) InitializeDummySchema() {
}

func (x *networkinstance) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{networkv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newNetworkInstanceList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetNetworkInstances() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *networkinstance) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {
		key0 := strings.ReplaceAll(*x.NetworkInstance.Name, "/", "-")

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(key0),
				strings.ToLower(deviceName)})

		if r, ok := resources[networkv1alpha1.NetworkInstanceKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *networkinstance) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[networkv1alpha1.NetworkInstanceKindKind]; ok {
		for resName := range res {
			o := &networkv1alpha1.NetworkNetworkInstance{
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

	// children

	return nil
}

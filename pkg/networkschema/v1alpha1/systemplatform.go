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
	errCreateSystemPlatform = "cannot create SystemPlatform"
	errDeleteSystemPlatform = "cannot delete SystemPlatform"
	errGetSystemPlatform    = "cannot get SystemPlatform"
)

type SystemPlatform interface {
	// methods children
	// methods data
	GetKey() []string
	Get() *networkv1alpha1.SystemPlatform
	Update(x *networkv1alpha1.SystemPlatform)
	// methods schema
	Print(key string, n int)
	DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
	ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error
	DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error
}

func NewSystemPlatform(c resource.ClientApplicator, p Device, key string) SystemPlatform {
	newSystemPlatformList := func() networkv1alpha1.IFNetworkSystemPlatformList {
		return &networkv1alpha1.NetworkSystemPlatformList{}
	}
	return &systemplatform{
		// k8s client
		client: c,
		// key
		key: key,
		// parent
		parent: p,
		// children
		// data key
		//SystemPlatform: &networkv1alpha1.SystemPlatform{
		//	Name: &name,
		//},
		newSystemPlatformList: newSystemPlatformList,
	}
}

type systemplatform struct {
	// k8s client
	client resource.ClientApplicator
	// key
	key string
	// parent
	parent Device
	// children
	// Data
	SystemPlatform        *networkv1alpha1.SystemPlatform
	newSystemPlatformList func() networkv1alpha1.IFNetworkSystemPlatformList
}

// key type/method

type SystemPlatformKey struct {
}

func WithSystemPlatformKey(key *SystemPlatformKey) string {
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
func (x *systemplatform) Update(d *networkv1alpha1.SystemPlatform) {
	x.SystemPlatform = d
}

// methods data
func (x *systemplatform) Get() *networkv1alpha1.SystemPlatform {
	return x.SystemPlatform
}

func (x *systemplatform) GetKey() []string {
	return strings.Split(x.key, ".")
}

// methods schema

func (x *systemplatform) Print(key string, n int) {
	if x.Get() != nil {
		d, err := json.Marshal(x.SystemPlatform)
		if err != nil {
			return
		}
		var x1 interface{}
		json.Unmarshal(d, &x1)
		fmt.Printf("%s SystemPlatform: %s Data: %v\n", strings.Repeat(" ", n), key, x1)
	} else {
		fmt.Printf("%s SystemPlatform: %s\n", strings.Repeat(" ", n), key)
	}

	n++
}

func (x *systemplatform) DeploySchema(ctx context.Context, mg resource.Managed, deviceName string, labels map[string]string) error {
	if x.Get() != nil {
		o := x.buildCR(mg, deviceName, labels)
		if err := x.client.Apply(ctx, o); err != nil {
			return errors.Wrap(err, errCreateSystemPlatform)
		}
	}

	return nil
}
func (x *systemplatform) buildCR(mg resource.Managed, deviceName string, labels map[string]string) *networkv1alpha1.NetworkSystemPlatform {

	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{
			strings.ToLower(deviceName)})

	labels[networkv1alpha1.LabelNddaDeploymentPolicy] = string(mg.GetDeploymentPolicy())
	labels[networkv1alpha1.LabelNddaOwner] = odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))
	labels[networkv1alpha1.LabelNddaDevice] = deviceName
	//labels[networkv1alpha1.LabelNddaItfce] = itfceName
	return &networkv1alpha1.NetworkSystemPlatform{
		ObjectMeta: metav1.ObjectMeta{
			Name:            resourceName,
			Namespace:       mg.GetNamespace(),
			Labels:          labels,
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: networkv1alpha1.SystemPlatformSpec{
			SystemPlatform: x.SystemPlatform,
		},
	}
}

func (x *systemplatform) InitializeDummySchema() {
}

func (x *systemplatform) ListResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR list
	opts := []client.ListOption{
		client.MatchingLabels{networkv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newSystemPlatformList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetSystemPlatforms() {
		if _, ok := resources[i.GetObjectKind().GroupVersionKind().Kind]; !ok {
			resources[i.GetObjectKind().GroupVersionKind().Kind] = make(map[string]interface{})
		}
		resources[i.GetObjectKind().GroupVersionKind().Kind][i.GetName()] = "dummy"

	}

	// children
	return nil
}

func (x *systemplatform) ValidateResources(ctx context.Context, mg resource.Managed, deviceName string, resources map[string]map[string]interface{}) error {
	// local CR validation
	if x.Get() != nil {

		resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
			[]string{
				strings.ToLower(deviceName)})

		if r, ok := resources[networkv1alpha1.SystemPlatformKindKind]; ok {
			delete(r, resourceName)
		}
	}

	// children
	return nil
}

func (x *systemplatform) DeleteResources(ctx context.Context, mg resource.Managed, resources map[string]map[string]interface{}) error {
	// local CR deletion
	if res, ok := resources[networkv1alpha1.SystemPlatformKindKind]; ok {
		for resName := range res {
			o := &networkv1alpha1.NetworkSystemPlatform{
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

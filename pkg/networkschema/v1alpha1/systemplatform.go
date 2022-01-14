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
	"strings"

	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type SystemPlatform interface {
	// methods children
	// methods data
	Update(x *networkv1alpha1.SystemPlatform)
	InitializeDummySchema()
	ListResources(ctx context.Context, mg resource.Managed, resources map[string]interface{}) error
}

func NewSystemPlatform(c resource.ClientApplicator, p Device, key string) SystemPlatform {
	newSystemPlatformList := func() networkv1alpha1.IFNetworkSystemPlatformList {
		return &networkv1alpha1.NetworkSystemPlatformList{}
	}
	return &systemplatform{
		client: c,
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
	client resource.ClientApplicator
	// parent
	parent Device
	// children
	// Data
	SystemPlatform *networkv1alpha1.SystemPlatform

	newSystemPlatformList func() networkv1alpha1.IFNetworkSystemPlatformList
}

// children
// Data
func (x *systemplatform) Update(d *networkv1alpha1.SystemPlatform) {
	x.SystemPlatform = d
}

func (x *systemplatform) InitializeDummySchema() {
}

func (x *systemplatform) ListResources(ctx context.Context, mg resource.Managed, resources map[string]interface{}) error {
	opts := []client.ListOption{
		client.MatchingLabels{networkv1alpha1.LabelNddaOwner: odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind))},
	}
	list := x.newSystemPlatformList()
	if err := x.client.List(ctx, list, opts...); err != nil {
		return err
	}

	for _, i := range list.GetSystemPlatforms() {
		name := i.GetName()
		kind := strings.ToLower(i.GetObjectKind().GroupVersionKind().Kind)
		resources[strings.Join([]string{name, kind}, "/")] = "dummy"
	}
	return nil
}

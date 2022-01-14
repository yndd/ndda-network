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
	errCreateNetworkInstance = "cannot create NetworkInstance"
	errDeleteNetworkInstance = "cannot delete NetworkInstance"
	errGetNetworkInstance    = "cannot get NetworkInstance"
)

type NetworkInstance interface {
	// methods children
	// methods data
	Update(x *networkv1alpha1.NetworkInstance)
	AddNetworkInstanceInterface(ai *networkv1alpha1.NetworkInstanceConfigInterface)

	Print(niName string, n int)
	ImplementSchema(ctx context.Context, mg resource.Managed, deviceName string) error
}

func NewNetworkInstance(c resource.ClientApplicator, p Device, key string) NetworkInstance {
	return &networkinstance{
		client: c,
		// parent
		parent: p,
		// children
		// data key
		//NetworkInstance: &networkv1alpha1.NetworkInstance{
		//	Name: &name,
		//},
	}
}

type networkinstance struct {
	client resource.ClientApplicator
	// parent
	parent Device
	// children
	// Data
	NetworkInstance *networkv1alpha1.NetworkInstance
}

// children
// Data
func (x *networkinstance) Update(d *networkv1alpha1.NetworkInstance) {
	x.NetworkInstance = d
}

// NetworkInstance interface network-instance-config NetworkInstance [network-instance config]
func (x *networkinstance) AddNetworkInstanceInterface(ai *networkv1alpha1.NetworkInstanceConfigInterface) {
	x.NetworkInstance.Config.Interface = append(x.NetworkInstance.Config.Interface, ai)
}

func (x *networkinstance) Print(niName string, n int) {
	fmt.Printf("%s Ni Name: %s Kind: %s\n", strings.Repeat(" ", n), niName, *x.NetworkInstance.Name)
	n++
	for _, itfce := range x.NetworkInstance.Config.Interface {
		fmt.Printf("%s %s\n", strings.Repeat(" ", n), *itfce.Name)
	}
}

func (x *networkinstance) ImplementSchema(ctx context.Context, mg resource.Managed, deviceName string) error {
	o := x.buildNddaNetworkInstance(mg, deviceName)
	if err := x.client.Apply(ctx, o); err != nil {
		return errors.Wrap(err, errCreateNetworkInstance)
	}
	return nil
}

func (x *networkinstance) buildNddaNetworkInstance(mg resource.Managed, deviceName string) *networkv1alpha1.NetworkNetworkInstance {
	resourceName := odns.GetOdnsResourceName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind),
		[]string{deviceName})

	return &networkv1alpha1.NetworkNetworkInstance{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourceName,
			Namespace: mg.GetNamespace(),
			Labels: map[string]string{
				networkv1alpha1.LabelNddaDeploymentPolicy: string(mg.GetDeploymentPolicy()),
				networkv1alpha1.LabelNddaOwner:            odns.GetOdnsResourceKindName(mg.GetName(), strings.ToLower(mg.GetObjectKind().GroupVersionKind().Kind)),
				networkv1alpha1.LabelNddaDevice:           deviceName,
			},
			OwnerReferences: []metav1.OwnerReference{meta.AsController(meta.TypedReferenceTo(mg, mg.GetObjectKind().GroupVersionKind()))},
		},
		Spec: networkv1alpha1.NetworkInstanceSpec{
			NetworkInstance: x.NetworkInstance,
		},
	}
}

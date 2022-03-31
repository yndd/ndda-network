package abstraction

import (
	"context"
	"fmt"
	"sync"

	"github.com/openconfig/ygot/ygot"
	"github.com/yndd/ndda-network/pkg/nodeitfceselector"
	nddov1 "github.com/yndd/nddo-runtime/apis/common/v1"
	"github.com/yndd/nddo-runtime/pkg/resource"
)

type Interface struct{
	name string
}

func Name(s string) Option {
	return func(x Object) {
		x.Name(s)
	}
}

func (x *Interface) IsAbstracted() bool

func (x *Interface) WithName(s string) {
	x.name= s
}

type Object interface{
	IsAbstracted() bool
	Name(s string)
}

type Option func(Object)

type Abstractor interface {
	Abstract(ygot.ValidatedGoStruct, ...Option)
	//GetInterfaceName(itfcName string) (string, error)
	GetSelectedNodeItfces(ctx context.Context, mg resource.Managed, epgSelectors []*nddov1.EpgInfo, nodeItfceSelectors map[string]*nddov1.ItfceInfo) (*nodeitfceselector.SelectedNodes, error)
}

func New(c resource.ClientApplicator, name string) *Compositeabstraction {
	return &Compositeabstraction{
		name: name,
		// k8s client
		client:       c,
		abstractions: make(map[string]Abstractor),
	}
}

type Compositeabstraction struct {
	name         string
	// k8s client
	client       resource.ClientApplicator
	m            sync.Mutex
	abstractions map[string]Abstractor
}

func (x *Compositeabstraction) AddChild(name string, a Abstractor) {
	x.m.Lock()
	defer x.m.Unlock()
	if _, ok := x.abstractions[name]; !ok {
		x.abstractions[name] = a
	}
}

func (x *Compositeabstraction) GetChild(name string) (Abstractor, error) {
	x.m.Lock()
	defer x.m.Unlock()
	a, ok := x.abstractions[name]
	if !ok {
		return nil, fmt.Errorf("abstraction does not exist: name %s", name)
	}
	return a, nil
}
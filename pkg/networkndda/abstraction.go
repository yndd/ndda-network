package networkndda

import (
	"context"
	"fmt"
	"sync"

	"github.com/yndd/ndda-network/pkg/nodeitfceselector"
	nddov1 "github.com/yndd/nddo-runtime/apis/common/v1"
	"github.com/yndd/nddo-runtime/pkg/resource"
)

type Abstraction interface {
	GetInterfaceName(itfcName string) (string, error)
	GetSelectedNodeItfces(ctx context.Context, mg resource.Managed, epgSelectors []*nddov1.EpgInfo, nodeItfceSelectors map[string]*nddov1.ItfceInfo) (*nodeitfceselector.SelectedNodes, error)
}

func New(c resource.ClientApplicator, name string) *Compositeabstraction {
	return &Compositeabstraction{
		name: name,
		// k8s client
		client:       c,
		abstractions: make(map[string]Abstraction),
	}
}

type Compositeabstraction struct {
	name         string
	// k8s client
	client       resource.ClientApplicator
	m            sync.Mutex
	abstractions map[string]Abstraction
}

func (x *Compositeabstraction) AddChild(name string, a Abstraction) {
	x.m.Lock()
	defer x.m.Unlock()
	if _, ok := x.abstractions[name]; !ok {
		x.abstractions[name] = a
	}
}

func (x *Compositeabstraction) GetChild(name string) (Abstraction, error) {
	x.m.Lock()
	defer x.m.Unlock()
	a, ok := x.abstractions[name]
	if !ok {
		return nil, fmt.Errorf("abstraction does not exist: name %s", name)
	}
	return a, nil
}

package ndda

import (
	"context"

	"github.com/yndd/ndd-runtime/pkg/logging"
	nddav1alpha1 "github.com/yndd/ndda-network/apis/ndda/v1alpha1"
	"github.com/yndd/nddo-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func New(opts ...Option) Handler {
	s := &handler{
		newNddaItfceList:        func() nddav1alpha1.IFNddaInterfaceList { return &nddav1alpha1.NddaInterfaceList{} },
		newNddaSubInterfaceList: func() nddav1alpha1.IFNddaInterfaceSubinterfaceList { return &nddav1alpha1.NddaInterfaceSubinterfaceList{} },
		newNddaNiList:           func() nddav1alpha1.IFNddaNetworkInstanceList { return &nddav1alpha1.NddaNetworkInstanceList{} },
		ctx:                     context.Background(),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (r *handler) WithLogger(log logging.Logger) {
	r.log = log
}

func (r *handler) WithClient(c client.Client) {
	r.client = resource.ClientApplicator{
		Client:     c,
		Applicator: resource.NewAPIPatchingApplicator(c),
	}
}

type handler struct {
	log logging.Logger
	// kubernetes
	client client.Client
	ctx    context.Context

	newNddaItfceList        func() nddav1alpha1.IFNddaInterfaceList
	newNddaSubInterfaceList func() nddav1alpha1.IFNddaInterfaceSubinterfaceList
	newNddaNiList           func() nddav1alpha1.IFNddaNetworkInstanceList
}

package ndda

import (
	"context"

	"github.com/yndd/ndd-runtime/pkg/logging"
	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
	"github.com/yndd/nddo-runtime/pkg/resource"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func New(opts ...Option) Handler {
	s := &handler{
		newNetworkItfceList: func() networkv1alpha1.IFNetworkInterfaceList { return &networkv1alpha1.NetworkInterfaceList{} },
		newNetworkSubInterfaceList: func() networkv1alpha1.IFNetworkInterfaceSubinterfaceList {
			return &networkv1alpha1.NetworkInterfaceSubinterfaceList{}
		},
		newNetworkNiList: func() networkv1alpha1.IFNetworkNetworkInstanceList {
			return &networkv1alpha1.NetworkNetworkInstanceList{}
		},
		ctx: context.Background(),
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

	newNetworkItfceList        func() networkv1alpha1.IFNetworkInterfaceList
	newNetworkSubInterfaceList func() networkv1alpha1.IFNetworkInterfaceSubinterfaceList
	newNetworkNiList           func() networkv1alpha1.IFNetworkNetworkInstanceList
}

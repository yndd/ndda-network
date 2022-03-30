package abstract

import (
	"context"

	"github.com/yndd/ndda-network/pkg/nodeitfceselector"
	nddov1 "github.com/yndd/nddo-runtime/apis/common/v1"
	"github.com/yndd/nddo-runtime/pkg/resource"
)

type Abstract interface {
	GetInterfaceName(itfcName string) (string, error)
	GetSelectedNodeItfces(ctx context.Context, mg resource.Managed, epgSelectors []*nddov1.EpgInfo, nodeItfceSelectors map[string]*nddov1.ItfceInfo) (*nodeitfceselector.SelectedNodes, error)
}

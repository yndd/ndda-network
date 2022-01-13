package ndda

import (
	networkv1alpha1 "github.com/yndd/ndda-network/apis/network/v1alpha1"
	"github.com/yndd/ndda-network/pkg/ndda/itfceinfo"
	networkschema "github.com/yndd/ndda-network/pkg/networkschema/v1alpha1"
	nddov1 "github.com/yndd/nddo-runtime/apis/common/v1"
	"github.com/yndd/nddo-runtime/pkg/odns"
	"github.com/yndd/nddo-runtime/pkg/resource"
)

func (r *handler) GetSelectedNodeItfces(mg resource.Managed, epgSelectors []*nddov1.EpgInfo, nodeItfceSelectors map[string]*nddov1.ItfceInfo) (map[string][]itfceinfo.ItfceInfo, error) {
	// get all ndda interfaces within the oda scope
	// oda is organization, deployement, availability zone
	opts := odns.GetClientListOptionFromResourceName(mg.GetName())
	nddaItfces := r.newNetworkItfceList()
	if err := r.client.List(r.ctx, nddaItfces, opts...); err != nil {
		return nil, err
	}

	sel := NewNodeItfceSelection()
	sel.GetNodeItfcesByEpgSelector(epgSelectors, nddaItfces)
	sel.GetNodeItfcesByNodeItfceSelector(nodeItfceSelectors, nddaItfces)
	return sel.GetSelectedNodeItfces(), nil

}

func (r *handler) GetSelectedNodeItfcesIrb(mg resource.Managed, s networkschema.Schema, niName string) (map[string][]itfceinfo.ItfceInfo, error) {
	// get all ndda interfaces within the oda scope
	// oda is organization, deployement, availability zone
	opts := odns.GetClientListOptionFromResourceName(mg.GetName())
	nddaItfces := r.newNetworkItfceList()
	if err := r.client.List(r.ctx, nddaItfces, opts...); err != nil {
		return nil, err
	}

	sel := NewNodeItfceSelection()
	sel.GetIrbNodeItfces(niName, s, nddaItfces)
	return sel.GetSelectedNodeItfces(), nil
}

func (r *handler) GetSelectedNodeItfcesVxlan(mg resource.Managed, s networkschema.Schema, niName string) (map[string][]itfceinfo.ItfceInfo, error) {
	// get all ndda interfaces within the oda scope
	// oda is organization, deployement, availability zone
	opts := odns.GetClientListOptionFromResourceName(mg.GetName())
	nddaItfces := r.newNetworkItfceList()
	if err := r.client.List(r.ctx, nddaItfces, opts...); err != nil {
		return nil, err
	}

	sel := NewNodeItfceSelection()
	sel.GetVxlanNodeItfces(niName, s, nddaItfces)
	return sel.GetSelectedNodeItfces(), nil
}

type NodeItfceSelection interface {
	GetSelectedNodeItfces() map[string][]itfceinfo.ItfceInfo
	GetNodeItfcesByEpgSelector([]*nddov1.EpgInfo, networkv1alpha1.IFNetworkInterfaceList)
	GetNodeItfcesByNodeItfceSelector(map[string]*nddov1.ItfceInfo, networkv1alpha1.IFNetworkInterfaceList)
	GetVxlanNodeItfces(string, networkschema.Schema, networkv1alpha1.IFNetworkInterfaceList)
	GetIrbNodeItfces(string, networkschema.Schema, networkv1alpha1.IFNetworkInterfaceList)
}

func NewNodeItfceSelection() NodeItfceSelection {
	return &selectedNodeItfces{
		nodes: make(map[string][]itfceinfo.ItfceInfo),
	}
}

type selectedNodeItfces struct {
	nodes map[string][]itfceinfo.ItfceInfo
}

func (x *selectedNodeItfces) GetSelectedNodeItfces() map[string][]itfceinfo.ItfceInfo {
	return x.nodes
}

func (x *selectedNodeItfces) GetNodeItfcesByEpgSelector(epgSelectors []*nddov1.EpgInfo, nddaItfceList networkv1alpha1.IFNetworkInterfaceList) {
	for _, nddaItfce := range nddaItfceList.GetInterfaces() {
		//fmt.Printf("getNodeItfcesByEpgSelector: epg: %s, itfceepg: %s, nodename: %s, itfcename: %s, lagmember: %t\n", epg, nddaItfce.GetEndpointGroup(), nddaItfce.GetNodeName(), nddaItfce.GetInterfaceName(), nddaItfce.GetLagMember())
		// TODO add specifc endpoint group selector
		for _, epgSelector := range epgSelectors {
			if epgSelector.EpgName != "" && nddaItfce.GetEndpointGroup() == epgSelector.EpgName {
				//fmt.Printf("getNodeItfcesByEpgSelector: %s\n", nddaItfce.GetName())
				// avoid selecting lag members
				if !nddaItfce.GetInterfaceConfigLagMember() {
					x.addNodeItfce(nddaItfce.GetDeviceName(), nddaItfce.GetInterfaceName(), itfceinfo.NewItfceInfo(
						itfceinfo.WithInnerVlanId(epgSelector.InnerVlanId),
						itfceinfo.WithOuterVlanId(epgSelector.OuterVlanId),
						itfceinfo.WithItfceKind(networkv1alpha1.E_InterfaceKind_INTERFACE),
						itfceinfo.WithIpv4Prefixes(epgSelector.Ipv4Prefixes),
						itfceinfo.WithIpv6Prefixes(epgSelector.Ipv6Prefixes),
					))
				}
			}
		}
	}
}

func (x *selectedNodeItfces) GetNodeItfcesByNodeItfceSelector(nodeItfceSelectors map[string]*nddov1.ItfceInfo, nddaItfceList networkv1alpha1.IFNetworkInterfaceList) {
	for _, nddaItfce := range nddaItfceList.GetInterfaces() {
		for deviceName, itfceInfo := range nodeItfceSelectors {
			//fmt.Printf("getNodeItfcesByNodeItfceSelector: nodename: %s, itfcename: %s, lagmember: %t, nodename: %s\n", nddaItfce.GetNodeName(), nddaItfce.GetInterfaceName(), nddaItfce.GetLagMember(), nodeName)
			// avoid selecting lag members
			if !nddaItfce.GetInterfaceConfigLagMember() && nddaItfce.GetDeviceName() == deviceName && nddaItfce.GetInterfaceName() == itfceInfo.ItfceName {
				//fmt.Printf("getNodeItfcesByNodeItfceSelector: nodename: %s, itfcename: %s, lagmember: %t, nodename: %s\n", nddaItfce.GetNodeName(), nddaItfce.GetInterfaceName(), nddaItfce.GetLagMember(), nodeName)
				x.addNodeItfce(nddaItfce.GetDeviceName(), nddaItfce.GetInterfaceName(), itfceinfo.NewItfceInfo(
					itfceinfo.WithInnerVlanId(itfceInfo.InnerVlanId),
					itfceinfo.WithOuterVlanId(itfceInfo.OuterVlanId),
					itfceinfo.WithItfceKind(networkv1alpha1.E_InterfaceKind_INTERFACE),
					itfceinfo.WithIpv4Prefixes(itfceInfo.Ipv4Prefixes),
					itfceinfo.WithIpv6Prefixes(itfceInfo.Ipv6Prefixes),
				))
			}
		}
	}
}

func (x *selectedNodeItfces) GetVxlanNodeItfces(niName string, s networkschema.Schema, nddaItfceList networkv1alpha1.IFNetworkInterfaceList) {
	for _, nddaItfce := range nddaItfceList.GetInterfaces() {
		for deviceName, d := range s.GetDevices() {
			for dniName := range d.GetNetworkInstances() {
				if dniName == niName {
					if nddaItfce.GetDeviceName() == deviceName && nddaItfce.GetInterfaceConfigKind() == networkv1alpha1.E_InterfaceKind_VXLAN {
						x.addNodeItfce(deviceName, nddaItfce.GetInterfaceName(), itfceinfo.NewItfceInfo(
							itfceinfo.WithItfceKind(networkv1alpha1.E_InterfaceKind_VXLAN),
							//WithItfceIndex(ni.GetIndex()), // we use the vxlan
							//WithIpv4Prefixes(make([]*string, 0)),
							//WithIpv6Prefixes(make([]*string, 0)),
						))
					}
				}
			}
		}
	}
}

func (x *selectedNodeItfces) GetIrbNodeItfces(niName string, s networkschema.Schema, nddaItfceList networkv1alpha1.IFNetworkInterfaceList) {
	for _, nddaItfce := range nddaItfceList.GetInterfaces() {
		for deviceName, d := range s.GetDevices() {
			for dniName := range d.GetNetworkInstances() {
				if dniName == niName {
					// we only select the irb interfaces to retain the index
					if nddaItfce.GetDeviceName() == deviceName && nddaItfce.GetInterfaceConfigKind() == networkv1alpha1.E_InterfaceKind_IRB {
						x.addNodeItfce(deviceName, nddaItfce.GetInterfaceName(), itfceinfo.NewItfceInfo(
							itfceinfo.WithItfceKind(networkv1alpha1.E_InterfaceKind_IRB),
							//WithItfceIndex(9999), // dummy
							//WithIpv4Prefixes(ipv4Prefixes),
							//WithIpv6Prefixes(ipv6Prefixes),
						))
					}
				}
			}
		}
	}
}

func (x *selectedNodeItfces) addNodeItfce(nodeName, intName string, ifInfo itfceinfo.ItfceInfo) {
	// check if node exists, if not initialize the node
	if _, ok := x.nodes[nodeName]; !ok {
		x.nodes[nodeName] = make([]itfceinfo.ItfceInfo, 0)
	}
	// check if the interfacename was already present on the node
	// if not add it to the list
	for _, itfceInfo := range x.nodes[nodeName] {
		if itfceInfo.GetItfceName() == intName {
			return
		}
	}
	ifInfo.SetItfceName(intName)
	x.nodes[nodeName] = append(x.nodes[nodeName], ifInfo)
}

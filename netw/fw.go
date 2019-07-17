package netw


import (
    "github.com/inwinstack/pango/netw/ikegw"
    "github.com/inwinstack/pango/netw/interface/eth"
    "github.com/inwinstack/pango/netw/interface/loopback"
    "github.com/inwinstack/pango/netw/interface/subinterface/layer2"
    "github.com/inwinstack/pango/netw/interface/subinterface/layer3"
    "github.com/inwinstack/pango/netw/interface/tunnel"
    vli "github.com/inwinstack/pango/netw/interface/vlan"
    "github.com/inwinstack/pango/netw/ipsectunnel"
    tpiv4 "github.com/inwinstack/pango/netw/ipsectunnel/proxyid/ipv4"
    "github.com/inwinstack/pango/netw/profile/bfd"
    "github.com/inwinstack/pango/netw/profile/ike"
    "github.com/inwinstack/pango/netw/profile/ipsec"
    "github.com/inwinstack/pango/netw/profile/mngtprof"
    "github.com/inwinstack/pango/netw/profile/monitor"
    redist4 "github.com/inwinstack/pango/netw/routing/profile/redist/ipv4"
    "github.com/inwinstack/pango/netw/routing/protocol/bgp"
    "github.com/inwinstack/pango/netw/routing/protocol/bgp/aggregate"
    agaf "github.com/inwinstack/pango/netw/routing/protocol/bgp/aggregate/filter/advertise"
    "github.com/inwinstack/pango/netw/routing/protocol/bgp/aggregate/filter/suppress"
    "github.com/inwinstack/pango/netw/routing/protocol/bgp/conadv"
    "github.com/inwinstack/pango/netw/routing/protocol/bgp/conadv/filter/advertise"
    "github.com/inwinstack/pango/netw/routing/protocol/bgp/conadv/filter/nonexist"
    "github.com/inwinstack/pango/netw/routing/protocol/bgp/exp"
    "github.com/inwinstack/pango/netw/routing/protocol/bgp/imp"
    "github.com/inwinstack/pango/netw/routing/protocol/bgp/peer"
    "github.com/inwinstack/pango/netw/routing/protocol/bgp/peer/group"
    "github.com/inwinstack/pango/netw/routing/protocol/bgp/profile/auth"
    "github.com/inwinstack/pango/netw/routing/protocol/bgp/profile/dampening"
    bgpredist "github.com/inwinstack/pango/netw/routing/protocol/bgp/redist"
    "github.com/inwinstack/pango/netw/routing/router"
    "github.com/inwinstack/pango/netw/routing/route/static/ipv4"
    "github.com/inwinstack/pango/netw/tunnel/gre"
    "github.com/inwinstack/pango/netw/vlan"
    "github.com/inwinstack/pango/netw/zone"
    "github.com/inwinstack/pango/util"
)


// Netw is the client.Network namespace.
type FwNetw struct {
    BfdProfile *bfd.FwBfd
    BgpAggregate *aggregate.FwAggregate
    BgpAggAdvertiseFilter *agaf.FwAdvertise
    BgpAggSuppressFilter *suppress.FwSuppress
    BgpAuthProfile *auth.FwAuth
    BgpConAdvAdvertiseFilter *advertise.FwAdvertise
    BgpConAdvNonExistFilter *nonexist.FwNonExist
    BgpConditionalAdv *conadv.FwConAdv
    BgpConfig *bgp.FwBgp
    BgpDampeningProfile *dampening.FwDampening
    BgpExport *exp.FwExp
    BgpImport *imp.FwImp
    BgpPeer *peer.FwPeer
    BgpPeerGroup *group.FwGroup
    BgpRedistRule *bgpredist.FwRedist
    EthernetInterface *eth.FwEth
    GreTunnel *gre.FwGre
    IkeCryptoProfile *ike.FwIke
    IkeGateway *ikegw.FwIkeGw
    IpsecCryptoProfile *ipsec.FwIpsec
    IpsecTunnel *ipsectunnel.FwIpsecTunnel
    IpsecTunnelProxyId *tpiv4.FwIpv4
    Layer2Subinterface *layer2.FwLayer2
    Layer3Subinterface *layer3.FwLayer3
    LoopbackInterface *loopback.FwLoopback
    ManagementProfile *mngtprof.FwMngtProf
    MonitorProfile *monitor.FwMonitor
    RedistributionProfile *redist4.FwIpv4
    StaticRoute *ipv4.FwIpv4
    TunnelInterface *tunnel.FwTunnel
    VirtualRouter *router.FwRouter
    Vlan *vlan.FwVlan
    VlanInterface *vli.FwVlan
    Zone *zone.FwZone
}

// Initialize is invoked on client.Initialize().
func (c *FwNetw) Initialize(i util.XapiClient) {
    c.BfdProfile = &bfd.FwBfd{}
    c.BfdProfile.Initialize(i)

    c.BgpAggregate = &aggregate.FwAggregate{}
    c.BgpAggregate.Initialize(i)

    c.BgpAggAdvertiseFilter = &agaf.FwAdvertise{}
    c.BgpAggAdvertiseFilter.Initialize(i)

    c.BgpAggSuppressFilter = &suppress.FwSuppress{}
    c.BgpAggSuppressFilter.Initialize(i)

    c.BgpAuthProfile = &auth.FwAuth{}
    c.BgpAuthProfile.Initialize(i)

    c.BgpConAdvAdvertiseFilter = &advertise.FwAdvertise{}
    c.BgpConAdvAdvertiseFilter.Initialize(i)

    c.BgpConAdvNonExistFilter = &nonexist.FwNonExist{}
    c.BgpConAdvNonExistFilter.Initialize(i)

    c.BgpConditionalAdv = &conadv.FwConAdv{}
    c.BgpConditionalAdv.Initialize(i)

    c.BgpConfig = &bgp.FwBgp{}
    c.BgpConfig.Initialize(i)

    c.BgpDampeningProfile = &dampening.FwDampening{}
    c.BgpDampeningProfile.Initialize(i)

    c.BgpExport = &exp.FwExp{}
    c.BgpExport.Initialize(i)

    c.BgpImport = &imp.FwImp{}
    c.BgpImport.Initialize(i)

    c.BgpPeer = &peer.FwPeer{}
    c.BgpPeer.Initialize(i)

    c.BgpPeerGroup = &group.FwGroup{}
    c.BgpPeerGroup.Initialize(i)

    c.BgpRedistRule = &bgpredist.FwRedist{}
    c.BgpRedistRule.Initialize(i)

    c.EthernetInterface = &eth.FwEth{}
    c.EthernetInterface.Initialize(i)

    c.GreTunnel = &gre.FwGre{}
    c.GreTunnel.Initialize(i)

    c.IkeCryptoProfile = &ike.FwIke{}
    c.IkeCryptoProfile.Initialize(i)

    c.IkeGateway = &ikegw.FwIkeGw{}
    c.IkeGateway.Initialize(i)

    c.IpsecCryptoProfile = &ipsec.FwIpsec{}
    c.IpsecCryptoProfile.Initialize(i)

    c.IpsecTunnel = &ipsectunnel.FwIpsecTunnel{}
    c.IpsecTunnel.Initialize(i)

    c.IpsecTunnelProxyId = &tpiv4.FwIpv4{}
    c.IpsecTunnelProxyId.Initialize(i)

    c.Layer2Subinterface = &layer2.FwLayer2{}
    c.Layer2Subinterface.Initialize(i)

    c.Layer3Subinterface = &layer3.FwLayer3{}
    c.Layer3Subinterface.Initialize(i)

    c.LoopbackInterface = &loopback.FwLoopback{}
    c.LoopbackInterface.Initialize(i)

    c.ManagementProfile = &mngtprof.FwMngtProf{}
    c.ManagementProfile.Initialize(i)

    c.MonitorProfile = &monitor.FwMonitor{}
    c.MonitorProfile.Initialize(i)

    c.RedistributionProfile = &redist4.FwIpv4{}
    c.RedistributionProfile.Initialize(i)

    c.StaticRoute = &ipv4.FwIpv4{}
    c.StaticRoute.Initialize(i)

    c.TunnelInterface = &tunnel.FwTunnel{}
    c.TunnelInterface.Initialize(i)

    c.VirtualRouter = &router.FwRouter{}
    c.VirtualRouter.Initialize(i)

    c.Vlan = &vlan.FwVlan{}
    c.Vlan.Initialize(i)

    c.VlanInterface = &vli.FwVlan{}
    c.VlanInterface.Initialize(i)

    c.Zone = &zone.FwZone{}
    c.Zone.Initialize(i)
}

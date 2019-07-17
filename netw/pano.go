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


// PanoNetw is the client.Network namespace.
type PanoNetw struct {
    BfdProfile *bfd.PanoBfd
    BgpAggregate *aggregate.PanoAggregate
    BgpAggAdvertiseFilter *agaf.PanoAdvertise
    BgpAggSuppressFilter *suppress.PanoSuppress
    BgpAuthProfile *auth.PanoAuth
    BgpConAdvAdvertiseFilter *advertise.PanoAdvertise
    BgpConAdvNonExistFilter *nonexist.PanoNonExist
    BgpConditionalAdv *conadv.PanoConAdv
    BgpConfig *bgp.PanoBgp
    BgpDampeningProfile *dampening.PanoDampening
    BgpExport *exp.PanoExp
    BgpImport *imp.PanoImp
    BgpPeer *peer.PanoPeer
    BgpPeerGroup *group.PanoGroup
    BgpRedistRule *bgpredist.PanoRedist
    EthernetInterface *eth.PanoEth
    GreTunnel *gre.PanoGre
    IkeCryptoProfile *ike.PanoIke
    IkeGateway *ikegw.PanoIkeGw
    IpsecCryptoProfile *ipsec.PanoIpsec
    IpsecTunnel *ipsectunnel.PanoIpsecTunnel
    IpsecTunnelProxyId *tpiv4.PanoIpv4
    Layer2Subinterface *layer2.PanoLayer2
    Layer3Subinterface *layer3.PanoLayer3
    LoopbackInterface *loopback.PanoLoopback
    ManagementProfile *mngtprof.PanoMngtProf
    MonitorProfile *monitor.PanoMonitor
    RedistributionProfile *redist4.PanoIpv4
    StaticRoute *ipv4.PanoIpv4
    TunnelInterface *tunnel.PanoTunnel
    VirtualRouter *router.PanoRouter
    Vlan *vlan.PanoVlan
    VlanInterface *vli.PanoVlan
    Zone *zone.PanoZone
}

// Initialize is invoked on client.Initialize().
func (c *PanoNetw) Initialize(i util.XapiClient) {
    c.BfdProfile = &bfd.PanoBfd{}
    c.BfdProfile.Initialize(i)

    c.BgpAggregate = &aggregate.PanoAggregate{}
    c.BgpAggregate.Initialize(i)

    c.BgpAggAdvertiseFilter = &agaf.PanoAdvertise{}
    c.BgpAggAdvertiseFilter.Initialize(i)

    c.BgpAggSuppressFilter = &suppress.PanoSuppress{}
    c.BgpAggSuppressFilter.Initialize(i)

    c.BgpAuthProfile = &auth.PanoAuth{}
    c.BgpAuthProfile.Initialize(i)

    c.BgpConAdvAdvertiseFilter = &advertise.PanoAdvertise{}
    c.BgpConAdvAdvertiseFilter.Initialize(i)

    c.BgpConAdvNonExistFilter = &nonexist.PanoNonExist{}
    c.BgpConAdvNonExistFilter.Initialize(i)

    c.BgpConditionalAdv = &conadv.PanoConAdv{}
    c.BgpConditionalAdv.Initialize(i)

    c.BgpConfig = &bgp.PanoBgp{}
    c.BgpConfig.Initialize(i)

    c.BgpDampeningProfile = &dampening.PanoDampening{}
    c.BgpDampeningProfile.Initialize(i)

    c.BgpExport = &exp.PanoExp{}
    c.BgpExport.Initialize(i)

    c.BgpImport = &imp.PanoImp{}
    c.BgpImport.Initialize(i)

    c.BgpPeer = &peer.PanoPeer{}
    c.BgpPeer.Initialize(i)

    c.BgpPeerGroup = &group.PanoGroup{}
    c.BgpPeerGroup.Initialize(i)

    c.BgpRedistRule = &bgpredist.PanoRedist{}
    c.BgpRedistRule.Initialize(i)

    c.EthernetInterface = &eth.PanoEth{}
    c.EthernetInterface.Initialize(i)

    c.GreTunnel = &gre.PanoGre{}
    c.GreTunnel.Initialize(i)

    c.IkeCryptoProfile = &ike.PanoIke{}
    c.IkeCryptoProfile.Initialize(i)

    c.IkeGateway = &ikegw.PanoIkeGw{}
    c.IkeGateway.Initialize(i)

    c.IpsecCryptoProfile = &ipsec.PanoIpsec{}
    c.IpsecCryptoProfile.Initialize(i)

    c.IpsecTunnel = &ipsectunnel.PanoIpsecTunnel{}
    c.IpsecTunnel.Initialize(i)

    c.IpsecTunnelProxyId = &tpiv4.PanoIpv4{}
    c.IpsecTunnelProxyId.Initialize(i)

    c.Layer2Subinterface = &layer2.PanoLayer2{}
    c.Layer2Subinterface.Initialize(i)

    c.Layer3Subinterface = &layer3.PanoLayer3{}
    c.Layer3Subinterface.Initialize(i)

    c.LoopbackInterface = &loopback.PanoLoopback{}
    c.LoopbackInterface.Initialize(i)

    c.ManagementProfile = &mngtprof.PanoMngtProf{}
    c.ManagementProfile.Initialize(i)

    c.MonitorProfile = &monitor.PanoMonitor{}
    c.MonitorProfile.Initialize(i)

    c.RedistributionProfile = &redist4.PanoIpv4{}
    c.RedistributionProfile.Initialize(i)

    c.StaticRoute = &ipv4.PanoIpv4{}
    c.StaticRoute.Initialize(i)

    c.TunnelInterface = &tunnel.PanoTunnel{}
    c.TunnelInterface.Initialize(i)

    c.VirtualRouter = &router.PanoRouter{}
    c.VirtualRouter.Initialize(i)

    c.Vlan = &vlan.PanoVlan{}
    c.Vlan.Initialize(i)

    c.VlanInterface = &vli.PanoVlan{}
    c.VlanInterface.Initialize(i)

    c.Zone = &zone.PanoZone{}
    c.Zone.Initialize(i)
}

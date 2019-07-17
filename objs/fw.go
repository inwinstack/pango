package objs


import (
    "github.com/inwinstack/pango/util"

    "github.com/inwinstack/pango/objs/addr"
    "github.com/inwinstack/pango/objs/addrgrp"
    "github.com/inwinstack/pango/objs/app"
    appgrp "github.com/inwinstack/pango/objs/app/group"
    "github.com/inwinstack/pango/objs/app/signature"
    "github.com/inwinstack/pango/objs/app/signature/andcond"
    "github.com/inwinstack/pango/objs/app/signature/orcond"
    "github.com/inwinstack/pango/objs/edl"
    "github.com/inwinstack/pango/objs/profile/logfwd"
    "github.com/inwinstack/pango/objs/profile/logfwd/matchlist"
    "github.com/inwinstack/pango/objs/profile/logfwd/matchlist/action"
    "github.com/inwinstack/pango/objs/srvc"
    "github.com/inwinstack/pango/objs/srvcgrp"
    "github.com/inwinstack/pango/objs/tags"
)


// FwObjs is the client.Objects namespace.
type FwObjs struct {
    Address *addr.FwAddr
    AddressGroup *addrgrp.FwAddrGrp
    Application *app.FwApp
    AppGroup *appgrp.FwGroup
    AppSignature *signature.FwSignature
    AppSigAndCond *andcond.FwAndCond
    AppSigAndCondOrCond *orcond.FwOrCond
    Edl *edl.FwEdl
    LogForwardingProfile *logfwd.FwLogFwd
    LogForwardingProfileMatchList *matchlist.FwMatchList
    LogForwardingProfileMatchListAction *action.FwAction
    Services *srvc.FwSrvc
    ServiceGroup *srvcgrp.FwSrvcGrp
    Tags *tags.FwTags
}

// Initialize is invoked on client.Initialize().
func (c *FwObjs) Initialize(i util.XapiClient) {
    c.Address = &addr.FwAddr{}
    c.Address.Initialize(i)

    c.AddressGroup = &addrgrp.FwAddrGrp{}
    c.AddressGroup.Initialize(i)

    c.Application = &app.FwApp{}
    c.Application.Initialize(i)

    c.AppGroup = &appgrp.FwGroup{}
    c.AppGroup.Initialize(i)

    c.AppSignature = &signature.FwSignature{}
    c.AppSignature.Initialize(i)

    c.AppSigAndCond = &andcond.FwAndCond{}
    c.AppSigAndCond.Initialize(i)

    c.AppSigAndCondOrCond = &orcond.FwOrCond{}
    c.AppSigAndCondOrCond.Initialize(i)

    c.Edl = &edl.FwEdl{}
    c.Edl.Initialize(i)

    c.LogForwardingProfile = &logfwd.FwLogFwd{}
    c.LogForwardingProfile.Initialize(i)

    c.LogForwardingProfileMatchList = &matchlist.FwMatchList{}
    c.LogForwardingProfileMatchList.Initialize(i)

    c.LogForwardingProfileMatchListAction = &action.FwAction{}
    c.LogForwardingProfileMatchListAction.Initialize(i)

    c.Services = &srvc.FwSrvc{}
    c.Services.Initialize(i)

    c.ServiceGroup = &srvcgrp.FwSrvcGrp{}
    c.ServiceGroup.Initialize(i)

    c.Tags = &tags.FwTags{}
    c.Tags.Initialize(i)
}

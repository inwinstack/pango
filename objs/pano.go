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


// PanoObjs is the client.Objects namespace.
type PanoObjs struct {
    Address *addr.PanoAddr
    AddressGroup *addrgrp.PanoAddrGrp
    Application *app.PanoApp
    AppGroup *appgrp.PanoGroup
    AppSignature *signature.PanoSignature
    AppSigAndCond *andcond.PanoAndCond
    AppSigOrCond *orcond.PanoOrCond
    Edl *edl.PanoEdl
    LogForwardingProfile *logfwd.PanoLogFwd
    LogForwardingProfileMatchList *matchlist.PanoMatchList
    LogForwardingProfileMatchListAction *action.PanoAction
    Services *srvc.PanoSrvc
    ServiceGroup *srvcgrp.PanoSrvcGrp
    Tags *tags.PanoTags
}

// Initialize is invoked on client.Initialize().
func (c *PanoObjs) Initialize(i util.XapiClient) {
    c.Address = &addr.PanoAddr{}
    c.Address.Initialize(i)

    c.AddressGroup = &addrgrp.PanoAddrGrp{}
    c.AddressGroup.Initialize(i)

    c.Application = &app.PanoApp{}
    c.Application.Initialize(i)

    c.AppGroup = &appgrp.PanoGroup{}
    c.AppGroup.Initialize(i)

    c.AppSignature = &signature.PanoSignature{}
    c.AppSignature.Initialize(i)

    c.AppSigAndCond = &andcond.PanoAndCond{}
    c.AppSigAndCond.Initialize(i)

    c.AppSigOrCond = &orcond.PanoOrCond{}
    c.AppSigOrCond.Initialize(i)

    c.Edl = &edl.PanoEdl{}
    c.Edl.Initialize(i)

    c.LogForwardingProfile = &logfwd.PanoLogFwd{}
    c.LogForwardingProfile.Initialize(i)

    c.LogForwardingProfileMatchList = &matchlist.PanoMatchList{}
    c.LogForwardingProfileMatchList.Initialize(i)

    c.LogForwardingProfileMatchListAction = &action.PanoAction{}
    c.LogForwardingProfileMatchListAction.Initialize(i)

    c.Services = &srvc.PanoSrvc{}
    c.Services.Initialize(i)

    c.ServiceGroup = &srvcgrp.PanoSrvcGrp{}
    c.ServiceGroup.Initialize(i)

    c.Tags = &tags.PanoTags{}
    c.Tags.Initialize(i)
}

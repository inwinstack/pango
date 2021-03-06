package peer

import (
    "fmt"
    "encoding/xml"

    "github.com/inwinstack/pango/util"
    "github.com/inwinstack/pango/version"
)


// PanoPeer is the client.Network.BgpPeer namespace.
type PanoPeer struct {
    con util.XapiClient
}

// Initialize is invoked by client.Initialize().
func (c *PanoPeer) Initialize(con util.XapiClient) {
    c.con = con
}

// ShowList performs SHOW to retrieve a list of values.
func (c *PanoPeer) ShowList(tmpl, ts, vr, pg string) ([]string, error) {
    c.con.LogQuery("(show) list of %s", plural)
    path := c.xpath(tmpl, ts, vr, pg, nil)
    return c.con.EntryListUsing(c.con.Show, path[:len(path) - 1])
}

// GetList performs GET to retrieve a list of values.
func (c *PanoPeer) GetList(tmpl, ts, vr, pg string) ([]string, error) {
    c.con.LogQuery("(get) list of %s", plural)
    path := c.xpath(tmpl, ts, vr, pg, nil)
    return c.con.EntryListUsing(c.con.Get, path[:len(path) - 1])
}

// Get performs GET to retrieve information for the given uid.
func (c *PanoPeer) Get(tmpl, ts, vr, pg, name string) (Entry, error) {
    c.con.LogQuery("(get) %s %q", singular, name)
    return c.details(c.con.Get, tmpl, ts, vr, pg, name)
}

// Show performs SHOW to retrieve information for the given uid.
func (c *PanoPeer) Show(tmpl, ts, vr, pg, name string) (Entry, error) {
    c.con.LogQuery("(show) %s %q", singular, name)
    return c.details(c.con.Show, tmpl, ts, vr, pg, name)
}

// Set performs SET to create / update one or more objects.
func (c *PanoPeer) Set(tmpl, ts, vr, pg string, e ...Entry) error {
    var err error

    if len(e) == 0 {
        return nil
    } else if tmpl == "" && ts == "" {
        return fmt.Errorf("tmpl or ts must be specified")
    } else if vr == "" {
        return fmt.Errorf("vr must be specified")
    } else if pg == "" {
        return fmt.Errorf("pg must be specified")
    }

    _, fn := c.versioning()
    names := make([]string, len(e))

    // Build up the struct.
    d := util.BulkElement{XMLName: xml.Name{Local: "peer"}}
    for i := range e {
        d.Data = append(d.Data, fn(e[i]))
        names[i] = e[i].Name
    }
    c.con.LogAction("(set) %s: %v", plural, names)

    // Set xpath.
    path := c.xpath(tmpl, ts, vr, pg, names)
    if len(e) == 1 {
        path = path[:len(path) - 1]
    } else {
        path = path[:len(path) - 2]
    }

    // Create the objects.
    _, err = c.con.Set(path, d.Config(), nil, nil)
    return err
}

// Edit performs EDIT to create / update one object.
func (c *PanoPeer) Edit(tmpl, ts, vr, pg string, e Entry) error {
    var err error

    if tmpl == "" && ts == "" {
        return fmt.Errorf("tmpl or ts must be specified")
    } else if vr == "" {
        return fmt.Errorf("vr must be specified")
    } else if pg == "" {
        return fmt.Errorf("pg must be specified")
    }

    _, fn := c.versioning()

    c.con.LogAction("(edit) %s %q", singular, e.Name)

    // Set xpath.
    path := c.xpath(tmpl, ts, vr, pg, []string{e.Name})

    // Edit the object.
    _, err = c.con.Edit(path, fn(e), nil, nil)
    return err
}

// Delete removes the given objects.
//
// Objects can be a string or an Entry object.
func (c *PanoPeer) Delete(tmpl, ts, vr, pg string, e ...interface{}) error {
    var err error

    if len(e) == 0 {
        return nil
    } else if vr == "" {
        return fmt.Errorf("vr must be specified")
    } else if pg == "" {
        return fmt.Errorf("pg must be specified")
    } else if tmpl == "" && ts == "" {
        return fmt.Errorf("tmpl or ts must be specified")
    }

    names := make([]string, len(e))
    for i := range e {
        switch v := e[i].(type) {
        case string:
            names[i] = v
        case Entry:
            names[i] = v.Name
        default:
            return fmt.Errorf("Unknown type sent to delete: %s", v)
        }
    }
    c.con.LogAction("(delete) %s: %v", plural, names)

    // Remove the objects.
    path := c.xpath(tmpl, ts, vr, pg, names)
    _, err = c.con.Delete(path, nil, nil)
    return err
}

/** Internal functions for this namespace struct **/

func (c *PanoPeer) versioning() (normalizer, func(Entry) (interface{})) {
    v := c.con.Versioning()

    if v.Gte(version.Number{8, 1, 0, ""}) {
        return &container_v4{}, specify_v4
    } else if v.Gte(version.Number{8, 0, 0, ""}) {
        return &container_v3{}, specify_v3
    } else if v.Gte(version.Number{7, 1, 0, ""}) {
        return &container_v2{}, specify_v2
    } else {
        return &container_v1{}, specify_v1
    }
}

func (c *PanoPeer) details(fn util.Retriever, tmpl, ts, vr, pg, name string) (Entry, error) {
    path := c.xpath(tmpl, ts, vr, pg, []string{name})
    obj, _ := c.versioning()
    if _, err := fn(path, nil, obj); err != nil {
        return Entry{}, err
    }
    ans := obj.Normalize()

    return ans, nil
}

func (c *PanoPeer) xpath(tmpl, ts, vr, pg string, vals []string) []string {
    ans := make([]string, 0, 17)
    ans = append(ans, util.TemplateXpathPrefix(tmpl, ts)...)
    ans = append(ans,
        "config",
        "devices",
        util.AsEntryXpath([]string{"localhost.localdomain"}),
        "network",
        "virtual-router",
        util.AsEntryXpath([]string{vr}),
        "protocol",
        "bgp",
        "peer-group",
        util.AsEntryXpath([]string{pg}),
        "peer",
        util.AsEntryXpath(vals),
    )

    return ans
}

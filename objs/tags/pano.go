package tags

import (
    "fmt"
    "encoding/xml"

    "github.com/inwinstack/pango/util"
)

// PanoTags is a namespace struct, included as part of pango.Client.
type PanoTags struct {
    con util.XapiClient
}

// Initialize is invoked when Initialize on the pango.Client is called.
func (c *PanoTags) Initialize(con util.XapiClient) {
    c.con = con
}

// GetList performs GET to retrieve a list of administrative tags.
func (c *PanoTags) GetList(dg string) ([]string, error) {
    c.con.LogQuery("(get) list of administrative tags")
    path := c.xpath(dg, nil)
    return c.con.EntryListUsing(c.con.Get, path[:len(path) - 1])
}

// ShowList performs SHOW to retrieve a list of administrative tags.
func (c *PanoTags) ShowList(dg string) ([]string, error) {
    c.con.LogQuery("(show) list of administrative tags")
    path := c.xpath(dg, nil)
    return c.con.EntryListUsing(c.con.Show, path[:len(path) - 1])
}

// Get performs GET to retrieve information for the given administrative tag.
func (c *PanoTags) Get(dg, name string) (Entry, error) {
    c.con.LogQuery("(get) administrative tag %q", name)
    return c.details(c.con.Get, dg, name)
}

// Get performs SHOW to retrieve information for the given administrative tag.
func (c *PanoTags) Show(dg, name string) (Entry, error) {
    c.con.LogQuery("(show) administrative tag %q", name)
    return c.details(c.con.Show, dg, name)
}

// Set performs SET to create / update one or more administrative tags.
func (c *PanoTags) Set(dg string, e ...Entry) error {
    var err error

    if len(e) == 0 {
        return nil
    }

    _, fn := c.versioning()
    names := make([]string, len(e))

    // Build up the struct with the given configs.
    d := util.BulkElement{XMLName: xml.Name{Local: "tag"}}
    for i := range e {
        d.Data = append(d.Data, fn(e[i]))
        names[i] = e[i].Name
    }
    c.con.LogAction("(set) administrative tags: %v", names)

    // Set xpath.
    path := c.xpath(dg, names)
    if len(e) == 1 {
        path = path[:len(path) - 1]
    } else {
        path = path[:len(path) - 2]
    }

    // Create the objects.
    _, err = c.con.Set(path, d.Config(), nil, nil)
    return err
}

// Edit performs EDIT to create / update an administrative tag.
func (c *PanoTags) Edit(dg string, e Entry) error {
    var err error

    _, fn := c.versioning()

    c.con.LogAction("(edit) administrative tag %q", e.Name)

    // Set xpath.
    path := c.xpath(dg, []string{e.Name})

    // Create the objects.
    _, err = c.con.Edit(path, fn(e), nil, nil)
    return err
}

// Delete removes the given administrative tags from the firewall.
//
// Administrative tags can be either a string or an Entry object.
func (c *PanoTags) Delete(dg string, e ...interface{}) error {
    var err error

    if len(e) == 0 {
        return nil
    }

    names := make([]string, len(e))
    for i := range e {
        switch v := e[i].(type) {
        case string:
            names[i] = v
        case Entry:
            names[i] = v.Name
        default:
            return fmt.Errorf("Unsupported type to delete: %s", v)
        }
    }
    c.con.LogAction("(delete) administrative tags: %v", names)

    path := c.xpath(dg, names)
    _, err = c.con.Delete(path, nil, nil)
    return err
}

/** Internal functions for the PanoTags struct **/

func (c *PanoTags) versioning() (normalizer, func(Entry) (interface{})) {
    return &container_v1{}, specify_v1
}

func (c *PanoTags) details(fn util.Retriever, dg, name string) (Entry, error) {
    path := c.xpath(dg, []string{name})
    obj, _ := c.versioning()
    _, err := fn(path, nil, obj)
    if err != nil {
        return Entry{}, err
    }
    ans := obj.Normalize()

    return ans, nil
}

func (c *PanoTags) xpath(dg string, vals []string) []string {
    if dg == "" {
        dg = "shared"
    }

    if dg == "shared" {
        return []string {
            "config",
            "shared",
            "tag",
            util.AsEntryXpath(vals),
        }
    }

    return []string {
        "config",
        "devices",
        util.AsEntryXpath([]string{"localhost.localdomain"}),
        "device-group",
        util.AsEntryXpath([]string{dg}),
        "tag",
        util.AsEntryXpath(vals),
    }
}

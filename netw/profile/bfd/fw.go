package bfd

import (
    "fmt"
    "encoding/xml"

    "github.com/inwinstack/pango/util"
)


// FwBfd is a namespace struct, included as part of pango.Client.
type FwBfd struct {
    con util.XapiClient
}

// Initialize is invoked when Initialize on the pango.Client is called.
func (c *FwBfd) Initialize(con util.XapiClient) {
    c.con = con
}

// GetList performs GET to retrieve a list of BFD profiles.
func (c *FwBfd) GetList() ([]string, error) {
    c.con.LogQuery("(get) list of bfd profiles")
    path := c.xpath(nil)
    return c.con.EntryListUsing(c.con.Get, path[:len(path) - 1])
}

// ShowList performs SHOW to retrieve a list of BFD profiles.
func (c *FwBfd) ShowList() ([]string, error) {
    c.con.LogQuery("(show) list of bfd profiles")
    path := c.xpath(nil)
    return c.con.EntryListUsing(c.con.Show, path[:len(path) - 1])
}

// Get performs GET to retrieve information for the given BFD profile.
func (c *FwBfd) Get(name string) (Entry, error) {
    c.con.LogQuery("(get) bfd profile %q", name)
    return c.details(c.con.Get, name)
}

// Get performs SHOW to retrieve information for the given BFD profile.
func (c *FwBfd) Show(name string) (Entry, error) {
    c.con.LogQuery("(show) bfd profile %q", name)
    return c.details(c.con.Show, name)
}

// Set performs SET to create / update one or more BFD profiles.
func (c *FwBfd) Set(e ...Entry) error {
    var err error

    if len(e) == 0 {
        return nil
    }

    _, fn := c.versioning()
    names := make([]string, len(e))

    // Build up the struct with the given configs.
    d := util.BulkElement{XMLName: xml.Name{Local: "bfd-profile"}}
    for i := range e {
        d.Data = append(d.Data, fn(e[i]))
        names[i] = e[i].Name
    }
    c.con.LogAction("(set) bfd profiles: %v", names)

    // Set xpath.
    path := c.xpath(names)
    if len(e) == 1 {
        path = path[:len(path) - 1]
    } else {
        path = path[:len(path) - 2]
    }

    // Create the profiles.
    _, err = c.con.Set(path, d.Config(), nil, nil)
    return err
}

// Edit performs EDIT to create / update an BFD profile.
func (c *FwBfd) Edit(e Entry) error {
    var err error

    _, fn := c.versioning()

    c.con.LogAction("(edit) bfd profile %q", e.Name)

    // Set xpath.
    path := c.xpath([]string{e.Name})

    // Edit the profile.
    _, err = c.con.Edit(path, fn(e), nil, nil)
    return err
}

// Delete removes the given BFD profiles from the firewall.
//
// Profiles can be either a string or an Entry object.
func (c *FwBfd) Delete(e ...interface{}) error {
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
    c.con.LogAction("(delete) bfd profiles: %v", names)

    path := c.xpath(names)
    _, err = c.con.Delete(path, nil, nil)
    return err
}

/** Internal functions for this namespace struct **/

func (c *FwBfd) versioning() (normalizer, func(Entry) (interface{})) {
    return &container_v1{}, specify_v1
}

func (c *FwBfd) details(fn util.Retriever, name string) (Entry, error) {
    path := c.xpath([]string{name})
    obj, _ := c.versioning()
    _, err := fn(path, nil, obj)
    if err != nil {
        return Entry{}, err
    }
    ans := obj.Normalize()

    return ans, nil
}

func (c *FwBfd) xpath(vals []string) []string {
    return []string {
        "config",
        "devices",
        util.AsEntryXpath([]string{"localhost.localdomain"}),
        "network",
        "profiles",
        "bfd-profile",
        util.AsEntryXpath(vals),
    }
}

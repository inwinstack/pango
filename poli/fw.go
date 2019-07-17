package poli

import (
    "github.com/inwinstack/pango/util"

    "github.com/inwinstack/pango/poli/nat"
    "github.com/inwinstack/pango/poli/pbf"
    "github.com/inwinstack/pango/poli/security"
)


// Poli is the client.Policies namespace.
type FwPoli struct {
    Nat *nat.FwNat
    PolicyBasedForwarding *pbf.FwPbf
    Security *security.FwSecurity
}

// Initialize is invoked on client.Initialize().
func (c *FwPoli) Initialize(i util.XapiClient) {
    c.Nat = &nat.FwNat{}
    c.Nat.Initialize(i)

    c.PolicyBasedForwarding = &pbf.FwPbf{}
    c.PolicyBasedForwarding.Initialize(i)

    c.Security = &security.FwSecurity{}
    c.Security.Initialize(i)
}

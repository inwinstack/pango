package v3

import (
    "github.com/inwinstack/pango/version"
)

type tc struct {
    desc string
    version version.Number
    conf Entry
}

func getTests() []tc {
    return []tc{
        {"basic check", version.Number{8, 0, 0, ""}, Entry{
            Name: "t1",
            Manager: "snmp.example.com",
            User: "jdoe",
            EngineId: "0A",
            AuthPassword: "auth",
            PrivPassword: "priv",
        }},
    }
}

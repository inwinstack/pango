package param

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
        {"basic", version.Number{7, 1, 0, ""}, Entry{
            Name: "SplunkId",
            Value: "secret",
        }},
    }
}

module github.com/manifoldco/go-connector

require (
	github.com/alecthomas/gometalinter v2.0.11+incompatible
	github.com/client9/misspell v0.3.4
	github.com/go-openapi/runtime v0.0.0-20170303002511-e66a4c440602
	github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	github.com/gordonklaus/ineffassign v0.0.0-20180909121442-1003c8bd00dc
	github.com/manifoldco/go-manifold v0.9.5
	github.com/tsenart/deadcode v0.0.0-20160724212837-210d2dc333e9
)

// This version of kingpin is incompatible with the released version of
// gometalinter until the next release of gometalinter, and possibly until it
// has go module support, we'll need this exclude, and perhaps some more.
//
// After that point, we should be able to remove it.
exclude gopkg.in/alecthomas/kingpin.v3-unstable v3.0.0-20180810215634-df19058c872c

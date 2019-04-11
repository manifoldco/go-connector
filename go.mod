module github.com/manifoldco/go-connector

require (
	github.com/alecthomas/gometalinter v3.0.0+incompatible
	github.com/client9/misspell v0.3.4
	github.com/go-openapi/runtime v0.0.0-20170303002511-e66a4c440602
	github.com/gordonklaus/ineffassign v0.0.0-20180909121442-1003c8bd00dc
	github.com/manifoldco/go-manifold v0.10.11
	github.com/tsenart/deadcode v0.0.0-20160724212837-210d2dc333e9
	golang.org/x/lint v0.0.0-20181217174547-8f45f776aaf1
)

// This version of kingpin is incompatible with the released version of
// gometalinter until the next release of gometalinter, and possibly until it
// has go module support, we'll need this exclude, and perhaps some more.
//
// After that point, we should be able to remove it.
exclude gopkg.in/alecthomas/kingpin.v3-unstable v3.0.0-20180810215634-df19058c872c

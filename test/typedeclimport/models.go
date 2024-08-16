package typedeclimport

import subpkg "github.com/cockroachdb/gogoproto/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}

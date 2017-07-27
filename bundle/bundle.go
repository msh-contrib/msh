package bundle

import (
	"github.com/hzlmn/graph"
)

// Bundle defines collection of packages and their relations
type Bundle struct {
	entryFile   string
	entryModule interface{}
	PkgsGraph   *graph.Graph
}

// New returns fresh Bundle instance
func New(entryFile string) *Bundle {
	bundle := &Bundle{}
	bundle.entryFile = entryFile

	return bundle
}

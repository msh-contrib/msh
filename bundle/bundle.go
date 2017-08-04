package bundle

import (
	//"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/hzlmn/graph"
	"github.com/hzlmn/msh/module"
	"github.com/hzlmn/msh/utils"
)

// Bundle defines collection of packages and their relations
type Bundle struct {
	EntryFile   string
	EntryModule *module.Module
	PkgsGraph   *graph.Graph
}

// New returns fresh Bundle instance
func New(entryFile string) *Bundle {
	bundle := &Bundle{}

	bundle.EntryFile = entryFile
	bundle.EntryModule = module.New(nil, entryFile, utils.ReadFile(entryFile))

	bundle.EntryModule.Parse()

	bundle.EntryModule.Lookup()

	spew.Dump(bundle.EntryModule.GetScope())

	//fmt.Print(res)

	return bundle
}

func (b *Bundle) Build() {
	fmt.Println("Building bundle...")
}

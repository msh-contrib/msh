package scope

import (
	"github.com/mvdan/sh/syntax"
)

type Scope interface {
	Get(decl string) *syntax.Node
	Set(decl string, node *syntax.Node) bool
}

type GlobaScope struct {
	parent     Scope
	references map[string]*syntax.Node
}

func NewModuleScope(node *syntax.Node, parent Scope) *ModuleScope {
	return &ModuleScope{
		parent:     parent,
		references: make(map[string]*syntax.Node),
	}
}

func (s *ModuleScope) Get(decl string) *syntax.Node {
	if s.references[decl] != nil {
		return s.references[decl]
	}

	if s.parent != nil {
		return s.parent.Get(decl)
	}

	return nil
}

func (s *ModuleScope) Set() {}

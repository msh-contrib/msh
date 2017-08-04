package module

import (
	"bytes"
	//"fmt"
	"github.com/hzlmn/msh/scope"
	"github.com/mvdan/sh/syntax"
)

// Module structure contains base information
// about sh package path, content, AST and more.
type Module struct {
	Importer *Module
	FilePath string
	Content  []byte
	Ast      *syntax.File
	scope    *scope.Scope
}

// New creates fresh instance of Module structure
func New(importer *Module, filepath string, content []byte) *Module {
	return &Module{
		Importer: importer,
		FilePath: filepath,
		Content:  content,
	}
}

func (m *Module) Parse() {
	parser := syntax.NewParser()
	file, _ := parser.Parse(bytes.NewReader(m.Content), m.FilePath)

	m.Ast = file

	// Allocate fresh scope
	m.scope = scope.New(nil)

	//fmt.Print(file)
}

func (m *Module) GetScope() *scope.Scope {
	return m.scope
}

func (m *Module) Lookup() {
	syntax.Walk(m.Ast, func(node syntax.Node) bool {
		switch t := node.(type) {
		case *syntax.Assign:
			m.scope.Set(t.Name.Value, new(interface{}))
		case *syntax.FuncDecl:
			m.scope.Set(t.Name.Value, t.Name.Value)
		}
		return true
	})
}

package module

import (
	"bytes"
	"fmt"
	//"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/hzlmn/msh/scope"
	"github.com/mvdan/sh/syntax"
	"text/scanner"
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

func (m *Module) ScanModule() {
	scan := &scanner.Scanner{}
	scan.Init(bytes.NewReader(m.Content))
	var token rune
	for token != scanner.EOF {
		token = scan.Scan()
		fmt.Println("token", scan.TokenText())
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
		//spew.Dump("Node", node)
		switch t := node.(type) {
		case *syntax.Assign:
			m.scope.Set(t.Name.Value, t.Value)

		case *syntax.FuncDecl:
			m.scope.Set(t.Name.Value, t.Name.Value)

		default:
			spew.Dump(node)
		}
		return true
	})
}

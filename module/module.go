package module

import (
	"github.com/hzlmn/msh/scope"
	"github.com/mvdan/sh/syntax"
)

type Module struct {
	filePath string
	content  []byte
	ast      *syntax.File
	scope    *scope.Scope
}

func New(filepath string, content []byte) *Module {
	return &Module{
		filePath: filepath,
		content:  content,
	}
}

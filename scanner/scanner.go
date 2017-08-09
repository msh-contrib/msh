package scanner

import (
	"io"
	"text/scanner"
)

// This is light wrapper around core text/scanner
// to make it real lexer that produces Token structures

var TokenType string

// Token single text unit structure
type Token struct {
	Type  string
	Value string
	Pos   *scanner.Position
}

// Scanner common structure for finding
// imports inside your sh script
type Scanner struct {
	entry     io.Reader
	engine    *scanner.Scanner
	currToken rune
	syntax    string
}

// New creates fresh instance of Scanner
func New(entry io.Reader) *Scanner {
	scan := &Scanner{entry: entry}

	scan.engine = &scanner.Scanner{}
	scan.engine.Init(scan.entry)

	return scan
}

func (s *Scanner) Next() *Token {
  char, charText := s.engine.Next(), s.engine.TokenText()
  if
}

// Scan module
func (s *Scanner) Scan() {
	//pass
}

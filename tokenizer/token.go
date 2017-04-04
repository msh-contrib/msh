package tokenizer

type Token interface {
  Type TokenType
  Literal string

  line int
  column int
}

package tokenizer

type Token struct {
	Type    int
	Literal string

	line   int
	column int
}

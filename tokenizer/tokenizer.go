package tokenizer

type TokenType int

type Tokenizer struct {
	file    []byte
	currPos int
}

func New(data []byte) *Tokenizer {
	return &Tokenizer{
		currPos: 0,
		file:    data,
	}
}

func (t *Tokenizer) NextToken() {
	return
}

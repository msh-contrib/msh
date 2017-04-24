package tokenizer

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want *Tokenizer
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizer_NextToken(t *testing.T) {
	type fields struct {
		file    []byte
		currPos int
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &Tokenizer{
				file:    tt.fields.file,
				currPos: tt.fields.currPos,
			}
			t.NextToken()
		})
	}
}

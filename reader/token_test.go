package reader

import (
	"testing"
)

func Test_TokenType(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  TokenKind
	}{
		{
			Name: "Reading Comma",
			In:   []byte(`, value`),
			Out:  SeparatorToken,
		},
		{
			Name: "Reading Colon",
			In:   []byte(`: value`),
			Out:  SeparatorToken,
		},
		{
			Name: "Reading Open Brace",
			In:   []byte(`{ key`),
			Out:  ObjectToken,
		},
		{
			Name: "Reading Open Bracket",
			In:   []byte(`[ value`),
			Out:  ArrayToken,
		},
		{
			Name: "Reading Close Bracket",
			In:   []byte(`}, key`),
			Out:  TerminatorToken,
		},
		{
			Name: "Reading Close Bracket",
			In:   []byte(`], key`),
			Out:  TerminatorToken,
		},
		{
			Name: "Reading Quote",
			In:   []byte(`"string"`),
			Out:  StringToken,
		},
		{
			Name: "Reading Zero",
			In:   []byte(`0,`),
			Out:  NumberToken,
		},
		{
			Name: "Reading Positive Number",
			In:   []byte(`6134431`),
			Out:  NumberToken,
		},
		{
			Name: "Reading Negative Number",
			In:   []byte(`-2394129`),
			Out:  NumberToken,
		},
		{
			Name: "Reading True Value",
			In:   []byte(`true, key`),
			Out:  BooleanToken,
		},
		{
			Name: "Reading False Value",
			In:   []byte(`false, key`),
			Out:  BooleanToken,
		},
		{
			Name: "Reading Null Value",
			In:   []byte(`null, key`),
			Out:  NullToken,
		},
		{
			Name: "Reading Invalid Token",
			In:   []byte(`g`),
			Out:  InvalidToken,
		},
		{
			Name: "Reading Nothing",
			In:   []byte(""),
			Out:  InvalidToken,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res := r.Token()
			if res != test.Out {
				t.Errorf("got result \"%d\", want \"%d\"", res, test.Out)
			}
		})
	}
}

func Test_IsNull(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  bool
	}{
		{
			Name: "Reading Null",
			In:   []byte(`null`),
			Out:  true,
		},
		{
			Name: "Reading Not Null",
			In:   []byte(`something`),
			Out:  false,
		},
		{
			Name: "Reading Noting",
			In:   []byte(""),
			Out:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res := r.IsNull()
			if res != test.Out {
				t.Errorf("got result \"%v\", want \"%v\"", res, test.Out)
			}
		})
	}
}

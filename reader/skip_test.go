package reader

import (
	"reflect"
	"testing"
)

func Test_Next(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  bool
		Pos  int
	}{
		{
			Name: "Comma",
			In:   []byte(`,key`),
			Out:  true, Pos: 1,
		},
		{
			Name: "Comma With Whitespace",
			In:   []byte(`, key`),
			Out:  true, Pos: 2,
		},
		{
			Name: "Not Comma",
			In:   []byte(`key`),
			Out:  false, Pos: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res := r.Next()
			if res != test.Out {
				t.Errorf("got result %v, want %v", res, test.Out)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_SkipWhiteSpace(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Pos  int
	}{
		{
			Name: "Whitespace",
			In:   []byte(" \t\n\r,"),
			Pos:  4,
		},
		{
			Name: "No Whitespace",
			In:   []byte(`,`),
			Pos:  0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			r.SkipWhiteSpace()
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_Skip(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Pos  int
		Err  error
	}{
		{
			Name: "Object",
			In:   []byte(`{"key":"value"}  ,`),
			Pos:  17, Err: nil,
		},
		{
			Name: "Array",
			In:   []byte(`["string",123,false]  ,`),
			Pos:  22, Err: nil,
		},
		{
			Name: "String",
			In:   []byte(`"Hello, World!"  ,`),
			Pos:  17, Err: nil,
		},
		{
			Name: "Number",
			In:   []byte(`123  ,`),
			Pos:  5, Err: nil,
		},
		{
			Name: "True",
			In:   []byte(`true  ,`),
			Pos:  6, Err: nil,
		},
		{
			Name: "False",
			In:   []byte(`false  ,`),
			Pos:  7, Err: nil,
		},
		{
			Name: "Null",
			In:   []byte(`null  ,`),
			Pos:  6, Err: nil,
		},
		{
			Name: "Invalid Character",
			In:   []byte(`abcdefg`),
			Pos:  0, Err: NewInvalidCharacterError('a', 0),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			err := r.Skip()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

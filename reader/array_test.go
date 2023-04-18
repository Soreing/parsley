package reader

import (
	"reflect"
	"testing"
)

func Test_OpenArray(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Pos  int
		Err  error
	}{
		{
			Name: "Opening Array",
			In:   []byte(`[value`),
			Pos:  1, Err: nil,
		},
		{
			Name: "Opening Array With Whitespace",
			In:   []byte(`[ value`),
			Pos:  2, Err: nil,
		},
		{
			Name: "Missing Opening Bracket",
			In:   []byte(`value`),
			Pos:  0, Err: NewInvalidCharacterError('v', 0),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			err := r.OpenArray()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_CloseArray(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Pos  int
		Err  error
	}{
		{
			Name: "Closing Array",
			In:   []byte(`], key`),
			Pos:  1, Err: nil,
		},
		{
			Name: "Closing Array With Whitespace",
			In:   []byte(`] , key`),
			Pos:  2, Err: nil,
		},
		{
			Name: "Missing Closing Bracket",
			In:   []byte(`,key`),
			Pos:  0, Err: NewInvalidCharacterError(',', 0),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			err := r.CloseArray()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_SkipArray(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Pos  int
		Err  error
	}{
		{
			Name: "Empty Array",
			In:   []byte(`[]`),
			Pos:  2, Err: nil,
		},
		{
			Name: "Array With One Element",
			In:   []byte(`["string"]`),
			Pos:  10, Err: nil,
		},
		{
			Name: "Array With Multiple Elements",
			In:   []byte(`["string",123,false]`),
			Pos:  20, Err: nil,
		},
		{
			Name: "Array With Whitespace",
			In:   []byte(`[ "string" , 123 , false ]  `),
			Pos:  26, Err: nil,
		},
		{
			Name: "Missing Opening Bracket",
			In:   []byte(`"string",123,false]`),
			Pos:  0, Err: NewInvalidCharacterError('"', 0),
		},
		{
			Name: "Missing Closing Bracket",
			In:   []byte(`["string",123,false`),
			Pos:  19, Err: NewEndOfFileError(),
		},
		{
			Name: "Incomplete Array",
			In:   []byte(`["string",123,false,`),
			Pos:  20, Err: NewEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			err := r.skipArray()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

package reader

import (
	"reflect"
	"testing"
)

func Test_ReadBytes(t *testing.T) {
	tests := []struct {
		Name    string
		In, Out []byte
		Pos     int
		Err     error
	}{
		{
			Name: "String",
			In:   []byte(`"Hello, World!"`),
			Out:  []byte("Hello, World!"),
			Pos:  15, Err: nil,
		},
		{
			Name: "String With Whitespace",
			In:   []byte(`"Hello, World!"  `),
			Out:  []byte("Hello, World!"),
			Pos:  17, Err: nil,
		},
		{
			Name: "String With Escape",
			In:   []byte(`"Hello,\tWorld!"`),
			Out:  []byte("Hello,\tWorld!"),
			Pos:  16, Err: nil,
		},
		{
			Name: "All Escape Characters",
			In:   []byte(`"\"\\\/\b\f\n\r\t\uBeeF"`),
			Out:  []byte("\"\\/\b\f\n\r\t\uBeeF"),
			Pos:  24, Err: nil,
		},
		{
			Name: "Missing Opening Quote",
			In:   []byte(`Hello, World!"`),
			Out:  nil,
			Pos:  0, Err: NewInvalidCharacterError('H', 0),
		},
		{
			Name: "Missing Closing Quote",
			In:   []byte(`"Hello, World!`),
			Out:  nil,
			Pos:  0, Err: NewEndOfFileError(),
		},
		{
			Name: "Invalid Escape",
			In:   []byte(`"Hello,\xWorld!"`),
			Out:  nil,
			Pos:  0, Err: NewInvalidCharacterError('x', 8),
		},
		{
			Name: "Invalid Unicode Sequence",
			In:   []byte(`"Hello,\uBeeTWorld!"`),
			Out:  nil,
			Pos:  0, Err: NewInvalidCharacterError('T', 12),
		},
		{
			Name: "Incomplete Escape At End",
			In:   []byte(`"Hello, World!\`),
			Out:  nil,
			Pos:  0, Err: NewEndOfFileError(),
		},
		{
			Name: "Incomplete Unicode Sequence At End",
			In:   []byte(`"Hello, World!\uA`),
			Out:  nil,
			Pos:  0, Err: NewEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.Bytes()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if string(res) != string(test.Out) {
				t.Errorf("got result \"%s\", want \"%s\"", res, test.Out)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_SkipString(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Pos  int
		Err  error
	}{
		{
			Name: "String",
			In:   []byte(`"Hello, World!"`),
			Pos:  15, Err: nil,
		}, {
			Name: "String With Whitespace",
			In:   []byte(`"Hello, World!"  `),
			Pos:  15, Err: nil,
		},
		{
			Name: "String With Escape",
			In:   []byte(`"Hello,\tWorld!"`),
			Pos:  16, Err: nil,
		},
		{
			Name: "All Escape Characters",
			In:   []byte(`"\"\\\/\b\f\n\r\t\uBeeF"`),
			Pos:  24, Err: nil,
		},
		{
			Name: "Missing Closing Quote",
			In:   []byte(`"Hello, World!`),
			Pos:  0, Err: NewEndOfFileError(),
		},
		{
			Name: "Invalid Escape",
			In:   []byte(`"Hello,\xWorld!"`),
			Pos:  0, Err: NewInvalidCharacterError('x', 8),
		},
		{
			Name: "Invalid Unicode Sequence",
			In:   []byte(`"Hello,\uBeeTWorld!"`),
			Pos:  0, Err: NewInvalidCharacterError('T', 12),
		},
		{
			Name: "Incomplete Escape At End",
			In:   []byte(`"Hello, World!\`),
			Pos:  0, Err: NewEndOfFileError(),
		},
		{
			Name: "Incomplete Unicode Sequence At End",
			In:   []byte(`"Hello, World!\uA`),
			Pos:  0, Err: NewEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			err := r.skipString()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_ReadStrings(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  []string
		Pos  int
		Err  error
	}{
		{
			Name: "Empty Slice",
			In:   []byte(`[]`),
			Out:  []string{},
			Pos:  2, Err: nil,
		},
		{
			Name: "Slice With One Element",
			In:   []byte(`["Hello"]`),
			Out:  []string{"Hello"},
			Pos:  9, Err: nil,
		},
		{
			Name: "Slice With Multiple Elements",
			In:   []byte(`["Hello","World"]`),
			Out:  []string{"Hello", "World"},
			Pos:  17, Err: nil,
		},
		{
			Name: "Slice With Whitespaces",
			In:   []byte(`[ "Hello" , "World" ]  `),
			Out:  []string{"Hello", "World"},
			Pos:  23, Err: nil,
		},
		{
			Name: "Missing Opening Bracket",
			In:   []byte(`"Hello","World"]`),
			Out:  nil,
			Pos:  0, Err: NewInvalidCharacterError('"', 0),
		},
		{
			Name: "Missing Closing Bracket",
			In:   []byte(`["Hello","World"`),
			Out:  []string{"Hello", "World"},
			Pos:  16, Err: NewEndOfFileError(),
		},
		{
			Name: "Incomplete Slice",
			In:   []byte(`["Hello","World",`),
			Out:  nil,
			Pos:  17, Err: NewEndOfFileError(),
		},
		{
			Name: "Missing Comma Between Elements",
			In:   []byte(`["Hello""World"]`),
			Out:  []string{"Hello"},
			Pos:  8, Err: NewInvalidCharacterError('"', 8),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.Strings()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if !reflect.DeepEqual(res, test.Out) {
				t.Errorf("got result %v, want %v", res, test.Out)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

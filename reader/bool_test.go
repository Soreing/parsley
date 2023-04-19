package reader

import (
	"reflect"
	"testing"
)

func Test_ReadBool(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  bool
		Pos  int
		Err  error
	}{
		{
			Name: "True",
			In:   []byte(`true`),
			Out:  true,
			Pos:  4, Err: nil,
		},
		{
			Name: "Not True",
			In:   []byte(`trale`),
			Out:  false,
			Pos:  0, Err: newInvalidCharacterError('a', 2),
		},
		{
			Name: "Short True",
			In:   []byte(`tr`),
			Out:  false,
			Pos:  0, Err: newEndOfFileError(),
		},
		{
			Name: "False",
			In:   []byte(`false`),
			Out:  false,
			Pos:  5, Err: nil,
		},
		{
			Name: "Wrong False",
			In:   []byte(`faulse`),
			Out:  false,
			Pos:  0, Err: newInvalidCharacterError('u', 2),
		},
		{
			Name: "Short False",
			In:   []byte(`fa`),
			Out:  false,
			Pos:  0, Err: newEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.Bool()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if res != test.Out {
				t.Errorf("got result %v, want %v", res, test.Out)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_ReadBools(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  []bool
		Pos  int
		Err  error
	}{
		{
			Name: "Empty Slice",
			In:   []byte(`[]`),
			Out:  []bool{},
			Pos:  2, Err: nil,
		},
		{
			Name: "Slice With One Element",
			In:   []byte(`[true]`),
			Out:  []bool{true},
			Pos:  6, Err: nil,
		},
		{
			Name: "Slice With Multiple Elements",
			In:   []byte(`[true,false]`),
			Out:  []bool{true, false},
			Pos:  12, Err: nil,
		},
		{
			Name: "Slice With Whitespaces",
			In:   []byte(`[ true , false ]  `),
			Out:  []bool{true, false},
			Pos:  18, Err: nil,
		},
		{
			Name: "Missing Opening Bracket",
			In:   []byte(`true,false]`),
			Out:  nil,
			Pos:  0, Err: newInvalidCharacterError('t', 0),
		},
		{
			Name: "Missing Closing Bracket",
			In:   []byte(`[true,false`),
			Out:  []bool{true, false},
			Pos:  11, Err: newEndOfFileError(),
		},
		{
			Name: "Incomplete Slice",
			In:   []byte(`[true,false,`),
			Out:  nil,
			Pos:  12, Err: newEndOfFileError(),
		},
		{
			Name: "Missing Comma Between Elements",
			In:   []byte(`[truefalse]`),
			Out:  []bool{true},
			Pos:  5, Err: newInvalidCharacterError('f', 5),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.Bools()
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

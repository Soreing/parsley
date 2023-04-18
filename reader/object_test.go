package reader

import (
	"reflect"
	"testing"
)

func Test_OpenObject(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Pos  int
		Err  error
	}{
		{
			Name: "Opening Object",
			In:   []byte(`{key`),
			Pos:  1, Err: nil,
		},
		{
			Name: "Opening Object With Whitespace",
			In:   []byte(`{ key`),
			Pos:  2, Err: nil,
		},
		{
			Name: "Missing Opening Brace",
			In:   []byte(`key`),
			Pos:  0, Err: NewInvalidCharacterError('k', 0),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			err := r.OpenObject()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_CloseObject(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Pos  int
		Err  error
	}{
		{
			Name: "Closing Object",
			In:   []byte(`}, key`),
			Pos:  1, Err: nil,
		},
		{
			Name: "Closing Object With Whitespace",
			In:   []byte(`} , key`),
			Pos:  2, Err: nil,
		},
		{
			Name: "Missing Closing Brace",
			In:   []byte(`,key`),
			Pos:  0, Err: NewInvalidCharacterError(',', 0),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			err := r.CloseObject()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_ReadKey(t *testing.T) {
	tests := []struct {
		Name    string
		In, Out []byte
		Pos     int
		Err     error
	}{
		{
			Name: "Reading Key",
			In:   []byte(`"key":`),
			Out:  []byte(`key`),
			Pos:  6, Err: nil,
		},
		{
			Name: "Reading Key With Whitespace",
			In:   []byte(`"key" : `),
			Out:  []byte(`key`),
			Pos:  8, Err: nil,
		},
		{
			Name: "Missing Colon",
			In:   []byte(`"key""value"`),
			Out:  nil,
			Pos:  5, Err: NewInvalidCharacterError('"', 5),
		},
		{
			Name: "Missing Colon End of Input",
			In:   []byte(`"key"`),
			Out:  nil,
			Pos:  5, Err: NewEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.Key()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if !reflect.DeepEqual(res, test.Out) {
				t.Errorf("got result %s, want %s", res, test.Out)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_SkipObject(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Pos  int
		Err  error
	}{
		{
			Name: "Empty Object",
			In:   []byte(`{}`),
			Pos:  2, Err: nil,
		},
		{
			Name: "Object With One Field",
			In:   []byte(`{"key":"value"}`),
			Pos:  15, Err: nil,
		},
		{
			Name: "Object With Multiple Fields",
			In:   []byte(`{"key":"value","key":"value"}`),
			Pos:  29, Err: nil,
		},
		{
			Name: "Object With Whitespace",
			In:   []byte(`{"key":"value","key":"value"}  `),
			Pos:  29, Err: nil,
		},
		{
			Name: "Missing Opening Brace",
			In:   []byte(`"key":"value","key":"value"}`),
			Pos:  0, Err: NewInvalidCharacterError('"', 0),
		},
		{
			Name: "Missing Closing Brace",
			In:   []byte(`{"key":"value","key":"value"`),
			Pos:  28, Err: NewEndOfFileError(),
		},
		{
			Name: "Incomplete Object",
			In:   []byte(`{"key":"value","key":"value",`),
			Pos:  29, Err: NewEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			err := r.skipObject()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

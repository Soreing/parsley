package writer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBoolLen tests if the length of a boolean is returned correctly
func TestBoolLen(t *testing.T) {
	tests := []struct {
		Name   string
		Value  bool
		Length int
	}{
		{
			Name:   "True",
			Value:  true,
			Length: 4,
		},
		{
			Name:   "False",
			Value:  false,
			Length: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res := BoolLen(test.Value)
			assert.Equal(t, test.Length, res)
		})
	}
}

// TestBoolpLen tests if the length of a boolean pointer is returned correctly
func TestBoolpLen(t *testing.T) {
	tests := []struct {
		Name   string
		Value  *bool
		Length int
	}{
		{
			Name:   "Nil",
			Value:  nil,
			Length: 4,
		},
		{
			Name:   "True",
			Value:  newp(true),
			Length: 4,
		},
		{
			Name:   "False",
			Value:  newp(false),
			Length: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res := BoolpLen(test.Value)
			assert.Equal(t, test.Length, res)
		})
	}
}

// TestBoolsLen tests if the length of a boolean slice is returned correctly
func TestBoolsLen(t *testing.T) {
	tests := []struct {
		Name   string
		Values []bool
		Length int
	}{
		{
			Name:   "Nil",
			Values: nil,
			Length: 4,
		},
		{
			Name:   "No values",
			Values: []bool{},
			Length: 2,
		},
		{
			Name:   "One value",
			Values: []bool{true},
			Length: 6,
		},
		{
			Name:   "Multiple values",
			Values: []bool{true, false, true, false},
			Length: 23,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res := BoolsLen(test.Values)
			assert.Equal(t, test.Length, res)
		})
	}
}

// TestWriteBool tests writing a boolean to a buffer
func TestWriteBool(t *testing.T) {
	tests := []struct {
		Name    string
		Space   int
		Value   bool
		Result  []byte
		Storage int
		Cursor  int
	}{
		{
			Name:    "True with enough space",
			Space:   10,
			Value:   true,
			Result:  []byte("true"),
			Storage: 0,
			Cursor:  4,
		},
		{
			Name:    "True without enough space",
			Space:   2,
			Value:   true,
			Result:  []byte("true"),
			Storage: 1,
			Cursor:  2,
		},
		{
			Name:    "False with enough space",
			Space:   10,
			Value:   false,
			Result:  []byte("false"),
			Storage: 0,
			Cursor:  5,
		},
		{
			Name:    "False without enough space",
			Space:   2,
			Value:   false,
			Result:  []byte("false"),
			Storage: 1,
			Cursor:  3,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := NewWriter(test.Space)
			w.Bool(test.Value)
			res := w.Build()

			assert.Equal(t, test.Result, res)
			assert.Equal(t, test.Cursor, w.Cursor)
			assert.Equal(t, test.Storage, len(w.Storage))
		})
	}
}

// TestWriteBoolp tests writing a boolean pointer to a buffer
func TestWriteBoolp(t *testing.T) {
	tests := []struct {
		Name   string
		Space  int
		Value  *bool
		Result []byte
	}{
		{
			Name:   "Nil",
			Space:  10,
			Value:  nil,
			Result: []byte("null"),
		},
		{
			Name:   "True",
			Space:  2,
			Value:  newp(true),
			Result: []byte("true"),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := NewWriter(test.Space)
			w.Boolp(test.Value)
			res := w.Build()

			assert.Equal(t, test.Result, res)
		})
	}
}

// TestWriteBoolp tests writing a boolean slice to a buffer
func TestWriteBools(t *testing.T) {
	tests := []struct {
		Name    string
		Space   int
		Value   []bool
		Result  []byte
		Storage int
		Cursor  int
	}{
		{
			Name:    "Nil",
			Space:   10,
			Value:   nil,
			Result:  []byte("null"),
			Storage: 0,
			Cursor:  4,
		},
		{
			Name:    "Empty slice",
			Space:   10,
			Value:   []bool{},
			Result:  []byte("[]"),
			Storage: 0,
			Cursor:  2,
		},
		{
			Name:    "Definitely enough size with one value",
			Space:   100,
			Value:   []bool{true},
			Result:  []byte("[true]"),
			Storage: 0,
			Cursor:  6,
		},
		{
			Name:    "Definitely enough size with multiple values",
			Space:   100,
			Value:   []bool{true, false, true},
			Result:  []byte("[true,false,true]"),
			Storage: 0,
			Cursor:  17,
		},
		{
			Name:    "Just enough size with one value",
			Space:   6,
			Value:   []bool{true},
			Result:  []byte("[true]"),
			Storage: 0,
			Cursor:  6,
		},
		{
			Name:    "Just enough size with multiple values",
			Space:   17,
			Value:   []bool{true, false, true},
			Result:  []byte("[true,false,true]"),
			Storage: 0,
			Cursor:  17,
		},
		{
			Name:    "Split buffer on opening bracket",
			Space:   0,
			Value:   []bool{true, false, true},
			Result:  []byte("[true,false,true]"),
			Storage: 1,
			Cursor:  17,
		},
		{
			Name:    "Split buffer after opening bracket",
			Space:   1,
			Value:   []bool{true, false, true},
			Result:  []byte("[true,false,true]"),
			Storage: 1,
			Cursor:  16,
		},
		{
			Name:    "Split buffer on value",
			Space:   8,
			Value:   []bool{true, false, true},
			Result:  []byte("[true,false,true]"),
			Storage: 1,
			Cursor:  9,
		},
		{
			Name:    "Split buffer on comma",
			Space:   11,
			Value:   []bool{true, false, true},
			Result:  []byte("[true,false,true]"),
			Storage: 1,
			Cursor:  6,
		},
		{
			Name:    "Split buffer after comma",
			Space:   12,
			Value:   []bool{true, false, true},
			Result:  []byte("[true,false,true]"),
			Storage: 1,
			Cursor:  5,
		},
		{
			Name:    "Split buffer on closing bracket",
			Space:   16,
			Value:   []bool{true, false, true},
			Result:  []byte("[true,false,true]"),
			Storage: 1,
			Cursor:  1,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := NewWriter(test.Space)
			w.Bools(test.Value)
			res := w.Build()

			assert.Equal(t, test.Result, res)
			assert.Equal(t, test.Cursor, w.Cursor)
			assert.Equal(t, test.Storage, len(w.Storage))
		})
	}
}

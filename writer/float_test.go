package writer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFloat32Len tests if the length of a float32 is returned correctly
func TestFloat32Len(t *testing.T) {
	tests := []struct {
		Name   string
		Value  float32
		Length int
	}{
		{
			Name:   "True",
			Value:  2.24,
			Length: 24,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res := Float32Len(test.Value)
			assert.Equal(t, test.Length, res)
		})
	}
}

// TestFloat32pLen tests if the length of a float32 pointer is returned correctly
func TestFloat32pLen(t *testing.T) {
	tests := []struct {
		Name   string
		Value  *float32
		Length int
	}{
		{
			Name:   "Nil",
			Value:  nil,
			Length: 4,
		},
		{
			Name:   "True",
			Value:  newp(float32(2.24)),
			Length: 24,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res := Float32pLen(test.Value)
			assert.Equal(t, test.Length, res)
		})
	}
}

// TestFloat32Len tests if the length of a float32 slice is returned correctly
func TestFloat32sLen(t *testing.T) {
	tests := []struct {
		Name   string
		Values []float32
		Length int
	}{
		{
			Name:   "Nil",
			Values: nil,
			Length: 4,
		},
		{
			Name:   "No values",
			Values: []float32{},
			Length: 2,
		},
		{
			Name:   "One value",
			Values: []float32{2.24},
			Length: 26,
		},
		{
			Name:   "Multiple values",
			Values: []float32{2.24, 5.43, 0.5523, 1},
			Length: 101,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			res := Float32sLen(test.Values)
			assert.Equal(t, test.Length, res)
		})
	}
}

// TestWriteFloat32 tests writing a float32 to a buffer
func TestWriteFloat32(t *testing.T) {
	tests := []struct {
		Name    string
		Space   int
		Value   float32
		Result  []byte
		Storage int
		Cursor  int
	}{
		{
			Name:    "Enough space",
			Space:   30,
			Value:   2.35,
			Result:  []byte("2.35"),
			Storage: 0,
			Cursor:  4,
		},
		{
			Name:    "Not enough space",
			Space:   2,
			Value:   2.35,
			Result:  []byte("2.35"),
			Storage: 1,
			Cursor:  4,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := NewWriter(test.Space)
			w.Float32(test.Value)
			res := w.Build()

			assert.Equal(t, test.Result, res)
			assert.Equal(t, test.Cursor, w.Cursor)
			assert.Equal(t, test.Storage, len(w.Storage))
		})
	}
}

func Test_WriteFloat32p(t *testing.T) {
	tests := []struct {
		Name   string
		Space  int
		Value  *float32
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
			Value:  newp(float32(2.56)),
			Result: []byte("2.56"),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := NewWriter(test.Space)
			w.Float32p(test.Value)
			res := w.Build()

			assert.Equal(t, test.Result, res)
		})
	}
}

func Test_WriteFloat32s(t *testing.T) {
	tests := []struct {
		Name    string
		Space   int
		Value   []float32
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
			Value:   []float32{},
			Result:  []byte("[]"),
			Storage: 0,
			Cursor:  2,
		},
		{
			Name:    "Definitely enough size with one value",
			Space:   100,
			Value:   []float32{2.74},
			Result:  []byte("[2.74]"),
			Storage: 0,
			Cursor:  6,
		},
		{
			Name:    "Definitely enough size with multiple values",
			Space:   100,
			Value:   []float32{2.74, 5.555, 0.01},
			Result:  []byte("[2.74,5.555,0.01]"),
			Storage: 0,
			Cursor:  17,
		},
		{
			// Weird because of fixed size assumption
			Name:    "Just enough size with one value",
			Space:   25,
			Value:   []float32{2.74},
			Result:  []byte("[2.74]"),
			Storage: 0,
			Cursor:  6,
		},
		{
			// Weird because of fixed size assumption
			Name:    "Just enough size with multiple values",
			Space:   37,
			Value:   []float32{2.74, 5.555, 0.01},
			Result:  []byte("[2.74,5.555,0.01]"),
			Storage: 0,
			Cursor:  17,
		},
		{
			Name:    "Split buffer on opening bracket",
			Space:   0,
			Value:   []float32{2.74, 5.555, 0.01},
			Result:  []byte("[2.74,5.555,0.01]"),
			Storage: 1,
			Cursor:  17,
		},
		{
			Name:    "Split buffer after opening bracket",
			Space:   1,
			Value:   []float32{2.74, 5.555, 0.01},
			Result:  []byte("[2.74,5.555,0.01]"),
			Storage: 1,
			Cursor:  16,
		},
		{
			// Weird because of fixed size
			Name:    "Split buffer on value",
			Space:   28,
			Value:   []float32{2.74, 5.555, 0.01},
			Result:  []byte("[2.74,5.555,0.01]"),
			Storage: 1,
			Cursor:  11,
		},
		// Not possible because of fixed size
		// {
		// 	Name:    "Split buffer on comma",
		// 	Space:   11,
		// 	Value:   []float32{2.74, 5.555, 0.01},
		// 	Result:  []byte("[2.74,5.555,0.01]"),
		// 	Storage: 1,
		// 	Cursor:  6,
		// },
		// {
		// 	Name:    "Split buffer after comma",
		// 	Space:   12,
		// 	Value:   []float32{2.74, 5.555, 0.01},
		// 	Result:  []byte("[2.74,5.555,0.01]"),
		// 	Storage: 1,
		// 	Cursor:  5,
		// },
		// {
		// 	Name:    "Split buffer on closing bracket",
		// 	Space:   16,
		// 	Value:   []float32{2.74, 5.555, 0.01},
		// 	Result:  []byte("[2.74,5.555,0.01]"),
		// 	Storage: 1,
		// 	Cursor:  1,
		// },
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			w := NewWriter(test.Space)
			w.Float32s(test.Value)
			res := w.Build()

			assert.Equal(t, test.Result, res)
			assert.Equal(t, test.Cursor, w.Cursor)
			assert.Equal(t, test.Storage, len(w.Storage))
		})
	}
}

// func Test_WriteFloat64(t *testing.T) {
// 	tests := []struct {
// 		Name   string
// 		In     float64
// 		Out    []byte
// 		Space  int
// 		Stored int
// 		Cur    int
// 		Err    error
// 	}{
// 		{
// 			Name: "One Buffer",
// 			In:   2.35, Out: []byte("2.35"),
// 			Space: 30, Stored: 0, Cur: 4,
// 		},
// 		{
// 			Name: "Split Buffer",
// 			In:   2.35, Out: []byte("2.35"),
// 			Space: 2, Stored: 1, Cur: 4,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			w := NewWriter(test.Space)
// 			w.Float64(test.In)
// 			res := w.Build()
// 			if string(res) != string(test.Out) {
// 				t.Errorf("got result %s, want %s", res, test.Out)
// 			}
// 			if w.Cursor != test.Cur {
// 				t.Errorf("got cursor %d, want %d", w.Cursor, test.Cur)
// 			}
// 			if len(w.Storage) != test.Stored {
// 				t.Errorf("got stored buffers %d, want %d", len(w.Storage), test.Stored)
// 			}
// 		})
// 	}
// }

// func Test_WriteFloat64p(t *testing.T) {
// 	tests := []struct {
// 		Name   string
// 		In     *float64
// 		Out    []byte
// 		Space  int
// 		Stored int
// 		Cur    int
// 		Err    error
// 	}{
// 		{
// 			Name: "One Buffer (nil)",
// 			In:   nil, Out: []byte("null"),
// 			Space: 10, Stored: 0, Cur: 4,
// 		},
// 		{
// 			Name: "Split Buffer (nil)",
// 			In:   nil, Out: []byte("null"),
// 			Space: 2, Stored: 1, Cur: 2,
// 		},
// 		{
// 			Name: "Pointer With Value",
// 			In:   newp(2.35), Out: []byte("2.35"),
// 			Space: 30, Stored: 0, Cur: 4,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			w := NewWriter(test.Space)
// 			w.Float64p(test.In)
// 			res := w.Build()
// 			if string(res) != string(test.Out) {
// 				t.Errorf("got result %s, want %s", res, test.Out)
// 			}
// 			if w.Cursor != test.Cur {
// 				t.Errorf("got cursor %d, want %d", w.Cursor, test.Cur)
// 			}
// 			if len(w.Storage) != test.Stored {
// 				t.Errorf("got stored buffers %d, want %d", len(w.Storage), test.Stored)
// 			}
// 		})
// 	}
// }

// func Test_WriteFloat64s(t *testing.T) {
// 	tests := []struct {
// 		Name   string
// 		In     []float64
// 		Out    []byte
// 		Space  int
// 		Stored int
// 		Cur    int
// 		Err    error
// 	}{
// 		{
// 			Name: "One Buffer (nil)",
// 			In:   nil, Out: []byte("null"),
// 			Space: 10, Stored: 0, Cur: 4,
// 		},
// 		{
// 			Name: "Split Buffer (nil)",
// 			In:   nil, Out: []byte("null"),
// 			Space: 2, Stored: 1, Cur: 2,
// 		},
// 		{
// 			Name: "One Buffer (empty)",
// 			In:   []float64{}, Out: []byte("[]"),
// 			Space: 10, Stored: 0, Cur: 2,
// 		},
// 		{
// 			Name: "Split Buffer (empty)",
// 			In:   []float64{}, Out: []byte("[]"),
// 			Space: 1, Stored: 1, Cur: 1,
// 		},
// 		{
// 			Name: "One Buffer (value)",
// 			In:   []float64{2.35}, Out: []byte("[2.35]"),
// 			Space: 30, Stored: 0, Cur: 6,
// 		},
// 		{
// 			Name: "One Buffer (values)",
// 			In:   []float64{2.35, 55.12, 0.12}, Out: []byte("[2.35,55.12,0.12]"),
// 			Space: 50, Stored: 0, Cur: 17,
// 		},
// 		{
// 			Name: "Split After [",
// 			In:   []float64{2.35, 55.12, 0.12}, Out: []byte("[2.35,55.12,0.12]"),
// 			Space: 1, Stored: 1, Cur: 16,
// 		},
// 		// Rest is dodgy because of no fixed size for floats
// 	}

// 	for _, test := range tests {
// 		t.Run(test.Name, func(t *testing.T) {
// 			w := NewWriter(test.Space)
// 			w.Float64s(test.In)
// 			res := w.Build()
// 			if string(res) != string(test.Out) {
// 				t.Errorf("got result %s, want %s", res, test.Out)
// 			}
// 			if w.Cursor != test.Cur {
// 				t.Errorf("got cursor %d, want %d", w.Cursor, test.Cur)
// 			}
// 			if len(w.Storage) != test.Stored {
// 				t.Errorf("got stored buffers %d, want %d", len(w.Storage), test.Stored)
// 			}
// 		})
// 	}
// }

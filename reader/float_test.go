package reader

import (
	"math"
	"reflect"
	"strconv"
	"testing"
)

func Test_ReadFloat64(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Ovr  float64
		Pos  int
		Err  error
	}{
		{
			Name: "Integer Zero",
			In:   []byte(`0`), Ovr: math.Inf(0),
			Pos: 1, Err: nil,
		},
		{
			Name: "Fraction Zero",
			In:   []byte(`0.0`), Ovr: math.Inf(0),
			Pos: 3, Err: nil,
		},
		{
			Name: "Zero Wit Exponent",
			In:   []byte(`0e+1`), Ovr: math.Inf(0),
			Pos: 4, Err: nil,
		},
		{
			Name: "Positive Integer",
			In:   []byte(`123`), Ovr: math.Inf(0),
			Pos: 3, Err: nil,
		},
		{
			Name: "Negative Integer",
			In:   []byte(`-123`), Ovr: math.Inf(0),
			Pos: 4, Err: nil,
		},
		{
			Name: "Number With Whitespace",
			In:   []byte(`123  `), Ovr: 123,
			Pos: 5, Err: nil,
		},
		{
			Name: "Fraction Above 1",
			In:   []byte(`12.3`), Ovr: math.Inf(0),
			Pos: 4, Err: nil,
		},
		{
			Name: "Fraction Below 1",
			In:   []byte(`0.025`), Ovr: math.Inf(0),
			Pos: 5, Err: nil,
		},
		{
			Name: "Integer Positive Exponent",
			In:   []byte(`12e+4`), Ovr: math.Inf(0),
			Pos: 5, Err: nil,
		},
		{
			Name: "Integer Negative Exponent",
			In:   []byte(`12e-4`), Ovr: math.Inf(0),
			Pos: 5, Err: nil,
		},
		{
			Name: "Fraction Above 1 Positive Exponent",
			In:   []byte(`12.3e+4`), Ovr: math.Inf(0),
			Pos: 7, Err: nil,
		},
		{
			Name: "Fraction Above 1 Negative Exponent",
			In:   []byte(`12.3e-4`), Ovr: math.Inf(0),
			Pos: 7, Err: nil,
		},
		{
			Name: "Fraction Below 1 Positive Exponent",
			In:   []byte(`0.025e+4`), Ovr: math.Inf(0),
			Pos: 8, Err: nil,
		},
		{
			Name: "Fraction Below 1 Negative Exponent",
			In:   []byte(`0.025e-4`), Ovr: math.Inf(0),
			Pos: 8, Err: nil,
		},
		{
			Name: "Truncated No Dot",
			In:   []byte(`123456789012345678901234567890`), Ovr: math.Inf(0),
			Pos: 30, Err: nil,
		},
		{
			Name: "Truncated Before Dot",
			In:   []byte(`1234567890123456789012345.67890`), Ovr: math.Inf(0),
			Pos: 31, Err: nil,
		},
		{
			Name: "Truncated After Dot",
			In:   []byte(`123456.789012345678901234567890`), Ovr: math.Inf(0),
			Pos: 31, Err: nil,
		},
		{
			Name: "Truncated After Dot",
			In:   []byte(`0.0123456789012345678901234567890`), Ovr: math.Inf(0),
			Pos: 33, Err: nil,
		},
		{
			Name: "Truncated Dot In Middle",
			In:   []byte(`12345678.9012345678901234567890`), Ovr: math.Inf(0),
			Pos: 31, Err: nil,
		},
		{
			Name: "Very Large Positive Number",
			In:   []byte(`123456789012345678901234567890e+350`), Ovr: math.Inf(0),
			Pos: 0, Err: NewNumberOutOfRangeError(
				[]byte(`123456789012345678901234567890e+350`), 0,
			),
		},
		{
			Name: "Very Large Negative Number",
			In:   []byte(`-123456789012345678901234567890e+350`), Ovr: math.Inf(0),
			Pos: 0, Err: NewNumberOutOfRangeError(
				[]byte(`123456789012345678901234567890e+350`), 0,
			),
		},
		{
			Name: "Very Small Number",
			In:   []byte(`123456789012345678901234567890e-350`), Ovr: math.Inf(0),
			Pos: 35, Err: nil,
		},
		{
			Name: "Leading Zero",
			In:   []byte(`0123`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError('1', 1),
		},
		{
			Name: "Trailing Sign",
			In:   []byte(`-,`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError(',', 1),
		},
		{
			Name: "Trailing Sign End of Input",
			In:   []byte(`-`), Ovr: 0,
			Pos: 0, Err: NewEndOfFileError(),
		},
		{
			Name: "Leading Dot",
			In:   []byte(`.5`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError('.', 0),
		},
		{
			Name: "Trailing Dot",
			In:   []byte(`2.,`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError(',', 2),
		},
		{
			Name: "Trailing Dot End of Input",
			In:   []byte(`2.`), Ovr: 0,
			Pos: 0, Err: NewEndOfFileError(),
		},
		{
			Name: "Trailing Exponent End of Input",
			In:   []byte(`2.0e`), Ovr: 0,
			Pos: 0, Err: NewEndOfFileError(),
		},
		{
			Name: "Trailing Exponent",
			In:   []byte(`2.0e,`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError(',', 4),
		},
		{
			Name: "Trailing Positive Exponent End of Input",
			In:   []byte(`2.0e+`), Ovr: 0,
			Pos: 0, Err: NewEndOfFileError(),
		},
		{
			Name: "Trailing Negative Exponent End of Input",
			In:   []byte(`2.0e-`), Ovr: 0,
			Pos: 0, Err: NewEndOfFileError(),
		},
		{
			Name: "Trailing Positive Exponent",
			In:   []byte(`2.0e+,`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError(',', 5),
		},
		{
			Name: "Trailing Negative Exponent",
			In:   []byte(`2.0e-,`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError(',', 5),
		},
		{
			Name: "Exponent No Sign",
			In:   []byte(`2.0e2`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError('2', 4),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			out, _ := strconv.ParseFloat(string(test.In), 64)
			res, err := r.Float64()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if test.Ovr == math.Inf(0) && res != out {
				t.Errorf("got result 1 \"%f\", want \"%f\"", res, out)
			}
			if test.Ovr != math.Inf(0) && res != test.Ovr {
				t.Errorf("got result 2 \"%f\", want \"%f\"", res, test.Ovr)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_ReadFloat32(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Ovr  float64
		Pos  int
		Err  error
	}{
		{
			Name: "Integer Zero",
			In:   []byte(`0`), Ovr: math.Inf(0),
			Pos: 1, Err: nil,
		},
		{
			Name: "Fraction Zero",
			In:   []byte(`0.0`), Ovr: math.Inf(0),
			Pos: 3, Err: nil,
		},
		{
			Name: "Zero Wit Exponent",
			In:   []byte(`0e+1`), Ovr: math.Inf(0),
			Pos: 4, Err: nil,
		},
		{
			Name: "Positive Integer",
			In:   []byte(`123`), Ovr: math.Inf(0),
			Pos: 3, Err: nil,
		},
		{
			Name: "Negative Integer",
			In:   []byte(`-123`), Ovr: math.Inf(0),
			Pos: 4, Err: nil,
		},
		{
			Name: "Number With Whitespace",
			In:   []byte(`123  `), Ovr: 123,
			Pos: 5, Err: nil,
		},
		{
			Name: "Fraction Above 1",
			In:   []byte(`12.3`), Ovr: math.Inf(0),
			Pos: 4, Err: nil,
		},
		{
			Name: "Fraction Below 1",
			In:   []byte(`0.025`), Ovr: math.Inf(0),
			Pos: 5, Err: nil,
		},
		{
			Name: "Integer Positive Exponent",
			In:   []byte(`12e+4`), Ovr: math.Inf(0),
			Pos: 5, Err: nil,
		},
		{
			Name: "Integer Negative Exponent",
			In:   []byte(`12e-4`), Ovr: math.Inf(0),
			Pos: 5, Err: nil,
		},
		{
			Name: "Fraction Above 1 Positive Exponent",
			In:   []byte(`12.3e+4`), Ovr: math.Inf(0),
			Pos: 7, Err: nil,
		},
		{
			Name: "Fraction Above 1 Negative Exponent",
			In:   []byte(`12.3e-4`), Ovr: math.Inf(0),
			Pos: 7, Err: nil,
		},
		{
			Name: "Fraction Below 1 Positive Exponent",
			In:   []byte(`0.025e+4`), Ovr: math.Inf(0),
			Pos: 8, Err: nil,
		},
		{
			Name: "Fraction Below 1 Negative Exponent",
			In:   []byte(`0.025e-4`), Ovr: math.Inf(0),
			Pos: 8, Err: nil,
		},
		{
			Name: "Truncated No Dot",
			In:   []byte(`123456789012345678901234567890`), Ovr: math.Inf(0),
			Pos: 30, Err: nil,
		},
		{
			Name: "Truncated Before Dot",
			In:   []byte(`1234567890123456789012345.67890`), Ovr: math.Inf(0),
			Pos: 31, Err: nil,
		},
		{
			Name: "Truncated After Dot",
			In:   []byte(`123456.789012345678901234567890`), Ovr: math.Inf(0),
			Pos: 31, Err: nil,
		},
		{
			Name: "Truncated After Dot",
			In:   []byte(`0.0123456789012345678901234567890`), Ovr: math.Inf(0),
			Pos: 33, Err: nil,
		},
		{
			Name: "Truncated Dot In Middle",
			In:   []byte(`12345678.9012345678901234567890`), Ovr: math.Inf(0),
			Pos: 31, Err: nil,
		},
		{
			Name: "Very Large Positive Number",
			In:   []byte(`123456789012345678901234567890e+150`), Ovr: math.Inf(0),
			Pos: 0, Err: NewNumberOutOfRangeError(
				[]byte(`123456789012345678901234567890e+150`), 0,
			),
		},
		{
			Name: "Very Large Negative Number",
			In:   []byte(`-123456789012345678901234567890e+150`), Ovr: math.Inf(0),
			Pos: 0, Err: NewNumberOutOfRangeError(
				[]byte(`123456789012345678901234567890e+150`), 0,
			),
		},
		{
			Name: "Very Small Number",
			In:   []byte(`123456789012345678901234567890e-150`), Ovr: math.Inf(0),
			Pos: 35, Err: nil,
		},
		{
			Name: "Leading Zero",
			In:   []byte(`0123`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError('1', 1),
		},
		{
			Name: "Trailing Sign",
			In:   []byte(`-,`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError(',', 1),
		},
		{
			Name: "Trailing Sign End of Input",
			In:   []byte(`-`), Ovr: 0,
			Pos: 0, Err: NewEndOfFileError(),
		},
		{
			Name: "Leading Dot",
			In:   []byte(`.5`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError('.', 0),
		},
		{
			Name: "Trailing Dot",
			In:   []byte(`2.,`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError(',', 2),
		},
		{
			Name: "Trailing Dot End of Input",
			In:   []byte(`2.`), Ovr: 0,
			Pos: 0, Err: NewEndOfFileError(),
		},
		{
			Name: "Trailing Exponent End of Input",
			In:   []byte(`2.0e`), Ovr: 0,
			Pos: 0, Err: NewEndOfFileError(),
		},
		{
			Name: "Trailing Exponent",
			In:   []byte(`2.0e,`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError(',', 4),
		},
		{
			Name: "Trailing Positive Exponent End of Input",
			In:   []byte(`2.0e+`), Ovr: 0,
			Pos: 0, Err: NewEndOfFileError(),
		},
		{
			Name: "Trailing Negative Exponent End of Input",
			In:   []byte(`2.0e-`), Ovr: 0,
			Pos: 0, Err: NewEndOfFileError(),
		},
		{
			Name: "Trailing Positive Exponent",
			In:   []byte(`2.0e+,`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError(',', 5),
		},
		{
			Name: "Trailing Negative Exponent",
			In:   []byte(`2.0e-,`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError(',', 5),
		},
		{
			Name: "Exponent No Sign",
			In:   []byte(`2.0e2`), Ovr: 0,
			Pos: 0, Err: NewInvalidCharacterError('2', 4),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			out, _ := strconv.ParseFloat(string(test.In), 32)
			res, err := r.Float32()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if test.Ovr == math.Inf(0) && res != float32(out) {
				t.Errorf("got result 1 \"%f\", want \"%f\"", res, out)
			}
			if test.Ovr != math.Inf(0) && res != float32(test.Ovr) {
				t.Errorf("got result 2 \"%f\", want \"%f\"", res, test.Ovr)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

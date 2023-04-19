package reader

import (
	"reflect"
	"testing"
)

func Test_SkipNumber(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Pos  int
		Err  error
	}{
		{
			Name: "Integer Zero",
			In:   []byte(`0`),
			Pos:  1, Err: nil,
		},
		{
			Name: "Fraction Zero",
			In:   []byte(`0.0`),
			Pos:  3, Err: nil,
		},
		{
			Name: "Zero Wit Exponent",
			In:   []byte(`0e+1`),
			Pos:  4, Err: nil,
		},
		{
			Name: "Positive Integer",
			In:   []byte(`123`),
			Pos:  3, Err: nil,
		},
		{
			Name: "Negative Integer",
			In:   []byte(`-123`),
			Pos:  4, Err: nil,
		},
		{
			Name: "Fraction Above 1",
			In:   []byte(`12.3`),
			Pos:  4, Err: nil,
		},
		{
			Name: "Fraction Below 1",
			In:   []byte(`0.025`),
			Pos:  5, Err: nil,
		},
		{
			Name: "Positive Exponent",
			In:   []byte(`12e+4`),
			Pos:  5, Err: nil,
		},
		{
			Name: "Negative Exponent",
			In:   []byte(`12e-4`),
			Pos:  5, Err: nil,
		},
		{
			Name: "Leading Zero",
			In:   []byte(`0123`),
			Pos:  0, Err: newInvalidCharacterError('1', 1),
		},
		{
			Name: "Trailing Sign",
			In:   []byte(`-,`),
			Pos:  0, Err: newInvalidCharacterError(',', 1),
		},
		{
			Name: "Trailing Sign End of Input",
			In:   []byte(`-`),
			Pos:  0, Err: newEndOfFileError(),
		},
		{
			Name: "Leading Dot",
			In:   []byte(`.5`),
			Pos:  0, Err: newInvalidCharacterError('.', 0),
		},
		{
			Name: "Trailing Dot",
			In:   []byte(`2.,`),
			Pos:  0, Err: newInvalidCharacterError(',', 2),
		},
		{
			Name: "Trailing Dot End of Input",
			In:   []byte(`2.`),
			Pos:  0, Err: newEndOfFileError(),
		},
		{
			Name: "Trailing Exponent End of Input",
			In:   []byte(`2.0e`),
			Pos:  0, Err: newEndOfFileError(),
		},
		{
			Name: "Trailing Exponent",
			In:   []byte(`2.0e,`),
			Pos:  0, Err: newInvalidCharacterError(',', 4),
		},
		{
			Name: "Trailing Positive Exponent End of Input",
			In:   []byte(`2.0e+`),
			Pos:  0, Err: newEndOfFileError(),
		},
		{
			Name: "Trailing Negative Exponent End of Input",
			In:   []byte(`2.0e-`),
			Pos:  0, Err: newEndOfFileError(),
		},
		{
			Name: "Trailing Positive Exponent",
			In:   []byte(`2.0e+,`),
			Pos:  0, Err: newInvalidCharacterError(',', 5),
		},
		{
			Name: "Trailing Negative Exponent",
			In:   []byte(`2.0e-,`),
			Pos:  0,
			Err:  newInvalidCharacterError(',', 5),
		},
		{
			Name: "Exponent No Sign",
			In:   []byte(`2.0e2`),
			Pos:  0,
			Err:  newInvalidCharacterError('2', 4),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			err := r.skipNumber()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_ReadInteger(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Num  uint64
		Neg  bool
		Pos  int
		Ok   bool
	}{
		{
			Name: "Integer Zero",
			In:   []byte(`0`),
			Num:  0, Neg: false,
			Pos: 1, Ok: true,
		},
		{
			Name: "Fraction Zero",
			In:   []byte(`0.0`),
			Num:  0, Neg: false,
			Pos: 3, Ok: true,
		},
		{
			Name: "Positive Integer",
			In:   []byte(`123`),
			Num:  123, Neg: false,
			Pos: 3, Ok: true,
		},
		{
			Name: "Negative Integer",
			In:   []byte(`-123`),
			Num:  123, Neg: true,
			Pos: 4, Ok: true,
		},
		{
			Name: "Fraction Above 1",
			In:   []byte(`12.3`),
			Num:  12, Neg: false,
			Pos: 4, Ok: true,
		},
		{
			Name: "Fraction Below 1",
			In:   []byte(`0.25`),
			Num:  0, Neg: false,
			Pos: 4, Ok: true,
		},
		{
			Name: "Integer With Positive Exponent",
			In:   []byte(`12e+4`),
			Num:  120000, Neg: false,
			Pos: 5, Ok: true,
		},
		{
			Name: "Integer With Negative Exponent",
			In:   []byte(`12e-4`),
			Num:  0, Neg: false,
			Pos: 5, Ok: true,
		},
		{
			Name: "Fraction With Positive Exponent",
			In:   []byte(`12.3e+4`),
			Num:  123000, Neg: false,
			Pos: 7, Ok: true,
		},
		{
			Name: "Fraction With Negative Exponent",
			In:   []byte(`12.3e-1`),
			Num:  1, Neg: false,
			Pos: 7, Ok: true,
		},
		{
			Name: "Fraction With Positive Exponent Large",
			In:   []byte(`12.3e+30`),
			Num:  0, Neg: true,
			Pos: 8, Ok: false,
		},
		{
			Name: "Fraction With Negative Exponent Zero",
			In:   []byte(`12.3e-4`),
			Num:  0, Neg: false,
			Pos: 7, Ok: true,
		},
		{
			Name: "Long Integer Part",
			In:   []byte(`1000000000000000000000000000000000000.1e-18`),
			Num:  1000000000000000000, Neg: false,
			Pos: 43, Ok: true,
		},
		{
			Name: "Leading Zeros",
			In:   []byte(`0.00000000001e+11`),
			Num:  1, Neg: false,
			Pos: 17, Ok: true,
		},
		{
			Name: "UInt64 Max",
			In:   []byte(`18446744073709551615`),
			Num:  18446744073709551615, Neg: false,
			Pos: 20, Ok: true,
		},
		{
			Name: "UInt64 Max Fraction With Exponent",
			In:   []byte(`1844674407370955.1615e+4`),
			Num:  18446744073709551615, Neg: false,
			Pos: 24, Ok: true,
		},
		{
			Name: "UInt64 Max +1",
			In:   []byte(`18446744073709551616`),
			Num:  0, Neg: true,
			Pos: 20, Ok: false,
		},
		{
			Name: "Leading Zero Integer",
			In:   []byte(`0123`),
			Num:  0, Neg: false,
			Pos: 1, Ok: false,
		},
		{
			Name: "Leading Dot",
			In:   []byte(`.5`),
			Num:  0, Neg: false,
			Pos: 0, Ok: false,
		},
		{
			Name: "Trailing Dot",
			In:   []byte(`2.,`),
			Num:  0, Neg: false,
			Pos: 2, Ok: false,
		},
		{
			Name: "Trailing Dot End of Input",
			In:   []byte(`2.`),
			Num:  0, Neg: false,
			Pos: 2, Ok: false,
		},
		{
			Name: "Trailing Exponent",
			In:   []byte(`2.0e`),
			Num:  0, Neg: false,
			Pos: 4, Ok: false,
		},
		{
			Name: "Trailing Positive Exponent",
			In:   []byte(`2.0e+,`),
			Num:  0, Neg: false,
			Pos: 5, Ok: false,
		},
		{
			Name: "Trailing Negative Exponent",
			In:   []byte(`2.0e-,`),
			Num:  0, Neg: false,
			Pos: 5, Ok: false,
		},
		{
			Name: "Trailing Positive Exponent End of Input",
			In:   []byte(`2.0e+`),
			Num:  0, Neg: false,
			Pos: 5, Ok: false,
		},
		{
			Name: "Trailing Negative Exponent End of Input",
			In:   []byte(`2.0e-`),
			Num:  0, Neg: false,
			Pos: 5, Ok: false,
		},
		{
			Name: "Exponent No Sign",
			In:   []byte(`2.0e2`),
			Num:  0, Neg: false,
			Pos: 4, Ok: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			num, neg, pos, ok := readInteger(test.In)
			if num != test.Num {
				t.Errorf("got number %d, want %d", num, test.Num)
			}
			if neg != test.Neg {
				t.Errorf("got negative %v, want %v", neg, test.Neg)
			}
			if pos != test.Pos {
				t.Errorf("got position %d, want %d", pos, test.Pos)
			}
			if ok != test.Ok {
				t.Errorf("got success %v, want %v", ok, test.Ok)
			}
		})
	}
}

func Test_ReadFloat(t *testing.T) {
	tests := []struct {
		Name      string
		In        []byte
		Man       uint64
		Dig, Exp  int
		Neg, Trc  bool
		Dp, Sp, I int
		Ok        bool
	}{
		{
			Name: "Integer Zero",
			In:   []byte(`0`),
			Man:  0, Dig: 0, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 0,
			I: 1, Ok: true,
		},
		{
			Name: "Fraction Zero",
			In:   []byte(`0.0`),
			Man:  0, Dig: 0, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 3,
			I: 3, Ok: true,
		},
		{
			Name: "Zero Wit Exponent",
			In:   []byte(`0e+1`),
			Man:  0, Dig: 0, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 0,
			I: 4, Ok: true,
		},
		{
			Name: "Positive Integer",
			In:   []byte(`123`),
			Man:  123, Dig: 3, Exp: 0, Neg: false, Trc: false, Dp: 3, Sp: 0,
			I: 3, Ok: true,
		},
		{
			Name: "Negative Integer",
			In:   []byte(`-123`),
			Man:  123, Dig: 3, Exp: 0, Neg: true, Trc: false, Dp: 4, Sp: 1,
			I: 4, Ok: true,
		},
		{
			Name: "Fraction Above 1",
			In:   []byte(`12.3`),
			Man:  123, Dig: 3, Exp: 0, Neg: false, Trc: false, Dp: 2, Sp: 0,
			I: 4, Ok: true,
		},
		{
			Name: "Fraction Below 1",
			In:   []byte(`0.025`),
			Man:  25, Dig: 2, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 3,
			I: 5, Ok: true,
		},
		{
			Name: "Positive Exponent",
			In:   []byte(`12e+4`),
			Man:  12, Dig: 2, Exp: 4, Neg: false, Trc: false, Dp: 2, Sp: 0,
			I: 5, Ok: true,
		},
		{
			Name: "Negative Exponent",
			In:   []byte(`12e-4`),
			Man:  12, Dig: 2, Exp: -4, Neg: false, Trc: false, Dp: 2, Sp: 0,
			I: 5, Ok: true,
		},
		{
			Name: "Truncated No Dot",
			In:   []byte(`1234567890123456789000000`),
			Man:  1234567890123456789, Dig: 19, Exp: 0, Neg: false, Trc: true, Dp: 25, Sp: 0,
			I: 25, Ok: true,
		},
		{
			Name: "Truncated Before Dot",
			In:   []byte(`1234567890123456789000000.0`),
			Man:  1234567890123456789, Dig: 19, Exp: 0, Neg: false, Trc: true, Dp: 25, Sp: 0,
			I: 27, Ok: true,
		},
		{
			Name: "Truncated After Dot",
			In:   []byte(`0.1234567890123456789000000`),
			Man:  1234567890123456789, Dig: 19, Exp: 0, Neg: false, Trc: true, Dp: 1, Sp: 2,
			I: 27, Ok: true,
		},
		{
			Name: "19 Digits No Dot",
			In:   []byte(`1234567890123456789`),
			Man:  1234567890123456789, Dig: 19, Exp: 0, Neg: false, Trc: false, Dp: 19, Sp: 0,
			I: 19, Ok: true,
		},
		{
			Name: "19 Digits Starting Before Dot",
			In:   []byte(`123456789012345678.9`),
			Man:  1234567890123456789, Dig: 19, Exp: 0, Neg: false, Trc: false, Dp: 18, Sp: 0,
			I: 20, Ok: true,
		},
		{
			Name: "19 Digits Starting After Dot",
			In:   []byte(`0.1234567890123456789`),
			Man:  1234567890123456789, Dig: 19, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 2,
			I: 21, Ok: true,
		},
		{
			Name: "Leading Zero Integer",
			In:   []byte(`0123`),
			Man:  0, Dig: 0, Exp: 0, Neg: false, Trc: false, Dp: 0, Sp: 0,
			I: 1, Ok: false,
		},
		{
			Name: "Trailing Sign",
			In:   []byte(`-,`),
			Man:  0, Dig: 0, Exp: 0, Neg: true, Trc: false, Dp: 0, Sp: 0,
			I: 1, Ok: false,
		},
		{
			Name: "Trailing Sign End of Input",
			In:   []byte(`-`),
			Man:  0, Dig: 0, Exp: 0, Neg: true, Trc: false, Dp: 0, Sp: 0,
			I: 1, Ok: false,
		},
		{
			Name: "Leading Dot",
			In:   []byte(`.5`),
			Man:  0, Dig: 0, Exp: 0, Neg: false, Trc: false, Dp: 0, Sp: 0,
			I: 0, Ok: false,
		},
		{
			Name: "Trailing Dot",
			In:   []byte(`2.,`),
			Man:  2, Dig: 1, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 0,
			I: 2, Ok: false,
		},
		{
			Name: "Trailing Dot End of Input",
			In:   []byte(`2.`),
			Man:  2, Dig: 1, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 0,
			I: 2, Ok: false,
		},
		{
			Name: "Trailing Exponent End of Input",
			In:   []byte(`2.0e`),
			Man:  20, Dig: 2, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 0,
			I: 4, Ok: false,
		},
		{
			Name: "Trailing Exponent",
			In:   []byte(`2.0e,`),
			Man:  20, Dig: 2, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 0,
			I: 4, Ok: false,
		},
		{
			Name: "Trailing Positive Exponent End of Input",
			In:   []byte(`2.0e+`),
			Man:  20, Dig: 2, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 0,
			I: 5, Ok: false,
		},
		{
			Name: "Trailing Negative Exponent End of Input",
			In:   []byte(`2.0e-`),
			Man:  20, Dig: 2, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 0,
			I: 5, Ok: false,
		},
		{
			Name: "Trailing Positive Exponent",
			In:   []byte(`2.0e+,`),
			Man:  20, Dig: 2, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 0,
			I: 5, Ok: false,
		},
		{
			Name: "Trailing Negative Exponent",
			In:   []byte(`2.0e-,`),
			Man:  20, Dig: 2, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 0,
			I: 5, Ok: false,
		},
		{
			Name: "Exponent No Sign",
			In:   []byte(`2.0e2`),
			Man:  20, Dig: 2, Exp: 0, Neg: false, Trc: false, Dp: 1, Sp: 0,
			I: 4, Ok: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			m, d, e, n, tr, dp, sp, i, ok := readFloat(test.In)
			if ok != test.Ok {
				t.Errorf("got success %v, want %v", ok, test.Ok)
			}
			if m != test.Man {
				t.Errorf("got mantissa %d, want %d", m, test.Man)
			}
			if d != test.Dig {
				t.Errorf("got digits %d, want %d", d, test.Dig)
			}
			if e != test.Exp {
				t.Errorf("got exponent %d, want %d", e, test.Exp)
			}
			if n != test.Neg {
				t.Errorf("got negative %v, want %v", n, test.Neg)
			}
			if tr != test.Trc {
				t.Errorf("got truncated %v, want %v", tr, test.Trc)
			}
			if dp != test.Dp {
				t.Errorf("got dot position %d, want %d", dp, test.Dp)
			}
			if sp != test.Sp {
				t.Errorf("got start position %d, want %d", sp, test.Sp)
			}
			if i != test.I {
				t.Errorf("got bytes read %d, want %d", i, test.I)
			}
		})
	}
}

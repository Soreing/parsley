package reader

import (
	"reflect"
	"testing"
)

func Test_ReadInt(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  int
		Pos  int
		Err  error
	}{
		{
			Name: "Integer Zero",
			In:   []byte(`0`),
			Out:  0, Pos: 1, Err: nil,
		},
		{
			Name: "Positive Integer",
			In:   []byte(`123`),
			Out:  123, Pos: 3, Err: nil,
		},
		{
			Name: "Negative Integer",
			In:   []byte(`-123`),
			Out:  -123, Pos: 4, Err: nil,
		},
		{
			Name: "Integer With Whitespace",
			In:   []byte(`123  `),
			Out:  123, Pos: 5, Err: nil,
		},
		{
			Name: "Int Max",
			In:   []byte(`2147483647`),
			Out:  2147483647, Pos: 10, Err: nil,
		},
		{
			Name: "Int Max +1",
			In:   []byte(`2147483648`),
			Out:  0, Pos: 0, Err: newNumberOutOfRangeError([]byte(`2147483648`), 0),
		},
		{
			Name: "Int Min",
			In:   []byte(`-2147483648`),
			Out:  -2147483648, Pos: 11, Err: nil,
		},
		{
			Name: "Int Min -1",
			In:   []byte(`-2147483649`),
			Out:  0, Pos: 0, Err: newNumberOutOfRangeError([]byte(`-2147483649`), 0),
		},
		{
			Name: "Syntax Error",
			In:   []byte(`01`),
			Out:  0, Pos: 0, Err: newInvalidCharacterError('1', 1),
		},
		{
			Name: "End of Input",
			In:   []byte(`12e`),
			Out:  0, Pos: 0, Err: newEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.Int()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if res != test.Out {
				t.Errorf("got result \"%d\", want \"%d\"", res, test.Out)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_ReadInt64(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  int64
		Pos  int
		Err  error
	}{
		{
			Name: "Integer Zero",
			In:   []byte(`0`),
			Out:  0, Pos: 1, Err: nil,
		},
		{
			Name: "Positive Integer",
			In:   []byte(`123`),
			Out:  123, Pos: 3, Err: nil,
		},
		{
			Name: "Negative Integer",
			In:   []byte(`-123`),
			Out:  -123, Pos: 4, Err: nil,
		},
		{
			Name: "Integer With Whitespace",
			In:   []byte(`123  `),
			Out:  123, Pos: 5, Err: nil,
		},
		{
			Name: "Int64 Max",
			In:   []byte(`9223372036854775807`),
			Out:  9223372036854775807, Pos: 19, Err: nil,
		},
		{
			Name: "Int64 Max +1",
			In:   []byte(`9223372036854775808`),
			Out:  0, Pos: 0, Err: newNumberOutOfRangeError([]byte(`9223372036854775808`), 0),
		},
		{
			Name: "Int64 Min",
			In:   []byte(`-9223372036854775808`),
			Out:  -9223372036854775808, Pos: 20, Err: nil,
		},
		{
			Name: "Int64 Min -1",
			In:   []byte(`-9223372036854775809`),
			Out:  0, Pos: 0, Err: newNumberOutOfRangeError([]byte(`-9223372036854775809`), 0),
		},
		{
			Name: "Syntax Error",
			In:   []byte(`01`),
			Out:  0, Pos: 0, Err: newInvalidCharacterError('1', 1),
		},
		{
			Name: "End of Input",
			In:   []byte(`12e`),
			Out:  0, Pos: 0, Err: newEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.Int64()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if res != test.Out {
				t.Errorf("got result \"%d\", want \"%d\"", res, test.Out)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_ReadInt32(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  int32
		Pos  int
		Err  error
	}{
		{
			Name: "Integer Zero",
			In:   []byte(`0`),
			Out:  0, Pos: 1, Err: nil,
		},
		{
			Name: "Positive Integer",
			In:   []byte(`123`),
			Out:  123, Pos: 3, Err: nil,
		},
		{
			Name: "Negative Integer",
			In:   []byte(`-123`),
			Out:  -123, Pos: 4, Err: nil,
		},
		{
			Name: "Integer With Whitespace",
			In:   []byte(`123  `),
			Out:  123, Pos: 5, Err: nil,
		},
		{
			Name: "Int32 Max",
			In:   []byte(`2147483647`),
			Out:  2147483647, Pos: 10, Err: nil,
		},
		{
			Name: "Int32 Max +1",
			In:   []byte(`2147483648`),
			Out:  0, Pos: 0, Err: newNumberOutOfRangeError([]byte(`2147483648`), 0),
		},
		{
			Name: "Int32 Min",
			In:   []byte(`-2147483648`),
			Out:  -2147483648, Pos: 11, Err: nil,
		},
		{
			Name: "Int32 Min -1",
			In:   []byte(`-2147483649`),
			Out:  0, Pos: 0, Err: newNumberOutOfRangeError([]byte(`-2147483649`), 0),
		},
		{
			Name: "Syntax Error",
			In:   []byte(`01`),
			Out:  0, Pos: 0, Err: newInvalidCharacterError('1', 1),
		},
		{
			Name: "End of Input",
			In:   []byte(`12e`),
			Out:  0, Pos: 0, Err: newEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.Int32()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if res != test.Out {
				t.Errorf("got result \"%d\", want \"%d\"", res, test.Out)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_ReadInt16(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  int16
		Pos  int
		Err  error
	}{
		{
			Name: "Integer Zero",
			In:   []byte(`0`),
			Out:  0, Pos: 1, Err: nil,
		},
		{
			Name: "Positive Integer",
			In:   []byte(`123`),
			Out:  123, Pos: 3, Err: nil,
		},
		{
			Name: "Negative Integer",
			In:   []byte(`-123`),
			Out:  -123, Pos: 4, Err: nil,
		},
		{
			Name: "Integer With Whitespace",
			In:   []byte(`123  `),
			Out:  123, Pos: 5, Err: nil,
		},
		{
			Name: "Int16 Max",
			In:   []byte(`32767`),
			Out:  32767, Pos: 5, Err: nil,
		},
		{
			Name: "Int16 Max +1",
			In:   []byte(`32768`),
			Out:  0, Pos: 0, Err: newNumberOutOfRangeError([]byte(`32768`), 0),
		},
		{
			Name: "Int16 Min",
			In:   []byte(`-32768`),
			Out:  -32768, Pos: 6, Err: nil,
		},
		{
			Name: "Int16 Min -1",
			In:   []byte(`-32769`),
			Out:  0, Pos: 0, Err: newNumberOutOfRangeError([]byte(`-32769`), 0),
		},
		{
			Name: "Syntax Error",
			In:   []byte(`01`),
			Out:  0, Pos: 0, Err: newInvalidCharacterError('1', 1),
		},
		{
			Name: "End of Input",
			In:   []byte(`12e`),
			Out:  0, Pos: 0, Err: newEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.Int16()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if res != test.Out {
				t.Errorf("got result \"%d\", want \"%d\"", res, test.Out)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

func Test_ReadInt8(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  int8
		Pos  int
		Err  error
	}{
		{
			Name: "Integer Zero",
			In:   []byte(`0`),
			Out:  0, Pos: 1, Err: nil,
		},
		{
			Name: "Positive Integer",
			In:   []byte(`123`),
			Out:  123, Pos: 3, Err: nil,
		},
		{
			Name: "Negative Integer",
			In:   []byte(`-123`),
			Out:  -123, Pos: 4, Err: nil,
		},
		{
			Name: "Integer With Whitespace",
			In:   []byte(`123  `),
			Out:  123, Pos: 5, Err: nil,
		},
		{
			Name: "Int8 Max",
			In:   []byte(`127`),
			Out:  127, Pos: 3, Err: nil,
		},
		{
			Name: "Int8 Max +1",
			In:   []byte(`128`),
			Out:  0, Pos: 0, Err: newNumberOutOfRangeError([]byte(`128`), 0),
		},
		{
			Name: "Int8 Min",
			In:   []byte(`-128`),
			Out:  -128, Pos: 4, Err: nil,
		},
		{
			Name: "Int8 Min -1",
			In:   []byte(`-129`),
			Out:  0, Pos: 0, Err: newNumberOutOfRangeError([]byte(`-129`), 0),
		},
		{
			Name: "Syntax Error",
			In:   []byte(`01`),
			Out:  0, Pos: 0, Err: newInvalidCharacterError('1', 1),
		},
		{
			Name: "End of Input",
			In:   []byte(`12e`),
			Out:  0, Pos: 0, Err: newEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.Int8()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if res != test.Out {
				t.Errorf("got result \"%d\", want \"%d\"", res, test.Out)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

// TODO: Write tests for Int Slices...

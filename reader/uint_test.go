package reader

import (
	"reflect"
	"testing"
)

func Test_ReadUInt(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  uint
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
			Out:  0, Pos: 0, Err: NewNumberOutOfRangeError([]byte(`-123`), 0),
		},
		{
			Name: "Integer With Whitespace",
			In:   []byte(`123  `),
			Out:  123, Pos: 5, Err: nil,
		},
		{
			Name: "UInt Max",
			In:   []byte(`4294967295`),
			Out:  4294967295, Pos: 10, Err: nil,
		},
		{
			Name: "UInt Max +1",
			In:   []byte(`4294967296`),
			Out:  0, Pos: 0, Err: NewNumberOutOfRangeError([]byte(`4294967296`), 0),
		},
		{
			Name: "Syntax Error",
			In:   []byte(`01`),
			Out:  0, Pos: 0, Err: NewInvalidCharacterError('1', 1),
		},
		{
			Name: "End of Input",
			In:   []byte(`12e`),
			Out:  0, Pos: 0, Err: NewEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.UInt()
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

func Test_ReadUInt64(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  uint64
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
			Out:  0, Pos: 0, Err: NewNumberOutOfRangeError([]byte(`-123`), 0),
		},
		{
			Name: "Integer With Whitespace",
			In:   []byte(`123  `),
			Out:  123, Pos: 5, Err: nil,
		},
		{
			Name: "UInt64 Max",
			In:   []byte(`18446744073709551615`),
			Out:  18446744073709551615, Pos: 20, Err: nil,
		},
		{
			Name: "UInt64 Max +1",
			In:   []byte(`18446744073709551616`),
			Out:  0, Pos: 0, Err: NewNumberOutOfRangeError([]byte(`18446744073709551616`), 0),
		},
		{
			Name: "Syntax Error",
			In:   []byte(`01`),
			Out:  0, Pos: 0, Err: NewInvalidCharacterError('1', 1),
		},
		{
			Name: "End of Input",
			In:   []byte(`12e`),
			Out:  0, Pos: 0, Err: NewEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.UInt64()
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

func Test_ReadUInt32(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  uint32
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
			Out:  0, Pos: 0, Err: NewNumberOutOfRangeError([]byte(`-123`), 0),
		},
		{
			Name: "Integer With Whitespace",
			In:   []byte(`123  `),
			Out:  123, Pos: 5, Err: nil,
		},
		{
			Name: "UInt32 Max",
			In:   []byte(`4294967295`),
			Out:  4294967295, Pos: 10, Err: nil,
		},
		{
			Name: "UInt32 Max +1",
			In:   []byte(`4294967296`),
			Out:  0, Pos: 0, Err: NewNumberOutOfRangeError([]byte(`4294967296`), 0),
		},
		{
			Name: "Syntax Error",
			In:   []byte(`01`),
			Out:  0, Pos: 0, Err: NewInvalidCharacterError('1', 1),
		},
		{
			Name: "End of Input",
			In:   []byte(`12e`),
			Out:  0, Pos: 0, Err: NewEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.UInt32()
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

func Test_ReadUInt16(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  uint16
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
			Out:  0, Pos: 0, Err: NewNumberOutOfRangeError([]byte(`-123`), 0),
		},
		{
			Name: "Integer With Whitespace",
			In:   []byte(`123  `),
			Out:  123, Pos: 5, Err: nil,
		},
		{
			Name: "UInt16 Max",
			In:   []byte(`65535`),
			Out:  65535, Pos: 5, Err: nil,
		},
		{
			Name: "UInt16 Max +1",
			In:   []byte(`65536`),
			Out:  0, Pos: 0, Err: NewNumberOutOfRangeError([]byte(`65536`), 0),
		},
		{
			Name: "Syntax Error",
			In:   []byte(`01`),
			Out:  0, Pos: 0, Err: NewInvalidCharacterError('1', 1),
		},
		{
			Name: "End of Input",
			In:   []byte(`12e`),
			Out:  0, Pos: 0, Err: NewEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.UInt16()
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

func Test_ReadUInt8(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  uint8
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
			Out:  0, Pos: 0, Err: NewNumberOutOfRangeError([]byte(`-123`), 0),
		},
		{
			Name: "Integer With Whitespace",
			In:   []byte(`123  `),
			Out:  123, Pos: 5, Err: nil,
		},
		{
			Name: "UInt8 Max",
			In:   []byte(`255`),
			Out:  255, Pos: 3, Err: nil,
		},
		{
			Name: "UInt8 Max +1",
			In:   []byte(`256`),
			Out:  0, Pos: 0, Err: NewNumberOutOfRangeError([]byte(`256`), 0),
		},
		{
			Name: "Syntax Error",
			In:   []byte(`01`),
			Out:  0, Pos: 0, Err: NewInvalidCharacterError('1', 1),
		},
		{
			Name: "End of Input",
			In:   []byte(`12e`),
			Out:  0, Pos: 0, Err: NewEndOfFileError(),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.UInt8()
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

func Test_ReadBase64(t *testing.T) {
	tests := []struct {
		Name string
		In   []byte
		Out  []byte
		Pos  int
		Err  error
	}{
		{
			Name: "Empty String",
			In:   []byte(`""`),
			Out:  []byte{},
			Pos:  2, Err: nil,
		},
		{
			Name: "One Group",
			In:   []byte(`"PyOr"`),
			Out:  []byte{0x3F, 0x23, 0xAB},
			Pos:  6, Err: nil,
		},
		{
			Name: "Multiple Groups",
			In:   []byte(`"PyOrTFVqms1X"`),
			Out:  []byte{0x3F, 0x23, 0xAB, 0x4C, 0x55, 0x6A, 0x9A, 0xCD, 0x57},
			Pos:  14, Err: nil,
		},
		{
			Name: "One Padding",
			In:   []byte(`"PyM="`),
			Out:  []byte{0x3F, 0x23},
			Pos:  6, Err: nil,
		},
		{
			Name: "Two Padding",
			In:   []byte(`"Pw=="`),
			Out:  []byte{0x3F},
			Pos:  6, Err: nil,
		},
		{
			Name: "Missing Opening Quote",
			In:   []byte(`PyOrTFVqms1X"`),
			Out:  nil,
			Pos:  0, Err: NewInvalidCharacterError('P', 0),
		},
		{
			Name: "Missing Closing Quote",
			In:   []byte(`"PyOrTFVqms1X`),
			Out:  nil,
			Pos:  0, Err: NewEndOfFileError(),
		},
		{
			Name: "Incomplete Group",
			In:   []byte(`"PyOrTFVqm"`),
			Out:  nil,
			Pos:  0, Err: NewBase64PaddingError(10),
		},
		{
			Name: "Invalid Digit",
			In:   []byte(`"PyOrTF,qms1X"`),
			Out:  nil,
			Pos:  0, Err: NewInvalidCharacterError(',', 7),
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := NewReader(test.In)
			res, err := r.UInt8s()
			if !reflect.DeepEqual(err, test.Err) {
				t.Errorf("got error %v, want error %v", err, test.Err)
			}
			if string(res) != string(test.Out) {
				t.Errorf("got result \"%v\", want \"%v\"", res, test.Out)
			}
			if r.pos != test.Pos {
				t.Errorf("got position %d, want %d", r.pos, test.Pos)
			}
		})
	}
}

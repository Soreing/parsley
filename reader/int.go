package reader

import (
	"math"
)

// intSeq extracts int values recursively until the closing bracket
// is found, then assigns the elements to the allocated slice.
func (r *Reader) intSeq(idx int) (res []int, err error) {
	var n int
	if n, err = r.Int(); err == nil {
		if r.Next() {
			res, err = r.intSeq(idx + 1)
		} else {
			res = make([]int, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

// Ints extracts an array of int values from the data and skips all whitespace
// after it. The values must be enclosed in square brackets "[...]" and the
// values must be separated by commas.
func (r *Reader) Ints() (res []int, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []int{}
			err = r.CloseArray()
		} else if res, err = r.intSeq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

// Int extracts the next int value from the data and skips all
// whitespace after it.
func (r *Reader) Int() (n int, err error) {
	dat := r.dat[r.pos:]
	num, neg, pos, ok := readInteger(dat)
	if !ok {
		if num == 0 && neg {
			return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
		} else if pos == len(dat) {
			return 0, newEndOfFileError()
		} else {
			return 0, newInvalidCharacterError(dat[pos], r.pos+pos)
		}
	} else if !neg && num <= math.MaxInt32 {
		n = int(num)
	} else if neg && num <= math.MaxInt32+1 {
		n = int(-int64(num))
	} else {
		return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

// Intp extracts the next int value and returns a pointer variable.
func (r *Reader) Intp() (res *int, err error) {
	if v, err := r.Int(); err == nil {
		res = &v
	}
	return
}

// int8Seq extracts int8 values recursively until the closing bracket
// is found, then assigns the elements to the allocated slice.
func (r *Reader) int8Seq(idx int) (res []int8, err error) {
	var n int8
	if n, err = r.Int8(); err == nil {
		if r.Next() {
			res, err = r.int8Seq(idx + 1)
		} else {
			res = make([]int8, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

// Int8s extracts an array of int8 values from the data and skips all
// whitespace after it. The values must be enclosed in square brackets "[...]"
// and the values must be separated by commas.
func (r *Reader) Int8s() (res []int8, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []int8{}
			err = r.CloseArray()
		} else if res, err = r.int8Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

// Int8 extracts the next int8 value from the data and skips all
// whitespace after it.
func (r *Reader) Int8() (n int8, err error) {
	dat := r.dat[r.pos:]
	num, neg, pos, ok := readInteger(dat)
	if !ok {
		if num == 0 && neg {
			return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
		} else if pos == len(dat) {
			return 0, newEndOfFileError()
		} else {
			return 0, newInvalidCharacterError(dat[pos], r.pos+pos)
		}
	} else if !neg && num <= math.MaxInt8 {
		n = int8(num)
	} else if neg && num <= math.MaxInt8+1 {
		n = int8(-int64(num))
	} else {
		return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

// Int8p extracts the next int8 value and returns a pointer variable.
func (r *Reader) Int8p() (res *int8, err error) {
	if v, err := r.Int8(); err == nil {
		res = &v
	}
	return
}

// int16Seq extracts int16 values recursively until the closing bracket
// is found, then assigns the elements to the allocated slice.
func (r *Reader) int16Seq(idx int) (res []int16, err error) {
	var n int16
	if n, err = r.Int16(); err == nil {
		if r.Next() {
			res, err = r.int16Seq(idx + 1)
		} else {
			res = make([]int16, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

// Int16s extracts an array of int16 values from the data and skips all
// whitespace after it. The values must be enclosed in square brackets "[...]"
// and the values must be separated by commas.
func (r *Reader) Int16s() (res []int16, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []int16{}
			err = r.CloseArray()
		} else if res, err = r.int16Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

// Int16 extracts the next int16 value from the data and skips all
// whitespace after it.
func (r *Reader) Int16() (n int16, err error) {
	dat := r.dat[r.pos:]
	num, neg, pos, ok := readInteger(dat)
	if !ok {
		if num == 0 && neg {
			return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
		} else if pos == len(dat) {
			return 0, newEndOfFileError()
		} else {
			return 0, newInvalidCharacterError(dat[pos], r.pos+pos)
		}
	} else if !neg && num <= math.MaxInt16 {
		n = int16(num)
	} else if neg && num <= math.MaxInt16+1 {
		n = int16(-int64(num))
	} else {
		return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

// Int16p extracts the next int16 value and returns a pointer variable.
func (r *Reader) Int16p() (res *int16, err error) {
	if v, err := r.Int16(); err == nil {
		res = &v
	}
	return
}

// int32Seq extracts int32 values recursively until the closing bracket
// is found, then assigns the elements to the allocated slice.
func (r *Reader) int32Seq(idx int) (res []int32, err error) {
	var n int32
	if n, err = r.Int32(); err == nil {
		if r.Next() {
			res, err = r.int32Seq(idx + 1)
		} else {
			res = make([]int32, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

// Int32s extracts an array of int32 values from the data and skips all
// whitespace after it. The values must be enclosed in square brackets "[...]"
// and the values must be separated by commas.
func (r *Reader) Int32s() (res []int32, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []int32{}
			err = r.CloseArray()
		} else if res, err = r.int32Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

// Int32 extracts the next int32 value from the data and skips all
// whitespace after it.
func (r *Reader) Int32() (n int32, err error) {
	dat := r.dat[r.pos:]
	num, neg, pos, ok := readInteger(dat)
	if !ok {
		if num == 0 && neg {
			return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
		} else if pos == len(dat) {
			return 0, newEndOfFileError()
		} else {
			return 0, newInvalidCharacterError(dat[pos], r.pos+pos)
		}
	} else if !neg && num <= math.MaxInt32 {
		n = int32(num)
	} else if neg && num <= math.MaxInt32+1 {
		n = int32(-int64(num))
	} else {
		return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

// Int32p extracts the next int32 value and returns a pointer variable.
func (r *Reader) Int32p() (res *int32, err error) {
	if v, err := r.Int32(); err == nil {
		res = &v
	}
	return
}

// int64Seq extracts int64 values recursively until the closing bracket
// is found, then assigns the elements to the allocated slice.
func (r *Reader) int64Seq(idx int) (res []int64, err error) {
	var n int64
	if n, err = r.Int64(); err == nil {
		if r.Next() {
			res, err = r.int64Seq(idx + 1)
		} else {
			res = make([]int64, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

// Int64s extracts an array of int64 values from the data and skips all
// whitespace after it. The values must be enclosed in square brackets "[...]"
// and the values must be separated by commas.
func (r *Reader) Int64s() (res []int64, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []int64{}
			err = r.CloseArray()
		} else if res, err = r.int64Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

// Int64 extracts the next int64 value from the data and skips all
// whitespace after it.
func (r *Reader) Int64() (n int64, err error) {
	dat := r.dat[r.pos:]
	num, neg, pos, ok := readInteger(dat)
	if !ok {
		if num == 0 && neg {
			return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
		} else if pos == len(dat) {
			return 0, newEndOfFileError()
		} else {
			return 0, newInvalidCharacterError(dat[pos], r.pos+pos)
		}
	} else if !neg && num <= math.MaxInt64 {
		n = int64(num)
	} else if neg && num <= math.MaxInt64 {
		n = -int64(num)
	} else if neg && num == math.MaxInt64+1 {
		n = math.MinInt64
	} else {
		return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

// Int64p extracts the next int64 value and returns a pointer variable.
func (r *Reader) Int64p() (res *int64, err error) {
	if v, err := r.Int64(); err == nil {
		res = &v
	}
	return
}

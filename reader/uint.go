package reader

import (
	"encoding/base64"
	"math"
)

// uintSeq extracts uint values recursively until the closing bracket
// is found, then assigns the elements to the allocated slice.
func (r *Reader) uintSeq(idx int) (res []uint, err error) {
	var n uint
	if n, err = r.UInt(); err == nil {
		if r.Next() {
			res, err = r.uintSeq(idx + 1)
		} else {
			res = make([]uint, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

// UInts extracts an array of uint values from the data and skips all whitespace
// after it. The values must be enclosed in square brackets "[...]" and the
// values must be separated by commas.
func (r *Reader) UInts() (res []uint, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []uint{}
			err = r.CloseArray()
		} else if res, err = r.uintSeq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

// UInt extracts the next uint value from the data and skips all
// whitespace after it.
func (r *Reader) UInt() (n uint, err error) {
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
	} else if !neg && num <= math.MaxUint32 {
		n = uint(num)
	} else {
		return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

// UIntp extracts the next uint value and returns a pointer variable.
func (r *Reader) UIntp() (res *uint, err error) {
	if v, err := r.UInt(); err == nil {
		res = &v
	}
	return
}

// UInt8s extracts the next base64 string value enclosed in quotes and returns
// the value in a byte array. Skips all whitespace after the string.
func (r *Reader) UInt8s() (res []uint8, err error) {
	dat := r.dat[r.pos:]

	if len(dat) < 2 {
		return nil, newEndOfFileError()
	} else if dat[0] != '"' {
		return nil, newInvalidCharacterError(dat[0], r.pos)
	}

	beg, end, c := 1, 1, dat[1]
	for dat[end] != '"' {
		if end == len(dat)-1 {
			return nil, newEndOfFileError()
		} else if c|0x20-'a' < 26 || c-'0' < 10 ||
			c == '+' || c == '/' || c == '=' {
			end++
			c = dat[end]
		} else {
			return nil, newInvalidCharacterError(c, r.pos+end)
		}
	}

	if (end-beg)%4 != 0 {
		return nil, newBase64PaddingError(r.pos + end)
	}

	bytes := (end - beg) / 4 * 3
	if bytes > 2 && dat[end-1] == '=' {
		bytes--
	}
	if bytes > 1 && dat[end-2] == '=' {
		bytes--
	}

	dst := make([]byte, bytes)
	if _, err := base64.StdEncoding.Decode(dst, dat[beg:end]); err != nil {
		return nil, err
	}

	r.pos += end + 1
	r.SkipWhiteSpace()
	return dst, nil
}

// UInt8 extracts the next uint8 value from the data and skips all
// whitespace after it.
func (r *Reader) UInt8() (n uint8, err error) {
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
	} else if !neg && num <= math.MaxUint8 {
		n = uint8(num)
	} else {
		return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

// UInt8p extracts the next uint8 value and returns a pointer variable.
func (r *Reader) UInt8p() (res *uint8, err error) {
	if v, err := r.UInt8(); err == nil {
		res = &v
	}
	return
}

// uint16Seq extracts uint16 values recursively until the closing bracket
// is found, then assigns the elements to the allocated slice.
func (r *Reader) uint16Seq(idx int) (res []uint16, err error) {
	var n uint16
	if n, err = r.UInt16(); err == nil {
		if r.Next() {
			res, err = r.uint16Seq(idx + 1)
		} else {
			res = make([]uint16, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

// UInt16s extracts an array of uint16 values from the data and skips all
// whitespace after it. The values must be enclosed in square brackets "[...]"
// and the values must be separated by commas.
func (r *Reader) UInt16s() (res []uint16, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []uint16{}
			err = r.CloseArray()
		} else if res, err = r.uint16Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

// UInt16 extracts the next uint16 value from the data and skips all
// whitespace after it.
func (r *Reader) UInt16() (n uint16, err error) {
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
	} else if !neg && num <= math.MaxUint16 {
		n = uint16(num)
	} else {
		return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

// UInt16p extracts the next uint16 value and returns a pointer variable.
func (r *Reader) UInt16p() (res *uint16, err error) {
	if v, err := r.UInt16(); err == nil {
		res = &v
	}
	return
}

// uint32Seq extracts uint32 values recursively until the closing bracket
// is found, then assigns the elements to the allocated slice.
func (r *Reader) uint32Seq(idx int) (res []uint32, err error) {
	var n uint32
	if n, err = r.UInt32(); err == nil {
		if r.Next() {
			res, err = r.uint32Seq(idx + 1)
		} else {
			res = make([]uint32, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

// UInt32s extracts an array of uint32 values from the data and skips all
// whitespace after it. The values must be enclosed in square brackets "[...]"
// and the values must be separated by commas.
func (r *Reader) UInt32s() (res []uint32, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []uint32{}
			err = r.CloseArray()
		} else if res, err = r.uint32Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

// UInt32 extracts the next uint32 value from the data and skips all
// whitespace after it.
func (r *Reader) UInt32() (n uint32, err error) {
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
	} else if !neg && num <= math.MaxUint32 {
		n = uint32(num)
	} else {
		return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

// UInt32p extracts the next uint32 value and returns a pointer variable.
func (r *Reader) UInt32p() (res *uint32, err error) {
	if v, err := r.UInt32(); err == nil {
		res = &v
	}
	return
}

// uint64Seq extracts uint64 values recursively  the closing bracket
// is found, then assigns the elements to the allocated slice.
func (r *Reader) uint64Seq(idx int) (res []uint64, err error) {
	var n uint64
	if n, err = r.UInt64(); err == nil {
		if r.Next() {
			res, err = r.uint64Seq(idx + 1)
		} else {
			res = make([]uint64, idx+1)
		}

		if err == nil {
			res[idx] = n
		}
	}
	return
}

// UInt64s extracts an array of uint64 values from the data and skips all
// whitespace after it. The values must be enclosed in square brackets "[...]"
// and the values must be separated by commas.
func (r *Reader) UInt64s() (res []uint64, err error) {
	if err = r.OpenArray(); err == nil {
		if r.Token() == TerminatorToken {
			res = []uint64{}
			err = r.CloseArray()
		} else if res, err = r.uint64Seq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

// UInt64 extracts the next uint64 value from the data and skips all
// whitespace after it.
func (r *Reader) UInt64() (n uint64, err error) {
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
	} else if !neg {
		n = uint64(num)
	} else {
		return 0, newNumberOutOfRangeError(dat[:pos], r.pos)
	}

	r.pos += pos
	r.SkipWhiteSpace()
	return
}

// UInt64p extracts the next uint64 value and returns a pointer variable.
func (r *Reader) UInt64p() (res *uint64, err error) {
	if v, err := r.UInt64(); err == nil {
		res = &v
	}
	return
}
